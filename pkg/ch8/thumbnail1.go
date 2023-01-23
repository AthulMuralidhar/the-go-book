package ch8

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Thumbnail1() {
	makeThumbnails(os.Args[2:])
}

func makeThumbnails(filenames []string) {
	ch := make(chan int)
	for _, filename := range filenames {
		go func(filename string) {
			// passing the filename arg here ensures that the filename from the for loop
			// is the current one and  not the last file in the filenames range
			_, err := ImageFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			ch <- 0
		}(filename)
	}
	for range filenames {
		// this waits for the go routine above to finish
		// and discards the ressult
		<-ch
	}
}

//src: https://github.com/adonovan/gopl.io/blob/master/ch8/thumbnail/thumbnail.go
func Image(src image.Image) image.Image {
	// Compute thumbnail size, preserving aspect ratio.
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // portrait
	} else {
		height = int(128 / aspect) // landscape
	}
	xScale := float64(xs) / float64(width)
	yScale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// a very crude scaling algorithm
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xScale)
			srcy := int(float64(y) * yScale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream reads an image from r and
// writes a thumbnail-size version of it to w.
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 reads an image from infile and writes
// a thumbnail-size version of it to outfile.
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g. "foo.thumb.jpeg".
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // e.g., ".jpg", ".JPEG"
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	log.Printf("made thumbnail: %s\n", outfile)
	return outfile, ImageFile2(outfile, infile)
}
