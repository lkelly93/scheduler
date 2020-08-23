package program

import (
	"os"
)

var supportedLanguages = map[string]func(*Program) (string, string){
	"python": Python,
}

//Python creates a runnerFile for python langues.
func Python(prog *Program) (string, string) {
	langCommand := "python3"
	outFileName := "../runnerFiles/PythonRunner.py"
	code := prog.Code

	runnerFile, err := os.Create(outFileName)
	if err == nil {
		runnerFile.WriteString(code)
		runnerFile.Close()
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
