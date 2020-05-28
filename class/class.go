package class

import (
	"bytes"
	"fmt"

	p "github.com/mikeqiao/codecreater/param"

	f "github.com/mikeqiao/codecreater/function"
)

type Class struct {
	name   string
	params []*p.Param
	funcs  []*f.Function

	buff *bytes.Buffer
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
	c.buff.WriteString("	sync\n")
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
	c.buff.WriteString("}\n\n")
}

func (c *Class) InitParamFunc() {
	for _, v := range c.params {
		if nil != v {
			c.CreateSetFunc(v.Name, v.Type)
			c.CreateGetFunc(v.Name, v.Type)
		}
	}
}

func (c *Class) CreateSetFunc(name, ctype string) {
	head := fmt.Sprintf("func(this *%v) Set%v(value %v){\n", c.name, name, ctype)
	c.buff.WriteString(head)
	body := fmt.Sprintf("	this.%v = value\n", name)
	c.buff.WriteString(body)
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateGetFunc(name, ctype string) {
	head := fmt.Sprintf("func(this *%v) Get%v() %v{\n", c.name, name, ctype)
	c.buff.WriteString(head)
	body := fmt.Sprintf("	return this.%v\n", name)
	c.buff.WriteString(body)
	c.buff.WriteString("}\n\n")
}
