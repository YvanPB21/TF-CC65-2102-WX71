package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var (
	x []float64
	y []float64

	data []Persona

	test []float64

	left  int
	rigth int
)

type Persona struct {
	Data1 float64 `json:"height"` //estatura
	Data2 float64 `json:"weight"` //peso
}

func leerCSV() {
	csvFile, err := os.Open("./data/TB_FALLECIDO_HOSP_VAC.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range csvLines {
		if i >= 1 && i < 10 {

			fmt.Println(line[0] + "," + line[1] + "," + line[2] + "," + line[3] + "," + line[4])

			data1, _ := strconv.ParseFloat(line[1], 64)
			data2, _ := strconv.ParseFloat(line[2], 64)
			per := Persona{
				Data1: data1,
				Data2: data2,
			}
			data = append(data, per)
		}
	}

}

func main() {
	leerCSV()
}
