package handler

import (
	"log"
	"strings"
)

//Java creates a runnerFile for java languages.
func createRunnerFileJava(code string, settings *HandlerSettings) (string, string) {
	langCommand := "java"
	outFileName := getRunnerFileLocation(settings.className + ".java")

	var formattedCode strings.Builder
	header(&formattedCode, settings)
	formattedCode.WriteString(code)
	footer(&formattedCode, settings)

	err := createFileAndAddCode(outFileName, formattedCode.String())

	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}

func header(formattedCode *strings.Builder, settings *HandlerSettings){
	formattedCode.WriteString(settings.leadingCode)
	formattedCode.WriteString("\n")
	formattedCode.WriteString("public class ")
	formattedCode.WriteString(settings.className)
	formattedCode.WriteString("{\n")
}

func footer(formattedCode *strings.Builder, settings *HandlerSettings){
	formattedCode.WriteString("\n")
	formattedCode.WriteString(settings.trailingCode)
	formattedCode.WriteString("\n}")
}
