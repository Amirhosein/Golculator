package modules

import (
	"fmt"
	"log"
	"net"
)

func RunServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server up and listening on port " + port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {

	log.Printf("Client %v connected.", c.RemoteAddr())
	sendMessage(c, "Welcome to the golculator!\nFirst send me a number and then an operator.\nThen send me another number.\n")

	for {
		clientMessage := readMessage(c)
		log.Printf("Client %v sent: %s", c.RemoteAddr(), clientMessage)

		sendMessage(c, clientMessage)
		log.Printf("Server response: %s", clientMessage)
	}
	log.Printf("Connection from %v closed.", c.RemoteAddr())
}

func sendMessage(c net.Conn, message string) {
	_, err := c.Write([]byte(message))
	if err != nil {
		log.Printf("Error writing to client: %s", err)
		c.Close()
	}
}

func readMessage(c net.Conn) string {
	buffer := make([]byte, 4096)
	n, err := c.Read(buffer)
	if err != nil || n == 0 {
		log.Printf("Error reading from client: %s", err)
		c.Close()
		return ""
	}
	return string(buffer[:n])
}
