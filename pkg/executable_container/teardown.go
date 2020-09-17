package main

import "os"

func (cs *configSettings) tearDownContainer() {
	os.Remove(cs.slashProcLoc)
}
