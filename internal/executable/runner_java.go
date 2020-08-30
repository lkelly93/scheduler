package executable

import (
	"log"
	"strings"
)

//Java creates a runnerFile for java languages.
func createRunnerFileJava(code string, settings *FileSettings) (string, string) {
	settings = fillRestOfFileSettings("java", settings)
	langCommand := "java"
	var fileName strings.Builder
	fileName.WriteString(settings.FileNamePrefix)
	fileName.WriteString(settings.ClassName)
	fileName.WriteString(".java")
	outFileName := getRunnerFileLocation(fileName.String())

	var formattedCode strings.Builder
	insertImportsJava(&formattedCode, settings)
	formattedCode.WriteString(code)
	insertTrailingCodeJava(&formattedCode, settings)

	err := createFileAndAddCode(outFileName, formattedCode.String())

	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
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
