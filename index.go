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

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("Handling new connection...")
		go handleConnection(conn)
	}
}

func Write(conn net.Conn, content string) (int, error) {
	writer := bufio.NewWriter(conn)
	number, err := writer.WriteString(content + "\n")
	if err == nil {
		err = writer.Flush()
	}

	return number, err
}

func WriteError(conn net.Conn, err error) (int, error) {
	ret := map[string]string{
		"type":    "error",
		"message": err.Error(),
	}

	retJSON, err := json.Marshal(ret)
	return Write(conn, string(retJSON))
}

func handleConnection(conn net.Conn) {
	bufReader := bufio.NewReader(conn)

	// Close connection when this function ends
	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}()

	for {
		bytes, err := bufReader.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			WriteError(conn, InternalError)
			return
		}

		var req map[string]interface{}
		err = json.Unmarshal(bytes, &req)
		if err != nil {
			fmt.Println(err)
			WriteError(conn, InvalidInputError)
		}

		err, res := Route(req)
		if err == nil {
			res["type"] = req["type"]
			resultJSON, _ := json.Marshal(res)
			Write(conn, string(resultJSON))
			continue
		}

		// Houston, we have a problem
		switch err {
		case InvalidInputError, IncopatibleVersions:
			WriteError(conn, err)
			return
		default:
			WriteError(conn, err)
		}

	}
}
