//Pacakge program represents a program written in a generic language.
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
	Code    string
	functor func(string) (string, string)
}

//NewProgram creates a new program struct and returns it.
//If the given language is not supported NewProgram will throw an error.
func NewProgram(lang string, code string) (*Program, error) {
	if runner.IsSupportedLanguage(lang) {
		prog := Program{
			Code:    code,
			functor: runner.GetFunctor(lang),
		}
		return &prog, nil
	}
	err := fmt.Sprintf("%s is not a supported language", lang)
	return nil, errors.New(err)
}

//Run runs the given program and then returns the output from that given program
//Run returns the result of the run and the err message. If err == nil then the
//run was successful
func (prog *Program) Run() (string, error) {
	//Create the file and get the data to run it
	sysCommand, fileLocation := prog.functor(prog.Code)
	//Remove the old files
	defer os.Remove(fileLocation)

	//Get the system resources to run the command
	command := exec.Command(sysCommand, fileLocation)

	//Run the command and get the stdOut/stdErr
	sysOut, err := command.CombinedOutput()

	//Return everything
	return string(sysOut), err
}
