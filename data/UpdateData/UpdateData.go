package updateData

import (
	"github.com/mikeqiao/codecreater/data/DataTest2"
)

type UpdateData struct {
	userData	*DataTest2.DataTest2
	uid	uint64
}

func NewUpdateData(uid uint64) *UpdateData{
	d := new(UpdateData)
	d.uid= uid
	return d
}

func(this *UpdateData) SetuserData(value *DataTest2.DataTest2){
	this.userData = value
}

func(this *UpdateData) GetuserData() *DataTest2.DataTest2{
	return this.userData
}

func(this *UpdateData) Destroy(){
 if nil != this.userData {
		this.userData.Destroy()
	}
}

