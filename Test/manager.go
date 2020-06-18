package Test

import (
	"sync"
	"time"
)

var Manager *TestManager

type TestManager struct {
	closed	bool
	mutex sync.RWMutex
	data map[uint64]*Test
}

func init(){
	Manager := new(TestManager)
	Manager.data= make(map[uint64]*Test)
	go Manager.Update()
}

func (this *TestManager)AddData(d *Test){
	if  nil == d{
		return
	}
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.data[d.uid]= d
}

func (this *TestManager)DelData(uid uint64)bool{
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

func (this *TestManager)GetData(uid uint64)*Test{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if  v,ok := this.data[uid];ok{
		return v
	}
	return nil
}

func AddData(data *Test){
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

func GetData(uid uint64) *Test {
	if nil== Manager {
		return nil
	}
	return Manager.GetData(uid)
}

func (this *TestManager)Close(){
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

func (this *TestManager)Update(){
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

