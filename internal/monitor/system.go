package monitor

import (
	"math"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type RealTimeStats struct {
	CPUUsage     float64 `json:"cpu_usage"`
	MemoryUsage  float64 `json:"memory_usage"` // percentage
	MemoryTotal  uint64  `json:"memory_total"`
	MemoryUsed   uint64  `json:"memory_used"`
	HostName     string  `json:"host_name"`
	OS           string  `json:"os"`
	Platform     string  `json:"platform"`
	PlatformVer  string  `json:"platform_ver"`
	KernelVer    string  `json:"kernel_ver"`
	CPUModel     string  `json:"cpu_model"`
	CPUCores     int     `json:"cpu_cores"`
	Uptime       uint64  `json:"uptime"`
	ResponseTime int     `json:"response_time"` // ms (mocked for now, or ping database)
	Stability    float64 `json:"stability"`     // mocked or calculated
	HealthScore  float64 `json:"health_score"`  // mocked or calculated
}

func GetRealTimeStats() (*RealTimeStats, error) {
	stats := &RealTimeStats{}

	// Host Info
	hostInfo, err := host.Info()
	if err == nil {
		stats.HostName = hostInfo.Hostname
		stats.OS = hostInfo.OS
		stats.Platform = hostInfo.Platform
		stats.PlatformVer = hostInfo.PlatformVersion
		stats.KernelVer = hostInfo.KernelVersion
		stats.Uptime = hostInfo.Uptime
	}

	// CPU Info
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err == nil && len(cpuPercent) > 0 {
		stats.CPUUsage = math.Round(cpuPercent[0]*10) / 10 // Round to 1 decimal
	}

	cpuInfo, err := cpu.Info()
	if err == nil && len(cpuInfo) > 0 {
		stats.CPUModel = cpuInfo[0].ModelName
		stats.CPUCores = runtime.NumCPU()
	}

	// Memory Info
	vmStat, err := mem.VirtualMemory()
	if err == nil {
		stats.MemoryUsage = math.Round(vmStat.UsedPercent*10) / 10
		stats.MemoryTotal = vmStat.Total
		stats.MemoryUsed = vmStat.Used
	}

	// Mock/Calculated fields
	stats.ResponseTime = 45 // Could be ping to DB
	stats.Stability = 99.9  // Could be based on uptime
	stats.HealthScore = 100 - (stats.CPUUsage/2 + stats.MemoryUsage/2)
	if stats.HealthScore < 0 {
		stats.HealthScore = 0
	} else {
		stats.HealthScore = math.Round(stats.HealthScore*10) / 10
	}

	return stats, nil
}
