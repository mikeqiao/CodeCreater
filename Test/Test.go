package Test

import (
	"sync"
	"github.com/mikeqiao/Db/redis"
	"strconv"
	"fmt"
)

type Test struct {
	Name	string
	Uid	uint64
	cdd	int32
	dasda	int32
	uid	uint64
	table	string
	changeData map[string]interface{}
}

func NewTest(uid uint64) *Test{
	data := new(Test)
	data.uid= uid
	data.changeData= make(map[string]interface{})
	data.table = "Test_" + fmt.Sprint(data.uid)
	return data
}

func (this *Test)InitData() {
	data, _:=redis.R.Hash_GetAllData(this.table)
	if d,ok:=data["Name"];ok{
		dv:=d
		this.Name= dv
	}

	if d,ok:=data["Uid"];ok{
		dv, _:=strconv.ParseUint(d,10,64)
		this.Uid= dv
	}

	if d,ok:=data["cdd"];ok{
		dd, _:=strconv.Atoi(d)
		dv:=int32(dd)
		this.cdd= dv
	}

	if d,ok:=data["dasda"];ok{
		dd, _:=strconv.Atoi(d)
		dv:=int32(dd)
		this.dasda= dv
	}

}

func (this *Test)UpdateData() {
	if len(this.changeData)>0{
		err:=redis.R.Hash_SetDataMap(this.table, this.changeData)
		if nil != err{
			return
		}
		this.changeData= make(map[string]interface{})
	}
}

func (this *Test)Close() {
	this.UpdateData()
}

func(this *Test) SetName(value string){
	this.Name = value
	this.changeData["Name"]= value
}

func(this *Test) GetName() string{
	return this.Name
}

func(this *Test) SetUid(value uint64){
	this.Uid = value
	this.changeData["Uid"]= value
}

func(this *Test) GetUid() uint64{
	return this.Uid
}

func(this *Test) Setcdd(value int32){
	this.cdd = value
	this.changeData["cdd"]= value
}

func(this *Test) Getcdd() int32{
	return this.cdd
}

func(this *Test) Setdasda(value int32){
	this.dasda = value
	this.changeData["dasda"]= value
}

func(this *Test) Getdasda() int32{
	return this.dasda
}

