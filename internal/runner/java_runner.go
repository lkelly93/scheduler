package runner

import (
	"log"
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
