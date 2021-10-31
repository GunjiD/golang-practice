package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

var port_num = flag.Int("port", 8000, "set the serber port number")

func main() {
	flag.Parse()

	serverPort := "localhost:" + strconv.Itoa(*port_num)

	conn, err := net.Dial("tcp", serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
