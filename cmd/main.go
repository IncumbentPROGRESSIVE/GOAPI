package main

import (
	"log"

	"github.com/incumbentPROGERESSIVE/GOAPI/cmd/api"
)
func main () {
	server := api.NewAPIsever(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
