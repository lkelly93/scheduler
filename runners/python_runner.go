package runners

import "log"

//Python creates a runnerFile for python langues.
func Python(code string) (string, string) {
	langCommand := "python3"
	outFileName := getRunnerFileLocation("PythonRunner.py")

	err := createFileAndAddCode(outFileName, code)
	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}
