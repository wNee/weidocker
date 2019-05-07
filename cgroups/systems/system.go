package systems

import (
	"error"
)

type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CupSet      string
}

type System interface {
	Name() string
	Set(path string, res *ResourceConfig)
	Apply(path string, pid int) error
	Remove(path string) error
}

var SystemsIns = []System{
	&CpuSubSystem{},
	&CpusetSystem{},
	&MemorySystem{},
}
