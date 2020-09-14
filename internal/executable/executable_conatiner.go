package executable

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

type configSettings struct {
	hostname string
	rootLoc  string
}

func (cs *configSettings) setupInternalContainer() {
	changeHostName(cs.hostname)
	changeRoot(cs.rootLoc)
}

func (cs *configSettings) tearDownInteralContainer() {
	fmt.Print("\n\n>>>Tearown happens here<<<\n\n")
	return
}

func changeHostName(name string) {
	must(syscall.Sethostname([]byte(name)))
}

func changeRoot(newRoot string) {
	must(syscall.Chroot(newRoot))
	must(os.Chdir("/"))
}

func mountProc() {
	return
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
