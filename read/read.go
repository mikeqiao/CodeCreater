package read

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/mikeqiao/codecreater/class"
	"github.com/mikeqiao/codecreater/file"
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
	//	log.Printf("Dmap:%v", Dmap)
}

func Create(path string) {
	for k, v := range Dmap {

		c := class.NewClass(k)
		c.Lock = false
		c.Path = path
		c.InitData(v)
		c.Init()
		c.ManagerInit()
		f := new(file.File)

		if c.IsData {
			path := k
			name := path + "/" + k + ".go"
			name2 := path + "/" + "manager.go"
			f.CreateDir(path)

			f.CreateFile(name)
			f.Write(c.GetBuff())
			f.Close()
			f2 := new(file.File)
			f2.CreateFile(name2)
			f2.Write(c.GetManagerBuff())
			f2.Close()
		} else {
			path := "common"
			name := path + "/" + k + ".go"

			f.CreateDir(path)

			f.CreateFile(name)
			f.Write(c.GetBuff())
			f.Close()

		}
	}
}
