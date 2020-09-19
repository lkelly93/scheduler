package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func setupAllCGroups() {
	pidGroup()
}

func pidGroup() {
	cgroup := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroup, "pids")

	err := ioutil.WriteFile(filepath.Join(pids, "runner/pids.max"), []byte("50"), 0700)
	if err != nil {
		log.Fatalf("Error writing to /sys/fs/cgroup/pids/runner/pids.max - Error Type:%T", err)
	}

	// Removes the new cgroup in place after the container exits
	err = ioutil.WriteFile(filepath.Join(pids, "runner/notify_on_release"), []byte("2"), 0700)
	if err != nil {
		log.Fatalf("Error writing to /sys/fs/cgroup/pids/runner/notify_on_release - Error Type:%T", err)
	}

	err = ioutil.WriteFile(filepath.Join(pids, "runner/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)
	if err != nil {
		log.Fatalf("Error writing to /sys/fs/cgroup/pids/runner/cgroups.procs - Error Type:%T", err)
	}
}
