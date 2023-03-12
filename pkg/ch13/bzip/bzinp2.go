package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include <bzlib.h>
#include <stdlib.h>
bz_stream* bz2alloc() { return calloc(1, sizeof(bz_stream)); }
int bz2compress(bz_stream *s, int action,
                char *in, unsigned *inlen, char *out, unsigned *outlen);
void bz2free(bz_stream* s) { free(s); }
*/

import "C"
import (
	"io"
	"unsafe"
)

type writer struct {
	w      io.Writer
	stream *C.bz_stream
	outBuf [64 * 1024]byte
}

func (w *writer) Write(data []byte) (n int, err error) {
	if w.stream == nil {
		panic("closed")
	}
	var total int
	for len(data) > 0 {
		inputLength, outputLength := C.uint(len(data), C.uint(cap(w.outBuf)))
		C.bzrcompress(w.stream, C.BZ_RUN, (*C.char)(unsafe.Pointer(&data[0])), &inputLength, (*C.char)(unsafe.Pointer(&w.outBuf)), &outputLength)

		total += int(inputLength)
		data = data[inputLength:]

		if _, err := w.w.Write(w.outBuf[:outputLength]); err != nil {
			return total, nil
		}

	}
	return total, nil
}

func (w *writer) Close() error {
	if w.stream == nil {
		panic("closed")
	}
	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		C.bz2free(w.stream)
		w.stream = nil
	}()
	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outBuf))
		r := C.bz2compress(w.stream, C.BZ_FINISH, nil, &inlen,
			(*C.char)(unsafe.Pointer(&w.outBuf)), &outlen)
		if _, err := w.w.Write(w.outBuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}
