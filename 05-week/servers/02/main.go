package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
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
			log.Println(err)
		}

		io.WriteString(conn, "You are connected!")

		bs, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(bs))
		io.WriteString(conn, "I see that you connected")
		conn.Close()

	}

}
