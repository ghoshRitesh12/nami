package lib

import (
	"flag"
	"log"
)

type RouteGeneratorFlags struct {
	RouteHandlerType       string // rht
	RouteHandlerTypeImport string // rhti
	RouterType             string // rt
	MainDirPath            string // mdp
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
		&flags.RouterType,
		FlagRouterType,
		"http",
		"stands for Router Type"+"\n"+
			"states the type of the router struct",
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
