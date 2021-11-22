package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/sajari/regression"
)

var localhost string
var remotehost string
var department string
var respuesta string

func procesar_csv() dataframe.DataFrame {
	csv, err := os.Open("../data/TB_FALLECIDO_HOSP_VAC.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(csv)
	df_by_department_vaccinated := df.Filter(
		dataframe.F{
			Colname:    "dpt_cdc",
			Comparator: series.Eq,
			Comparando: department,
		},
	)

	df_not_vaccinated := df_by_department_vaccinated.Filter(
		dataframe.F{
			Colname:    "fabricante_dosis2",
			Comparator: series.Eq,
			Comparando: "",
		},
	)

	group := df_not_vaccinated.GroupBy("mes_desde_inicio_covid")
	aggre := group.Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"id_persona"})

	sorted := aggre.Arrange(
		dataframe.Sort("mes_desde_inicio_covid"),
	)

	output, err := os.Create("../data/df_by_department_vaccinated.csv")
	if err != nil {
		log.Fatal(err)
	}
	sorted.WriteCSV(output)

	return sorted
}

func linear_regression() {

	procesar_csv()

	f, err := os.Open("../data/df_by_department_vaccinated.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := csv.NewReader(f)
	data.FieldsPerRecord = 2

	records, err := data.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var r regression.Regression

	r.SetObserved("dead")
	r.SetVar(0, "month-covid")
	r.SetVar(1, "month-covid2")

	for i, record := range records {
		if i == 0 {
			continue
		}

		dead, err := strconv.ParseFloat(records[i][0], 64)

		if err != nil {
			log.Fatal(err)
		}

		monthCovid, err := strconv.ParseFloat(record[1], 64)

		if err != nil {
			log.Fatal(err)
		}
		r.Train(regression.DataPoint(dead, []float64{monthCovid, math.Pow(monthCovid, 2)}))
	}

	r.Run()

	pred, err := r.Predict([]float64{21, math.Pow(21, 2)})
	if err != nil {
		log.Fatal(err)
	}

	helper := -1

	for i, record := range records {
		if i == 0 {
			continue
		}
		record = record
		helper++
	}

	if pred < 0 {
		var predFloat, _ = strconv.ParseFloat(records[helper][0], 64)
		respuesta = strconv.Itoa(int(predFloat))
	} else {
		respuesta = strconv.Itoa(int(pred))
	}
}

func main() {

	localhost = fmt.Sprintf("localhost:%s", "6944")
	ln, err := net.Listen("tcp", localhost) //punto de conexión
	if err != nil {
		fmt.Println("Falla al resolver la dirección de red:", err.Error())
		os.Exit(1)
	}

	//diferidos
	defer ln.Close()

	for {
		//aceptar conexión (una)
		con, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//usamos manejador
		receiver(con)
		linear_regression()
		sender()
	}
}

func receiver(con net.Conn) {
	//lectura de lo que llega al server
	bufferIn := bufio.NewReader(con) //objeto de lectura de conexión

	msg, err := bufferIn.ReadString('\n') //leemos del cliente
	if err != nil {
		log.Fatal(err)
	}
	department = strings.TrimSpace(msg)
	fmt.Println(department) //se imprime mensaje del cliente
}

func sender() {
	remotehost = fmt.Sprintf("localhost:%s", "8090")
	//Conectándonos a nodo regresión
	con, err := net.Dial("tcp", remotehost)
	if err != nil {
		log.Fatal(err)
	}

	//enviamos datos al server

	fmt.Fprintln(con, respuesta+localhost[len(localhost)-4:len(localhost)])
	//fmt.Fprintln(con, respuesta+"8090")
	defer con.Close()
}
