package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
)

func main() {
	f, err := os.Open("./df_by_department.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	deadData := csv.NewReader(f)
	deadData.FieldsPerRecord = 2
	records, err := deadData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	header := records[0]
	shuffled := make([][]string, len(records)-1)
	perm := rand.Perm(len(records) - 1)
	for i, v := range perm {
		shuffled[v] = records[i+1]
	}
	trainingIdx := (len(shuffled)) * 3 / 4
	trainingSet := shuffled[1 : trainingIdx+1]
	testingSet := records[len(records)/2:]

	sets := map[string][][]string{
		"./data/training.csv": trainingSet,
		"./data/testing.csv":  testingSet,
	}
	for fn, dataset := range sets {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		out := csv.NewWriter(f)
		if err := out.Write(header); err != nil {
			log.Fatal(err)
		}

		if err := out.WriteAll(dataset); err != nil {
			log.Fatal(err)
		}
		out.Flush()
	}
}
