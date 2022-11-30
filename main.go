package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

const Host = "http://localhost:5601"

func main() {
	res, err := http.Get(Host)
	if err != nil || res.StatusCode >= 500 {
		restart()
	}
}

func restart() {
	cmd := exec.Command("sudo", "service", "elasticsearch", "restart")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Elasticsearch restarting...")
}
