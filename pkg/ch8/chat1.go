package ch8

import (
	"log"
	"net"
)

type client chan string

func Chat1() {
	enteringChan := make(chan client)
	leavingChan := make(chan client)
	messagesChan := make(chan string)

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

}

func broadcaster(messagesChan chan string, enteringChan client) {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messagesChan:
			for c := range clients {
				c <- msg
			}
		case c := <-enteringChan:
			clients[c] = true
		case <-leavingChan:

		}
	}
}
