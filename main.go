package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

type Result struct {
	Method string `json:"method"`
	ID     string `json:"id"`
	Error  string `json:"error,omitempty"`
}

func outputJSON(res Result) {
	jsonRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println("{\"error\":\"JSON encoding failed\"}")
		return
	}
	fmt.Println(string(jsonRes))
}

func generateIDBasedOnCPU() {
	cpuInfo, err := cpu.Info()
	if err != nil {
		outputJSON(Result{Method: "cpu", Error: err.Error()})
		return
	}
	hash := md5.Sum([]byte(cpuInfo[0].ModelName))
	uniqueID := hex.EncodeToString(hash[:])
	outputJSON(Result{Method: "cpu", ID: uniqueID})
}

func generateIDBasedOnMAC() {
	interfaces, err := net.Interfaces()
	if err != nil {
		outputJSON(Result{Method: "mac", Error: err.Error()})
		return
	}
	for _, inter := range interfaces {
		if inter.HardwareAddr != "" {
			hash := md5.Sum([]byte(inter.HardwareAddr))
			uniqueID := hex.EncodeToString(hash[:])
			outputJSON(Result{Method: "mac", ID: uniqueID})
			return
		}
	}
	outputJSON(Result{Method: "mac", Error: "No valid MAC address found"})
}

func generateIDBasedOnDisk() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		outputJSON(Result{Method: "disk", Error: err.Error()})
		return
	}
	hash := md5.Sum([]byte(partitions[0].Device))
	uniqueID := hex.EncodeToString(hash[:])
	outputJSON(Result{Method: "disk", ID: uniqueID})
}

func generateTimestamp() {
	now := time.Now()
	milliseconds := now.UnixNano() / 1e6
	timestampStr := fmt.Sprintf("%d", milliseconds)
	outputJSON(Result{Method: "time", ID: timestampStr})
}

func main() {
	method := flag.String("method", "cpu", "The method to use for generating the unique ID. Options are: cpu, mac, disk, time")
	flag.Parse()

	switch *method {
	case "cpu":
		generateIDBasedOnCPU()
	case "mac":
		generateIDBasedOnMAC()
	case "disk":
		generateIDBasedOnDisk()
	case "time":
		generateTimestamp()
	default:
		outputJSON(Result{Method: *method, Error: "Invalid method. Available options are: cpu, mac, disk"})
	}
}
