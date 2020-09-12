//Package executable represents a program written in a generic language.
//This package can run the given program and return the result
package executable

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/docker/docker/pkg/reexec"
)

//NewExecutable creates a new executable with the given settings and code.
//If the given language is not supported NewProgram will throw an error.
//If FileSettings is nil the default settings will be used for that language.
func NewExecutable(lang string, code string, settings *FileSettings) (Executable, error) {
	function := getFileCreationFunction(lang)
	if function != nil {
		return &executableState{
			code:       code,
			lang:       lang,
			settings:   settings,
			createFile: function,
		}, nil
	}
	return nil, &UnsupportedLanguageError{
		lang: lang,
	}
}

//Run runs the given program and then returns the output, this could be the
//output from a successful run or the error message from an unsuccessful run.
//Run attempts to put all runner files in a folder called runner_files, which
//should be located in the same directory as the file that class run.
//If this is not the case, Run will just put it in the same directory
func (state *executableState) Run() (string, error) {
	timeoutInSeconds := 15
	state.settings = fillRestOfFileSettings(state.lang, state.settings)
	//Create the file and get the data to run it. If sys command is an empty
	//string then we had a compilation error and the error is stored in the
	//fileLocation variable.
	sysCommand, fileLocation, err := state.createFile(state.code, state.settings)
	if err != nil {
		return "", err
	}
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
	err = command.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return "", &TimeLimitExceededError{maxTime: timeoutInSeconds}
	}
	if err != nil {
		errorMessage := removeFilePath(stErr.String(), fileLocation)
		//Remove FileNamePrefix as well.
		errorMessage = strings.ReplaceAll(errorMessage, state.settings.FileNamePrefix, "")
		return "", &RuntimeError{errMessage: errorMessage}
	}

	return string(stOut.String()), nil
}

func initReexec() {
	reexec.Register("initContainer", initContainer)
	if reexec.Init() {
		os.Exit(0)
	}
}

func initContainer() {

}

func runProgramInContainer(sysCommand string, fileLocation string) {
	cmd := exec.Command(sysCommand, fileLocation)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Print(err)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
