package common

import (
	"github.com/mikeqiao/newworld/data"
	"strconv"
	"fmt"
)

type Test struct {
	uid	uint64
	cdd	int32
	dasda	int32
	myData	map[uint64]uint32
	name	string
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

func(this *Test) GetmyDataDataByKey(key uint64) (value uint32) {
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

func (this *Test)InitDataParam(ks []string, d string) {
	if nil == this.update{
		return
	}
		if len(ks) <= 0 {
			return
		}
		tkey := ks[0]
		switch tkey {
		case "uid":
		dv, _:=strconv.ParseUint(d,10,64)
			this.uid= dv
		case "cdd":
		dd, _:=strconv.Atoi(d)
		dv:=int32(dd)
			this.cdd= dv
		case "dasda":
		dd, _:=strconv.Atoi(d)
		dv:=int32(dd)
			this.dasda= dv
			if nil == this.myData {
				this.myData=make(map[uint64]uint32)
			}
			if len(ks) == 2 {
				d1 := ks[1]
		dv1, _:=strconv.ParseUint(d1,10,64)
		dd, _:=strconv.ParseUint(d,10,64)
		dv:=uint32(dd)
				this.myData[dv1]= dv
			}
		case "name":
		dv:=d
			this.name= dv
		}
}

func(this *Test) Destroy(){
	this.update.DelData(this.prefix + "uid")
	this.update.DelData(this.prefix + "cdd")
	this.update.DelData(this.prefix + "dasda")
	for k,_:=range this.myData{
		key := this.prefix + fmt.Sprint(k)
		this.update.DelData(key)
	}
	this.update.DelData(this.prefix + "name")
}

