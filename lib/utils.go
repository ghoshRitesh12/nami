package lib

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type packages map[string]bool

func (pkgs *packages) add(pkgPath string) {
	(*pkgs)[pkgPath] = true
}

type (
	route struct {
		Verb    string
		Handler string // <package.GET>
	}
	routes   []route
	routeMap map[string]routes
)

func (rmap *routeMap) add(endpoint, httpVerb, pkgName, routeHandlerName string) {
	((*rmap)[endpoint]) = append(((*rmap)[endpoint]), route{
		Verb:    httpVerb,
		Handler: pkgName + "." + routeHandlerName,
	})
}

var fset = token.NewFileSet()

func getRouteHandler(path string) (string, error) {
	routeHandlerName := ""

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return "", nil
	}

	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if isValidRouteHandlerName(fn.Name.Name) {
			routeHandlerName = fn.Name.Name
			break
		}

		fmt.Println(fn.Name.Name)
	}

	if routeHandlerName == "" {
		return "", errors.New("no exported route handler found")
	}

	return routeHandlerName, nil
}

// sd

// func GetDirPath() (string, error) {
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		return "", err
// 	}

// 	return filepath.Join(cwd, DIR_NAME), nil
// }
