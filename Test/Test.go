package Test

import (
)

type Test struct {
	Name	string
	Uid	uint64
	cdd	int32
	dasda	int32
}

func(this *Test) SetName(value string){
	this.Name = value
}

func(this *Test) GetName() string{
	return this.Name
}

func(this *Test) SetUid(value uint64){
	this.Uid = value
}

func(this *Test) GetUid() uint64{
	return this.Uid
}

func(this *Test) Setcdd(value int32){
	this.cdd = value
}

func(this *Test) Getcdd() int32{
	return this.cdd
}

func(this *Test) Setdasda(value int32){
	this.dasda = value
}

func(this *Test) Getdasda() int32{
	return this.dasda
}

