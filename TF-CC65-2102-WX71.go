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
	return df
}

func filter_by_department(df dataframe.DataFrame, department string) dataframe.DataFrame {
	df_by_department := df.Filter(
		dataframe.F{
			Colname:    "dpt_cdc",
			Comparator: series.Eq,
			Comparando: department,
		},
	)

	group := df_by_department.GroupBy("mes_desde_inicio_covid")
	aggre := group.Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"id_persona"})

	sorted := aggre.Arrange(
		dataframe.Sort("mes_desde_inicio_covid"),
	)
	return sorted
}

func main() {
	df := read_csv()
	department := "LIMA"
	df_department := filter_by_department(df, department)
	fmt.Println(df_department)

	output, err := os.Create("df_by_department.csv")
	if err != nil {
		log.Fatal(err)
	}
	df_department.WriteCSV(output)

}
