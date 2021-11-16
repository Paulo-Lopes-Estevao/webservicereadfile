package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Paulo-Lopes-Estevaogochallenge_web/loadfile"
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
	w.Write(loadfile.ReadFile("file.json"))
}
