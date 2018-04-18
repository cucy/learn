package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}

	li.Close()
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		bs := []byte(line)
		r := rot13(bs)
		fmt.Fprintf(conn, "%s - %s\n\n", line, r)
	}

}

func rot13(b []byte) []byte {
	r13 := make([]byte, len(b))
	for i, v := range b {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
