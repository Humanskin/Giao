package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte

		b, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("error from buf")
			break
		}

		recvStr := string(buf[:b])
		fmt.Println("收到客户端发来的数据：", recvStr)
		//conn.Write([]byte(recvStr))
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(string(inputInfo)) == "Q" {
			return
		}
		conn.Write([]byte(inputInfo))

	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("error from listen")
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("error from Accept")
			continue
		}

		go process(conn)
	}


}
