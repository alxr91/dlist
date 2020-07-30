package main

import (
	"context"
	"fmt"
	"log"
	"syscall"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func main() {

	getMem()
}

func getMem() {

	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("Total Mem: ", info.Totalram/1024)
	fmt.Println("Free Mem: ", info.Freeram/1024)

	getCpu()

}

func getCpu() {

	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}

	fmt.Println("")

	fmt.Println(stat.CPUStatAll)

	getContainers()
}

func getContainers() error {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) > 0 {
		for _, container := range containers {

			fmt.Println("Running containers")
			fmt.Println("Container name: ", container.Names)

		}
	} else {

		fmt.Println("No containers to list..")
	}
	return nil
}
