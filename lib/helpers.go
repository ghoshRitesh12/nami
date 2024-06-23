package lib

import (
	"os/exec"
	"path/filepath"
	"strings"
)

func getGoModuleName() (string, error) {
	cmd := exec.Command("go", "list", "-m")

	outputInBytes, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(outputInBytes), nil
}

func getPathParams(path string) (PathParams, error) {
	separator := string(filepath.Separator)
	firstOccurence := strings.Index(path, separator)
	lastOccurence := strings.LastIndex(path, separator)

	if firstOccurence == -1 || lastOccurence == -1 {
		return "", ErrSeparatorMissing
	}

	return PathParams(path[firstOccurence:lastOccurence]), nil
}

func getHTTPVerb(filename string) string {
	const DELIMETER string = "."
	splitText := strings.Split(filename, DELIMETER)

	return splitText[1]
}

func getPackagePath(moduleName, path string) string {
	return strings.ReplaceAll(
		filepath.Join(moduleName, MAIN_DIR_NAME, filepath.Clean(path)),
		"\n",
		"",
	)
}

func getPackageName(endpoint string) string {
	lastOccurence := strings.LastIndex(endpoint, string(filepath.Separator))
	return endpoint[lastOccurence+1:]
}

func isValidFile(filename string) bool {
	_, ok := POSSIBLE_ROUTE_FILENAME[filename]
	return ok
}

func isValidRouteHandlerFileName(functionName string) bool {
	_, ok := POSSIBLE_ROUTE_HANDLER_NAMES[functionName]
	return ok
}

func autoFormatFile(filepath string) error {
	fmtCmd := exec.Command("go", "fmt", filepath)
	if err := fmtCmd.Run(); err != nil {
		return err
	}
	return nil
}