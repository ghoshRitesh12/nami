package nami

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type templData struct {
	Packages         packages
	RouteMap         routeMap
	RouteHandlerType string
}

type RouteGenerator struct {
	templData   templData
	mainDirPath string
	// routeHandlerPkgName string
}

func NewRouteGenerator() *RouteGenerator {
	return &RouteGenerator{
		mainDirPath: "./" + MAIN_DIR_NAME,
		templData: templData{
			Packages:         make(packages),
			RouteMap:         make(routeMap),
			RouteHandlerType: "any",
		},
	}
}

func (rg *RouteGenerator) GenerateRoutes() {
	if err := rg.traverseAndParse(); err != nil {
		log.Fatalf(
			"error occured while traversing and parsing %s directory\n error -> %v",
			MAIN_DIR_NAME,
			err,
		)
	}

	if err := rg.generate(); err != nil {
		log.Fatalf(
			"error occured while generating routes for %s directory\n error -> %v",
			MAIN_DIR_NAME,
			err,
		)
	}
}

func (rg *RouteGenerator) AddRouteHandlerType(packageName, routeHandlerType string) *RouteGenerator {
	const dot = "."
	firstSeparatorOccurence := strings.Index(routeHandlerType, dot)

	if filepath.Base(packageName) != routeHandlerType[:firstSeparatorOccurence] {
		log.Fatalln(ErrInvalidPkgOrRouteHandlerType)
	}

	rg.templData.Packages.add(packageName)
	rg.templData.RouteHandlerType = routeHandlerType
	return rg
}

// func (rg *RouteGenerator) AddRouteHandlerImport(packageName string) *RouteGenerator {
// 	const dot = "."
// 	firstSeparatorOccurence := strings.Index(routeHandlerType, dot)

// 	if filepath.Base(packageName) != routeHandlerType[:firstSeparatorOccurence] {
// 		log.Fatalln(ErrInvalidPkgOrRouteHandlerType)
// 	}

// 	rg.templData.Packages.add(packageName)
// 	rg.templData.RouteHandlerType = routeHandlerType
// 	return rg
// }

func (rg *RouteGenerator) AddMainDirPath(dirPath string) *RouteGenerator {
	if filepath.Base(dirPath) != MAIN_DIR_NAME {
		log.Fatalln(ErrInvalidMainDirPath)
	}

	rg.mainDirPath = dirPath
	return rg
}

func (rg *RouteGenerator) traverseAndParse() error {
	if err := rg.findMainDir(); err != nil {
		return ErrNonExistentMainDir
	}

	moduleName, err := getGoModuleName()
	if err != nil {
		return err
	}

	walkErr := filepath.WalkDir(rg.mainDirPath, func(path string, file fs.DirEntry, e error) error {
		if file.IsDir() {
			return nil
		}

		filename := file.Name()
		if !isValidFile(filename) {
			return nil
		}

		endpoint, err := getEndpoint(path)
		if err != nil || endpoint.String() == "" {
			return nil
		}

		routeHandler, err := getRouteHandler(path)
		if err != nil {
			return nil
		}

		fmt.Println("PATH ->", path)
		fmt.Println("ENDPOINT ->", endpoint)
		fmt.Println("handler ->", routeHandler)
		fmt.Println("")

		rg.templData.RouteMap.add(
			endpoint.GetPath(),
			getHTTPVerb(filename),
			getPackageName(endpoint.String()),
			routeHandler,
		)

		rg.templData.Packages.add(getPackagePath(moduleName, endpoint.String()))

		return nil
	})

	if walkErr != nil {
		return err
	}

	return nil
}

func (rg *RouteGenerator) generate() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filepath.Join(cwd, rg.mainDirPath, OUTPUT_FILENAME))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	tmpl, err := template.New(TMPL_FILENAME).ParseFiles(TMPL_FILENAME)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(outputFile, rg.templData); err != nil {
		return err
	}

	return nil
}

// finds the "routes" dir, if not present returns an error
func (rg *RouteGenerator) findMainDir() error {
	_, err := os.Open(rg.mainDirPath)
	if os.IsNotExist(err) {
		return ErrNonExistentMainDir
	}

	return nil
}
