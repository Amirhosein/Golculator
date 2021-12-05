package modules

import (
	"fmt"
	"net"
)

func RunClient(port string) {
	var msg string

	hostName := "localhost"

	conn, err := net.Dial("tcp", hostName+":"+port)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Connection established between %s and localhost.\n", hostName)
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

	for {
		fmt.Print("Enter message: ")
		fmt.Scanln(&msg)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}

		buf := make([]byte, 4096)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Print("Message received:")
			fmt.Println(string(buf))
		}
	}

}
