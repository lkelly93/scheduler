package runner

import (
	"log"
	"strings"
)

//javaSettings represents the optional settings for creating a Java runnerFile.
type javaSettings struct {
	leadingCode, className, trailingCode string	
}

//Replaces empty settings with the default values.
func (settings *javaSettings) replaceEmptyWithDefaults(defaults javaSettings) {
	if settings.leadingCode == "" {
		settings.leadingCode = defaults.leadingCode
	}
	if settings.className == "" {
		settings.className = defaults.className
	}
	if settings.trailingCode == "" {
		settings.trailingCode = defaults.trailingCode	
	}
}

//Creates a runner string for Java langues.
func createRunnerStringJava(code string, settings javaSettings) string {
	settings.replaceEmptyWithDefaults(javaSettings{"import java.util.*;","JavaRunner",""})
	
	var formattedCode strings.Builder
	
	formattedCode.WriteString(settings.leadingCode+"\n")
	formattedCode.WriteString("public class "+settings.className+"{\n")
	
	formattedCode.WriteString(code)
	
	formattedCode.WriteString("\n}")
	if settings.trailingCode != "" {
		formattedCode.WriteString("\n"+settings.trailingCode)
	}
	
	return formattedCode.String()
}

//Creates a runnerFile for java langues with optional settings.
func createRunnerFileJavaWithSettings(code string, settings javaSettings) (string, string) {
	langCommand := "java"
	outFileName := getRunnerFileLocation("JavaRunner.java")

	runnerString := createRunnerStringJava(code, settings)
	err := createFileAndAddCode(outFileName, runnerString)

	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}

//Creates a runnerFile for java langues.
func createRunnerFileJava(code string) (string, string) {
	return createRunnerFileJavaWithSettings(code, javaSettings{})
}
