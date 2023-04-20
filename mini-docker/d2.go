package main

import (
	"os"
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func run() {
	cmd := exec.Command(os.Args[2])

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("failed to exec the command ", err)
	}
}

func main() {
	arg := os.Args

	switch arg[1] {
	case "run":
		run()
	default:
		panic("the arg is not defined")
	}
	log.Println("################")

}
