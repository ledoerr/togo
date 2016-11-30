package main

import (
	_ "expvar"
	"fmt"
	"net"
	"net/http"
)

func main() {
	sock, err := net.Listen("tcp", "localhost:9123")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("HTTP debug now available at port 9123")
	http.Serve(sock, nil)

}
