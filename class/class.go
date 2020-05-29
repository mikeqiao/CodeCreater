package class

import (
	"bytes"
	"fmt"
	"strconv"

	p "github.com/mikeqiao/codecreater/param"

	f "github.com/mikeqiao/codecreater/function"
)

type Class struct {
	name   string
	params []*p.Param
	funcs  []*f.Function
	Lock   bool
	buff   *bytes.Buffer
}

func NewClass(name string) *Class {
	c := new(Class)
	c.name = name
	c.buff = new(bytes.Buffer)

	return c
}

func (c *Class) GetBuff() (b *bytes.Buffer) {
	return c.buff
}

func (c *Class) InitData(d map[string]string) {
	for k, v := range d {
		np := new(p.Param)
		np.Name = k
		np.Type = v
		c.params = append(c.params, np)
	}
}

func (c *Class) Init() {
	c.InitPackage()
	c.InitImport()
	c.InitParam()
	c.CreateNewFunc()
	c.CreateInitDataFunc()
	c.InitParamFunc()
}

func (c *Class) InitPackage() {
	str := "package " + c.name
	c.buff.WriteString(str)
	c.buff.WriteString("\n\n")
}

func (c *Class) InitImport() {
	c.buff.WriteString("import (\n")
	//添加包含文件
	pak := strconv.Quote("sync")
	c.buff.WriteString("	" + pak + "\n")
	pak2 := strconv.Quote("github.com/mikeqiao/Db/redis")
	c.buff.WriteString("	" + pak2 + "\n")
	c.buff.WriteString(")\n\n")
}

func (c *Class) InitParam() {
	str := fmt.Sprintf("type %v struct {\n", c.name)
	c.buff.WriteString(str)

	for _, v := range c.params {
		if nil != v {
			c.buff.WriteString("	")
			c.buff.WriteString(v.Name)
			c.buff.WriteString("	")
			c.buff.WriteString(v.Type)
			c.buff.WriteString("\n")
		}
	}
	c.AddMap()
	c.AddLock()
	c.buff.WriteString("}\n\n")
}

func (c *Class) AddLock() {
	if c.Lock {
		c.buff.WriteString("	mutex sync.RWMutex\n")
	}
}

func (c *Class) AddMap() {

	c.buff.WriteString("	changeData map[string]interface{}\n")

}

func (c *Class) InitParamFunc() {
	for _, v := range c.params {
		if nil != v {
			c.CreateSetFunc(v.Name, v.Type)
			c.CreateGetFunc(v.Name, v.Type)
		}
	}
}

func (c *Class) CreateNewFunc() {
	head := fmt.Sprintf("func New%v() *%v{\n", c.name, c.name)
	c.buff.WriteString(head)

	body := fmt.Sprintf("	data := new(%v)\n", c.name)
	c.buff.WriteString(body)
	c.buff.WriteString("	data.changeData= make(map[string]interface{})\n")
	c.buff.WriteString("	return data\n")
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateInitDataFunc() {
	head := fmt.Sprintf("func (this *%v)InitData() {\n", c.name)
	c.buff.WriteString(head)
	t := fmt.Sprintf("	table = %v + _fmt.Sprint(this.uid)\n", c.name)
	c.buff.WriteString(t)
	c.buff.WriteString("	data, _:=redis.R.Hash_GetAllData(table)\n")
	for _, v := range c.params {
		if nil != v {
			namestr := strconv.Quote(v.Name)
			key := fmt.Sprintf("	if d,ok:=data[%v];ok{\n", namestr)
			c.buff.WriteString(key)
			dvalue := ""
			have := true
			switch v.Type {
			case "string":
				dvalue = fmt.Sprintf("		dv:=d\n")
				c.buff.WriteString(dvalue)
			case "uint64":
				dvalue = fmt.Sprintf("		dv, _:=strconv.ParseUint(d,10,64)\n") //strconv.ParseFloat() ParseUint(d,10,64)
				c.buff.WriteString(dvalue)
			case "uint32":
				dvalue = fmt.Sprintf("		dd, _:=strconv.ParseUint(d,10,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("		dv:=uint32(dd)\n")
				c.buff.WriteString(nvalue)
			case "int32":
				dvalue = fmt.Sprintf("		dd, _:=strconv.Atoi(d)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("		dv:=int32(dd)\n")
				c.buff.WriteString(nvalue)
			case "int64":
				dvalue = fmt.Sprintf("		dv, _:=strconv.ParseInt(d,10,64)\n")
				c.buff.WriteString(dvalue)
			case "float64":
				dvalue = fmt.Sprintf("		dv, _:=strconv.ParseFloat(d,64)\n")
				c.buff.WriteString(dvalue)
			case "float32":
				dvalue = fmt.Sprintf("		dd, _:=strconv.ParseFloat(d,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("		dv:=float32(dd)\n")
				c.buff.WriteString(nvalue)
			case "bool":
				dvalue = fmt.Sprintf("		dv, _:=strconv.ParseBool(d)\n")
				c.buff.WriteString(dvalue)
			default:
				have = false
			}
			if have {
				value := fmt.Sprintf("		this.%v= dv\n", v.Name)
				c.buff.WriteString(value)

			}

			c.buff.WriteString("	}\n\n")
		}
	}
	body := fmt.Sprintf("	data := new(%v)\n", c.name)
	c.buff.WriteString(body)
	c.buff.WriteString("	data.changeData= make(map[string]interface{})\n")
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateSetFunc(name, ctype string) {
	head := fmt.Sprintf("func(this *%v) Set%v(value %v){\n", c.name, name, ctype)
	c.buff.WriteString(head)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}
	body := fmt.Sprintf("	this.%v = value\n", name)
	c.buff.WriteString(body)
	namestr := strconv.Quote(name)
	add := fmt.Sprintf("	this.changeData[%v]= value\n", namestr)
	c.buff.WriteString(add)
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateGetFunc(name, ctype string) {
	head := fmt.Sprintf("func(this *%v) Get%v() %v{\n", c.name, name, ctype)
	c.buff.WriteString(head)

	if c.Lock {
		c.buff.WriteString("	this.mutex.RLock()\n")
		c.buff.WriteString("	defer this.mutex.RUnlock()\n")
	}

	body := fmt.Sprintf("	return this.%v\n", name)
	c.buff.WriteString(body)
	c.buff.WriteString("}\n\n")
}
