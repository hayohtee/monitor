package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hayohtee/monitor/internal/hardware"
)

func main() {
	log.Println("starting system monitor")
	go func() {
		for {
			systemSection, err := hardware.GetSystemSection()
			if err != nil {
				log.Fatal(err)
			}

			cpuSection, err := hardware.GetCPUSection()
			if err != nil {
				log.Fatal(err)
			}

			diskSection, err := hardware.GetDiskSection()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(systemSection)
			fmt.Println(diskSection)
			fmt.Println(cpuSection)

			time.Sleep(3 * time.Second)
		}
	}()
	time.Sleep(5 * time.Minute)
}
