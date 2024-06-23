package internal

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type config struct {
	RouteHandlerType string
	mainDirPath      string
}

type TemplateGenerator struct {
	Packages packages
	RouteMap routeMap
	Config   config
}

var TEMPLATE_GENERATOR = TemplateGenerator{
	Config: config{
		RouteHandlerType: "any",
	},
	Packages: make(packages),
	RouteMap: make(routeMap),
}

func GetTemplateGenerator() *TemplateGenerator {
	return &TEMPLATE_GENERATOR
}

func (tg *TemplateGenerator) GenerateRoutes() {
	if err := tg.traverseAndParse(); err != nil {
		log.Fatalf(
			"error occured while traversing and parsing %s directory\n error -> %v",
			MAIN_DIR_NAME,
			err,
		)
	}

	if err := tg.generateRoutesFile(); err != nil {
		log.Fatalf(
			"error occured while generating routes for %s directory\n error -> %v",
			MAIN_DIR_NAME,
			err,
		)
	}
}

func (tg *TemplateGenerator) traverseAndParse() error {
	if err := tg.findMainDir(); err != nil {
		return ErrNonExistentMainDir
	}

	moduleName, err := getGoModuleName()
	if err != nil {
		return err
	}

	walkErr := filepath.WalkDir(tg.Config.mainDirPath, func(path string, file fs.DirEntry, e error) error {
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

		tg.RouteMap.add(
			pathParams.String(),
			getHTTPVerb(filename),
			getPackageName(pathParams.String()),
			routeHandler,
		)

		tg.Packages.add(getPackagePath(moduleName, pathParams.String()))

		return nil
	})

	if walkErr != nil {
		return err
	}

	return nil
}

func (tg *TemplateGenerator) generateRoutesFile() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputFilePath := filepath.Join(cwd, tg.Config.mainDirPath, OUTPUT_FILENAME)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	tmpl, err := template.New(ROUTES_TMPL_FILE).ParseFiles(ROUTES_TMPL_FILE)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(outputFile, tg); err != nil {
		return err
	}

	return autoFormatFile(outputFilePath)
}

// finds the "routes" dir, if not present returns an error
func (tg *TemplateGenerator) findMainDir() error {
	_, err := os.Open(tg.Config.mainDirPath)
	if os.IsNotExist(err) {
		return ErrNonExistentMainDir
	}

	return nil
}

func (tg *TemplateGenerator) AddRouteHandlerType(packageImport, routeHandlerType string) *TemplateGenerator {
	const dot = "."
	firstSeparatorOccurence := strings.Index(routeHandlerType, dot)

	if filepath.Base(packageImport) != routeHandlerType[:firstSeparatorOccurence] {
		log.Fatalln(ErrInvalidPkgOrRouteHandlerType)
	}

	tg.Packages.add(packageImport)
	tg.Config.RouteHandlerType = routeHandlerType
	return tg
}

func (tg *TemplateGenerator) AddMainDirPath(dirPath string) *TemplateGenerator {
	if filepath.Base(dirPath) != MAIN_DIR_NAME {
		log.Fatalln(ErrInvalidMainDirPath)
	}

	tg.Config.mainDirPath = dirPath
	return tg
}
