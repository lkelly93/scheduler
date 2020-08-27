package handler

import (
	"log"
)

//Python creates a runnerFile for python languages.
func createRunnerFilePython(code string, settings *HandlerSettings) (string, string) {
	langCommand := "python3"
	outFileName := getRunnerFileLocation(settings.className + ".py")

	err := createFileAndAddCode(outFileName, code)
	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}
