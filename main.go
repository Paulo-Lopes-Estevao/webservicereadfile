package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	// 1. Create a struct for storing CSV lines and annotate it with JSON struct field tags
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
}

func main() {

	route := mux.NewRouter()

	route.HandleFunc("/", Loadfile)

	fmt.Println("Server started at port 9000")
	err := http.ListenAndServe(":9000", route)

	if err != nil {
		log.Println("Not Running Server...", err.Error())
	}
}

func Loadfile(w http.ResponseWriter, r *http.Request) {

	w.Write(ReadFile("file.csv"))
}

func ReadFile(dataJSON string) []byte {
	data := loadData(dataJSON)

	result := createJSONList(data)

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return []byte(jsonData)

}

func createJSONList(data [][]string) []Person {
	// convert csv lines to array of structs
	var person []Person
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Person
			for j, field := range line {
				if j == 0 {
					rec.Nome = field
				} else if j == 1 {
					rec.Email = field
				} else if j == 2 {
					var err error
					rec.Idade, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				}
			}
			person = append(person, rec)
		}
	}
	return person
}

func loadData(dataJSON string) [][]string {
	jsonFile, err := os.Open(dataJSON)

	if err != nil {
		log.Println("Not found File", err.Error())
	}
	defer jsonFile.Close()

	csvReader := csv.NewReader(jsonFile)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data

}
