package runner

import (
	"log"
	"strings"
)

//Java creates a runnerFile for java langues.
func java(code string) (string, string) {
	langCommand := "java"
	outFileName := getRunnerFileLocation("JavaRunner.java")

	var formattedCode strings.Builder
	formattedCode.WriteString("import java.util.*;")
	formattedCode.WriteString("public class JavaRunner{")
	formattedCode.WriteString(code)
	formattedCode.WriteString("}")

	err := createFileAndAddCode(outFileName, formattedCode.String())

	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}
