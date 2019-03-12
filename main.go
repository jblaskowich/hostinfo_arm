package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	arch := runtime.GOARCH
	operatingSystem := runtime.GOOS
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hostname:", host, "\nArch:", arch, "\nOS:", operatingSystem)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
