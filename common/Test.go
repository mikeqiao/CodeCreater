package common

import (
	"github.com/mikeqiao/newworld/data"
	"strconv"
	"strings"
	"fmt"
)

type Test struct {
	myData	map[uint64]uint32
	name	string
	uid	uint64
	cdd	int32
	dasda	int32
	prefix string
	update *data.UpdateMod
}

func NewTest(uid uint64, prefix string, update *data.UpdateMod) *Test{
	d := new(Test)
	d.uid= uid
	d.prefix= prefix
	if nil != update{
		d.update= update
	}else{
		d.update= new(data.UpdateMod)
		table := "Test_" + fmt.Sprint(d.uid)
		d.update.Init(table)
	}
	return d
}

func(this *Test) AddmyDataData(key uint64,value uint32){
 	keystr := fmt.Sprint(key)
	this.update.AddData(this.prefix + "myData." + keystr, value)
	this.myData[key] = value
}

func(this *Test) DelmyDataData(key uint64){
 	keystr := fmt.Sprint(key)
	this.update.DelData(this.prefix + "myData." + keystr)
	if _,ok:=this.myData[key]; ok{
		delete(this.myData, key)
	}
}

func(this *Test) GeTmyDataDataByKey(key uint64) (value uint32) {
	if v,ok:=this.myData[key]; ok{
		value = v
	}
	return
}

func(this *Test) GetmyDataDataAll() (d map[uint64]uint32){
	d = make(map[uint64]uint32)
	for k, v := range this.myData{
		d[k] = v
	}
	return
}

func(this *Test) Setname(value string){
	this.update.AddData(this.prefix+"name", value)
	this.name = value
}

func(this *Test) Getname() string{
	return this.name
}

func(this *Test) Setuid(value uint64){
	this.update.AddData(this.prefix+"uid", value)
	this.uid = value
}

func(this *Test) Getuid() uint64{
	return this.uid
}

func(this *Test) Setcdd(value int32){
	this.update.AddData(this.prefix+"cdd", value)
	this.cdd = value
}

func(this *Test) Getcdd() int32{
	return this.cdd
}

func(this *Test) Setdasda(value int32){
	this.update.AddData(this.prefix+"dasda", value)
	this.dasda = value
}

func(this *Test) Getdasda() int32{
	return this.dasda
}

func (this *Test)InitData(data map[string]string) {
	if nil == data{
		return
	}
	if d,ok:=data[this.prefix+"name"];ok{
		dv:=d
		this.name= dv
	}

	if d,ok:=data[this.prefix+"uid"];ok{
		dv, _:=strconv.ParseUint(d,10,64)
		this.uid= dv
	}

	if d,ok:=data[this.prefix+"cdd"];ok{
		dd, _:=strconv.Atoi(d)
		dv:=int32(dd)
		this.cdd= dv
	}

	if d,ok:=data[this.prefix+"dasda"];ok{
		dd, _:=strconv.Atoi(d)
		dv:=int32(dd)
		this.dasda= dv
	}

	this.myData=make(map[uint64]uint32)
	for k,v:=range data{
		if strings.HasPrefix(k, this.prefix + "myData."){
			d := strings.TrimLeft(k, this.prefix + "myData.")
			dv, _:=strconv.ParseUint(d,10,64)
			dd2, _:=strconv.ParseUint(v,10,64)
			dv2:=uint32(dd2)
			this.myData[dv]= dv2
		}
	}
}

