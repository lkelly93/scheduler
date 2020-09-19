package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func (cs *containerSettings) setupInternalContainer() {
	// setupAllCGroups()

	cs.mountProc()
	cs.mountSys()

	//Change hostname and Chroot.
	cs.changeHostName()
	cs.changeRoot()

}

func (cs *containerSettings) changeHostName() {
	err := syscall.Sethostname([]byte(cs.hostname))
	if err != nil {
		message := fmt.Sprintf("Error setting hostname - Error Type:%T", err)
		cs.serverFatal(message)
	}
}

func (cs *containerSettings) changeRoot() {
	err := syscall.Chroot(cs.rootLoc)
	if err != nil {
		message := fmt.Sprintf("Error changing root to /securefs - Error Type:%T", err)
		cs.serverFatal(message)
	}
	err = os.Chdir("/")
	if err != nil {
		message := fmt.Sprintf("Error changing dir to \"/\" - Error Type:%T", err)
		cs.serverFatal(message)
	}
}

func (cs *containerSettings) mountProc() {
	source := "proc"
	fstype := "proc"
	target := filepath.Join(cs.rootLoc, "/proc")
	flags := uintptr(0)
	data := ""

	err := syscall.Mount(source, target, fstype, flags, data)
	if err != nil {
		message := fmt.Sprintf("Error mounting /securefs/proc - Error Type:%T", err)
		cs.serverFatal(message)
	}
}

func (cs *containerSettings) mountSys() {
	source := "sysfs"
	target := filepath.Join(cs.rootLoc, "/sys")
	fstype := "sysfs"
	flags := uintptr(0)
	data := ""

	err := syscall.Mount(source, target, fstype, flags, data)
	if err != nil {
		message := fmt.Sprintf("Error mounting /securefs/sys - Error Type:%T", err)
		cs.serverFatal(message)
	}
}
