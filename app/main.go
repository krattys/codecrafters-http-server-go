package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	reqLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
	}
	reqLine = strings.TrimRight(reqLine, "\r\n")
	reqSlice := strings.Split(reqLine, " ")
	var statusCode string = "404 Not Found"
	if reqSlice[1] == "/" {
		statusCode = "200 OK"
	}

	resp := fmt.Sprintf("HTTP/1.1 %s\r\n\r\n", statusCode)
	_, err = conn.Write([]byte(resp))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		os.Exit(1)
	}
}
