package ch8

import (
	"io"
	"log"
	"net"
	"os"
)

func Ex8_3() {
	conn, err := net.DialTCP("localhost:8000", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	go func() {
		// output bit
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("done")
		err = conn.CloseWrite()
		if err != nil {
			log.Fatal(err)
		}
		err = conn.CloseWrite()
		if err != nil {
			log.Fatal(err)
		}
	}()
	// input bit
	mustCopy(conn, os.Stdin)
}
