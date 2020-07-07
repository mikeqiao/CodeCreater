package updateData

import (
	dataTest2 "github.com/mikeqiao/codecreater/data/DataTest2"
)

type UpdateData struct {
	userData  *dataTest2.DataTest2
	userDatab *dataTest2.DataTest2
	uid       uint64
}

func NewUpdateData(uid uint64) *UpdateData {
	d := new(UpdateData)
	d.uid = uid
	return d
}

func (this *UpdateData) SetuserData(value *dataTest2.DataTest2) {
	this.userData = value
}

func (this *UpdateData) GetuserData() *dataTest2.DataTest2 {
	return this.userData
}

func (this *UpdateData) SetuserDatab(value *dataTest2.DataTest2) {
	this.userDatab = value
}

func (this *UpdateData) GetuserDatab() *dataTest2.DataTest2 {
	return this.userDatab
}

func (this *UpdateData) UpdateData() {
	if nil != this.userData {
		this.userData.UpdateData()
	}
	if nil != this.userDatab {
		this.userDatab.UpdateData()
	}
}

func (this *UpdateData) Close() {
	this.UpdateData()
}

func (this *UpdateData) Destroy() {
	if nil != this.userData {
		this.userData.Destroy()
	}
	if nil != this.userDatab {
		this.userDatab.Destroy()
	}
}
