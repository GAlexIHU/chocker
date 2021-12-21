package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func main() {
	fmt.Println(os.Args[1])
	switch os.Args[1] {
	case "run":
		run()
	case "fork":
		fork()
	default:
		panic("UNKNOWN COMMAND")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"fork"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func fork() {
	fmt.Println(os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	dir, _ := os.Getwd()

	syscall.Chroot(filepath.Join(dir, "ubuntu"))
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
