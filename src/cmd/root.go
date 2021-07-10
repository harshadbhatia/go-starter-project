package cmd 

import (
	"fmt"
	"os"
  
	"github.com/spf13/cobra"
)


  
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(process)
}

func initConfig() {
	// Use this in future
}

var rootCmd = &cobra.Command{
	Use:   "dcot",
	Short: "DCOT is Data Upload Tool",
	Long: `Allows upload to your favourite cloud.
	Complete documentation is available at XXX`,
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}