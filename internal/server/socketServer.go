package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

type SocketServer struct {
	Port string
}

func (ss SocketServer) RunServer() {
	ln, err := net.Listen("tcp", ":"+ss.Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server up and listening on port " + ss.Port)

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
	var num float64
	var operator string

	log.Printf("Client %v connected.", c.RemoteAddr())
	err := sendMessage(c, "Welcome to the golculator!\nFirst send me a number and then an operator.\nThen send me another number.\n")
	if err != nil {
		log.Printf("Error sending message to client: %s", err)
		c.Close()
		return
	}
	for {
		clientMessage, err := readMessage(c)
		if err != nil {
			break
		}
		log.Printf("Client %v sent: %s", c.RemoteAddr(), clientMessage)

		if clientMessage[0] == '+' || clientMessage[0] == '-' || clientMessage[0] == '*' || clientMessage[0] == '/' {

			operator = clientMessage[0:1]
			secondNum, err := ParseNumber(clientMessage[1:])
			if err != nil {
				log.Print(err)
				return
			}
			num = calculate(num, secondNum, operator)

			err = sendMessage(c, "RESULT: "+fmt.Sprintf("%g", num))
			if err != nil {
				break
			}
			log.Printf("Server response: %s", "RESULT: "+fmt.Sprintf("%f", num))

		} else {
			num, err = ParseNumber(clientMessage)
			if err != nil {
				log.Print(err)
				return
			}

			err = sendMessage(c, clientMessage)
			if err != nil {
				break
			}
			log.Printf("Server response: %s", clientMessage)
		}
	}
	log.Printf("Connection from %v closed.", c.RemoteAddr())
}

func sendMessage(c net.Conn, message string) error {
	_, err := c.Write([]byte(message))
	if err != nil {
		log.Printf("Error writing to client: %s", err)
		c.Close()
		return err
	}
	return nil
}

func readMessage(c net.Conn) (string, error) {
	buffer := make([]byte, 4096)
	n, err := c.Read(buffer)
	if err != nil || n == 0 {
		log.Printf("Error reading from client: %s", err)
		c.Close()
		return "", err
	}
	return string(buffer[:n]), nil
}

func ParseNumber(message string) (float64, error) {
	num, err := strconv.ParseFloat(message, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func calculate(num1 float64, num2 float64, operator string) float64 {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0
}
