package read

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/mikeqiao/codecreater/class"
	"github.com/mikeqiao/codecreater/file"
)

var Dmap map[string]map[string]string
var Smap map[string]map[string]map[string]string
var MsgMap map[string]string

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
			path := class.DataPath + k
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
			tpath := class.DataPath + "common"
			name := tpath + "/" + k + ".go"

			f.CreateDir(tpath)

			f.CreateFile(name)
			f.Write(c.GetBuff())
			f.Close()

		}
	}
}

func InitService() {
	MsgMap = make(map[string]string)
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
	mf := new(file.File)
	name := class.ProtoPath + "msg.go"

	mf.CreateDir(class.ProtoPath)
	mf.CreateFile(name)
	Msgbuff := new(bytes.Buffer)
	CreateMsgHead(Msgbuff, path)
	Msgbuff.WriteString("func init(){\n")
	for k, v := range Smap {
		Msg(v, Msgbuff)
		c := class.NewMod(k)
		c.Path = path
		c.InitData(v)
		c.Init()
		f := new(file.File)

		path := class.ModPath + k
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
	Msgbuff.WriteString("}\n\n")
	mf.Write(Msgbuff)
	mf.Close()
}

func Msg(d map[string]map[string]string, buff *bytes.Buffer) {
	for k1, v1 := range d {
		if req, ok := v1["Req"]; ok {
			name := strFirstToUpper(req)
			if _, ok := MsgMap[name]; !ok {

				b := fmt.Sprintf("	m.DefaultProcessor.RegisterMsg(%v, reflect.TypeOf(proto.%v{}))\n", strconv.Quote(name), name)
				buff.WriteString(b)
				MsgMap[name] = name
			}
		} else {
			fmt.Printf("this service have no req info, service:%v\n", k1)
			return
		}
		if res, ok := v1["Res"]; ok {
			name := strFirstToUpper(res)
			if _, ok := MsgMap[name]; !ok {
				b := fmt.Sprintf("	m.DefaultProcessor.RegisterMsg(%v,  reflect.TypeOf(proto.%v{}))\n", strconv.Quote(name), name)
				buff.WriteString(b)
				MsgMap[name] = name
			}
		} else {
			fmt.Printf("this service have no res info, service:%v\n", k1)
			return
		}
	}
}

func CreateMsgHead(buff *bytes.Buffer, path string) {
	str := "package msg"
	buff.WriteString(str)
	buff.WriteString("\n\n")
	buff.WriteString("import (\n")
	pak := strconv.Quote("github.com/mikeqiao/newworld/manager")
	buff.WriteString("	m" + pak + "\n")
	pak2 := strconv.Quote("reflect")
	buff.WriteString("	" + pak2 + "\n")

	tpath := path + "/proto"
	pak5 := strconv.Quote(tpath)
	buff.WriteString("	" + pak5 + "\n")

	buff.WriteString(")\n\n")
}

func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
