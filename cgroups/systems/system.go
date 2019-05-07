package systems

type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

type System interface {
	Name() string
	Set(path string, res *ResourceConfig) error
	Apply(path string, pid int) error
	Remove(path string) error
}

var SystemsIns = []System{
	&CpuSubSystem{},
	&CpusetSystem{},
	&MemorySystem{},
}
