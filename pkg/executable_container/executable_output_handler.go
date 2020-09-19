package main

import (
	"strings"
)

func parseOutput(message string, fileLocation string, fileNamePrefix string) string {
	output := removeFilePath(message, fileLocation)
	return removeFileNamePrefix(output, fileNamePrefix)

}

//removeFilePath removes the file path from the error text of an executable.
func removeFilePath(message string, fileLocation string) string {
	indexSlash := strings.LastIndex(fileLocation, "/") + 1
	message = strings.ReplaceAll(message, fileLocation, fileLocation[indexSlash:])
	return message
}

func removeFileNamePrefix(message string, fileNamePrefix string) string {
	return strings.ReplaceAll(message, fileNamePrefix, "")
}
