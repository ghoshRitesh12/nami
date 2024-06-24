package lib

import (
	"errors"
	"fmt"
)

const (
	MAIN_DIR_NAME         string = "routes"
	DEFAULT_MAIN_DIR_PATH string = "./routes"
)

const MAIN_OUTPUT_FILE string = "main.gen.go"

const (
	MAIN_TMPL_NAME           string = "main"
	HTTP_SERVE_MUX_TMPL_NAME string = "main"
)

const TEMPLATE_GLOB_PATTERN string = "./templates/*.tmpl"

const (
	FlagRouteHandlerType        string = "rht"
	FlagRouteHandlerTypeImport  string = "rhti"
	FlagRouterStructPointerType string = "rspt"
	FlagMainDirPath             string = "mdp"
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
	ErrNoExportedRouteHandlerFound     = errors.New("no exported route handler found")
	ErrSeparatorMissing                = errors.New("last or first separator occurence missing")
	ErrNonExistentMainDir              = fmt.Errorf("main directory %q does not exist", MAIN_DIR_NAME)
	ErrInvalidMainDirPath              = errors.New("the main directory path provided is invalid")
	ErrInvalidRouterType               = errors.New("the router struct pointer type provided is invalid")
	ErrInvalidRouteHandlerTypeOrImport = errors.New("the package name or route handler type is invalid")
	ErrParsingCLIArgs                  = errors.New("error while parsing CLI arguments/flags")
)
