package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"socket-connection/util"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	util.LogFatalIfErr("error on connecting to port 8000", err)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("text to send: ")

		text, err := reader.ReadString('\n')
		util.LogFatalIfErr("error on read string at client", err)

		log.Printf("text is : %v", text)

		// send to server over socket
		fmt.Fprintf(conn, text+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		util.LogFatalIfErr("error on reading reply", err)

		fmt.Printf("received message: %v\n", message)

		conn.Close()
		break
	}
}
