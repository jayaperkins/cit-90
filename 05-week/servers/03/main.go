package main

import (
	"bufio"
	"fmt"
	"io"
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

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				break
			}
		}

		io.WriteString(conn, "hello")

		conn.Close()

	}

}
