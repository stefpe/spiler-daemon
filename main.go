package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/franela/goreq"
)

const (
	Protocol   = "tcp"
	Address    = "127.0.0.1:9001"
	BufferSize = 1024
)

var dataChannel = make(chan string)

//reads from a connection and sends the data to the event loop
func handleClient(conn net.Conn) {
	defer conn.Close()

	var buffer bytes.Buffer

	for {
		tmp := make([]byte, BufferSize)
		n, err := conn.Read(tmp)
		if err != nil && err != io.EOF {
			return
		}

		if n == 0 {
			break
		}
		buffer.Write(tmp) //add chunk to buffer
	}

	dataChannel <- buffer.String()
}

//post data to an endpoint
func postData(data, endPoint string) {
	_, err := goreq.Request{
		Method:      http.MethodPost,
		ContentType: "application/json",
		Accept:      "application/json",
		Uri:         endPoint,
		Timeout:     500 * time.Millisecond,
		Body:        data,
	}.Do()

	if err != nil {
		log.Println(err.Error())
	}
}

//handles incoming data
func handleEvents(endPoint string) {
	for {
		data := <-dataChannel
		log.Println("received message", data)
		postData(data, endPoint)
	}
}

//run tcp server
func runTcpServer(address string) {
	listener, err := net.Listen(Protocol, address)
	defer listener.Close()
	if err != nil {
		log.Fatal("listener error:", err.Error())
	}

	log.Printf("Listen on address %s\n", address)

	for {
		conn, err := listener.Accept() //new client connection
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

//main entry point
func main() {

	address := flag.String("address", Address, "e.g. 127.0.0.1:9001")
	endPoint := flag.String("endpoint", "", "e.g. Portal endpoint 127.0.0.1:8080")
	flag.Parse()

	if len(*endPoint) == 0 {
		log.Fatal("please provide an endpoint which will accept posted data")
	}

	go handleEvents(*endPoint)
	runTcpServer(*address)
}
