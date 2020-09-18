package main

import (
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type configSettings struct {
	hostname string
	rootLoc  string
}

func (cs *configSettings) setupInternalContainer() {
	setupAllCGroups()

	mountProc(cs.rootLoc)
	mountSys(cs.rootLoc)

	//Change hostname and Chroot.
	changeHostName(cs.hostname)
	changeRoot(cs.rootLoc)

}

func changeHostName(name string) {
	err := syscall.Sethostname([]byte(name))
	if err != nil {
		log.Fatalf("Error setting hostname - Error Type:%T", err)
	}
}

func changeRoot(newRoot string) {
	err := syscall.Chroot(newRoot)
	if err != nil {
		log.Fatalf("Error changing root to /securefs - Error Type:%T", err)
	}
	err = os.Chdir("/")
	if err != nil {
		log.Fatalf("Error changing dir to \"/\" - Error Type:%T", err)
	}
}

func mountProc(newRoot string) {
	source := "proc"
	fstype := "proc"
	target := filepath.Join(newRoot, "/proc")
	flags := uintptr(0)
	data := ""

	checkMkdirErrors(
		os.MkdirAll(target, 0755),
		"/securesfs/proc",
	)

	err := syscall.Mount(source, target, fstype, flags, data)
	if err != nil {
		log.Fatalf("Error mounting /securefs/proc - Error Type:%T", err)
	}
}

func mountSys(rootLocation string) {
	source := "sysfs"
	target := filepath.Join(rootLocation, "/sys")
	fstype := "sysfs"
	flags := uintptr(0)
	data := ""

	checkMkdirErrors(
		os.MkdirAll(target, 0755),
		"/securesfs/sys",
	)

	err := syscall.Mount(source, target, fstype, flags, data)
	if err != nil {
		log.Fatalf("Error mounting /securefs/sys - Error Type:%T", err)
	}
}
