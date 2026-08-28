package system

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// SystemInfo represents the system information collected by DevSpeed
type SystemInfo struct {
	OS           string `json:"os"`
	OSVersion    string `json:"osVersion"`
	Arch         string `json:"arch"`
	CPUModel     string `json:"cpuModel"`
	LogicalCores int    `json:"logicalCores"`
	MemoryBytes  uint64 `json:"memoryBytes"`
}

// GetSystemInfo collects system information
func GetSystemInfo() (*SystemInfo, error) {
	info := &SystemInfo{
		OS:           runtime.GOOS,
		Arch:         runtime.GOARCH,
		LogicalCores: runtime.NumCPU(),
	}

	// Get OS version from /etc/os-release
	if info.OS == "linux" {
		if content, err := os.ReadFile("/etc/os-release"); err == nil {
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "PRETTY_NAME=") {
					info.OSVersion = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"\n\r")
					break
				}
			}
		}
	}

	// Get CPU information from lscpu
	if out, err := exec.Command("lscpu").Output(); err == nil {
		scanner := bufio.NewScanner(strings.NewReader(string(out)))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "Model name:") {
				info.CPUModel = strings.TrimSpace(strings.Split(line, ":")[1])
				break
			}
		}
	}

	// Get memory information from /proc/meminfo
	if content, err := os.ReadFile("/proc/meminfo"); err == nil {
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "MemTotal:") {
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					if memKB, err := strconv.ParseUint(parts[1], 10, 64); err == nil {
						info.MemoryBytes = memKB * 1024 // Convert KB to bytes
					}
				}
				break
			}
		}
	}

	return info, nil
}
