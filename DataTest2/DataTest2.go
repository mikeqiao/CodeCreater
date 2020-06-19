package DataTest2

import (
	"github.com/mikeqiao/newworld/data"
	"strconv"
	"github.com/mikeqiao/codecreater/common"
	"strings"
	"fmt"
)

type DataTest2 struct {
	lang	string
	name2	string
	name3	string
	myTest	*common.Test
	myData	map[uint32]string
	uid	uint64
	prefix string
	update *data.UpdateMod
}

func NewDataTest2(uid uint64, prefix string, update *data.UpdateMod) *DataTest2{
	d := new(DataTest2)
	d.uid= uid
	d.prefix= prefix
	if nil != update{
		d.update= update
	}else{
		d.update= new(data.UpdateMod)
		table := "DataTest2_" + fmt.Sprint(d.uid)
		d.update.Init(table)
	}
	return d
}

func(this *DataTest2) Setlang(value string){
	this.update.AddData(this.prefix+"lang", value)
	this.lang = value
}

func(this *DataTest2) Getlang() string{
	return this.lang
}

func(this *DataTest2) Setname2(value string){
	this.update.AddData(this.prefix+"name2", value)
	this.name2 = value
}

func(this *DataTest2) Getname2() string{
	return this.name2
}

func(this *DataTest2) Setname3(value string){
	this.update.AddData(this.prefix+"name3", value)
	this.name3 = value
}

func(this *DataTest2) Getname3() string{
	return this.name3
}

func(this *DataTest2) SetmyTest(value *common.Test){
	this.myTest = value
}

func(this *DataTest2) GetmyTest() *common.Test{
	return this.myTest
}

func(this *DataTest2) AddmyDataData(key uint32,value string){
 	keystr := fmt.Sprint(key)
	this.update.AddData(this.prefix + "myData." + keystr, value)
	this.myData[key] = value
}

func(this *DataTest2) DelmyDataData(key uint32){
 	keystr := fmt.Sprint(key)
	this.update.DelData(this.prefix + "myData." + keystr)
	if _,ok:=this.myData[key]; ok{
		delete(this.myData, key)
	}
}

func(this *DataTest2) GeTmyDataDataByKey(key uint32) (value string) {
	if v,ok:=this.myData[key]; ok{
		value = v
	}
	return
}

func(this *DataTest2) GetmyDataDataAll() (d map[uint32]string){
	d = make(map[uint32]string)
	for k, v := range this.myData{
		d[k] = v
	}
	return
}

func(this *DataTest2) Setuid(value uint64){
	this.update.AddData(this.prefix+"uid", value)
	this.uid = value
}

func(this *DataTest2) Getuid() uint64{
	return this.uid
}

func (this *DataTest2)InitData() {
	if nil == this.update{
		return
	}
	data:= this.update.GetAllData()
	if d,ok:=data["lang"];ok{
		dv:=d
		this.lang= dv
	}

	if d,ok:=data["name2"];ok{
		dv:=d
		this.name2= dv
	}

	if d,ok:=data["name3"];ok{
		dv:=d
		this.name3= dv
	}

	if d,ok:=data["uid"];ok{
		dv, _:=strconv.ParseUint(d,10,64)
		this.uid= dv
	}

	this.myTest= common.NewTest(this.uid, "myTest.", this.update)
	this.myTest.InitData(data)
	this.myData=make(map[uint32]string)
	for k,v:=range data{
		if strings.HasPrefix(k, "myData."){
			d := strings.TrimLeft(k, "myData.")
			dd, _:=strconv.ParseUint(d,10,64)
			dv:=uint32(dd)
			dv2:=v
			this.myData[dv]= dv2
		}
	}
}

func (this *DataTest2)UpdateData() {
	if nil != this.update{
		this.update.Update()
	}
}

func (this *DataTest2)Close() {
	this.UpdateData()
}

