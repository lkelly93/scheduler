package executable

import (
	"os"
	"strings"
)

//getFileHandler returns a FileHandler for the given language.
func getFileCreationFunction(lang string) fileCreationFunction {
	function, found := supportedLanguages[lang]
	if found {
		return function
	}
	return nil
}

//defaultFileSettings takes a FileSettings that is either empty, nil, or
//partially filled and returns a FileSettings with everything needed to create
//a runner file
func fillRestOfFileSettings(lang string, settings *FileSettings) *FileSettings {
	defaults := fileSettingsDefaults[lang]
	if settings == nil {
		return &defaults
	}

	if settings.Imports == "" {
		settings.Imports = defaults.Imports
	}

	if settings.ClassName == "" {
		settings.ClassName = defaults.ClassName
	}

	if settings.TrailingCode == "" {
		settings.TrailingCode = defaults.TrailingCode
	}

	if settings.FileNamePrefix == "" {
		settings.FileNamePrefix = defaults.FileNamePrefix
	}

	return settings
}

//removeFilePath removes the file path from the error text of an executable.
func removeFilePath(stdErr string, fileLocation string) string {
	indexSlash := strings.LastIndex(fileLocation, "/") + 1
	stdErr = strings.ReplaceAll(stdErr, fileLocation, fileLocation[indexSlash:])
	return stdErr
}

//createFileAndAddCode creates the runner file and adds the code to the file
func createFileAndAddCode(outFileName string, code string) error {
	runnerFile, err := os.Create(outFileName)
	if err == nil {
		runnerFile.WriteString(code)
		runnerFile.Close()
	}
	return err
}

//getRunnerFileLocation returns a the string used to create the file runner in
//the system.
func getRunnerFileLocation(fileName string) string {
	dir := "/securefs/runnerfiles/"
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		var location strings.Builder
		location.WriteString(dir)
		location.WriteString(fileName)
		return location.String()
	}
	return fileName
}
