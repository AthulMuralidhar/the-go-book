package ch8

import (
	"io"
	"log"
	"net"
	"time"
)

func ClockServer1() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn) // adding go here makes it a concurrent function call
		// try removing it and running in the terminal, u will find that the server can only handle one netcat call after the other

	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:01:05\n"))
		if err != nil {
			log.Println("client disconnected")
			return
		}
		time.Sleep(time.Second * 1)
	}
}
