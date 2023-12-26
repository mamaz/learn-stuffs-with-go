package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"socket-connection/util"
	"time"
)

func main() {
	// create socket server
	ln, err := net.Listen("tcp", ":8000")
	util.LogFatalIfErr("error on listening to tcp port 8000", err)
	log.Println("listening to port :8000")

	for {
		conn, err := ln.Accept()
		util.LogFatalIfErr("error on waiting for connection", err)

		log.Printf("accepting new connection from %s\n", conn.RemoteAddr().String())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	message, err := bufio.NewReader(conn).ReadString('\n')
	if errors.Is(err, io.EOF) {
		util.LogFatalIfErr("gets EOF", err)
	}

	log.Println("message received: ", message)

	conn.Write([]byte(fmt.Sprintf("%s is received with timestamp: %d", message, time.Now().Unix())))
}
