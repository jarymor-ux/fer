package main

import (
	"log"

	"braces.dev/errtrace"
	app "github.com/jarymor-ux/fer/internal/app"
)

func main() {
	if err := app.StartServer(); err != nil {
		log.Fatal(errtrace.Wrap(err))
	}
}
