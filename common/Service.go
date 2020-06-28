package common

import (
	"github.com/mikeqiao/newworld/data"
	"strconv"
	"fmt"
)

type Service struct {
	name	TestFunc
	req	req
	res	res
	uid	uint64
	prefix string
	update *data.UpdateMod
}

func NewService(uid uint64, prefix string, update *data.UpdateMod) *Service{
	d := new(Service)
	d.uid= uid
	d.prefix= prefix
	if nil != update{
		d.update= update
	}else{
		d.update= new(data.UpdateMod)
		table := "Service_" + fmt.Sprint(d.uid)
		d.update.Init(table)
	}
	return d
}

func(this *Service) Setname(value ){
	this.name = value
}

func(this *Service) Getname() {
	return this.name
}

func(this *Service) Setreq(value ){
	this.req = value
}

func(this *Service) Getreq() {
	return this.req
}

func(this *Service) Setres(value ){
	this.res = value
}

func(this *Service) Getres() {
	return this.res
}

func (this *Service)InitDataParam(ks []string, d string) {
	if nil == this.update{
		return
	}
		if len(ks) <= 0 {
			return
		}
		tkey := ks[0]
		switch tkey {
		}
}

func(this *Service) Destroy(){
}

