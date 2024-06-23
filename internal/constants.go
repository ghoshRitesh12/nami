package internal

import (
	"errors"
	"fmt"
)

const (
	MAIN_DIR_NAME = "routes"
)

const (
	ROUTES_TMPL_FILE string = "routes.gen.tmpl"
	OUTPUT_FILENAME  string = "routes.gen.go"
	MAIN_TMPL_FILE   string = "main.gen.go"
)

var (
	POSSIBLE_ROUTE_FILENAME = map[string]bool{
		"route.get.go":    true,
		"route.put.go":    true,
		"route.post.go":   true,
		"route.head.go":   true,
		"route.patch.go":  true,
		"route.delete.go": true,
	}

	POSSIBLE_ROUTE_HANDLER_NAMES = map[string]bool{
		"GET":    true,
		"PUT":    true,
		"POST":   true,
		"HEAD":   true,
		"PATCH":  true,
		"DELETE": true,
	}
)

var (
	ErrNoExportedRouteHandlerFound  = errors.New("no exported route handler found")
	ErrSeparatorMissing             = errors.New("last or first separator occurence missing")
	ErrNonExistentMainDir           = fmt.Errorf("main directory %q does not exist", MAIN_DIR_NAME)
	ErrInvalidMainDirPath           = errors.New("the main directory path provided is invalid")
	ErrInvalidPkgOrRouteHandlerType = errors.New("the package name or route handler type is invalid")
)
