package program

import (
	"log"
	"os"
	"strings"
)

var supportedLanguages = map[string]func(*Program) (string, string){
	"python": Python,
	"java":   Java,
}

//Python creates a runnerFile for python langues.
func Python(prog *Program) (string, string) {
	langCommand := "python3"
	outFileName := getRunnerFileLocation("PythonRunner.py")
	code := prog.Code

	err := createFileAndAddCode(outFileName, code)
	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}

//Java creates a runnerFile for java langues.
func Java(prog *Program) (string, string) {
	langCommand := "java"
	outFileName := getRunnerFileLocation("JavaRunner.java")

	var code strings.Builder
	code.WriteString("import java.util.*;")
	code.WriteString("public class JavaRunner{")
	code.WriteString(prog.Code)
	code.WriteString("}")

	err := createFileAndAddCode(outFileName, code.String())
	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}

//IsSupportedLanguage checks if the given language is supported
func IsSupportedLanguage(lang string) bool {
	_, found := supportedLanguages[lang]
	return found
}

//GetFunctor returns the create file function for the given language
func GetFunctor(lang string) func(*Program) (string, string) {
	if IsSupportedLanguage(lang) {
		return supportedLanguages[lang]
	}
	return nil
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
	location.WriteString("../runnerFiles/")
	location.WriteString(suffix)
	return location.String()
}
