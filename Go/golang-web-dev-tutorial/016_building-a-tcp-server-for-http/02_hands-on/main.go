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
		log.Fatalln(err.Error())
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {

			log.Fatalln(err.Error())
			continue
		}

		go handle(conn)

	}

}

func handle(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()

		if i == 0 {
			uri := strings.Fields(ln)

			fmt.Println("Print URI:", uri[1])
		}
		i++
	}

}
