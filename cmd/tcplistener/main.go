// implement the http server functions
package main

import (
	"fmt"
	"log"
	"net"
	"sort"

	"_http_protocol_1.1/internal/request"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("connection accepted")
		req, err := request.RequestFromReader(conn)
		if err != nil {
			log.Println(err)
			_ = conn.Close()
			continue
		}

		fmt.Println("Request line:")
		fmt.Printf("- Method: %s\n", req.RequestLine.Method)
		fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)

		fmt.Println("Headers:")
		keys := make([]string, 0, len(req.Headers))
		for k := range req.Headers {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Printf("- %s: %s\n", k, req.Headers[k])
		}

		fmt.Println("Body:")
		fmt.Println(string(req.Body))

		body := "received " + req.RequestLine.RequestTarget + "\n"
		if _, err := fmt.Fprintf(
			conn,
			"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s",
			len(body),
			body,
		); err != nil {
			log.Println(err)
		}
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}
}
