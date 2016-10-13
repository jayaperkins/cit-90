package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}

	io.WriteString(conn, "hello from calling in!\n")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		break
	}
	conn.Close()

}

