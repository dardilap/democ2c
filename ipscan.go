package main

import (
	"fmt"
	"net"
)

func main() {

	var targetHostAndPort string

	for i := 1; i <= 1024; i++ {

		targetHostAndPort = fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", targetHostAndPort)

		if err != nil {
			continue

		}

		conn.Close()
		fmt.Printf("%d open\n", i)

	}

}
