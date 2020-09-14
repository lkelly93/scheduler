package executable

import (
	"log"
	"strings"
)

//Python creates a runnerFile for python languages.
func createRunnerFilePython(code string, settings *FileSettings) (string, string, error) {
	langCommand := "python3"
	var runnerFileName strings.Builder
	runnerFileName.WriteString(settings.FileNamePrefix)
	runnerFileName.WriteString(settings.ClassName)
	runnerFileName.WriteString(".py")
	outFileName := getRunnerFileLocation(runnerFileName.String())

	var formattedCode strings.Builder
	insertImportsPython(&formattedCode, settings)
	formattedCode.WriteString(code)
	insertTrailingCodePython(&formattedCode, settings)

	err := createFileAndAddCode(outFileName, formattedCode.String())

	if err != nil {
		log.Println("Could not create runner file!")
		log.Println(err.Error())
	}
	return langCommand, outFileName, nil
}

func insertImportsPython(formattedCode *strings.Builder, settings *FileSettings) {
	formattedCode.WriteString(settings.Imports)
	formattedCode.WriteString("\n")
}

func insertTrailingCodePython(formattedCode *strings.Builder, settings *FileSettings) {
	formattedCode.WriteString("\n")
	formattedCode.WriteString(settings.TrailingCode)
}
