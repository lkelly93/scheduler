package runner

import (
	"fmt"
	"log"
	"strings"
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

func parsePythonStandardErr(stErr string) string {
	lenthOfWordLine := 4
	//Add lengthOfWordLine becuase it returns the index to the l and we want
	//the line number after that. We add one more becuse this is a space
	//between "line" and the line number
	indexFirstNewline := strings.Index(stErr, "line") + lenthOfWordLine + 1
	indexEndOfFirstLine := strings.Index(stErr, "\n") + 1
	//If we slice between these two indecies we will get our line number
	errorLineNumber := stErr[indexFirstNewline:indexEndOfFirstLine]
	//Remove the first line from stdErr and add the rest into our output
	stErr = stErr[indexEndOfFirstLine:]

	var output strings.Builder

	output.WriteString(fmt.Sprintf("Error on line number %s", errorLineNumber))
	output.WriteString(stErr)

	return output.String()
}
