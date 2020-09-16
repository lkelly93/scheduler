//Package executable represents a program written in a generic language.
//This package can run the given program and return the result
package executable

import (
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
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
	state.settings = fillRestOfFileSettings(state.lang, state.settings)
	//Create the file and get the data to run it. If sys command is an empty
	//string then we had a compilation error and the error is stored in the
	//fileLocation variable.
	sysCommand, fileLocation, err := state.createFile(state.code, state.settings)
	if err != nil {
		log.Println(err)
		return "", &CompilationError{errMessage: err.Error()}
	}
	//Remove the old files
	defer os.Remove(fileLocation)

	timeoutInSeconds := 15
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeoutInSeconds)*time.Second,
	)
	defer cancel()

	//Get the system resources to run the command
	cmd := exec.CommandContext(
		ctx,
		"executable_container",
		sysCommand,
		fileLocation,
	)

	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	//Run the command and get the stdOut/stdErr
	// log.Panicln(cmd.Path)
	err = cmd.Run()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return stdOut.String(), &TimeLimitExceededError{maxTime: timeoutInSeconds}
		}
		log.Println(err)
		return "", &RuntimeError{errMessage: err.Error()}
	}

	if stdErr.Len() != 0 {
		errorMessage := strings.ReplaceAll(stdErr.String(), state.settings.FileNamePrefix, "")

		sizeOfDateTiemStamp := 20
		errorMessage = errorMessage[sizeOfDateTiemStamp:]

		return stdOut.String(), &RuntimeError{errMessage: errorMessage}
	}
	return string(stdOut.String()), nil
}
