package param

type Param struct {
	Name  string
	Type  string
	MTye  string // 如果是map 而且map元素为struct 类型 那 MType 为 map的元素类型
	TType uint32 // 1 基础类型  2 struct 3 map 元素为基础类型 4 map 元素为struct指针类型  先不支持数组类型
	Ktype string //map 的key值 类型
	Vtype string //map 的元素值类型
}

type Server struct {
	Name string
	Req  string
	Res  string
}
