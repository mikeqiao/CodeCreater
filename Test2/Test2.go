package Test2

import (
	"sync"
	"github.com/mikeqiao/Db/redis"
)

type Test2 struct {
	Name2	string
	Name3	string
	Uid	uint64
	Lang	string
	changeData map[string]interface{}
	mutex sync.RWMutex
}

func NewTest2() *Test2{
	data := new(Test2)
	data.changeData= make(map[string]interface{})
	return data
}

func (this *Test2)InitData() {
	table = Test2 + _fmt.Sprint(this.uid)
	data, _:=redis.R.Hash_GetAllData(table)
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

	if d,ok:=data["Lang"];ok{
		dv:=d
		this.Lang= dv
	}

	data := new(Test2)
	data.changeData= make(map[string]interface{})
}

func(this *Test2) SetName2(value string){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.Name2 = value
	this.changeData["Name2"]= value
}

func(this *Test2) GetName2() string{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Name2
}

func(this *Test2) SetName3(value string){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.Name3 = value
	this.changeData["Name3"]= value
}

func(this *Test2) GetName3() string{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Name3
}

func(this *Test2) SetUid(value uint64){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.Uid = value
	this.changeData["Uid"]= value
}

func(this *Test2) GetUid() uint64{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Uid
}

func(this *Test2) SetLang(value string){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.Lang = value
	this.changeData["Lang"]= value
}

func(this *Test2) GetLang() string{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Lang
}

