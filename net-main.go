// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
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

		//writeString function write the string to a writer
		io.WriteString(conn, "hellllo, I write to the server")
		//alternatively you can use fmt.Fprintf function
		fmt.Fprintf(conn, "hoow are you")
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
