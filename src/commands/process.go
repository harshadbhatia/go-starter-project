package commands

import (
	"os"
	"fmt"
	"log"
	
	"github.com/urfave/cli"
	"github.com/rocketlaunchr/dataframe-go/imports"
)

func processCSV(c *cli.Context) {
	csvFile, err := os.Open("data/example.csv")
	if err != nil { 
		log.Fatal(err)	
	}

	df, err := imports.LoadFromCSV(ctx, csvFile, imports.CSVLoadOptions{InferDataTypes: true, NilValue: &[]string{"NA"}[0]})
	if err != nil {
		panic(err)
	}

	fmt.Print(df.Table())
}