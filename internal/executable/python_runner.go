package executable

import (
	"log"
	"strings"
)

//Python creates a runnerFile for python languages.
func createRunnerFilePython(code string, settings *FileSettings) (string, string) {
	settings = fillRestOfFileSettings("python", settings)
	langCommand := "python3"
	outFileName := getRunnerFileLocation(settings.ClassName + ".py")

	var formattedCode strings.Builder
	insertImportsPython(&formattedCode, settings)
	formattedCode.WriteString(code)
	insertTrailingCodePython(&formattedCode, settings)

	err := createFileAndAddCode(outFileName, formattedCode.String())

	if err != nil {
		log.Fatal("Could not create runner file!")
	}
	return langCommand, outFileName
}

func insertImportsPython(formattedCode *strings.Builder, settings *FileSettings) {
	formattedCode.WriteString(settings.Imports)
	formattedCode.WriteString("\n")
}

func insertTrailingCodePython(formattedCode *strings.Builder, settings *FileSettings) {
	formattedCode.WriteString("\n")
	formattedCode.WriteString(settings.TrailingCode)
}
