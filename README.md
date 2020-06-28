# codecreater
	结构化数据
	{
		基础数据类型 直接存放hash
		struct  数据类型
		Base_name {
			key_name  value   hash field
			
		}
		Base_name:key_name value
		map[uint32]mtype
		{
			hash field   map:key  如果 mtype 是基础类型 就是map:key     如果mtyp 是 复杂类型 就是 map:key.字段名字
		}
		
		所以 一个struct 的所有成员的 hash  field
		{
			基础类型          成员名
			struct类型       成员名  作为  前缀+ "." 链接内部结构的 命名
			map类型         成员名 作为前缀 +":" + key 值 （如果map 元素不是基础类型 + "." 链接内部结构的 命名）
		}
		 
		限制 struct  内部不能使用interface{} 类型  不推荐使用嵌套的map类型
	}
	
	//规则
	1 struct   首字母大写
	2 param    首字母小写
	3 存库的节点 struct  Data开头
		创建 InitData UpdateData Close 方法
	4 不是data开头的struct  别的 struct 的子节点
	
	5 map  涉及 add get del 操作
	
	自动生成服务
	{
		不能修改原来存在的文件
		只能新增
		1判断mod 是否存在 不存在 创建
		2添加func
		3注册到mod
		
		service
		{
			ModName
			Name
			Req
			Res
		}
		
		common/module/ mod 名字 路径
		modXXX.go 文件
		{
			创建mod 注册到 框架
		}
		handleXXX.go 文件
		{
			注册到 mod
			定义方法	
		}
		
		
	}