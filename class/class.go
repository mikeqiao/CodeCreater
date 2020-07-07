package class

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	p "github.com/mikeqiao/codecreater/param"

	f "github.com/mikeqiao/codecreater/function"
)

type Class struct {
	name        string
	managername string
	Path        string
	params      []*p.Param
	funcs       []*f.Function
	Lock        bool
	IsData      bool
	IsUpdate    bool
	HaveChield  bool
	HaveMap     bool
	buff        *bytes.Buffer
	managerbuff *bytes.Buffer
}

func NewClass(name string) *Class {
	name = strFirstToUpper(name)
	c := new(Class)
	c.name = name
	c.managername = name + "Manager"
	c.buff = new(bytes.Buffer)
	c.managerbuff = new(bytes.Buffer)
	c.CheckName()
	return c
}

func (c *Class) CheckName() {
	if strings.HasPrefix(c.name, "Data") {
		c.IsData = true
	}
	if strings.HasPrefix(c.name, "Update") {
		c.IsUpdate = true
	}
}

func (c *Class) GetBuff() (b *bytes.Buffer) {
	return c.buff
}

func (c *Class) GetManagerBuff() (b *bytes.Buffer) {
	return c.managerbuff
}

func (c *Class) InitData(d map[string]string) {
	for k, v := range d {
		np := new(p.Param)
		nstr := strFirstToLower(k)
		np.Name = nstr
		np.Type = v
		if CheckBaseType(v) {
			np.TType = 1
			np.MTye = v
			np.UType = v
		} else {

			is, mtype := CheckStruct(v)
			if is {
				ctype := strings.TrimLeft(mtype, "*common.")
				ctype = "*" + ctype + "." + ctype
				np.TType = 2
				np.MTye = mtype
				np.UType = ctype
				c.HaveChield = true
			} else if CheckMap(v) {
				c.HaveMap = true
				np.TType = 3
				is, ktype, vtype, mtype := CheckMapStruct(v)
				if is {
					np.TType = 4
					c.HaveChield = true
				}
				np.MTye = mtype
				np.Vtype = vtype
				np.Ktype = ktype
			}
		}
		c.params = append(c.params, np)
	}
}

func (c *Class) Init() {
	c.InitPackage()

	if c.IsUpdate {
		c.InitUpdateImport()
		c.InitUpdateParam()
		c.CreateNewUpdateFunc()
		c.InitUpdateParamFunc()

	} else {
		c.InitImport()
		c.InitParam()
		c.CreateNewFunc()
		c.InitParamFunc()
	}
	if c.IsData {
		c.CreateInitDataFuncNew()
		c.CreateUpdateFunc()
		c.CreateClose()
	} else if !c.IsUpdate {
		c.CreateInitDataParamFunc()
	}
	c.CreateDestroy()
}

func (c *Class) InitPackage() {

	str := "package " + strFirstToLower(c.name)
	if !c.IsData && !c.IsUpdate {
		str = "package common"
	}
	c.buff.WriteString(str)
	c.buff.WriteString("\n\n")
}

func (c *Class) InitImport() {
	c.buff.WriteString("import (\n")
	//添加包含文件
	if c.Lock {
		pak := strconv.Quote("sync")
		c.buff.WriteString("	" + pak + "\n")
	}
	pak2 := strconv.Quote("github.com/mikeqiao/newworld/data")
	c.buff.WriteString("	" + pak2 + "\n")

	pak3 := strconv.Quote("strconv")
	c.buff.WriteString("	" + pak3 + "\n")

	if c.HaveChield {
		path := c.Path + "/common"
		pak5 := strconv.Quote(path)
		c.buff.WriteString("	" + pak5 + "\n")
	}
	if c.IsData {
		pak6 := strconv.Quote("strings")
		c.buff.WriteString("	" + pak6 + "\n")
	}
	pak4 := strconv.Quote("fmt")
	c.buff.WriteString("	" + pak4 + "\n")
	c.buff.WriteString(")\n\n")
}

func (c *Class) InitUpdateImport() {
	c.buff.WriteString("import (\n")
	//添加包含文件
	for _, v := range c.params {
		if 2 == v.TType {
			ctype := strings.TrimLeft(v.MTye, "*common.")
			path := c.Path + "/" + ctype
			pak5 := strconv.Quote(path)
			c.buff.WriteString("	" + pak5 + "\n")
		}

	}
	c.buff.WriteString(")\n\n")
}

func (c *Class) InitParam() {
	str := fmt.Sprintf("type %v struct {\n", c.name)
	c.buff.WriteString(str)
	have := false
	for _, v := range c.params {
		if nil != v {
			//先判断是否是基础类型
			ctype := v.Type
			if 2 == v.TType {
				ctype = v.MTye
			}
			if 4 == v.TType {
				ctype = v.MTye
			}
			c.buff.WriteString("	")
			c.buff.WriteString(v.Name)
			c.buff.WriteString("	")
			c.buff.WriteString(ctype)
			c.buff.WriteString("\n")
			if "uid" == v.Name {
				have = true
			}
			if 2 == v.TType {

			}
		}
	}
	if !have {
		c.buff.WriteString("	uid	uint64\n")
	}

	c.AddUpdate()
	c.AddLock()
	c.buff.WriteString("}\n\n")
}

func (c *Class) InitUpdateParam() {
	str := fmt.Sprintf("type %v struct {\n", c.name)
	c.buff.WriteString(str)
	have := false
	for _, v := range c.params {
		if nil != v {
			//先判断是否是基础类型
			ctype := v.Type
			if 2 == v.TType {
				ctype = strings.TrimLeft(v.MTye, "*common.")
				ctype = "*" + ctype + "." + ctype
			}

			c.buff.WriteString("	")
			c.buff.WriteString(v.Name)
			c.buff.WriteString("	")
			c.buff.WriteString(ctype)
			c.buff.WriteString("\n")
			if "uid" == v.Name {
				have = true
			}
			if 2 == v.TType {

			}
		}
	}
	if !have {
		c.buff.WriteString("	uid	uint64\n")
	}

	c.AddUpdate()
	c.AddLock()
	c.buff.WriteString("}\n\n")
}

func (c *Class) AddLock() {
	if c.Lock {
		c.buff.WriteString("	mutex sync.RWMutex\n")
	}
}

func (c *Class) AddUpdate() {
	if !c.IsUpdate {
		c.buff.WriteString("	prefix string\n")
		c.buff.WriteString("	update *data.UpdateMod\n")
	}
}

func (c *Class) InitParamFunc() {
	for _, v := range c.params {
		if nil != v && 3 != v.TType && 4 != v.TType {
			c.CreateSetFunc(v.Name, v.MTye, v.TType)
			c.CreateGetFunc(v.Name, v.MTye)
		}
		if nil != v && 3 == v.TType {
			c.CreateMapFunc(v.Name, v.Ktype, v.Vtype)
		}
		if nil != v && 4 == v.TType {
			c.CreateMapStructFunc(v.Name, v.Ktype, v.Vtype)
		}
	}
}

func (c *Class) InitUpdateParamFunc() {
	for _, v := range c.params {
		if nil != v && 3 != v.TType && 4 != v.TType {
			c.CreateSetFunc(v.Name, v.UType, 2)
			c.CreateGetFunc(v.Name, v.UType)
		}

	}
}

func (c *Class) CreateNewFunc() {
	head := fmt.Sprintf("func New%v(uid uint64, prefix string, update *data.UpdateMod) *%v{\n", c.name, c.name)
	c.buff.WriteString(head)

	body := fmt.Sprintf("	d := new(%v)\n", c.name)
	c.buff.WriteString(body)
	c.buff.WriteString("	d.uid= uid\n")
	c.buff.WriteString("	d.prefix= prefix\n")
	c.buff.WriteString("	if nil != update{\n")
	c.buff.WriteString("		d.update= update\n")
	c.buff.WriteString("	}else{\n")
	c.buff.WriteString("		d.update= new(data.UpdateMod)\n")
	namestr := strconv.Quote(c.name + "_")
	t := fmt.Sprintf("		table := %v + fmt.Sprint(d.uid)\n", namestr)
	c.buff.WriteString(t)
	c.buff.WriteString("		d.update.Init(table)\n")
	c.buff.WriteString("	}\n")

	c.buff.WriteString("	return d\n")
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateNewUpdateFunc() {
	head := fmt.Sprintf("func New%v(uid uint64) *%v{\n", c.name, c.name)
	c.buff.WriteString(head)

	body := fmt.Sprintf("	d := new(%v)\n", c.name)
	c.buff.WriteString(body)
	c.buff.WriteString("	d.uid= uid\n")
	c.buff.WriteString("	return d\n")
	c.buff.WriteString("}\n\n")
}

/*
func (c *Class) CreateInitData2Func() {
	head := fmt.Sprintf("func (this *%v)InitData(data map[string]string) {\n", c.name)
	c.buff.WriteString(head)
	c.buff.WriteString("	if nil == data{\n")
	c.buff.WriteString("		return\n")
	c.buff.WriteString("	}\n")

	for _, v := range c.params {
		if nil != v && 1 == v.TType {
			namestr := strconv.Quote(v.Name)
			key := fmt.Sprintf("	if d,ok:=data[this.prefix+%v];ok{\n", namestr)
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

	for _, v := range c.params {
		if 2 == v.TType {
			stype := strings.TrimLeft(v.Type, "*")
			prefix := strconv.Quote("this.prefix" + v.Name + ".")
			value := fmt.Sprintf("	this.%v= common.New%v(this.uid, %v, this.update)\n", v.Name, stype, prefix)
			//	value := fmt.Sprintf("		this.%v.\n", v.Name)
			c.buff.WriteString(value)
			value2 := fmt.Sprintf("	this.%v.InitData(data)\n", v.Name)
			c.buff.WriteString(value2)

		}
	}

	for _, v := range c.params {
		if 3 == v.TType {
			tvalue := fmt.Sprintf("	this.%v=make(%v)\n", v.Name, v.MTye)
			c.buff.WriteString(tvalue)
			prefix := strconv.Quote(v.Name + ".")
			c.buff.WriteString("	for k,v:=range data{\n")

			value := fmt.Sprintf("		if strings.HasPrefix(k, this.prefix + %v){\n", prefix)
			c.buff.WriteString(value)
			value2 := fmt.Sprintf("			d := strings.TrimLeft(k, this.prefix + %v)\n", prefix)
			c.buff.WriteString(value2)
			//key
			have := true
			switch v.Ktype {
			case "string":
				dvalue := fmt.Sprintf("			dv:=d\n")
				c.buff.WriteString(dvalue)
			case "uint64":
				dvalue := fmt.Sprintf("			dv, _:=strconv.ParseUint(d,10,64)\n") //strconv.ParseFloat() ParseUint(d,10,64)
				c.buff.WriteString(dvalue)
			case "uint32":
				dvalue := fmt.Sprintf("			dd, _:=strconv.ParseUint(d,10,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv:=uint32(dd)\n")
				c.buff.WriteString(nvalue)
			case "int32":
				dvalue := fmt.Sprintf("			dd, _:=strconv.Atoi(d)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv:=int32(dd)\n")
				c.buff.WriteString(nvalue)
			case "int64":
				dvalue := fmt.Sprintf("			dv, _:=strconv.ParseInt(d,10,64)\n")
				c.buff.WriteString(dvalue)
			case "float64":
				dvalue := fmt.Sprintf("			dv, _:=strconv.ParseFloat(d,64)\n")
				c.buff.WriteString(dvalue)
			case "float32":
				dvalue := fmt.Sprintf("			dd, _:=strconv.ParseFloat(d,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv:=float32(dd)\n")
				c.buff.WriteString(nvalue)
			case "bool":
				dvalue := fmt.Sprintf("			dv, _:=strconv.ParseBool(d)\n")
				c.buff.WriteString(dvalue)
			default:
				have = false
			}

			have2 := true
			//value
			switch v.Vtype {
			case "string":
				dvalue := fmt.Sprintf("			dv2:=v\n")
				c.buff.WriteString(dvalue)
			case "uint64":
				dvalue := fmt.Sprintf("			dv2, _:=strconv.ParseUint(v,10,64)\n") //strconv.ParseFloat() ParseUint(d,10,64)
				c.buff.WriteString(dvalue)
			case "uint32":
				dvalue := fmt.Sprintf("			dd2, _:=strconv.ParseUint(v,10,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv2:=uint32(dd2)\n")
				c.buff.WriteString(nvalue)
			case "int32":
				dvalue := fmt.Sprintf("			dd2, _:=strconv.Atoi(v)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv2:=int32(dd2)\n")
				c.buff.WriteString(nvalue)
			case "int64":
				dvalue := fmt.Sprintf("			dv2, _:=strconv.ParseInt(v,10,64)\n")
				c.buff.WriteString(dvalue)
			case "float64":
				dvalue := fmt.Sprintf("			dv2, _:=strconv.ParseFloat(v,64)\n")
				c.buff.WriteString(dvalue)
			case "float32":
				dvalue := fmt.Sprintf("			dd2, _:=strconv.ParseFloat(v,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv2:=float32(dd2)\n")
				c.buff.WriteString(nvalue)
			case "bool":
				dvalue := fmt.Sprintf("			dv2, _:=strconv.ParseBool(v)\n")
				c.buff.WriteString(dvalue)
			default:
				have2 = false
			}
			if have && have2 {
				value3 := fmt.Sprintf("			this.%v[dv]= dv2\n", v.Name)
				c.buff.WriteString(value3)

			}

			c.buff.WriteString("		}\n")
			c.buff.WriteString("	}\n")
		}
	}
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateInitDataFunc() {
	head := fmt.Sprintf("func (this *%v)InitData() {\n", c.name)
	c.buff.WriteString(head)
	c.buff.WriteString("	if nil == this.update{\n")
	c.buff.WriteString("		return\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("	data:= this.update.GetAllData()\n")
	for _, v := range c.params {
		if nil != v && 1 == v.TType {
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
	for _, v := range c.params {
		if 2 == v.TType {
			stype := strings.TrimLeft(v.Type, "*")
			prefix := strconv.Quote(v.Name + ".")
			value := fmt.Sprintf("	this.%v= common.New%v(this.uid, %v, this.update)\n", v.Name, stype, prefix)
			//	value := fmt.Sprintf("		this.%v.\n", v.Name)
			c.buff.WriteString(value)
			value2 := fmt.Sprintf("	this.%v.InitData(data)\n", v.Name)
			c.buff.WriteString(value2)

		}
	}

	for _, v := range c.params {
		if 3 == v.TType {
			tvalue := fmt.Sprintf("	this.%v=make(%v)\n", v.Name, v.MTye)
			c.buff.WriteString(tvalue)
			prefix := strconv.Quote(v.Name + ".")
			c.buff.WriteString("	for k,v:=range data{\n")

			value := fmt.Sprintf("		if strings.HasPrefix(k, %v){\n", prefix)
			c.buff.WriteString(value)
			value2 := fmt.Sprintf("			d := strings.TrimLeft(k, %v)\n", prefix)
			c.buff.WriteString(value2)
			//key
			have := true
			switch v.Ktype {
			case "string":
				dvalue := fmt.Sprintf("			dv:=d\n")
				c.buff.WriteString(dvalue)
			case "uint64":
				dvalue := fmt.Sprintf("			dv, _:=strconv.ParseUint(d,10,64)\n") //strconv.ParseFloat() ParseUint(d,10,64)
				c.buff.WriteString(dvalue)
			case "uint32":
				dvalue := fmt.Sprintf("			dd, _:=strconv.ParseUint(d,10,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv:=uint32(dd)\n")
				c.buff.WriteString(nvalue)
			case "int32":
				dvalue := fmt.Sprintf("			dd, _:=strconv.Atoi(d)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv:=int32(dd)\n")
				c.buff.WriteString(nvalue)
			case "int64":
				dvalue := fmt.Sprintf("			dv, _:=strconv.ParseInt(d,10,64)\n")
				c.buff.WriteString(dvalue)
			case "float64":
				dvalue := fmt.Sprintf("			dv, _:=strconv.ParseFloat(d,64)\n")
				c.buff.WriteString(dvalue)
			case "float32":
				dvalue := fmt.Sprintf("			dd, _:=strconv.ParseFloat(d,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv:=float32(dd)\n")
				c.buff.WriteString(nvalue)
			case "bool":
				dvalue := fmt.Sprintf("			dv, _:=strconv.ParseBool(d)\n")
				c.buff.WriteString(dvalue)
			default:
				have = false
			}

			have2 := true
			//value
			switch v.Vtype {
			case "string":
				dvalue := fmt.Sprintf("			dv2:=v\n")
				c.buff.WriteString(dvalue)
			case "uint64":
				dvalue := fmt.Sprintf("			dv2, _:=strconv.ParseUint(v,10,64)\n") //strconv.ParseFloat() ParseUint(d,10,64)
				c.buff.WriteString(dvalue)
			case "uint32":
				dvalue := fmt.Sprintf("			dd2, _:=strconv.ParseUint(v,10,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv2:=uint32(dd2)\n")
				c.buff.WriteString(nvalue)
			case "int32":
				dvalue := fmt.Sprintf("			dd2, _:=strconv.Atoi(v)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv2:=int32(dd2)\n")
				c.buff.WriteString(nvalue)
			case "int64":
				dvalue := fmt.Sprintf("			dv2, _:=strconv.ParseInt(v,10,64)\n")
				c.buff.WriteString(dvalue)
			case "float64":
				dvalue := fmt.Sprintf("			dv2, _:=strconv.ParseFloat(v,64)\n")
				c.buff.WriteString(dvalue)
			case "float32":
				dvalue := fmt.Sprintf("			dd2, _:=strconv.ParseFloat(v,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("			dv2:=float32(dd2)\n")
				c.buff.WriteString(nvalue)
			case "bool":
				dvalue := fmt.Sprintf("			dv2, _:=strconv.ParseBool(v)\n")
				c.buff.WriteString(dvalue)
			default:
				have2 = false
			}
			if have && have2 {
				value3 := fmt.Sprintf("			this.%v[dv]= dv2\n", v.Name)
				c.buff.WriteString(value3)

			}

			c.buff.WriteString("		}\n")
			c.buff.WriteString("	}\n")
		}
	}

	for _, v := range c.params {
		if 4 == v.TType {
			tvalue := fmt.Sprintf("	this.%v=make(%v)\n", v.Name, v.MTye)
			c.buff.WriteString(tvalue)
			prefix := strconv.Quote(v.Name + ".")
			c.buff.WriteString("	for k,v:=range data{\n")

			value := fmt.Sprintf("		if strings.HasPrefix(k, %v){\n", prefix)
			c.buff.WriteString(value)
			value2 := fmt.Sprintf("			ad := strings.TrimLeft(k, %v)\n", prefix)
			c.buff.WriteString(value2)
			dian := strconv.Quote(".")
			ks := fmt.Sprintf("		dl := strings.Split(ad, %v)\n", dian)
			c.buff.WriteString(ks)
			c.buff.WriteString("		if len(dl)<=1 {\n")
			c.buff.WriteString("			continue\n")
			c.buff.WriteString("		}\n")
			c.buff.WriteString("		d := dl[0]\n")
			//key
			have := true
			switch v.Ktype {
			case "string":
				dvalue := fmt.Sprintf("		dv:=d\n")
				c.buff.WriteString(dvalue)
			case "uint64":
				dvalue := fmt.Sprintf("		dv, _:=strconv.ParseUint(d,10,64)\n") //strconv.ParseFloat() ParseUint(d,10,64)
				c.buff.WriteString(dvalue)
			case "uint32":
				dvalue := fmt.Sprintf("		dd, _:=strconv.ParseUint(d,10,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("		dv:=uint32(dd)\n")
				c.buff.WriteString(nvalue)
			case "int32":
				dvalue := fmt.Sprintf("		dd, _:=strconv.Atoi(d)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("		dv:=int32(dd)\n")
				c.buff.WriteString(nvalue)
			case "int64":
				dvalue := fmt.Sprintf("		dv, _:=strconv.ParseInt(d,10,64)\n")
				c.buff.WriteString(dvalue)
			case "float64":
				dvalue := fmt.Sprintf("		dv, _:=strconv.ParseFloat(d,64)\n")
				c.buff.WriteString(dvalue)
			case "float32":
				dvalue := fmt.Sprintf("		dd, _:=strconv.ParseFloat(d,64)\n")
				c.buff.WriteString(dvalue)
				nvalue := fmt.Sprintf("		dv:=float32(dd)\n")
				c.buff.WriteString(nvalue)
			case "bool":
				dvalue := fmt.Sprintf("		dv, _:=strconv.ParseBool(d)\n")
				c.buff.WriteString(dvalue)
			default:
				have = false
			}

			if have {
				ifstr := fmt.Sprintf("		if s,ok := this.%v[dv]; ok {\n", v.Name)
				c.buff.WriteString(ifstr)
				c.buff.WriteString("			s.InitDataSingle(dl[1:], v)\n")
				c.buff.WriteString("		}else{\n")

				stype := strings.TrimLeft(v.Vtype, "*common.")
				prefix := strconv.Quote(v.Name + ".")
				value := fmt.Sprintf("	s := common.New%v(this.uid, %v+d, this.update)\n", stype, prefix)
				//	value := fmt.Sprintf("		this.%v.\n", v.Name)
				c.buff.WriteString(value)

				c.buff.WriteString("			s.InitDataSingle(dl[1:], v)\n")
				c.buff.WriteString("		}\n")

			}

			c.buff.WriteString("		}\n")
			c.buff.WriteString("	}\n")
		}
	}

	c.buff.WriteString("}\n\n")
}
*/
func (c *Class) CreateInitDataParamFunc() {
	head := fmt.Sprintf("func (this *%v)InitDataParam(ks []string, d string) {\n", c.name)
	c.buff.WriteString(head)
	c.buff.WriteString("	if nil == this.update{\n")
	c.buff.WriteString("		return\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("		if len(ks) <= 0 {\n")
	c.buff.WriteString("			return\n")
	c.buff.WriteString("		}\n")
	c.buff.WriteString("		tkey := ks[0]\n")
	c.buff.WriteString("		switch tkey {\n")
	for _, v := range c.params {
		if nil != v {
			if 1 == v.TType {
				namestr := strconv.Quote(v.Name)
				key := fmt.Sprintf("		case %v:\n", namestr)
				c.buff.WriteString(key)
				have := CheckValueType(v.Type, c)
				if have {
					value := fmt.Sprintf("			this.%v= dv\n", v.Name)
					c.buff.WriteString(value)
				}

			}
			if 2 == v.TType {
				ifcheck := fmt.Sprintf("			if nil == this.%v {\n", v.Name)
				c.buff.WriteString(ifcheck)
				stype := strings.TrimLeft(v.Type, "*")
				prefix := strconv.Quote(v.Name + ".")
				value := fmt.Sprintf("				this.%v= common.New%v(this.uid, %v, this.update)\n", v.Name, stype, prefix)
				c.buff.WriteString(value)
				c.buff.WriteString("			}\n")
				value2 := fmt.Sprintf("			this.%v.InitDataParam(ks[1:],d)\n", v.Name)
				c.buff.WriteString(value2)
			}
			if 3 == v.TType {
				ifcheck := fmt.Sprintf("			if nil == this.%v {\n", v.Name)
				c.buff.WriteString(ifcheck)
				tvalue := fmt.Sprintf("				this.%v=make(%v)\n", v.Name, v.MTye)
				c.buff.WriteString(tvalue)
				c.buff.WriteString("			}\n")
				c.buff.WriteString("			if len(ks) == 2 {\n")
				c.buff.WriteString("				d1 := ks[1]\n")
				have := CheckValueType2(v.Ktype, c)
				have2 := CheckValueType(v.Vtype, c)
				if have && have2 {
					value := fmt.Sprintf("				this.%v[dv1]= dv\n", v.Name)
					c.buff.WriteString(value)
				}
				c.buff.WriteString("			}\n")
			}
			if 4 == v.TType {
				ifcheck := fmt.Sprintf("			if nil == this.%v {\n", v.Name)
				c.buff.WriteString(ifcheck)
				tvalue := fmt.Sprintf("				this.%v=make(%v)\n", v.Name, v.MTye)
				c.buff.WriteString(tvalue)
				c.buff.WriteString("			}\n")
				c.buff.WriteString("			if len(ks) > 2 {\n")
				c.buff.WriteString("				d1 := ks[1]\n")
				have := CheckValueType2(v.Ktype, c)
				if have {
					ifcheck := fmt.Sprintf("			ts,ok := this.%v[dv1]\n", v.Name)
					c.buff.WriteString(ifcheck)
					c.buff.WriteString("			if !ok || nil == ts {\n")
					stype := strings.TrimLeft(v.Type, "*")
					prefix := strconv.Quote(v.Name + ".")
					dian := strconv.Quote(".")
					value := fmt.Sprintf("				ts = common.New%v(this.uid, %v+dl+%v, this.update)\n", stype, prefix, dian)
					c.buff.WriteString(value)
					c.buff.WriteString("			}\n")
					value2 := fmt.Sprintf("				this.%v[dv1]= ts\n", v.Name)
					c.buff.WriteString(value2)
					c.buff.WriteString("			ts.InitDataParam(ks[2:],d)\n")
				}
				c.buff.WriteString("			}\n")
			}
		}
	}
	c.buff.WriteString("		}\n")
	c.buff.WriteString("}\n\n")

}

func (c *Class) CreateInitDataFuncNew() {
	head := fmt.Sprintf("func (this *%v)InitData() {\n", c.name)
	c.buff.WriteString(head)
	c.buff.WriteString("	if nil == this.update{\n")
	c.buff.WriteString("		return\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("	data:= this.update.GetAllData()\n")
	c.buff.WriteString("	for k,d := range data {\n")
	dian := strconv.Quote(".")
	ks := fmt.Sprintf("			ks := strings.Split(k, %v)\n", dian)
	c.buff.WriteString(ks)
	c.buff.WriteString("		if len(ks) <= 0 {\n")
	c.buff.WriteString("			continue\n")
	c.buff.WriteString("		}\n")
	c.buff.WriteString("		tkey := ks[0]\n")
	c.buff.WriteString("		switch tkey {\n")
	for _, v := range c.params {
		if nil != v {
			if 1 == v.TType {
				namestr := strconv.Quote(v.Name)
				key := fmt.Sprintf("		case %v:\n", namestr)
				c.buff.WriteString(key)
				have := CheckValueType(v.Type, c)
				if have {
					value := fmt.Sprintf("			this.%v= dv\n", v.Name)
					c.buff.WriteString(value)
				}

			}
			if 2 == v.TType {
				namestr := strconv.Quote(v.Name)
				key := fmt.Sprintf("		case %v:\n", namestr)
				c.buff.WriteString(key)
				ifcheck := fmt.Sprintf("			if nil == this.%v {\n", v.Name)
				c.buff.WriteString(ifcheck)
				stype := strings.TrimLeft(v.Type, "*")
				prefix := strconv.Quote(v.Name + ".")
				value := fmt.Sprintf("				this.%v= common.New%v(this.uid, %v, this.update)\n", v.Name, stype, prefix)
				c.buff.WriteString(value)
				c.buff.WriteString("			}\n")
				value2 := fmt.Sprintf("			this.%v.InitDataParam(ks[1:],d)\n", v.Name)
				c.buff.WriteString(value2)
			}
			if 3 == v.TType {
				namestr := strconv.Quote(v.Name)
				key := fmt.Sprintf("		case %v:\n", namestr)
				c.buff.WriteString(key)
				ifcheck := fmt.Sprintf("			if nil == this.%v {\n", v.Name)
				c.buff.WriteString(ifcheck)
				tvalue := fmt.Sprintf("				this.%v=make(%v)\n", v.Name, v.MTye)
				c.buff.WriteString(tvalue)
				c.buff.WriteString("			}\n")
				c.buff.WriteString("			if len(ks) == 2 {\n")
				c.buff.WriteString("				d1 := ks[1]\n")
				have := CheckValueType2(v.Ktype, c)
				have2 := CheckValueType(v.Vtype, c)
				if have && have2 {
					value := fmt.Sprintf("				this.%v[dv1]= dv\n", v.Name)
					c.buff.WriteString(value)
				}
				c.buff.WriteString("			}\n")
			}
			if 4 == v.TType {
				namestr := strconv.Quote(v.Name)
				key := fmt.Sprintf("		case %v:\n", namestr)
				c.buff.WriteString(key)
				ifcheck := fmt.Sprintf("			if nil == this.%v {\n", v.Name)
				c.buff.WriteString(ifcheck)
				tvalue := fmt.Sprintf("				this.%v=make(%v)\n", v.Name, v.MTye)
				c.buff.WriteString(tvalue)
				c.buff.WriteString("			}\n")
				c.buff.WriteString("			if len(ks) > 2 {\n")
				c.buff.WriteString("				d1 := ks[1]\n")
				have := CheckValueType2(v.Ktype, c)
				if have {
					ifcheck := fmt.Sprintf("			ts,ok := this.%v[dv1]\n", v.Name)
					c.buff.WriteString(ifcheck)
					c.buff.WriteString("			if !ok || nil == ts {\n")
					stype := strings.TrimLeft(v.Vtype, "*common.")
					prefix := strconv.Quote(v.Name + ".")
					dian := strconv.Quote(".")
					value := fmt.Sprintf("				ts = common.New%v(this.uid, %v + d1 + %v, this.update)\n", stype, prefix, dian)
					c.buff.WriteString(value)
					c.buff.WriteString("			}\n")
					value2 := fmt.Sprintf("				this.%v[dv1]= ts\n", v.Name)
					c.buff.WriteString(value2)
					c.buff.WriteString("			ts.InitDataParam(ks[2:],d)\n")
				}
				c.buff.WriteString("			}\n")
			}
		}
	}
	c.buff.WriteString("		}\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("}\n\n")

}

func (c *Class) CreateSetFunc(name, ctype string, ttype uint32) {
	//先判断是否是基础类型

	head := fmt.Sprintf("func(this *%v) Set%v(value %v){\n", c.name, name, ctype)
	c.buff.WriteString(head)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}
	if 1 == ttype { //需要判断是否是基础类型
		//	newname := fmt.Sprintf("this.prefix%v", name)
		namestr := strconv.Quote(name)
		add := fmt.Sprintf("	this.update.AddData(this.prefix+%v, value)\n", namestr)
		c.buff.WriteString(add)
	}
	body := fmt.Sprintf("	this.%v = value\n", name)
	c.buff.WriteString(body)
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateGetFunc(name, ctype string) {
	//先判断是否是基础类型

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

func (c *Class) CreateUpdateFunc() {
	head := fmt.Sprintf("func (this *%v)UpdateData() {\n", c.name)
	c.buff.WriteString(head)
	if c.Lock {
		c.buff.WriteString("	this.mutex.RLock()\n")
		c.buff.WriteString("	defer this.mutex.RUnlock()\n")
	}
	c.buff.WriteString("	if nil != this.update{\n")
	c.buff.WriteString("		this.update.Update()\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateClose() {
	head := fmt.Sprintf("func (this *%v)Close() {\n", c.name)
	c.buff.WriteString(head)
	c.buff.WriteString("	this.UpdateData()\n")
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateDestroy() {
	head2 := fmt.Sprintf("func(this *%v) Destroy(){\n", c.name)
	c.buff.WriteString(head2)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}
	for _, v := range c.params {
		if nil != v && 1 == v.TType {
			namestr := strconv.Quote(v.Name)
			add2 := fmt.Sprintf("	this.update.DelData(this.prefix + %v)\n", namestr)
			c.buff.WriteString(add2)
		}

		if 2 == v.TType {
			key := fmt.Sprintf(" if nil != this.%v {\n", v.Name)
			c.buff.WriteString(key)
			value2 := fmt.Sprintf("		this.%v.Destroy()\n", v.Name)
			c.buff.WriteString(value2)
			c.buff.WriteString("	}\n")
		}
		if 3 == v.TType {
			roll := fmt.Sprintf("	for k,_:=range this.%v{\n", v.Name)
			c.buff.WriteString(roll)
			c.buff.WriteString("		key := this.prefix + fmt.Sprint(k)\n")
			c.buff.WriteString("		this.update.DelData(key)\n")
			c.buff.WriteString("	}\n")

		}
		if 4 == v.TType {
			roll := fmt.Sprintf("	for _,v:=range this.%v{\n", v.Name)
			c.buff.WriteString(roll)
			c.buff.WriteString("		if nil != v {\n")
			c.buff.WriteString("			v.Destroy()\n")
			c.buff.WriteString("		}\n")
			c.buff.WriteString("	}\n")
		}

	}
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateMapFunc(name, ktype, vtype string) {
	//add func
	head := fmt.Sprintf("func(this *%v) Add%vData(key %v,value %v){\n", c.name, name, ktype, vtype)
	c.buff.WriteString(head)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}

	prefix := fmt.Sprintf("%v.", name)
	prefix = strconv.Quote(prefix)
	c.buff.WriteString(" 	keystr := fmt.Sprint(key)\n")
	add := fmt.Sprintf("	this.update.AddData(this.prefix + %v + keystr, value)\n", prefix)
	c.buff.WriteString(add)

	body := fmt.Sprintf("	this.%v[key] = value\n", name)
	c.buff.WriteString(body)
	c.buff.WriteString("}\n\n")
	//def func
	head2 := fmt.Sprintf("func(this *%v) Del%vData(key %v){\n", c.name, name, ktype)
	c.buff.WriteString(head2)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}

	c.buff.WriteString(" 	keystr := fmt.Sprint(key)\n")
	add2 := fmt.Sprintf("	this.update.DelData(this.prefix + %v + keystr)\n", prefix)
	c.buff.WriteString(add2)
	body1 := fmt.Sprintf("	if _,ok:=this.%v[key]; ok{\n", name)
	c.buff.WriteString(body1)
	body2 := fmt.Sprintf("		delete(this.%v, key)\n", name)
	c.buff.WriteString(body2)
	c.buff.WriteString("	}\n")
	c.buff.WriteString("}\n\n")

	//get by key
	head3 := fmt.Sprintf("func(this *%v) Get%vDataByKey(key %v) (value %v) {\n", c.name, name, ktype, vtype)
	c.buff.WriteString(head3)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}

	body3 := fmt.Sprintf("	if v,ok:=this.%v[key]; ok{\n", name)
	c.buff.WriteString(body3)
	c.buff.WriteString("		value = v\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("	return\n")
	c.buff.WriteString("}\n\n")
	//get all
	head4 := fmt.Sprintf("func(this *%v) Get%vDataAll() (d map[%v]%v){\n", c.name, name, ktype, vtype)
	c.buff.WriteString(head4)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}
	body4 := fmt.Sprintf("	d = make(map[%v]%v)\n", ktype, vtype)
	c.buff.WriteString(body4)
	body5 := fmt.Sprintf("	for k, v := range this.%v{\n", name)
	c.buff.WriteString(body5)
	c.buff.WriteString("		d[k] = v\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("	return\n")
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateMapStructFunc(name, ktype, vtype string) {
	//add func
	head := fmt.Sprintf("func(this *%v) Create%vNewData(key %v)(value %v){\n", c.name, name, ktype, vtype)
	c.buff.WriteString(head)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}

	stype := strings.TrimLeft(vtype, "*common.")
	prefix := strconv.Quote("this.prefix" + name + ".")
	value := fmt.Sprintf("	newdata := common.New%v(this.uid, %v, this.update)\n", stype, prefix)
	c.buff.WriteString(value)
	body := fmt.Sprintf("	this.%v[key] = newdata\n", name)
	c.buff.WriteString(body)
	c.buff.WriteString("	value = newdata\n")
	c.buff.WriteString("	return\n")
	c.buff.WriteString("}\n\n")
	//def func
	head2 := fmt.Sprintf("func(this *%v) Del%vData(key %v){\n", c.name, name, ktype)
	c.buff.WriteString(head2)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}

	//	c.buff.WriteString(" 	keystr := fmt.Sprint(key)\n")
	//	add2 := fmt.Sprintf("	this.update.DelData(this.prefix + %v + keystr)\n", prefix)
	//	c.buff.WriteString(add2)
	body1 := fmt.Sprintf("	if v,ok:=this.%v[key]; ok{\n", name)
	c.buff.WriteString(body1)
	body2 := fmt.Sprintf("		delete(this.%v, key)\n", name)
	c.buff.WriteString(body2)
	c.buff.WriteString("		v.Destroy()\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("}\n\n")

	//get by key
	head3 := fmt.Sprintf("func(this *%v) Get%vDataByKey(key %v) (value %v) {\n", c.name, name, ktype, vtype)
	c.buff.WriteString(head3)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}

	body3 := fmt.Sprintf("	if v,ok:=this.%v[key]; ok{\n", name)
	c.buff.WriteString(body3)
	c.buff.WriteString("		value = v\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("	return\n")
	c.buff.WriteString("}\n\n")
	//get all
	head4 := fmt.Sprintf("func(this *%v) Get%vDataAll() (d map[%v]%v){\n", c.name, name, ktype, vtype)
	c.buff.WriteString(head4)
	if c.Lock {
		c.buff.WriteString("	this.mutex.Lock()\n")
		c.buff.WriteString("	defer this.mutex.Unlock()\n")
	}
	body4 := fmt.Sprintf("	d = make(map[%v]%v)\n", ktype, vtype)
	c.buff.WriteString(body4)
	body5 := fmt.Sprintf("	for k, v := range this.%v{\n", name)
	c.buff.WriteString(body5)
	c.buff.WriteString("		d[k] = v\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("	return\n")
	c.buff.WriteString("}\n\n")
}

//manager
func (c *Class) ManagerInit() {
	if c.IsUpdate {
		c.ManagerInitPackage()
		c.ManagerInitImport()
		c.ManagerInitParam()
		c.ManagerInitNew()
		c.ManagerCreateAddFunc()
		c.ManagerCreateDelFunc()
		c.ManagerCreateGetFunc()
		c.ManagerCreateFunc()
		c.ManagerCreateClose()
		c.ManagerCreateUpdate()
	}
}

func (c *Class) ManagerInitPackage() {
	str := "package " + strFirstToLower(c.name)
	c.managerbuff.WriteString(str)
	c.managerbuff.WriteString("\n\n")
}

func (c *Class) ManagerInitImport() {
	c.managerbuff.WriteString("import (\n")
	//添加包含文件
	pak := strconv.Quote("sync")
	c.managerbuff.WriteString("	" + pak + "\n")

	pak2 := strconv.Quote("time")
	c.managerbuff.WriteString("	" + pak2 + "\n")
	/*
		pak3 := strconv.Quote("strconv")
		c.managerbuff.WriteString("	" + pak3 + "\n")
		pak4 := strconv.Quote("fmt")
		c.managerbuff.WriteString("	" + pak4 + "\n")*/
	c.managerbuff.WriteString(")\n\n")
}

func (c *Class) ManagerInitParam() {
	str := fmt.Sprintf("var Manager *%v\n\n", c.managername)
	str2 := fmt.Sprintf("type %v struct {\n", c.managername)
	c.managerbuff.WriteString(str)
	c.managerbuff.WriteString(str2)
	c.managerbuff.WriteString("	closed	bool\n")
	c.managerbuff.WriteString("	mutex sync.RWMutex\n")
	str3 := fmt.Sprintf("	data map[uint64]*%v\n", c.name)
	c.managerbuff.WriteString(str3)
	c.managerbuff.WriteString("}\n\n")
}

func (c *Class) ManagerInitNew() {
	head := fmt.Sprintf("func init(){\n")
	c.managerbuff.WriteString(head)
	body := fmt.Sprintf("	Manager := new(%v)\n", c.managername)
	c.managerbuff.WriteString(body)
	str := fmt.Sprintf("	Manager.data= make(map[uint64]*%v)\n", c.name)
	c.managerbuff.WriteString(str)
	c.managerbuff.WriteString("	go Manager.Update()\n")
	c.managerbuff.WriteString("}\n\n")
}

func (c *Class) ManagerCreateFunc() {
	head := fmt.Sprintf("func AddData(data *%v){\n", c.name)
	c.managerbuff.WriteString(head)
	c.managerbuff.WriteString("	if nil== Manager || nil == data{\n")
	c.managerbuff.WriteString("		return\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("	Manager.AddData(data)\n")
	c.managerbuff.WriteString("}\n\n")

	head = fmt.Sprintf("func DelData(uid uint64)bool{\n")
	c.managerbuff.WriteString(head)
	c.managerbuff.WriteString("	if nil== Manager {\n")
	c.managerbuff.WriteString("		return false\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("	return Manager.DelData(uid)\n")
	c.managerbuff.WriteString("}\n\n")

	head = fmt.Sprintf("func GetData(uid uint64) *%v {\n", c.name)
	c.managerbuff.WriteString(head)
	c.managerbuff.WriteString("	if nil== Manager {\n")
	c.managerbuff.WriteString("		return nil\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("	return Manager.GetData(uid)\n")
	c.managerbuff.WriteString("}\n\n")
}

func (c *Class) ManagerCreateAddFunc() {
	head := fmt.Sprintf("func (this *%v)AddData(d *%v){\n", c.managername, c.name)
	c.managerbuff.WriteString(head)
	c.managerbuff.WriteString("	if  nil == d{\n")
	c.managerbuff.WriteString("		return\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("	this.mutex.Lock()\n")
	c.managerbuff.WriteString("	defer this.mutex.Unlock()\n")
	c.managerbuff.WriteString("	this.data[d.uid]= d\n")
	c.managerbuff.WriteString("}\n\n")
}

func (c *Class) ManagerCreateDelFunc() {
	head := fmt.Sprintf("func (this *%v)DelData(uid uint64)bool{\n", c.managername)
	c.managerbuff.WriteString(head)
	c.managerbuff.WriteString("	this.mutex.Lock()\n")
	c.managerbuff.WriteString("	defer this.mutex.Unlock()\n")
	c.managerbuff.WriteString("	if  v,ok := this.data[uid];ok{\n")
	c.managerbuff.WriteString("		if nil !=v{\n")
	c.managerbuff.WriteString("			v.Close()\n")
	c.managerbuff.WriteString("		}\n")
	c.managerbuff.WriteString("		delete(this.data, uid)\n")
	c.managerbuff.WriteString("		return true\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("	return false\n")
	c.managerbuff.WriteString("}\n\n")
}

func (c *Class) ManagerCreateGetFunc() {
	head := fmt.Sprintf("func (this *%v)GetData(uid uint64)*%v{\n", c.managername, c.name)
	c.managerbuff.WriteString(head)
	c.managerbuff.WriteString("	this.mutex.RLock()\n")
	c.managerbuff.WriteString("	defer this.mutex.RUnlock()\n")
	c.managerbuff.WriteString("	if  v,ok := this.data[uid];ok{\n")
	c.managerbuff.WriteString("		return v\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("	return nil\n")
	c.managerbuff.WriteString("}\n\n")
}

func (c *Class) ManagerCreateClose() {
	head := fmt.Sprintf("func (this *%v)Close(){\n", c.managername)
	c.managerbuff.WriteString(head)
	c.managerbuff.WriteString("	this.mutex.Lock()\n")
	c.managerbuff.WriteString("	defer this.mutex.Unlock()\n")
	c.managerbuff.WriteString("	this.closed= true\n")
	c.managerbuff.WriteString("	for k, v := range this.data{\n")
	c.managerbuff.WriteString("		if nil !=v{\n")
	c.managerbuff.WriteString("			v.Close()\n")
	c.managerbuff.WriteString("		}\n")
	c.managerbuff.WriteString("		delete(this.data, k)\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("}\n\n")
}

func (c *Class) ManagerCreateUpdate() {
	head := fmt.Sprintf("func (this *%v)Update(){\n", c.managername)
	c.managerbuff.WriteString(head)
	c.managerbuff.WriteString("	t := time.Tick(500 * time.Millisecond)\n")
	c.managerbuff.WriteString("	for _ = range t {\n")
	c.managerbuff.WriteString("		this.mutex.RLock()\n")
	c.managerbuff.WriteString("		if true == this.closed{\n")
	c.managerbuff.WriteString("			this.mutex.RUnlock()\n")
	c.managerbuff.WriteString("			break\n")
	c.managerbuff.WriteString("		}\n")
	c.managerbuff.WriteString("		for _, v := range this.data{\n")
	c.managerbuff.WriteString("			if nil !=v{\n")
	c.managerbuff.WriteString("				v.UpdateData()\n")
	c.managerbuff.WriteString("			}\n")
	c.managerbuff.WriteString("		}\n")
	c.managerbuff.WriteString("		this.mutex.RUnlock()\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("}\n\n")
}
