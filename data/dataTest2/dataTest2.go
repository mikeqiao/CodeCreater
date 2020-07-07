package DataTest2

import (
	"github.com/mikeqiao/newworld/data"
	"strconv"
	"codecreater/data/common"
	"strings"
	"fmt"
)

type DataTest2 struct {
	name3	string
	myTest	*common.Test
	myData	map[uint32]*common.Test
	mySData	map[uint64]uint32
	uid	uint64
	lang	string
	name2	string
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

func(this *DataTest2) CreatemyDataNewData(key uint32)(value *common.Test){
	newdata := common.NewTest(this.uid, "this.prefixmyData.", this.update)
	this.myData[key] = newdata
	value = newdata
	return
}

func(this *DataTest2) DelmyDataData(key uint32){
	if v,ok:=this.myData[key]; ok{
		delete(this.myData, key)
		v.Destroy()
	}
}

func(this *DataTest2) GeTmyDataDataByKey(key uint32) (value *common.Test) {
	if v,ok:=this.myData[key]; ok{
		value = v
	}
	return
}

func(this *DataTest2) GetmyDataDataAll() (d map[uint32]*common.Test){
	d = make(map[uint32]*common.Test)
	for k, v := range this.myData{
		d[k] = v
	}
	return
}

func(this *DataTest2) AddmySDataData(key uint64,value uint32){
 	keystr := fmt.Sprint(key)
	this.update.AddData(this.prefix + "mySData." + keystr, value)
	this.mySData[key] = value
}

func(this *DataTest2) DelmySDataData(key uint64){
 	keystr := fmt.Sprint(key)
	this.update.DelData(this.prefix + "mySData." + keystr)
	if _,ok:=this.mySData[key]; ok{
		delete(this.mySData, key)
	}
}

func(this *DataTest2) GeTmySDataDataByKey(key uint64) (value uint32) {
	if v,ok:=this.mySData[key]; ok{
		value = v
	}
	return
}

func(this *DataTest2) GetmySDataDataAll() (d map[uint64]uint32){
	d = make(map[uint64]uint32)
	for k, v := range this.mySData{
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

func (this *DataTest2)InitData() {
	if nil == this.update{
		return
	}
	data:= this.update.GetAllData()
	for k,d := range data {
			ks := strings.Split(k, ".")
		if len(ks) <= 0 {
			continue
		}
		tkey := ks[0]
		switch tkey {
		case "name3":
		dv:=d
			this.name3= dv
		case "myTest":
			if nil == this.myTest {
				this.myTest= common.NewTest(this.uid, "myTest.", this.update)
			}
			this.myTest.InitDataParam(ks[1:],d)
		case "myData":
			if nil == this.myData {
				this.myData=make(map[uint32]*common.Test)
			}
			if len(ks) > 2 {
				d1 := ks[1]
		dd1, _:=strconv.ParseUint(d1,10,64)
		dv1:=uint32(dd1)
			ts,ok := this.myData[dv1]
			if !ok || nil == ts {
				ts = common.NewTest(this.uid, "myData." + d1 + ".", this.update)
			}
				this.myData[dv1]= ts
			ts.InitDataParam(ks[2:],d)
			}
		case "mySData":
			if nil == this.mySData {
				this.mySData=make(map[uint64]uint32)
			}
			if len(ks) == 2 {
				d1 := ks[1]
		dv1, _:=strconv.ParseUint(d1,10,64)
		dd, _:=strconv.ParseUint(d,10,64)
		dv:=uint32(dd)
				this.mySData[dv1]= dv
			}
		case "uid":
		dv, _:=strconv.ParseUint(d,10,64)
			this.uid= dv
		case "lang":
		dv:=d
			this.lang= dv
		case "name2":
		dv:=d
			this.name2= dv
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

func(this *DataTest2) Destroy(){
	this.update.DelData(this.prefix + "name3")
 if nil != this.myTest {
		this.myTest.Destroy()
	}
	for _,v:=range this.myData{
		if nil != v {
			v.Destroy()
		}
	}
	for k,_:=range this.mySData{
		key := this.prefix + fmt.Sprint(k)
		this.update.DelData(key)
	}
	this.update.DelData(this.prefix + "uid")
	this.update.DelData(this.prefix + "lang")
	this.update.DelData(this.prefix + "name2")
}

