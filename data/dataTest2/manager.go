package DataTest2

import (
	"sync"
	"time"
)

var Manager *DataTest2Manager

type DataTest2Manager struct {
	closed	bool
	mutex sync.RWMutex
	data map[uint64]*DataTest2
}

func init(){
	Manager := new(DataTest2Manager)
	Manager.data= make(map[uint64]*DataTest2)
	go Manager.Update()
}

func (this *DataTest2Manager)AddData(d *DataTest2){
	if  nil == d{
		return
	}
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.data[d.uid]= d
}

func (this *DataTest2Manager)DelData(uid uint64)bool{
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

func (this *DataTest2Manager)GetData(uid uint64)*DataTest2{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if  v,ok := this.data[uid];ok{
		return v
	}
	return nil
}

func AddData(data *DataTest2){
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

func GetData(uid uint64) *DataTest2 {
	if nil== Manager {
		return nil
	}
	return Manager.GetData(uid)
}

func (this *DataTest2Manager)Close(){
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

func (this *DataTest2Manager)Update(){
	t := time.Tick(500 * time.Millisecond)
	for _ = range t {
		this.mutex.RLock()
		if true == this.closed{
			this.mutex.RUnlock()
			break
		}
		for _, v := range this.data{
			if nil !=v{
				v.UpdateData()
			}
		}
		this.mutex.RUnlock()
	}
}

