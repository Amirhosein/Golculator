package client

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type SocketClient struct {
	Port string
}

func (sc SocketClient) RunClient() {
	var msg string
	var operator string
	var num bool

	hostName := "localhost"

	conn, err := net.Dial("tcp", hostName+":"+sc.Port)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Connection established between %s and localhost.\n", hostName)
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())

	msg, err = readMessage(conn)
	if err != nil {
		return
	}
	fmt.Print(msg)

	for {
		fmt.Scanln(&msg)

		if isOperator(msg) && !num {
			fmt.Println("You can't use operators without a number.")
			continue
		} else if isOperator(msg) && num {
			operator = msg
			fmt.Scanln(&msg)
			msg = operator + msg
		} else if _, err := strconv.ParseFloat(msg, 64); err != nil {
			fmt.Println("Invalid message. Please enter a number.")
			continue
		}
		num = true

		err = sendMessage(conn, msg)
		if err != nil {
			return
		}

		msg, err = readMessage(conn)
		if err != nil {
			return
		}
		// if msg contains "RESULT", then we have the result and print the message
		if strings.Contains(msg, "RESULT") {
			fmt.Println(msg)
		}
	}
}

func isOperator(msg string) bool {
	return msg == "+" || msg == "-" || msg == "*" || msg == "/"
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
