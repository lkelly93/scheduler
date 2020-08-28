//Package program represents a program written in a generic language.
//This package can run the given program and return the result
package executable

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/lkelly93/scheduler/internal/handler"
)

//Executable represents program that is ready to execute
type Executable interface {
	Run() string
}

//Program represents a Program that needs to be run
type program struct {
	code    string
	handler handler.FileHandler
}

//NewExecutable creates a new executable and then return it.
//If the given language is not supported NewProgram will throw an error.
func NewExecutable(lang string, code string) (Executable, error) {
	handler := handler.GetFileHandler(lang, nil)
	if handler != nil {
		prog := program{
			code:    code,
			handler: handler,
		}
		return &prog, nil
	}
	err := fmt.Sprintf("%s is not a supported language", lang)
	return nil, errors.New(err)
}

//Run runs the given program and then returns the output from that given program
//Run returns the result of the run and the err message. If err == nil then the
//run was successful
func (prog *program) Run() string {
	timeoutInSeconds := 15
	//Create the file and get the data to run it
	sysCommand, fileLocation := prog.handler.CreateRunnerFile(prog.code)
	//Remove the old files
	defer os.Remove(fileLocation)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(15)*time.Second)

	defer cancel()

	//Get the system resources to run the command
	command := exec.CommandContext(ctx, sysCommand, fileLocation)

	var stOut bytes.Buffer
	var stErr bytes.Buffer

	command.Stdout = &stOut
	command.Stderr = &stErr

	//Run the command and get the stdOut/stdErr
	err := command.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Sprintf("Time Limit Exceeded %ds", timeoutInSeconds)
	}
	if err != nil {
		return handler.RemoveFilePath(stErr.String(), fileLocation)
	}

	return string(stOut.String())
}
