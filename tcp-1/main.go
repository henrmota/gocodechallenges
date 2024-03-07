package main

import (
	"fmt"
	"log"
	"net"
)

func launchClient(c <-chan interface{}) {

	for {
		conn, err := net.Dial("tcp", "localhost:1234")
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received %d, message %s", n, string(buf))
	}
}

func setupServer() {
	list, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	defer list.Close()

	for {
		con, err := list.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(con net.Conn) {
			defer con.Close()

			n, err := con.Write([]byte(con.RemoteAddr().String()))

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Wrote %d bytes to %s\n ", n, con.RemoteAddr().String())
		}(con)
	}
}

func main() {

}
