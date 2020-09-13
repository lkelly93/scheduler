package executable

import (
	"log"
	"os"
	"syscall"
)

func setupInternalContainer() {
	changeHostName("runner")
	changeRoot("/securefs")
}

func changeHostName(name string) {
	must(syscall.Sethostname([]byte(name)))
}

func changeRoot(newRoot string) {
	must(syscall.Chroot(newRoot))
	must(os.Chdir("/"))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
