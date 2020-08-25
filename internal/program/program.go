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

//Executable represents program that is ready to execute
type Executable interface {
	Run() (string, error)
}

//Program represents a Program that needs to be run
type program struct {
	Code    string
	functor func(string) (string, string)
}

//NewExecutable creates a new executable and then return it.
//If the given language is not supported NewProgram will throw an error.
func NewExecutable(lang string, code string) (Executable, error) {
	functor := runner.GetFunctor(lang)
	if functor != nil {
		prog := program{
			Code:    code,
			functor: functor,
		}
		return &prog, nil
	}
	err := fmt.Sprintf("%s is not a supported language", lang)
	return nil, errors.New(err)
}

//Run runs the given program and then returns the output from that given program
//Run returns the result of the run and the err message. If err == nil then the
//run was successful
func (prog *program) Run() (string, error) {
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
