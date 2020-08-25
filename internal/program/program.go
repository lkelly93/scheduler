//A package to represent a program written in a generic language.
//This package can run the given program and return the result
package program

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/lkelly93/scheduler/internal/runner"
)

//Program represents a Program that needs to be run
type Program struct {
	Lang string
	Code string
}

//NewProgram creates a new program struct and returns it.
//If the given language is not supported NewProgram will throw an error.
func NewProgram(lang string, code string) (*Program, error) {
	if runner.IsSupportedLanguage(lang) {
		prog := Program{
			Lang: lang,
			Code: code,
		}
		return &prog, nil
	}
	err := fmt.Sprintf("%s is not a supported language", lang)
	return nil, errors.New(err)
}

//Run runs the given program and then returns the output from that given program
//Run returns the result of the run and the err message. If err == nil then the
//run was successful
func Run(prog *Program) (string, error) {
	//Get the function to create the runner file
	createRunnerFunctor := runner.GetFunctor(prog.Lang)

	//Create the file and get the data to run it
	sysCommand, fileLocation := createRunnerFunctor(prog.Code)

	//Get the system resources to run the command
	command := exec.Command(sysCommand, fileLocation)

	//Run the command and get the stdOut/stdErr
	sysOut, err := command.CombinedOutput()

	//Remove the old files
	os.Remove(fileLocation)

	//Return everything
	return string(sysOut), err
}
