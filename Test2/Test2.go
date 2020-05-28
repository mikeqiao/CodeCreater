package Test2

import (
)

type Test2 struct {
	Uid	uint64
	Lang	string
	Name2	string
	Name3	string
}

func(this *Test2) SetUid(value uint64){
	this.Uid = value
}

func(this *Test2) GetUid() uint64{
	return this.Uid
}

func(this *Test2) SetLang(value string){
	this.Lang = value
}

func(this *Test2) GetLang() string{
	return this.Lang
}

func(this *Test2) SetName2(value string){
	this.Name2 = value
}

func(this *Test2) GetName2() string{
	return this.Name2
}

func(this *Test2) SetName3(value string){
	this.Name3 = value
}

func(this *Test2) GetName3() string{
	return this.Name3
}

