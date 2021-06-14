package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

const (
	PORT = ":8888"
)

func main() {
	// Start listening to port 8888 for TCP connection
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		go handleConnection(conn)
	}
}

func Write(conn net.Conn, content string) (int, error) {
	writer := bufio.NewWriter(conn)
	number, err := writer.WriteString(content)
	if err == nil {
		err = writer.Flush()
	}

	return number, err
}

func handleConnection(conn net.Conn) {
	fmt.Println("Handling new connection...")

	// Close connection when this function ends
	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}()

	bufReader := bufio.NewReader(conn)

	for {
		bytes, err := bufReader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("Connection Closed")
			return
		}
		if err != nil {
			fmt.Println("Buffer problem", err)
			return
		}

		var req map[string]interface{}
		err = json.Unmarshal(bytes, &req)
		if err != nil {
			fmt.Println("Cannot Parse JSON")
			continue
		}

		err, res := Route(req)
		if err != nil {
			fmt.Println("Error in handling message")
			return
		}

		fltB, _ := json.Marshal(res)
		Write(conn, string(fltB))
	}
}
