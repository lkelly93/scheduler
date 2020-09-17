package main

type programError struct {
	fileNamePrefix string
	fileLocation   string
	errMess        string
}

func (pe *programError) Error() string {
	return parseOutput(pe.errMess, pe.fileLocation, pe.fileNamePrefix)
}
