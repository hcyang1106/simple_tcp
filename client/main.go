package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString error, err=", err)
		}

		line = strings.Trim(line, " \n")
		if line == "exit" {
			return
		}

		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write error, err=", err)
		}
		fmt.Printf("client sends %d bytes of data\n", n)
	}
}
