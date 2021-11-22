package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var localhost string
var remotehost string
var departamento string
var respuesta1 string
var respuesta2 string

//TODO: CREAR STRUCT GLOBAL CON RESPUESTA 1 Y RESPUESTA 2

func main() {
	//servidor
	respuesta1 = "null"
	respuesta2 = "null"

	localhost = fmt.Sprintf("localhost:%s", "8090")
	ln, err := net.Listen("tcp", localhost) //punto de conexión
	if err != nil {
		fmt.Println("Falla al resolver la dirección de red:", err.Error())
		os.Exit(1)
	}

	//diferidos
	defer ln.Close()

	for {
		log.Println("antes del accept")
		//aceptar conexión (una)
		con, err := ln.Accept()
		log.Println("despues del accept")

		if err != nil {
			log.Fatal(err)
		}
		//usamos manejador
		go receiver(con)
		//go sender()
	}
}

func receiver(con net.Conn) {

	//lectura de lo que llega al server
	bufferIn := bufio.NewReader(con) //objeto de lectura de conexión

	msg, err := bufferIn.ReadString('\n') //leemos del cliente
	if err != nil {
		log.Fatal(err)
	}

	var puertoEmisor = strings.TrimSpace(msg)[len(strings.TrimSpace(msg))-4 : len(strings.TrimSpace(msg))]
	msg = strings.TrimSpace(msg)[0 : len(strings.TrimSpace(msg))-4]

	if puertoEmisor == "6943" { //Si vino del back, emitir los dos mensajes a los nodos regresiones
		respuesta1 = msg
	}

	if puertoEmisor == "6944" { //Si vino del back, emitir los dos mensajes a los nodos regresiones
		respuesta2 = msg
	}

	if respuesta1 != "null" && respuesta2 != "null" {
		//LÓGICA PARA ENVIAR RESPUESTA A BACK
		answer := "En " + departamento + ", falleceran " + respuesta1 + " personas, de las cuales " + respuesta2 + " no estuvieron vacunadas con dos dosis."
		respuesta1 = "null"
		respuesta2 = "null"
		fmt.Println(answer)
		go sender(answer, "8072")
	}
	//Emisión de respuestas

	if puertoEmisor == "8071" { //Si vino del back, emitir los dos mensajes a los nodos regresiones
		departamento = msg
		go sender(msg, "6943")
		go sender(msg, "6944")
	}

}

func sender(msg string, port string) {
	log.Println("enviando a ", port)

	remotehost = fmt.Sprintf("localhost:%s", port)
	//Conectándonos a nodo regresión
	con, err := net.Dial("tcp", remotehost)
	if err != nil {
		log.Fatal(err)
	}

	//enviamos datos al server
	fmt.Fprintln(con, msg)
	defer con.Close()
}
