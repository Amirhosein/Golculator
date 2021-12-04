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

	buffer := make([]byte, 4096)

	for {
		n, err := c.Read(buffer)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		_, err = c.Write(buffer[0:n])
		if err != nil {
			c.Close()
			break
		}
	}
	log.Printf("Connection from %v closed.", c.RemoteAddr())
}
