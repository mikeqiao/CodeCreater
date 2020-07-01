package TestMod

import (
	m"github.com/mikeqiao/newworld/manager"
	mod"github.com/mikeqiao/newworld/module"
	"github.com/mikeqiao/codecreater/proto"
)

var Mod *mod.Mod

func Init(){
	Mod := m.NewMod("TestMod")
	Register()
	m.ModManager.Registe(Mod)
}

func Register(){
	Mod.Register("Service1", Service1, proto.Req{}, proto.Res{})
	Mod.Register("Service2", Service2, proto.Req{}, proto.Res{})
}

