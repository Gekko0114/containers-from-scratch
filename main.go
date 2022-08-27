package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// docker         run image <cmd> <params>
// go run main.go run       <cmd> <params>

func main() {
	if len(os.Args) < 2 {
		panic("You must have at least two command line arguments")
	}
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	/*	cmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		}
	*/

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func child() {
	fmt.Printf("Running child %v as %d\n", os.Args[2:], os.Getpid())

	cg()

	err := syscall.Sethostname([]byte("container"))
	if err != nil {
		fmt.Println(err)
	}
	err = syscall.Chdir("/")
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	must(ioutil.WriteFile(filepath.Join(cgroups, "pids.max"), []byte("20"), 0700))
	must(ioutil.WriteFile(filepath.Join(cgroups, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
