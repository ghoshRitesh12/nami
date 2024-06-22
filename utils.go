package nami

import (
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

func (rmap *routeMap) add(pathName, httpVerb, pkgName, routeHandlerName string) {
	((*rmap)[pathName]) = append(((*rmap)[pathName]), route{
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

		if isValidRouteHandlerFileName(fn.Name.Name) {
			routeHandlerName = fn.Name.Name
			break
		}
	}

	if routeHandlerName == "" {
		return "", ErrNoExportedRouteHandlerFound
	}

	return routeHandlerName, nil
}

type PathParams string

func (pp *PathParams) GetPath() string {
	switch string(*pp) {
	case "/index":
		return "/"
	default:
		return string(*pp)
	}
}

func (pp *PathParams) String() string {
	return string(*pp)
}

// sd

// func GetDirPath() (string, error) {
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		return "", err
// 	}

// 	return filepath.Join(cwd, MAIN_DIR_NAME), nil
// }
