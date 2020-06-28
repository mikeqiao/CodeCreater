package read

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/mikeqiao/codecreater/class"
	"github.com/mikeqiao/codecreater/file"
)

var Dmap map[string]map[string]string
var Smap map[string]map[string]map[string]string

func InitData() {
	//	Dmap = make(map[string]string)
	data, err := ioutil.ReadFile("./data.json")
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

func CreateData(path string) {
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

func InitService() {
	//	Dmap = make(map[string]string)
	data, err := ioutil.ReadFile("./service.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	//	log.Printf("data:%v", data)
	err = json.Unmarshal(data, &Smap)
	if err != nil {
		log.Fatal("%v", err)
	}
	//	log.Printf("Dmap:%v", Dmap)
}

func CreateService(path string) {
	for k, v := range Smap {

		c := class.NewMod(k)
		c.Path = path
		c.InitData(v)
		c.Init()
		f := new(file.File)

		path := k
		name := path + "/" + k + ".go"
		f.CreateDir(path)

		f.CreateFile(name)
		f.Write(c.GetModBuff())
		f.Close()
		for _, s := range c.Params {
			if nil != s {
				c.CreateServiceFile(s)
				name2 := path + "/" + s.Name + ".go"
				f2 := new(file.File)
				f2.CreateFile(name2)
				f2.Write(c.GetServiceBuff())
				f2.Close()
			}
		}
	}
}
