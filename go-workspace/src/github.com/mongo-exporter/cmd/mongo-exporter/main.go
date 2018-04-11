package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Starting server")
	// TODO: Chain handlers and have a logging handler
	http.Handle("/", exportHandler{})
	http.ListenAndServe(":8080", nil)
}

type exportHandler struct{}

func (exportHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running export")
	runMongoExport()
	runls()
}

func runls() {
	command := "ls"
	cmd, err := exec.Command(command, "-al").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(cmd))
}

func runMongoExport() {
	t := time.Now()
	timestamp := t.Format("01-02-06-3:04PM")
	commandString := fmt.Sprintf("mongoexport -d test -c accelData -u main_admin -p abc123 --authenticationDatabase admin --jsonArray --host mongodb-service:27017 > %s", timestamp)

	cmd := exec.Command("bash", "-c", commandString)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("error!")
		fmt.Println(err.Error())
		fmt.Println(stderr.String())

	}
	// fmt.Println(stderr.String())
	fmt.Println(out.String())
}
