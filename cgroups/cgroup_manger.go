package cgroups

import (
	"fmt"
	"weidocker/cgroups/systems"
)

type CgroupManger struct {
	Path     string
	Resource *systems.ResourceConfig
}

func NewCgroupManger(path string) *CgroupManger {
	return &CgroupManger{
		Path: path,
	}
}

// add pid to cgroup
func (c *CgroupManger) Apply(pid int) error {
	for _, sysIns := range systems.SystemsIns {
		sysIns.Apply(c.Path, pid)
	}
	return nil
}

// set cgroup
func (s *CgroupManger) Set(res *systems.ResourceConfig) error {
	for _, sysIns := range systems.SystemsIns {
		sysIns.Set(s.Path, res)
	}
	return nil
}

// destory cgroup
func (s *CgroupManger) Destory() error {
	for _, sysIns := range systems.SystemsIns {
		if err := sysIns.Remove(s.Path); err != nil {
			fmt.Println("remove cgroup fail %v", err)
			return err
		}
	}
	return nil
}
