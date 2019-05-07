package systems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuSubSystem struct{}

func (s *CpuSubSystem) Set(cgPatg string, res *ResourceConfig) error {
	if cgPath, err := GetCgroupPath(s.Name(), cgPatg, true); err == nil {
		if res.CpuShare != "" {
			if err := ioutil.WriteFile(path.Join(cgPath, "cpu.shares"), []byte(res.CpuShare), 0644); err != nil {
				return fmt.Errorf("set cgroup cpu share fail %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}

func (s *CpuSubSystem) Remove(cgroupPath string) error {
	sysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	return os.RemoveAll(sysCgroupPath)
}
func (s *CpuSubSystem) Apply(cgroupPath string, pid int) error {
	sysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
	}
	if err := ioutil.WriteFile(path.Join(sysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("set cgroup proc fail %v", err)
	}
	return nil
}

func (s *CpuSubSystem) Name() string {
	return "cpu"
}
