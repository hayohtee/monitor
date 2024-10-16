package hardware

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// GetSystemSection retrieve information about the system
// (operating system runtime, virtual memory, host name...).
func GetSystemSection() (string, error) {
	runtimeOS := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	hostStat, err := host.Info()
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Hostname: %s\nTotal Memory: %d\nUsed Memory: %d\nOS: %s",
		hostStat.Hostname, vmStat.Total, vmStat.Used, runtimeOS,
	)
	return output, nil
}

// GetCPUSection retrieve information about CPU
// (number of cores, model names...)
func GetCPUSection() (string, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("CPU: %s\nCores: %d", cpuStat[0].ModelName, cpuStat[0].Cores)
	return output, nil
}

// GetDiskSection retrieve information about the storage
// (total disk space, free disk space...)
func GetDiskSection() (string, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("Total Disk Space: %d\nFree Disk Space: %d", diskStat.Total, diskStat.Free)
	return output, nil
}
