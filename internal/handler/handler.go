//Package handler supports the creation of runner files for executables.
//A runner file is a file properly formatted in a given language that can be
//executed by the program package.
package handler

import (
	"os"
	"strings"
)

//FileHandler is a interface that handles the creation of all runner files
type FileHandler interface {
	CreateRunnerFile(string) (string, string)
}

//HandlerSettings are all the settings for a Handler
type HandlerSettings struct {
	Imports      string
	ClassName    string
	TrailingCode string
}

type fileData struct {
	createFileFunctor fileCreationFunction
	settings          HandlerSettings
}

type fileCreationFunction func(string, *HandlerSettings) (string, string)

var supportedLanguages = map[string]fileData{
	"python": fileData{
		createFileFunctor: createRunnerFilePython,
		settings: HandlerSettings{
			Imports:   "import numpy as np",
			ClassName: "PythonRunner",
		},
	},
	"java": fileData{
		createFileFunctor: createRunnerFileJava,
		settings: HandlerSettings{
			Imports:   "import java.util.*;",
			ClassName: "JavaRunner",
		},
	},
}

//GetFileHandler returns a FileHandler for the given language. If settings
//is nil it will use the default settings for that language
func GetFileHandler(lang string, settings *HandlerSettings) FileHandler {
	handler, found := supportedLanguages[lang]
	if found {
		if settings != nil {
			handler.settings = *settings
		}
		return &handler
	}
	return nil
}

//CreateRunnerFile creates a runner file from the given FileHandler,
//it returns the system command required to run the file, and then then the file
//to be run.
func (data *fileData) CreateRunnerFile(code string) (string, string) {
	return data.createFileFunctor(code, &data.settings)
}

//RemoveFilePath removes the file path from the error text of an executable.
func RemoveFilePath(stdErr string, fileLocation string) string {
	indexSlash := strings.LastIndex(fileLocation, "/") + 1
	stdErr = strings.ReplaceAll(stdErr, fileLocation, fileLocation[indexSlash:])
	return stdErr
}

/***** Supporting Methods ******/
func createFileAndAddCode(outFileName string, code string) error {
	runnerFile, err := os.Create(outFileName)
	if err == nil {
		runnerFile.WriteString(code)
		runnerFile.Close()
	}

	return err
}

func getRunnerFileLocation(suffix string) string {
	var location strings.Builder
	location.WriteString("../runner_files/")
	location.WriteString(suffix)
	return location.String()
}
