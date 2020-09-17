package main

import (
	"log"
	"os"
	"syscall"
)

type configSettings struct {
	hostname     string
	rootLoc      string
	slashProcLoc string
}

func (cs *configSettings) setupInternalContainer() {
	mountProc(cs.slashProcLoc)
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

func mountProc(procLocation string) {
	source := "proc"
	fstype := "proc"
	flags := uintptr(0)
	data := ""

	must(os.MkdirAll(procLocation, 0755))
	must(syscall.Mount(source, procLocation, fstype, flags, data))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
