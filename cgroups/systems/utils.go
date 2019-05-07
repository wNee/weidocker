package systems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func GetCgroupPath(system string, cgPath string, autoCreate bool) (string, error) {
	cgRoot := FindCgroupMountpoint(system)
	if _, err := os.Stat(path.Join(cgRoot, cgPath)); err != nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgRoot, cgPath), 0755); err != nil {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(cgRoot, cgPath), nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}

func FindCgroupMountpoint(system string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == system {
				return fields[4]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}
