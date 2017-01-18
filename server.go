/* 18/01/2017 */
/* ®Marian Bergerot*/
package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {

	ln, err := net.Listen("tcp", ":4321")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}


var Clients []*net.Conn
func handleConnection(conn net.Conn) {
	Clients = append(Clients, &conn)
	for {
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
		}
		go SendAll(status)
	}
}

func SendAll(msg string) {
	for _, client := range Clients {
		fmt.Fprintf(*client, msg)
	}
}