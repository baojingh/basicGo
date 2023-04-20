package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"
)

const cgroupMemoryHierarchyMout = "/sys/fs/cgroup/memory"

func main() {
	arg := os.Args
	if arg[0] == "/proc/self/exe" {
		log.Println("current pid is ", syscall.Getpid())
		log.Println()
		// cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd := exec.Command("sh", "-c", `df -h`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Println(err)
			os.Exit(1)
		}

	}
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	} else {
		log.Println("%v", cmd.Process.Pid)
		os.Mkdir(path.Join(cgroupMemoryHierarchyMout, "testmemorylimit"), 0755)
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMout, "testmemorylimit", "tasks"),
			[]byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMout, "testmemorylimit",
			"memory.limit in bytes"), []byte("lOOm"), 0644)

	}

	cmd.Process.Wait()

	log.Println("################")

}
