package main

import (
	"github.com/c12s/stellar/model"
	"github.com/c12s/stellar/server"
	"log"
)

func main() {

	// Load configurations
	conf, err := model.ConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	//Start Server
	server.Run(conf)
}
