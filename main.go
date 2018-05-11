package main

import (
	"log"

	"github.com/josebur86/osmose/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
