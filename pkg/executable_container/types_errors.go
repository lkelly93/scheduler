package main

import "strings"

type programError struct {
	fileNamePrefix string
	fileLocation   string
	errMess        string
}

func (pe *programError) Error() string {
	output := removeFilePath(pe.errMess, pe.fileLocation)
	output = strings.ReplaceAll(output, pe.fileNamePrefix, "")
	return output
}

//removeFilePath removes the file path from the error text of an executable.
func removeFilePath(stdErr string, fileLocation string) string {
	indexSlash := strings.LastIndex(fileLocation, "/") + 1
	stdErr = strings.ReplaceAll(stdErr, fileLocation, fileLocation[indexSlash:])
	return stdErr
}
