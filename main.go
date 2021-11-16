package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Version string = "/v1"

func main() {

	route := mux.NewRouter()

	route.HandleFunc(Version+"/file", Loadfile)

	fmt.Println("Server started at port 9000")
	err := http.ListenAndServe(":9000", route)

	if err != nil {
		log.Println("Not Running Server...", err.Error())
	}
}

func Loadfile(w http.ResponseWriter, r *http.Request) {
	w.Write(ReadFile("file.json"))
}

func ReadFile(dataJSON string) []byte {
	data := loadData(dataJSON)
	return []byte(data)

}

func loadData(dataJSON string) []byte {
	jsonFile, err := os.Open(dataJSON)

	if err != nil {
		log.Println("Not found File", err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Println("Not Read File", err.Error())
	}

	return data

}
