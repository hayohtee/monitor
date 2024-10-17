package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hayohtee/monitor/internal/hardware"
)

func main() {
	log.Println("starting system monitor")
	srv := NewServer()

	go func(s *server) {
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

			s.broadcast([]byte(systemSection))
			s.broadcast([]byte(cpuSection))
			s.broadcast([]byte(diskSection))

			time.Sleep(3 * time.Second)
		}
	}(srv)

	err := http.ListenAndServe(":4000", &srv.mux)
	if err != nil {
		log.Fatal(err)
	}
}
