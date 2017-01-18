package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

var name string

func main() {

	fmt.Print("Pseudo: ")
	fmt.Scanln(&name)
	fmt.Println(name +": ")

	conn, err := net.Dial("tcp", "localhost:4321")
	if err != nil {
	}
	go SendMsg(conn)
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
		}
		fmt.Print(msg)
	}
}

func SendMsg(conn net.Conn) {
	for  {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, name +": "+ text)
	}
}
