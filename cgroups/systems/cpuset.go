package systems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpusetSystem struct{}

func (s *CpusetSystem) Set(cgPath string, res *ResourceConfig) error {
	sysPath, err := GetCgroupPath(s.Name(), cgPath, true)
	if err != nil {
		return err
	}
	if res.CpuSet != "" {
		if err := ioutil.WriteFile(path.Join(sysPath, "cpuset.cpus"), []byte(res.CpuSet), 0644); err != nil {
			return fmt.Errorf("set cgroup cpuset faill %v", err)
		}
	}
	return nil
}

func (s *CpusetSystem) Remove(cgroupPath string) error {
	sysPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return err
	}

	return os.RemoveAll(sysPath)
}

func (s *CpusetSystem) Apply(cgroupPath string, pid int) error {
	sysPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return fmt.Errorf("get cgroup proc fail %v", err)
	}
	if err := ioutil.WriteFile(path.Join(sysPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("set cgroup proc fail %v", err)
	}
	return nil
}

func (s *CpusetSystem) Name() string {
	return "cpuset"
}
