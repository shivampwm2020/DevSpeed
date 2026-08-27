package system

import (
	"runtime"
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

	// For development, we'll use placeholder values
	// In a real implementation, these would be determined by system calls
	switch runtime.GOOS {
	case "darwin":
		info.CPUModel = "Apple M1"
		info.OSVersion = "macOS 14.0"
		info.MemoryBytes = 16 * 1024 * 1024 * 1024 // 16 GB
	case "linux":
		info.CPUModel = "AMD Ryzen 9"
		info.OSVersion = "Ubuntu 22.04"
		info.MemoryBytes = 32 * 1024 * 1024 * 1024 // 32 GB
	case "windows":
		info.CPUModel = "Intel i7"
		info.OSVersion = "Windows 11"
		info.MemoryBytes = 16 * 1024 * 1024 * 1024 // 16 GB
	default:
		info.CPUModel = "Unknown CPU"
		info.OSVersion = "Unknown Version"
		info.MemoryBytes = 8 * 1024 * 1024 * 1024 // 8 GB
	}

	return info, nil
}
