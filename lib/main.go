package lib

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type (
	config struct {
		RouteHandlerType      string
		RouteHandlerPkgImport string
		RouterObjType         string

		mainDirPath string
	}

	RouteGenerator struct {
		Packages packages
		RouteMap routeMap
		Config   config
	}
)

var ROUTE_GENERATOR = RouteGenerator{
	Config: config{
		RouteHandlerType: "any",
		mainDirPath:      "./" + MAIN_DIR_NAME,
	},
	Packages: make(packages),
	RouteMap: make(routeMap),
}

func GetRouteGenerator() *RouteGenerator {
	return &ROUTE_GENERATOR
}

func (rg *RouteGenerator) GenerateRoutes() {
	if err := rg.traverseAndParse(); err != nil {
		log.Fatalf(
			"error occured while traversing and parsing %s directory\n error -> %v",
			MAIN_DIR_NAME,
			err,
		)
	}

	if err := rg.generateRoutesFile(); err != nil {
		log.Fatalf(
			"error occured while generating routes for %s directory\n error -> %v",
			MAIN_DIR_NAME,
			err,
		)
	}
}

func (rg *RouteGenerator) traverseAndParse() error {
	if err := rg.findMainDir(); err != nil {
		return ErrNonExistentMainDir
	}

	moduleName, err := getGoModuleName()
	if err != nil {
		return err
	}

	walkErr := filepath.WalkDir(rg.Config.mainDirPath, func(path string, file fs.DirEntry, e error) error {
		if file.IsDir() {
			return nil
		}

		filename := file.Name()
		if !isValidFile(filename) {
			return nil
		}

		pathParams, err := getPathParams(path)
		if err != nil || pathParams.String() == "" {
			return nil
		}

		routeHandler, err := getRouteHandler(path)
		if err != nil {
			return nil
		}

		fmt.Println("PATH ->", path)
		fmt.Println("PATH PARAMS ->", pathParams)
		fmt.Println("handler ->", routeHandler)
		fmt.Println("")

		rg.RouteMap.add(
			pathParams.String(),
			getHTTPVerb(filename),
			getPackageName(pathParams.String()),
			routeHandler,
		)

		rg.Packages.add(getPackagePath(moduleName, pathParams.String()))

		return nil
	})

	if walkErr != nil {
		return err
	}

	return nil
}

func (rg *RouteGenerator) generateRoutesFile() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputFilePath := filepath.Join(cwd, rg.Config.mainDirPath, ROUTES_OUTPUT_FILE)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	tmpl, err := template.New(ROUTES_TMPL_FILE).ParseFiles(ROUTES_TMPL_FILE)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(outputFile, rg); err != nil {
		return err
	}

	return autoFormatFile(outputFilePath)
}

// finds the "routes" dir, if not present returns an error
func (rg *RouteGenerator) findMainDir() error {
	_, err := os.Open(rg.Config.mainDirPath)
	if os.IsNotExist(err) {
		return ErrNonExistentMainDir
	}

	return nil
}

func (rg *RouteGenerator) AddRouteHandlerInfo(routeHandlerTypeImport, routeHandlerType string) *RouteGenerator {
	const dot = "."
	firstSeparatorOccurence := strings.Index(routeHandlerType, dot)

	if filepath.Base(routeHandlerTypeImport) != routeHandlerType[:firstSeparatorOccurence] {
		log.Fatalln(ErrInvalidRouteHandlerTypeOrImport)
	}

	rg.Packages.add(routeHandlerTypeImport)
	rg.Config.RouteHandlerType = routeHandlerType
	return rg
}

func (rg *RouteGenerator) AddMainDirPath(dirPath string) *RouteGenerator {
	if filepath.Base(dirPath) != MAIN_DIR_NAME {
		log.Fatalln(ErrInvalidMainDirPath)
	}

	rg.Config.mainDirPath = dirPath
	return rg
}
