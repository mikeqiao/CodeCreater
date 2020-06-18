package Test2

import (
	"sync"
	"time"
)

var Manager *Test2Manager

type Test2Manager struct {
	closed	bool
	mutex sync.RWMutex
	data map[uint64]*Test2
}

func init(){
	Manager := new(Test2Manager)
	Manager.data= make(map[uint64]*Test2)
	go Manager.Update()
}

func (this *Test2Manager)AddData(d *Test2){
	if  nil == d{
		return
	}
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.data[d.uid]= d
}

func (this *Test2Manager)DelData(uid uint64)bool{
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if  v,ok := this.data[uid];ok{
		if nil !=v{
			v.Close()
		}
		delete(this.data, uid)
		return true
	}
	return false
}

func (this *Test2Manager)GetData(uid uint64)*Test2{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if  v,ok := this.data[uid];ok{
		return v
	}
	return nil
}

func AddData(data *Test2){
	if nil== Manager || nil == data{
		return
	}
	Manager.AddData(data)
}

func DelData(uid uint64)bool{
	if nil== Manager {
		return false
	}
	return Manager.DelData(uid)
}

func GetData(uid uint64) *Test2 {
	if nil== Manager {
		return nil
	}
	return Manager.GetData(uid)
}

func (this *Test2Manager)Close(){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.closed= true
	for k, v := range this.data{
		if nil !=v{
			v.Close()
		}
		delete(this.data, k)
	}
}

func (this *Test2Manager)Update(){
	t := time.Tick(500 * time.Millisecond)
	for _ = range t {
		this.mutex.RLock()
		if true == this.closed{
			this.mutex.Unlock()
			break
		}
		for _, v := range this.data{
			if nil !=v{
				v.UpdateData()
			}
		}
		this.mutex.Unlock()
	}
}

