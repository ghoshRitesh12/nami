package nami_test

import (
	"testing"

	"github.com/ghoshRitesh12/nami"
)

func TestGenerateRoutes(t *testing.T) {
	// fmt.Println(filepath.Base("./something/routes/adad.go"))

	nami := nami.NewRouteGenerator()
	nami.AddMainDirPath("./routes").
		AddRouteHandlerType("net/http", "http.HandlerFunc").
		GenerateRoutes()
}
