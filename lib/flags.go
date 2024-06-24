package lib

import (
	"flag"
	"log"
	"net/http"
)

type RouteGeneratorFlags struct {
	RouteHandlerType        string // rht
	RouteHandlerTypeImport  string // rhti
	RouterStructPointerType string // rspt
	MainDirPath             string // mdp
}

func GetParsedFlags() RouteGeneratorFlags {
	flags := RouteGeneratorFlags{}

	flag.StringVar(
		&flags.RouteHandlerType,
		FlagRouteHandlerType,
		"http.HandlerFunc",
		"stands for Route Handler Type"+"\n"+
			"states the type of the route handler function",
	)

	flag.StringVar(
		&flags.RouteHandlerTypeImport,
		FlagRouteHandlerTypeImport,
		"net/http",
		"stands for Route Handler Type Import"+"\n"+
			"states the import path of the route handler function",
	)

	flag.StringVar(
		&flags.RouterStructPointerType,
		FlagRouterStructPointerType,
		// "http",
		"*http.ServeMux",
		"stands for Router Struct Pointer Type"+"\n"+
			"states the type of the router struct pointer",
	)

	flag.StringVar(
		&flags.MainDirPath,
		FlagMainDirPath,
		DEFAULT_MAIN_DIR_PATH,
		"stands for Main Directory Path"+"\n"+
			"states the main directory path of routes directory",
	)

	flag.Parse()

	if !flag.Parsed() {
		log.Fatal(ErrParsingCLIArgs)
	}

	return flags
}

func RegisterRoutes(router *http.ServeMux) {

}
