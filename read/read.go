package read

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Dmap map[string]map[string]string

func Init() {
	//	Dmap = make(map[string]string)
	data, err := ioutil.ReadFile("./table.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	//	log.Printf("data:%v", data)
	err = json.Unmarshal(data, &Dmap)
	if err != nil {
		log.Fatal("%v", err)
	}
	log.Printf("Dmap:%v", Dmap)
}
