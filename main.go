package main

import (
	"log"

	"github.com/DeneesK/packet-manager/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
