package status

import (
	log "github.com/liuzheng712/golog"
	"fmt"
	//"runtime"
	json "github.com/pquerna/ffjson/ffjson"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

func Mem() {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Error("status", "%v", err)
	}
	//var m runtime.MemStats
	//runtime.ReadMemStats(&m)
	res, err := json.Marshal(v)
	if err != nil {
		log.Error("status", "%v", err)
	}
	fmt.Println(string(res))
}
func CPU() {
	v, err := cpu.Info()
	if err != nil {
		log.Error("status", "%v", err)
	}
	res, err := json.Marshal(v)
	if err != nil {
		log.Error("status", "%v", err)
	}
	fmt.Println(string(res))
}
func Load() {
	v, err := load.Avg()
	if err != nil {
		log.Error("status", "%v", err)
	}
	res, err := json.Marshal(v)
	if err != nil {
		log.Error("status", "%v", err)
	}
	fmt.Println(string(res))
}
func Disk() {
	v, err := disk.Partitions(true)
	if err != nil {
		log.Error("status", "%v", err)
	}
	res, err := json.Marshal(v)
	if err != nil {
		log.Error("status", "%v", err)
	}
	fmt.Println(string(res))
}
func Net(){
		v, err := net.IOCounters(true)
	if err != nil {
		log.Error("status", "%v", err)
	}
	res, err := json.Marshal(v)
	if err != nil {
		log.Error("status", "%v", err)
	}
	fmt.Println(string(res))
}
