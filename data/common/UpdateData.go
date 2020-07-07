package common

import (
	"github.com/mikeqiao/newworld/data"
	"strconv"
	"codecreater/data/common"
	"fmt"
)

type UpdateData struct {
	userData	*common.dataTest2
	uid	uint64
	prefix string
	update *data.UpdateMod
}

func NewUpdateData(uid uint64, prefix string, update *data.UpdateMod) *UpdateData{
	d := new(UpdateData)
	d.uid= uid
	d.prefix= prefix
	if nil != update{
		d.update= update
	}else{
		d.update= new(data.UpdateMod)
		table := "UpdateData_" + fmt.Sprint(d.uid)
		d.update.Init(table)
	}
	return d
}

func(this *UpdateData) SetuserData(value *common.dataTest2){
	this.userData = value
}

func(this *UpdateData) GetuserData() *common.dataTest2{
	return this.userData
}

func (this *UpdateData)InitDataParam(ks []string, d string) {
	if nil == this.update{
		return
	}
		if len(ks) <= 0 {
			return
		}
		tkey := ks[0]
		switch tkey {
			if nil == this.userData {
				this.userData= common.NewdataTest2(this.uid, "userData.", this.update)
			}
			this.userData.InitDataParam(ks[1:],d)
		}
}

func(this *UpdateData) Destroy(){
 if nil != this.userData {
		this.userData.Destroy()
	}
}

