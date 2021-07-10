package commands

import (
	"github.com/urfave/cli"
)

func CommandList() {
	App.Commands = []cli.Command{
	  {
		Name:    "process",
		Aliases: []string{"p"},
		Usage:   "Preprocess file before upload",
		Action: processCSV,
	  },

	}
  }