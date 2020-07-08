package class

import (
	"bytes"
	"fmt"
	"strconv"

	p "github.com/mikeqiao/codecreater/param"
)

type Mod struct {
	modName     string
	Path        string
	Params      []*p.Server
	Servicebuff *bytes.Buffer
	Modbuff     *bytes.Buffer
}

func NewMod(name string) *Mod {
	name = strFirstToUpper(name)
	c := new(Mod)
	c.modName = name
	c.Servicebuff = new(bytes.Buffer)
	c.Modbuff = new(bytes.Buffer)
	return c
}

func (m *Mod) InitData(d map[string]map[string]string) {
	for k, v := range d {
		ns := new(p.Server)
		ns.Name = k
		if req, ok := v["Req"]; ok {
			ns.Req = req
		} else {
			fmt.Printf("this service have no req info, service:%v\n", k)
			return
		}
		if res, ok := v["Res"]; ok {
			ns.Res = res
		} else {
			fmt.Printf("this service have no res info, service:%v\n", k)
			return
		}
		m.Params = append(m.Params, ns)
	}
}

func (m *Mod) Init() {
	m.CreateModFile()
}

func (m *Mod) GetModBuff() *bytes.Buffer {
	return m.Modbuff
}

func (m *Mod) GetServiceBuff() *bytes.Buffer {
	return m.Servicebuff
}

func (m *Mod) CreateModFile() {
	m.CreateModHead()
	m.CreateModFunc()
}

func (m *Mod) CreateModHead() {
	str := "package " + m.modName
	m.Modbuff.WriteString(str)
	m.Modbuff.WriteString("\n\n")
	m.Modbuff.WriteString("import (\n")
	pak := strconv.Quote("github.com/mikeqiao/newworld/manager")
	m.Modbuff.WriteString("	m" + pak + "\n")
	pak2 := strconv.Quote("github.com/mikeqiao/newworld/module")
	m.Modbuff.WriteString("	mod" + pak2 + "\n")
	pak3 := strconv.Quote(m.Path + "/proto")
	m.Modbuff.WriteString("	" + pak3 + "\n")
	m.Modbuff.WriteString(")\n\n")
	m.Modbuff.WriteString("var Mod *mod.Mod\n\n")
}

func (m *Mod) CreateModFunc() {
	m.Modbuff.WriteString("func Init(){\n")

	body := fmt.Sprintf("	Mod = m.NewMod(0,%v)\n", strconv.Quote(m.modName))
	m.Modbuff.WriteString(body)
	m.Modbuff.WriteString("	Register()\n")
	m.Modbuff.WriteString("	m.ModManager.Registe(Mod)\n")
	m.Modbuff.WriteString("}\n\n")

	m.Modbuff.WriteString("func Register(){\n")
	for _, v := range m.Params {
		if nil != v {
			name := strconv.Quote(v.Name)
			s := fmt.Sprintf("	Mod.Register(%v, %v, proto.%v{}, proto.%v{})\n", name, v.Name, strFirstToUpper(v.Req), strFirstToUpper(v.Res))
			m.Modbuff.WriteString(s)
		}
	}
	m.Modbuff.WriteString("}\n\n")
}

func (m *Mod) CreateServiceFile(v *p.Server) {
	m.Servicebuff = new(bytes.Buffer)
	m.CreateServiceHead()
	m.CreateServiceFunc(v)
}

func (m *Mod) CreateServiceHead() {
	str := "package " + m.modName
	m.Servicebuff.WriteString(str)
	m.Servicebuff.WriteString("\n\n")
	m.Servicebuff.WriteString("import (\n")
	pak2 := strconv.Quote("github.com/mikeqiao/newworld/module")
	m.Servicebuff.WriteString("	mod" + pak2 + "\n")
	path := m.Path + "/msg"
	pak5 := strconv.Quote(path)
	m.Servicebuff.WriteString("	" + pak5 + "\n")

	m.Servicebuff.WriteString(")\n\n")
}

func (m *Mod) CreateServiceFunc(v *p.Server) {

	if nil != v {
		h := fmt.Sprintf("func %v(call *mod.CallInfo) {", v.Name)
		m.Servicebuff.WriteString(h)
	}

	m.Servicebuff.WriteString("}\n\n")
}
