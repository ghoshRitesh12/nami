package lib

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

// will use this method later if necessary
// func (pp *PathParams) GetPath() string {
// 	if strings.HasPrefix(pp.String(), "/index") {
// 		_, v, _ := strings.Cut(pp.String(), "/index")
// 		fmt.Println(pp.String(), v)
// 		// return pp.String()[firstOccurence:]
// 	}

// 	return pp.String()
// }

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
