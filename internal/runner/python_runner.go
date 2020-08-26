package runner

import (
	"log"
)

//Python creates a runnerFile for python langues.
func createRunnerFilePython(code string) (string, string) {
	langCommand := "python3"
	outFileName := getRunnerFileLocation("PythonRunner.py")

	err := createFileAndAddCode(outFileName, code)
	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}
