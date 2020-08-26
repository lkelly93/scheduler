package runner

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

//Java creates a runnerFile for java langues.
func createRunnerFileJava(code string) (string, string) {
	langCommand := "java"
	outFileName := getRunnerFileLocation("JavaRunner.java")

	var formattedCode strings.Builder
	formattedCode.WriteString("import java.util.*;\n")
	formattedCode.WriteString("public class JavaRunner{\n")
	formattedCode.WriteString(code)
	formattedCode.WriteString("}")

	err := createFileAndAddCode(outFileName, formattedCode.String())

	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}

func parseJavaStandardErr(stErr string) string {
	expression, err := regexp.Compile("[0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	errorLineNumber := expression.FindString(stErr)

	// indexFirstNewline := strings.Index(stErr, ":") + 1
	// firstLine := stErr[:indexFirstNewline]
	indexEndOfFirstLine := strings.Index(stErr, "\n") + 1
	stErr = stErr[indexEndOfFirstLine:]
	// errorLineNumber := firstLine[len(firstLine)-1:]

	var output strings.Builder

	output.WriteString(fmt.Sprintf("Error on line number %s\n", errorLineNumber))
	output.WriteString(stErr)

	return output.String()
}
