package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Starting server")
	// TODO: Chain handlers and have a logging handler
	http.Handle("/export", exportHandler{})
	http.ListenAndServe(":8080", nil)
}

type exportHandler struct{}

func (exportHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running export")
	runMongoExport(w)
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

func runMongoExport(w http.ResponseWriter) {
	t := time.Now()
	timestamp := t.Format("01-02-06-3:04PM")
	commandString := fmt.Sprintf("mongoexport -d test -c accelData -u main_admin -p abc123 --authenticationDatabase admin --jsonArray --host mongodb-service:27017 > %s", timestamp)

	cmd := exec.Command("bash", "-c", commandString)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	respJSON := Response{}

	err := cmd.Run()
	if err != nil {
		fmt.Println("error!")
		fmt.Println(err.Error())
		fmt.Println(stderr.String())
		respJSON.Status = err.Error()
	} else {
		statusString := fmt.Sprintf("Success! %s", timestamp)
		respJSON.Status = statusString
	}
	// fmt.Println(stderr.String())
	fmt.Println(out.String())
	json.NewEncoder(w).Encode(respJSON)
}

// DBCountResponse is the json data type we respond to for Android that tells how many entries in the AccelData collection
type Response struct {
	// The `json` struct tag maps between the json name
	// and actual name of the field
	Status string `json:"status"`
}
