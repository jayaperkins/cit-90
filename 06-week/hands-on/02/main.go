package main

import (
	"io"
	"log"
	"net"
	"io/ioutil"
	"fmt"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "You are connected!")

		bs, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(bs))

		//This line is never written because ioutil is constantly reading from the connection.
		// You never break out of the loop and reach io.WriteString
		io.WriteString(conn, "Hello")
		conn.Close()

	}

}

