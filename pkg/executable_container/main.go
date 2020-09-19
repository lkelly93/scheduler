package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		log.Print("This file should only be run through the executable package. ")
		log.Fatal("Failure to do so could cause damage to the system it is running on.\n")
	}
	initContainerAndRunProgram()
}

func initContainerAndRunProgram() {
	rootLoc := "/securefs"

	sysCommand := os.Args[1]
	fileLocation := strings.ReplaceAll(
		os.Args[2],
		rootLoc,
		"",
	)
	fileNamePrefix := os.Args[3]

	containerSettings := configSettings{
		hostname: "runner",
		rootLoc:  rootLoc,
	}

	containerSettings.setupInternalContainer()
	runProgramInContainer(sysCommand, fileLocation, fileNamePrefix)
}

func runProgramInContainer(sysCommand string, fileLocation string, fileNamePrefix string) {
	cmd := exec.Command(sysCommand, fileLocation)

	var stdErr bytes.Buffer
	var stdOut bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil || stdErr.Len() != 0 {
		log.Print(&programError{
			fileNamePrefix: fileNamePrefix,
			fileLocation:   fileLocation,
			errMess:        stdErr.String(),
		})
	}

	fmt.Print(parseOutput(stdOut.String(), fileLocation, fileNamePrefix))
}
