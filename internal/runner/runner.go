//Package runner supports the creation of runner files for languages.
//A runner file is a file properly formatted in a given language that can be
//executed by the program package.
package runner

import (
	"os"
	"strings"
)

//FileCreationFunctor represents a function that will create a runner file for
//an executable.
type FileCreationFunctor func(string) (string, string)

var supportedLanguages = map[string]FileCreationFunctor{
	"python": createRunnerFilePython,
	"java":   createRunnerFileJava,
}

//IsSupportedLanguage checks if the given language is supported
func IsSupportedLanguage(lang string) bool {
	_, found := supportedLanguages[lang]
	return found
}

//GetNeededFunctions returns a NeededFunctions struct with all the
//functions needed by an executable. It retuns nil if the language
//is not supported
func GetCreateFileFunctor(lang string) FileCreationFunctor {
	if IsSupportedLanguage(lang) {
		functions := supportedLanguages[lang]
		return functions
	}
	return nil

}

//RemoveFilePath removes the file path from the error text of an executable.
func RemoveFilePath(stdErr string, fileLocation string) string {
	indexSlash := strings.LastIndex(fileLocation, "/") + 1
	stdErr = strings.ReplaceAll(stdErr, fileLocation, fileLocation[indexSlash:])
	return stdErr
}

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
