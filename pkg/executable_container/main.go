package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

type configSettings struct {
	hostname string
	rootLoc  string
}

func main() {
	initContainerAndRunProgram()
}

func initContainerAndRunProgram() {
	containerSettings := configSettings{
		hostname: "runner",
		rootLoc:  "/securefs",
	}

	containerSettings.setupInternalContainer()

	sysCommand := os.Args[1]
	fileLocation := os.Args[2]
	fileLocation = strings.ReplaceAll(fileLocation,
		containerSettings.rootLoc,
		"")
	runProgramInContainer(sysCommand, fileLocation)
}

func runProgramInContainer(sysCommand string, fileLocation string) {
	cmd := exec.Command(sysCommand, fileLocation)

	var stdErr bytes.Buffer
	var stdOut bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()

	if err != nil {
		log.Print(removeFilePath(stdErr.String(), fileLocation))
	}

	fmt.Print(removeFilePath(stdOut.String(), fileLocation))
}

func (cs *configSettings) setupInternalContainer() {
	mountProc(cs.rootLoc)
	changeHostName(cs.hostname)
	changeRoot(cs.rootLoc)
}

func changeHostName(name string) {
	must(syscall.Sethostname([]byte(name)))
}

func changeRoot(newRoot string) {
	must(syscall.Chroot(newRoot))
	must(os.Chdir("/"))
}

func mountProc(rootLocation string) {
	source := "proc"
	target := filepath.Join(rootLocation, "/proc")
	fstype := "proc"
	flags := uintptr(0)
	data := ""

	must(os.MkdirAll(target, 0755))
	must(syscall.Mount(source, target, fstype, flags, data))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//removeFilePath removes the file path from the error text of an executable.
func removeFilePath(stdErr string, fileLocation string) string {
	indexSlash := strings.LastIndex(fileLocation, "/") + 1
	stdErr = strings.ReplaceAll(stdErr, fileLocation, fileLocation[indexSlash:])
	return stdErr
}
