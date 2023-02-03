package ch8

import (
	"io"
	"log"
	"net"
	"os"
)

func NetCat3a() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	done := make(chan int)
	go func() {
		// output bit
		_, err := io.Copy(os.Stdout, conn) // copy values from connection to stdout
		if err != nil {
			log.Fatal(err)
		}
		log.Println("done") // this line prints done when the server is killed
		done <- 0           //signal the main that its done
	}()
	// input bit
	mustCopy(conn, os.Stdin) // read happens here, from terminal to the connection and is copied to the connection (tcp server)
	<-done                   // here it waits
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
