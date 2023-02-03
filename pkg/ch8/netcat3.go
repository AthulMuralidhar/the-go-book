package ch8

import (
	"io"
	"log"
	"net"
	"os"
)

func NetCat3() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	<-done // wait for background goroutine to finish
}
