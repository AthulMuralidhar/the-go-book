package ch8

import (
	"bufio"
	"log"
	"net"
	"time"
)

func Reverb2() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("connection aborted")
			continue
		}
		go handleReverb(conn)
	}
}

func handleReverb(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 1*time.Second)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
