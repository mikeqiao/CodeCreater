package Test

import (
	"sync"
	"github.com/mikeqiao/Db/redis"
)

type Test struct {
	dasda	int32
	Name	string
	Uid	uint64
	cdd	int32
	changeData map[string]interface{}
	mutex sync.RWMutex
}

func NewTest() *Test{
	data := new(Test)
	data.changeData= make(map[string]interface{})
	return data
}

func (this *Test)InitData() {
	table = Test + _fmt.Sprint(this.uid)
	data, _:=redis.R.Hash_GetAllData(table)
	if d,ok:=data["dasda"];ok{
		dd, _:=strconv.Atoi(d)
		dv:=int32(dd)
		this.dasda= dv
	}

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

	data := new(Test)
	data.changeData= make(map[string]interface{})
}

func(this *Test) Setdasda(value int32){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.dasda = value
	this.changeData["dasda"]= value
}

func(this *Test) Getdasda() int32{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.dasda
}

func(this *Test) SetName(value string){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.Name = value
	this.changeData["Name"]= value
}

func(this *Test) GetName() string{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Name
}

func(this *Test) SetUid(value uint64){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.Uid = value
	this.changeData["Uid"]= value
}

func(this *Test) GetUid() uint64{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Uid
}

func(this *Test) Setcdd(value int32){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.cdd = value
	this.changeData["cdd"]= value
}

func(this *Test) Getcdd() int32{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.cdd
}

