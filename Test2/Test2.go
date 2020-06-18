package Test2

import (
	"sync"
	"github.com/mikeqiao/Db/redis"
	"strconv"
	"fmt"
)

type Test2 struct {
	Lang	string
	Name2	string
	Name3	string
	Uid	uint64
	uid	uint64
	table	string
	changeData map[string]interface{}
}

func NewTest2(uid uint64) *Test2{
	data := new(Test2)
	data.uid= uid
	data.changeData= make(map[string]interface{})
	data.table = "Test2_" + fmt.Sprint(data.uid)
	return data
}

func (this *Test2)InitData() {
	data, _:=redis.R.Hash_GetAllData(this.table)
	if d,ok:=data["Lang"];ok{
		dv:=d
		this.Lang= dv
	}

	if d,ok:=data["Name2"];ok{
		dv:=d
		this.Name2= dv
	}

	if d,ok:=data["Name3"];ok{
		dv:=d
		this.Name3= dv
	}

	if d,ok:=data["Uid"];ok{
		dv, _:=strconv.ParseUint(d,10,64)
		this.Uid= dv
	}

}

func (this *Test2)UpdateData() {
	if len(this.changeData)>0{
		err:=redis.R.Hash_SetDataMap(this.table, this.changeData)
		if nil != err{
			return
		}
		this.changeData= make(map[string]interface{})
	}
}

func (this *Test2)Close() {
	this.UpdateData()
}

func(this *Test2) SetLang(value string){
	this.Lang = value
	this.changeData["Lang"]= value
}

func(this *Test2) GetLang() string{
	return this.Lang
}

func(this *Test2) SetName2(value string){
	this.Name2 = value
	this.changeData["Name2"]= value
}

func(this *Test2) GetName2() string{
	return this.Name2
}

func(this *Test2) SetName3(value string){
	this.Name3 = value
	this.changeData["Name3"]= value
}

func(this *Test2) GetName3() string{
	return this.Name3
}

func(this *Test2) SetUid(value uint64){
	this.Uid = value
	this.changeData["Uid"]= value
}

func(this *Test2) GetUid() uint64{
	return this.Uid
}

