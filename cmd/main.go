package main

import (
	"github.com/ghoshRitesh12/nami/internal"
)

func main() {
	tg := internal.GetTemplateGenerator()
	tg.GenerateRoutes()

	// nami := nami.NewRouteGenerator()
}
