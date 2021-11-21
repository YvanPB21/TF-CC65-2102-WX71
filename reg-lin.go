package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {

	f, err := os.Open("./df_by_department copy.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	friosData := csv.NewReader(f)
	friosData.FieldsPerRecord = 2

	records, err := friosData.ReadAll()
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
		fmt.Println("dead", records[i][0], "month", record[1])
		helper++
	}

	if pred < 0 {
		fmt.Println("Prediccion->", records[helper][0])
	} else {
		fmt.Println("Prediccion->", pred)
	}

}
