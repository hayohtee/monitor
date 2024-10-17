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

			timestamp := time.Now().Format("2006-01-02 15:04:05")

			htmlContainer := `
			<div hx-swap-oob="innerHTML:#update-timestamp">
				<p><i style="color: green" class="fa fa-circle"></i> ` + timestamp + `</p>
		 	 </div>
		  	<div hx-swap-oob="innerHTML:#system-data">` + systemSection + `</div>
		  	<div hx-swap-oob="innerHTML:#cpu-data">` + cpuSection + `</div>
		  	<div hx-swap-oob="innerHTML:#disk-data">` + diskSection + `</div>`

			s.broadcast([]byte(htmlContainer))

			time.Sleep(3 * time.Second)
		}
	}(srv)

	err := http.ListenAndServe(":4000", &srv.mux)
	if err != nil {
		log.Fatal(err)
	}
}
