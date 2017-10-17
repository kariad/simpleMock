package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	PnrRecordLocator string `json:"pnrRecordLocator"`
}

func main() {
	setup()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setup() {
	http.HandleFunc("/response", func(responseWriter http.ResponseWriter, request *http.Request) {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("error")
		}
		byte, err := ioutil.ReadFile(dir + "/fixtures/response.json")
		if err != nil {
			print("could not read json file")
		}

		json := string(byte)

		fmt.Fprintf(responseWriter, string(json))
	})
}
