package main_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ghoshRitesh12/nami/lib"
)

func TestMain(t *testing.T) {
	flags := lib.GetParsedFlags()
	router := http.NewServeMux()
	router.HandleFunc("/asd", func(w http.ResponseWriter, r *http.Request) {})

	// routeGenerator := lib.GetRouteGenerator()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	//
	fmt.Printf("%+v\n", flags)
}
