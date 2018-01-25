package main

import (
	"flag"
	"net/http"

	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"os/exec"
)

var addr = flag.String("listen-address", ":8080", "The address to listen to HTTP requests")

func main() {
	cmd := exec.Command("tdtool", "-l")
	out, _ := cmd.Output()
	fmt.Printf(string(out))
	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
