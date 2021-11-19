package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func read_csv() dataframe.DataFrame {
	csv, err := os.Open("./data/TB_FALLECIDO_HOSP_VAC.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(csv)
	df_by_department := df.Filter(
		dataframe.F{"dpt_cdc", series.Eq, "department"},
		dataframe.F{"sexo", series.Eq, "M"},
	)
	return df_by_department
}

/*
func filter_by_department(df dataframe.DataFrame, department string) dataframe.DataFrame {
	df_by_department := df.Filter(
		// dataframe.F{"dpt_cdc", series.Eq, "department"},
	)
	return df_by_department
}
*/

func main() {
	df := read_csv()
	fmt.Println(df)
}
