package executable

import (
	"log"
	"syscall"
)

func setupInternalContainer() {
	must(syscall.Sethostname([]byte("runner")))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
