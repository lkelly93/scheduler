package program

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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
	if IsSupportedLanguage(lang) {
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
	runnerFileFunctor := GetFunctor(prog.Lang)
	sysCommand, fileLocation := runnerFileFunctor(prog)
	command := exec.Command(sysCommand, fileLocation)
	sysOut, err := command.CombinedOutput()
	os.Remove(fileLocation)
	return string(sysOut), err
}
