package lib

const (
	DIR_NAME    = "routes"
	FILE_PREFIX = "route."
)

var POSSIBLE_ROUTE_FILENAME = map[string]bool{
	"route.get.go":    true,
	"route.put.go":    true,
	"route.post.go":   true,
	"route.head.go":   true,
	"route.patch.go":  true,
	"route.delete.go": true,
}

var POSSIBLE_ROUTE_HANDLER_NAMES = map[string]bool{
	"GET":    true,
	"PUT":    true,
	"POST":   true,
	"HEAD":   true,
	"PATCH":  true,
	"DELETE": true,
}
