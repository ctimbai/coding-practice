package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed! err: ", err)
		return
	}

	defer listen.Close()
	fmt.Println("tcp server is running...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed! err: ", err);
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		for {
			n, err := reader.Read(buf[:])
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("read from client failed, err: ", err)
				break
			}
			recvStr := string(buf[:n])
			fmt.Println("recv data from client: ", recvStr)
			//conn.Write([]byte(recvStr))
		}
	}
}
