package main

import (
	"os"
	"log"

	c "dcot/commands"

)

func main() {
	c.Info()
	c.CommandList()
	err := c.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

