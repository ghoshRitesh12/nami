package lib

import (
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

type FBR struct {
	Packages packages
	RouteMap routeMap
}

func NewFileBasedRouter() *FBR {
	return &FBR{
		Packages: make(packages),
		RouteMap: make(routeMap),
	}
}

func (fbr *FBR) TraverseAndSet() error {
	moduleName, err := getGoModuleName()
	if err != nil {
		return err
	}

	walkErr := filepath.WalkDir("./"+DIR_NAME, func(path string, file fs.DirEntry, e error) error {
		if file.IsDir() {
			return nil
		}

		filename := file.Name()
		if !isValidFile(filename) {
			return nil
		}

		endpoint, err := getEndpoint(path)
		if err != nil {
			return nil
		}

		routeHandler, err := getRouteHandler(path)
		if err != nil {
			return nil
		}

		fbr.RouteMap.add(
			endpoint,
			getHTTPVerb(filename),
			getPackageName(endpoint),
			routeHandler,
		)
		fbr.Packages.add(getPackagePath(moduleName, endpoint))

		return nil
	})

	if walkErr != nil {
		return err
	}

	return nil
}

func (fbr *FBR) WriteOut() error {
	const TMPL_FILENAME = "routes.tmpl"
	const OUTPUT_FILENAME = "routes/main.gen.go"

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filepath.Join(cwd, OUTPUT_FILENAME))
	if err != nil {
		return err
	}

	defer outputFile.Close()

	tmpl, err := template.New(TMPL_FILENAME).ParseFiles(TMPL_FILENAME)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(outputFile, fbr); err != nil {
		return err
	}

	return nil
}
