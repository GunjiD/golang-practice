package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

// 問題8.1 ポート番号を受け付けるようにする

var port_num = flag.Int("port", 8000, "set the serber port number")

func main() {
	flag.Parse()

	serverPort := "localhost:" + strconv.Itoa(*port_num)

	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例: 接続が切れた
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // クライアントとの接続が切れた
		}
		time.Sleep(1 * time.Second)
	}
}
