package ch8

import (
	"bufio"
	"log"
	"net"
	"golang.org/x/sync/errgroup"
)

type client chan string

func Chat1() {
	// these are channels of chans
	var eg errgroup.Group
	enteringChan := make(chan client)
	leavingChan := make(chan client)
	messagesChan := make(chan string)

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster(messagesChan, enteringChan, leavingChan)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		err := go handleConn(conn) // per client routine
	}

}

func handleConn(conn net.Conn, enteringChan, leavingChan chan client) error {
	ch := make(chan string)
	go clientWriter(conn, ch)

	connAddr := conn.RemoteAddr().String()
	ch <- "you are: " + connAddr
	enteringChan <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- connAddr + ":" + input.Text()
	}
	leavingChan <- ch
	messages <- connAddr + "has left"

	return conn.Close()
}

func clientWriter(conn net.Conn, msgsChan chan string) {

}

func broadcaster(messagesChan chan string, enteringChan, leavingChan chan client) {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messagesChan:
			for c := range clients {
				c <- msg // broadcast on a common msg channel
			}
		case c := <-enteringChan:
			clients[c] = true // update the client set
		case c := <-leavingChan:
			delete(clients, c)
			close(c) // close the client set

		}
	}
}
