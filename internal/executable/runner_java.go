package executable

import (
	"log"
	"strings"
)

//Java creates a runnerFile for java languages.
func createRunnerFileJava(code string, settings *FileSettings) (string, string, error) {
	langCommand := "java"
	var runnerFileName strings.Builder
	runnerFileName.WriteString(settings.FileNamePrefix)
	runnerFileName.WriteString(settings.ClassName)
	runnerFileName.WriteString(".java")
	runnerFileLocation := getRunnerFileLocation(runnerFileName.String())

	var formattedCode strings.Builder
	insertImportsJava(&formattedCode, settings)
	formattedCode.WriteString(code)
	insertTrailingCodeJava(&formattedCode, settings)

	err := createFileAndAddCode(runnerFileLocation, formattedCode.String())
	if err != nil {
		log.Fatal("Could not create runner file!")
	}

	return langCommand,
		runnerFileLocation,
		nil
}

func insertImportsJava(formattedCode *strings.Builder, settings *FileSettings) {
	formattedCode.WriteString(settings.Imports)
	formattedCode.WriteString("\n")
	formattedCode.WriteString("public class ")
	formattedCode.WriteString(settings.ClassName)
	formattedCode.WriteString("{\n")
}

func insertTrailingCodeJava(formattedCode *strings.Builder, settings *FileSettings) {
	formattedCode.WriteString("\n")
	formattedCode.WriteString(settings.TrailingCode)
	formattedCode.WriteString("\n}")
}
