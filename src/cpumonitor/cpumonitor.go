package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// TODO
// security??
// have some sort of chan of percentage
// trigger start, either via a handler or something
// get results, either via a handler or something

func start(w http.ResponseWriter, r *http.Request) {
	// start a goroutine for percentage
}

func stop(w http.ResponseWriter, r *http.Request) {
	// return the serialized stuff
}

func getCPU(w http.ResponseWriter, r *http.Request) {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		log.Fatal("Couldn't get percentage")
	}

	body := fmt.Sprintf("%#v", percent)
	w.Write([]byte(body))
}

func main() {
	http.HandleFunc("/", getCPU)             // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
