package templates

import "embed"

//go:embed *.tmpl
var templateFS embed.FS

func GetTemplateFS() embed.FS {
	return templateFS
}
