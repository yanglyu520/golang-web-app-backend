package main

import (
	//"bufio"
	"fmt"
	"net"
	"io"
)

func main() {
	//create a listener that will listen for any incoming calls from the client
	li, err := net.Listen("tcp", ":8080")
	checkErr(err)

	//need to close the listener after the connection is finished
	defer li.Close()

	for {
		//create a connection, which is a writer and a reader as well
		conn, err := li.Accept()
		checkErr(err)
		//check the address of the link
		li.Addr()
    defer conn.Close()
		//write to the connection
		io.WriteString(conn, "hello from the server")

	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}