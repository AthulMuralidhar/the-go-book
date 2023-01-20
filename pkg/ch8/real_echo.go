package ch8

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func TheREALEchoServer1() {
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
		go handleEchoConnection(conn)

	}
}

func handleEchoConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
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

func echo(conn net.Conn, text string, duration time.Duration) {
	_, err := fmt.Fprintln(conn, "\t", strings.ToUpper(text))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(duration)
	_, err = fmt.Fprintln(conn, "\t", text)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(duration)
	_, err = fmt.Fprintln(conn, "\t", strings.ToLower(text))
	if err != nil {
		log.Fatal(err)
	}
}
