package systems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySystem struct{}

func (s *MemorySystem) Set(cgPath string, res *ResourceConfig) error {
	sysPath, err := GetCgroupPath(s.Name(), cgPath, true)
	if err != nil {
		return err
	}
	if res.MemoryLimit != "" {
		if err := ioutil.WriteFile(path.Join(sysPath, "memmory.limit_in_bytes"), []byte(res.CpuSet), 0644); err != nil {
			return fmt.Errorf("set cgroup memmory faill %v", err)
		}
	}
	return nil
}

func (s *MemorySystem) Remove(cgroupPath string) error {
	sysPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return err
	}

	return os.RemoveAll(sysPath)
}

func (s *MemorySystem) Apply(cgroupPath string, pid int) error {
	sysPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return fmt.Errorf("get cgroup fail %v", err)
	}
	if err := ioutil.WriteFile(path.Join(sysPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("set cgroup proc fail %v", err)
	}
	return nil
}

func (s *MemorySystem) Name() string {
	return "memory"
}
