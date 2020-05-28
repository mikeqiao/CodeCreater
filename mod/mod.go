package mod

import (
	"strconv"
	"sync"
)

type Mod struct {
	//包含字段
	uid  uint64
	name string
	//生成字段
	mutex     sync.RWMutex
	tableName string //表名  Mod_uid
	changed   bool
	change    map[string]interface{} //记录修改的字段 值
}

func (m *Mod) Init() {
	m.change = make(map[string]interface{})
}

func (m *Mod) InitData() { //去读redis 获取用户数据
	//hash  getall
	data := make(map[string]string)
	if v, ok := data["uid"]; ok {
		m.uid, _ = strconv.ParseUint(v, 10, 64) //格式转换
	}

}

func (m *Mod) Run() {

}

func (m *Mod) Update() {
	if m.changed {
		//hash set （table, m.change）
		m.change = make(map[string]interface{})
		m.changed = false
	}
}

func (m *Mod) Close() {

}

//set
func (m *Mod) SetUid(id uint64) {
	m.uid = id
	m.change["uid"] = id
	m.changed = true
}
func (m *Mod) SetName(name string) {
	m.name = name
	m.change["name"] = name
	m.changed = true
}

//get
func (m *Mod) GetUid() uint64 {
	return m.uid
}

func (m *Mod) GetName() string {
	return m.name
}
