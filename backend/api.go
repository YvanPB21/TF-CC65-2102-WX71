package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/rs/cors"
)

//estructura
type Prediction struct {
	Predict string `json:"prediction"`
}

var listPredictions []Prediction
var localhost string
var remotehost string

func loadData(predi string) {
	listPredictions = []Prediction{
		{predi},
	}
}

func predictDeathsNumber(response http.ResponseWriter, request *http.Request) {
	log.Println("Se departamento para predecir")
	//recuperar el parámetro de envío
	dep := request.FormValue("dep")

	log.Println(dep)
	sendToController(dep)
	receiver()
	//respuesta
	response.Header().Set("Content-Type", "application/json")

	for _, oPrediction := range listPredictions {

		jsonBytes, _ := json.MarshalIndent(oPrediction, "", " ")
		io.WriteString(response, string(jsonBytes))

	}

}

func receiver() {
	localhost = fmt.Sprintf("localhost:%s", "8072")
	ln, err := net.Listen("tcp", localhost) //punto de conexión
	if err != nil {
		fmt.Println("Falla al resolver la dirección de red:", err.Error())
		os.Exit(1)
	}
	//diferidos
	defer ln.Close()

	//aceptar conexión (una)
	con, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	//lectura de lo que llega al server
	bufferIn := bufio.NewReader(con) //objeto de lectura de conexión

	msg, err := bufferIn.ReadString('\n') //leemos del cliente
	if err != nil {
		log.Fatal(err)
	}

	loadData(msg)
}

func sendToController(msg string) {
	log.Println("enviando a controller")

	port := "8090"
	remotehost = fmt.Sprintf("localhost:%s", port)
	//Conectándonos a nodo regresión
	con, err := net.Dial("tcp", remotehost)
	if err != nil {
		log.Fatal(err)
	}

	//enviamos datos al server
	fmt.Fprintln(con, msg+"8071")
	defer con.Close()
}

func requestsHandler() {
	//enrutador
	mux := http.NewServeMux()
	//endpoints
	mux.HandleFunc("/prediction", predictDeathsNumber)
	handler := cors.AllowAll().Handler(mux)
	log.Fatal(http.ListenAndServe(":8071", handler))
}

func main() {
	requestsHandler()
}
