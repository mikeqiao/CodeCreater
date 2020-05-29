package class

import (
	"bytes"
	"fmt"
	"strconv"

	p "github.com/mikeqiao/codecreater/param"

	f "github.com/mikeqiao/codecreater/function"
)

type Class struct {
	name        string
	managername string
	params      []*p.Param
	funcs       []*f.Function
	Lock        bool
	buff        *bytes.Buffer
	managerbuff *bytes.Buffer
}

func NewClass(name string) *Class {
	c := new(Class)
	c.name = name
	c.managername = name + "Manager"
	c.buff = new(bytes.Buffer)
	c.managerbuff = new(bytes.Buffer)
	return c
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
	c.CreateUpdateFunc()
	c.CreateClose()
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
	pak3 := strconv.Quote("strconv")
	c.buff.WriteString("	" + pak3 + "\n")
	pak4 := strconv.Quote("fmt")
	c.buff.WriteString("	" + pak4 + "\n")
	c.buff.WriteString(")\n\n")
}

func (c *Class) InitParam() {
	str := fmt.Sprintf("type %v struct {\n", c.name)
	c.buff.WriteString(str)
	have := false
	for _, v := range c.params {
		if nil != v {
			c.buff.WriteString("	")
			c.buff.WriteString(v.Name)
			c.buff.WriteString("	")
			c.buff.WriteString(v.Type)
			c.buff.WriteString("\n")
			if "uid" == v.Name {
				have = true
			}
		}
	}
	if !have {
		c.buff.WriteString("	uid	uint64\n")
	}
	c.buff.WriteString("	table	string\n")
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
	head := fmt.Sprintf("func New%v(uid uint64) *%v{\n", c.name, c.name)
	c.buff.WriteString(head)

	body := fmt.Sprintf("	data := new(%v)\n", c.name)
	c.buff.WriteString(body)
	c.buff.WriteString("	data.uid= uid\n")
	c.buff.WriteString("	data.changeData= make(map[string]interface{})\n")
	namestr := strconv.Quote(c.name + "_")
	t := fmt.Sprintf("	data.table = %v + fmt.Sprint(data.uid)\n", namestr)
	c.buff.WriteString(t)

	c.buff.WriteString("	return data\n")
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateInitDataFunc() {
	head := fmt.Sprintf("func (this *%v)InitData() {\n", c.name)
	c.buff.WriteString(head)

	c.buff.WriteString("	data, _:=redis.R.Hash_GetAllData(this.table)\n")
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

func (c *Class) CreateUpdateFunc() {
	head := fmt.Sprintf("func (this *%v)UpdateData() {\n", c.name)
	c.buff.WriteString(head)
	if c.Lock {
		c.buff.WriteString("	this.mutex.RLock()\n")
		c.buff.WriteString("	defer this.mutex.RUnlock()\n")
	}
	c.buff.WriteString("	if len(this.changeData)>0{\n")
	c.buff.WriteString("		err:=redis.R.Hash_SetDataMap(this.table, this.changeData)\n")
	c.buff.WriteString("		if nil != err{\n")
	c.buff.WriteString("			return\n")
	c.buff.WriteString("		}\n")
	c.buff.WriteString("		this.changeData= make(map[string]interface{})\n")
	c.buff.WriteString("	}\n")
	c.buff.WriteString("}\n\n")
}

func (c *Class) CreateClose() {
	head := fmt.Sprintf("func (this *%v)Close() {\n", c.name)
	c.buff.WriteString(head)
	c.buff.WriteString("	this.UpdateData()\n")
	c.buff.WriteString("}\n\n")
}

//manager
func (c *Class) ManagerInit() {
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

func (c *Class) ManagerInitPackage() {
	str := "package " + c.name
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
	c.managerbuff.WriteString("			this.mutex.Unlock()\n")
	c.managerbuff.WriteString("			break\n")
	c.managerbuff.WriteString("		}\n")
	c.managerbuff.WriteString("		for _, v := range this.data{\n")
	c.managerbuff.WriteString("			if nil !=v{\n")
	c.managerbuff.WriteString("				v.UpdateData()\n")
	c.managerbuff.WriteString("			}\n")
	c.managerbuff.WriteString("		}\n")
	c.managerbuff.WriteString("		this.mutex.Unlock()\n")
	c.managerbuff.WriteString("	}\n")
	c.managerbuff.WriteString("}\n\n")
}
