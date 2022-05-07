package main

import (
	"fmt"
	_ "io"
	"net"
)

func process(conn net.Conn) {
	//循環接收client的訊息
	defer conn.Close() // close conn

	for {

		fmt.Printf("server在等待client %s 發送訊息\n", conn.RemoteAddr().String())

		buf := make([]byte, 1024)
		fmt.Println("blocking...")
		n, err := conn.Read(buf) // 阻塞
		if err != nil {
			fmt.Printf("client退出 err=%v\n", err)
			return
		}
		// show message
		fmt.Println(string(buf[:n]))
	}

}

func main() {
	fmt.Println("server listening....")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen error, err=", err)
		return
	}
	defer listen.Close()


	for {
		//等待client連接
		fmt.Println("等待client連接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() error, err=", err)
		} else {
			fmt.Printf("Accept() success conn=%v, client ip=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
