package program

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/lkelly93/code-runner/runners"
	runner_files "github.com/lkelly93/code-runner/runners"
)

//Program represents a program that needs to be run
type Program struct {
	Lang       string
	Code       string
	FileRunner string
	OutputFile string
}

//NewProgram creates a new program struct and returns it.
//If the given language is not supported NewProgram will throw an error.
func NewProgram(lang string, code string) (*Program, error) {
	if runners.IsSupportedLanguage(lang) {
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
	runnerFileFunctor := runner_files.GetFunctor(prog.Lang)
	sysCommand, fileLocation := runnerFileFunctor(prog.Code)
	command := exec.Command(sysCommand, fileLocation)
	sysOut, err := command.CombinedOutput()
	os.Remove(fileLocation)
	return string(sysOut), err
}
