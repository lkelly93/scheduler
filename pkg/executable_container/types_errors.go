package main

import (
	"log"
	"os"
)

type programError struct {
	fileNamePrefix string
	fileLocation   string
	errMess        string
}

func (pe *programError) Error() string {
	return parseOutput(pe.errMess, pe.fileLocation, pe.fileNamePrefix)
}

func checkMkdirErrors(err error, dirLocation string) {
	if err != nil && !os.IsExist(err) {
		if os.IsPermission(err) {
			log.Fatalf("Permission denied creating %s", dirLocation)
		}
		log.Fatalf("Error creating %s - Error Type:%T", dirLocation, err)
	}
}
