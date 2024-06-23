package main

import (
	"fmt"

	"github.com/ghoshRitesh12/nami/lib"
)

func main() {
	flags := lib.GetParsedFlags()
	routeGenerator := lib.GetRouteGenerator()

	routeGenerator.
		AddMainDirPath(flags.MainDirPath).
		AddRouteHandlerInfo(
			flags.RouteHandlerTypeImport,
			flags.RouteHandlerType,
		)
		// GenerateRoutes()

	fmt.Printf("%+v\n", flags)
}
