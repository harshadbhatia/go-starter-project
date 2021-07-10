package cmd

import (
	"os"
	"fmt"
	"log"
	"context"

	"strings"
	
	"github.com/spf13/cobra"
	"github.com/rocketlaunchr/dataframe-go/imports"
)

var ctx = context.Background()

func processCSV(cmd *cobra.Command, args []string) {
	csvFile, err := os.Open(strings.Join(args, ""))
	if err != nil { 
		log.Fatal(err)	
	}

	df, err := imports.LoadFromCSV(ctx, csvFile, imports.CSVLoadOptions{InferDataTypes: true, NilValue: &[]string{"NA"}[0]})
	if err != nil {
		panic(err)
	}

	fmt.Print(df.Table())
}

var process = &cobra.Command{
	Use:   "process [csv file to process]",
    Short: "Process CSV file for upload",
    Long: `process is to process CSV file before upload.
		It does bunch of preprocessing.`,
    Args: cobra.MinimumNArgs(1),
    Run: processCSV,
}