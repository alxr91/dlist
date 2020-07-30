package main

import (
	"fmt"
	"syscall"
	linuxproc "github.com/c9s/goprocinfo/linux"
	"log"
)

func main() {

	fmt.Println("Hello, you!")
	getMem()
}


func getMem() {

	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Total Mem: ",info.Totalram / 1024)
	fmt.Println("Free Mem: ", info.Freeram / 1024)

}

func getCpu() {

	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}

	for _, s := range stat.CPUStats {

		fmt.Println("CPU for user: ", s.User)
		fmt.Println("CPU for system: ", s.System)
	}
}
