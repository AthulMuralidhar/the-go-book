package ch8

import (
	"bufio"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
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
		// one way to handle errors is through error groups as shown:
		eg.Go(func() error {
			return handleConn(conn, messagesChan, enteringChan, leavingChan) // per client routine
		})

		if err := eg.Wait(); err != nil {
			log.Println("error during handleConn")
			log.Fatal(err)
		}

	}

}

func handleConn(conn net.Conn, messagesChan chan<- string, enteringChan, leavingChan chan client) error {
	log.Println("handle-connection is up ")
	// the other way to handle errors is to have an error channel
	//errCh := make(chan error, 1) // make it buffered so that we only have 1 err
	ch := make(chan string)
	//go func() {
	//	errCh <- clientWriter(conn, ch)
	//}()
	//if err := <-errCh; err != nil {
	//	return err
	//}
	go clientWriter(conn, ch)

	connAddr := conn.RemoteAddr().String()
	ch <- "you are: " + connAddr
	enteringChan <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		log.Printf("incoming input: %s\n", input.Text())
		messagesChan <- connAddr + ":" + input.Text()
	}
	leavingChan <- ch
	messagesChan <- connAddr + "has left"

	return conn.Close()
}

func clientWriter(conn net.Conn, msgChan chan string) {
	for msg := range msgChan {
		_, _ = fmt.Fprint(conn, msg)
	}
}

func broadcaster(messagesChan chan string, enteringChan, leavingChan chan client) {
	log.Println("broadcaster is up")
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messagesChan:
			log.Println("ready for printing")
			for c := range clients {
				c <- msg // broadcast on a common msg channel
			}
		case c := <-enteringChan:
			log.Println("ready for entering chan")
			clients[c] = true // update the client set
		case c := <-leavingChan:
			log.Println("ready for leaving chan")
			delete(clients, c)
			close(c) // close the client set

		}
	}
}
