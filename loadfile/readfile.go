package loadfile

import (
	"io/ioutil"
	"log"
	"os"
)

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
