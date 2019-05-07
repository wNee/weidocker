package main

import (
	"fmt"
	"os"
	"strings"
	"weidocker/cgroups"
	"weidocker/cgroups/systems"
	"weidocker/container"
)

func Run(tty bool, comArray []string, res *systems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		fmt.Println("new parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		fmt.Println(err)
		return
	}

	// user cgroup
	cgroupManger := cgroups.NewCgroupManger("mydocker-cgroup")
	defer cgroupManger.Destory()
	cgroupManger.Set(res)
	cgroupManger.Apply(parent.Process.Pid)

	sendInitCommand(comArray, writePipe)
	parent.Wait()
}
func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	fmt.Println("command all is: %s ", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
