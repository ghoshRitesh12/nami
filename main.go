package nami

import (
	"github.com/ghoshRitesh12/nami/internal"
)

// only hosts route generator config
type RouteGenerator struct {
	RouteHandlerType string
	mainDirPath      string
	// routeHandlerPkgName string
}

func NewRouteGenerator() *RouteGenerator {
	return &RouteGenerator{
		mainDirPath:      "./" + internal.MAIN_DIR_NAME,
		RouteHandlerType: "any",
	}
}
