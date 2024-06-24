package main

import (
	"github.com/ghoshRitesh12/nami/lib"
)

func main() {
	flags := lib.GetParsedFlags()
	routeGenerator := lib.NewRouteGenerator()

	// fmt.Printf("%+v\n", flags)

	routeGenerator.
		AddMainDirPath(flags.MainDirPath).
		AddRouteHandlerInfo(
			flags.RouteHandlerTypeImport,
			flags.RouteHandlerType,
		).
		AddRouterStructPointerType(
			flags.RouterStructPointerType,
		).
		GenerateRoutes()
}
