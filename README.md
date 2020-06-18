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
	背包
	{
		
		
	}
	装备
	{
		
	}
	
	buff
	{
		map
		{
			
		}
	}
	技能
	次数进度
	
	
	
	排行榜
	
	聊天
	好友
	队伍
	公会