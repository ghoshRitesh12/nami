package lib

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ghoshRitesh12/nami/templates"
)

type (
	config struct {
		RouteHandlerType        string
		RouteHandlerPkgImport   string
		RouterStructPointerType string

		mainDirPath string
	}

	RouteGenerator struct {
		Packages packages
		RouteMap routeMap
		Config   config
	}
)

func NewRouteGenerator() *RouteGenerator {
	return &RouteGenerator{
		Packages: make(packages),
		RouteMap: make(routeMap),
	}
}

func (rg *RouteGenerator) GenerateRoutes() {
	if err := rg.traverseAndParse(); err != nil {
		log.Fatalf(
			"error occured while traversing and parsing %s directory\n"+"error -> %v",
			MAIN_DIR_NAME,
			err,
		)
	}

	if err := rg.generateRoutesFile(); err != nil {
		log.Fatalf(
			"error occured while generating routes for %s directory\n"+"error -> %v",
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

	fmt.Printf("Parsing file based routes\n\n")

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

		verb := getHTTPVerb(filename)

		rg.RouteMap.add(
			pathParams.String(),
			verb,
			getPackageName(pathParams.String()),
			routeHandler,
		)

		rg.Packages.add(getPackagePath(moduleName, pathParams.String()))

		fmt.Println("├ ƒ", verb, pathParams.String())

		return nil
	})

	if walkErr != nil {
		return err
	}

	fmt.Println()

	return nil
}

func (rg *RouteGenerator) generateRoutesFile() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputFilePath := filepath.Join(cwd, rg.Config.mainDirPath, MAIN_OUTPUT_FILE)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	templateFS := templates.GetTemplateFS()
	tmpl, err := template.ParseFS(templateFS, TEMPLATE_GLOB_PATTERN)
	if err != nil {
		return err
	}
	if err := tmpl.ExecuteTemplate(outputFile, MAIN_TMPL_NAME, rg); err != nil {
		return err
	}

	if fmtErr := autoFormatFile(outputFilePath); fmtErr != nil {
		return fmtErr
	}

	fmt.Printf("Generated route file `%s` in `%s` directory \n\n", MAIN_OUTPUT_FILE, rg.Config.mainDirPath)
	return nil
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
	firstSeparatorOccurence := strings.Index(routeHandlerType, ".")

	if filepath.Base(routeHandlerTypeImport) != routeHandlerType[:firstSeparatorOccurence] {
		log.Fatalln(ErrInvalidRouteHandlerTypeOrImport)
	}

	rg.Packages.add(routeHandlerTypeImport)
	rg.Config.RouteHandlerPkgImport = routeHandlerTypeImport
	rg.Config.RouteHandlerType = routeHandlerType
	return rg
}

func (rg *RouteGenerator) AddRouterStructPointerType(structPointerType string) *RouteGenerator {
	pointerTypePkg := structPointerType

	if structPointerType[:1] != "*" {
		log.Fatalln(ErrInvalidRouterType)
	}

	pointerStrippedType := structPointerType[1:]
	pointerTypePkg = pointerStrippedType[:strings.Index(pointerStrippedType, ".")]

	if filepath.Base(rg.Config.RouteHandlerPkgImport) != pointerTypePkg {
		log.Fatalln(ErrInvalidRouterType)
	}

	rg.Config.RouterStructPointerType = structPointerType
	return rg
}

func (rg *RouteGenerator) AddMainDirPath(dirPath string) *RouteGenerator {
	if filepath.Base(dirPath) != MAIN_DIR_NAME {
		log.Fatalln(ErrInvalidMainDirPath)
	}

	rg.Config.mainDirPath = dirPath
	return rg
}
