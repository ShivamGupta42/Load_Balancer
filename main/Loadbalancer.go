package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func initLb() {
	var cfg Config
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return
	}

	ln, err := net.Listen("tcp", ":1994")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Printf("\nListening for connections\n")
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		panic(err)
	}

	//Sending a get request to connection back
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	if err != nil {
		return
	}
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connection status : %s\n", status)
}

func main() {
	initLb()
}
