package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"syscall"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func main() {

	getMem()
	getCpu()

	getContainers()
}

func getMem() {

	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	memTotal := strconv.Itoa(int(info.Totalram / 1024))
	memFree := strconv.Itoa(int(info.Freeram / 1024))

	fmt.Println("Total Mem: " + memTotal + " Available Mem: " + memFree)

}

func getCpu() {

	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}

	fmt.Println("")

	fmt.Println(stat.CPUStatAll)

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

		fmt.Println("Running containers")
		for _, container := range containers {

			fmt.Println("Container name: ", container.Names)

		}
	} else {

		fmt.Println("No running containers")
	}

	args := filters.NewArgs(filters.KeyValuePair{"status", "exited"})
	stoppedContainer, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Filters: args})

	if err != nil {
		panic(err)
	}

	if len(stoppedContainer) > 0 {

		fmt.Println("")
		fmt.Println("Stopped containers")

		for _, exited := range stoppedContainer {

			fmt.Println("Stopped container: ", exited.Names)
		}

	} else {

		fmt.Println("You don't have any stopped containers.")
	}

	return nil
}
