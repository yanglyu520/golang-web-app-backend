// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello, 世界")
	//create a listener that will listen for any incoming calls from the client
	li, err := net.Listen("tcp", ":8080")

	checkErr(err)

	//need to close the listener after the connection is finished
	defer li.Close()

	for {
		//create a connection, which is a writer and a reader as well
		conn, err := li.Accept()
		checkErr(err)
		li.Addr()
		go handle(conn)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("meow")
}
