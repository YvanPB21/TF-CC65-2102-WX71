package main

import (
	"fmt"
	"log"
	"net"
)

var localhost string
var remotehost string

func main() {
	//Este nodo simula un request del api

	localhost = fmt.Sprintf("localhost:%s", "8071")

	remotehost = fmt.Sprintf("localhost:%s", "8090")

	//Conectándonos a nodo server
	con, err := net.Dial("tcp", remotehost)
	if err != nil {
		log.Fatal(err)
	}

	//enviamos datos al server
	fmt.Fprintln(con, "PUNO"+localhost[len(localhost)-4:len(localhost)])

	defer con.Close()

}
