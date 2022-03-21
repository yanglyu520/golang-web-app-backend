package main

import (
	"bufio"
	"fmt"
	"net"
	//"io"
	"strings"
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

		//write to the connection
		//io.WriteString(conn, "hello message from the server\n")??????

		//go routine to handle this
		go handle(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	
	//read the http request
	request(conn)

	//write the http response
	respond(conn)
}

func request(conn net.Conn) {
	//set up a flag
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		//print all the http request information
	  fmt.Println(ln)

		if i == 0 {
			//request line METHOD+URL+HTTPversion
			m := strings.Fields(ln)[0]
			fmt.Println("**request method is**", m)
			u := strings.Fields(ln)[1]
			fmt.Println("**request url is**", u)
		}
		if ln == "" {
      //after the request line and headers, there is a blank line
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
  body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	//write to the connection the http response line HTTPVERSION + statusCode + reason
	fmt.Fprint(conn, "HTTP/1.1 200 ok\n")

	//write to the connection the http response header
	fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\n")

  //the separator between header and body
	fmt.Fprint(conn, "\n")

	//write to the connection the http response body
	fmt.Fprint(conn, body)
}
