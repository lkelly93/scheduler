package main

import (
	"log"
	"os"
	"strings"
)

func (cs *containerSettings) serverPrintln(message string) {
	filePtr := getOutputFile(cs.fileNamePrefix)

	_, err := filePtr.Write([]byte(message + "\n"))

	if err != nil {
		log.Fatalf("Could not write output to log file. - %s", err.Error())
	}
}

func (cs *containerSettings) serverFatal(message string) {
	cs.serverPrintln(message)
	os.Exit(1)
}

func getOutputFile(fileNamePrefix string) *os.File {
	var builder strings.Builder
	builder.WriteString("/serverOutput/")
	builder.WriteString(fileNamePrefix)
	builder.WriteString(".log")
	filePtr, err := os.OpenFile(
		builder.String(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0744,
	)

	if err != nil {
		log.Fatalf("Couldn't create server logging output file - %s", err.Error())
	}

	return filePtr
}
