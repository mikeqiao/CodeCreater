package class

import (
	"fmt"
	"strings"
)

const (
	DataPath  = "data/"
	ModPath   = "module/"
	ProtoPath = "msg/"
)

func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

func strFirstToLower(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 65 && strArry[0] <= 90 {
		strArry[0] += 32
	}
	return string(strArry)
}

func CheckBaseType(ttype string) bool {
	switch ttype {
	case "string":
		return true
	case "uint64":
		return true
	case "uint32":
		return true
	case "int32":
		return true
	case "int64":
		return true
	case "uint":
		return true
	case "int":
		return true
	case "float64":
		return true
	case "float32":
		return true
	case "bool":
		return true
	default:
		return false
	}
	return false
}

func CheckStruct(ttype string) (is bool, mtype string) {
	if strings.HasPrefix(ttype, "*") {
		is = true
		stype := strings.TrimLeft(ttype, "*")
		mtype = "*common." + stype
	}
	return
}

func CheckMap(ttype string) bool {
	if strings.HasPrefix(ttype, "map") {
		return true
	}
	return false
}

func CheckMapStruct(ttype string) (is bool, ktype, vtype, mtype string) {
	d := strings.Split(ttype, "]")
	if 2 == len(d) {
		sis, smt := CheckStruct(d[1])
		if sis {
			is = true

			mtype = d[0] + "]" + smt
			vtype = smt
		} else {
			vtype = d[1]
			mtype = ttype
		}
		ktype = strings.TrimLeft(d[0], "map[")
	}

	return
}

func CheckValueType(ttype string, c *Class) (have bool) {
	have = true
	dvalue := ""
	switch ttype {
	case "string":
		dvalue = fmt.Sprintf("		dv:=d\n")
		c.buff.WriteString(dvalue)
	case "uint64":
		dvalue = fmt.Sprintf("		dv, _:=strconv.ParseUint(d,10,64)\n") //strconv.ParseFloat() ParseUint(d,10,64)
		c.buff.WriteString(dvalue)
	case "uint32":
		dvalue = fmt.Sprintf("		dd, _:=strconv.ParseUint(d,10,64)\n")
		c.buff.WriteString(dvalue)
		nvalue := fmt.Sprintf("		dv:=uint32(dd)\n")
		c.buff.WriteString(nvalue)
	case "int32":
		dvalue = fmt.Sprintf("		dd, _:=strconv.Atoi(d)\n")
		c.buff.WriteString(dvalue)
		nvalue := fmt.Sprintf("		dv:=int32(dd)\n")
		c.buff.WriteString(nvalue)
	case "int64":
		dvalue = fmt.Sprintf("		dv, _:=strconv.ParseInt(d,10,64)\n")
		c.buff.WriteString(dvalue)
	case "float64":
		dvalue = fmt.Sprintf("		dv, _:=strconv.ParseFloat(d,64)\n")
		c.buff.WriteString(dvalue)
	case "float32":
		dvalue = fmt.Sprintf("		dd, _:=strconv.ParseFloat(d,64)\n")
		c.buff.WriteString(dvalue)
		nvalue := fmt.Sprintf("		dv:=float32(dd)\n")
		c.buff.WriteString(nvalue)
	case "bool":
		dvalue = fmt.Sprintf("		dv, _:=strconv.ParseBool(d)\n")
		c.buff.WriteString(dvalue)
	default:
		have = false
	}
	return
}

func CheckValueType2(ttype string, c *Class) (have bool) {
	have = true
	dvalue := ""
	switch ttype {
	case "string":
		dvalue = fmt.Sprintf("		dv1:=d1\n")
		c.buff.WriteString(dvalue)
	case "uint64":
		dvalue = fmt.Sprintf("		dv1, _:=strconv.ParseUint(d1,10,64)\n") //strconv.ParseFloat() ParseUint(d,10,64)
		c.buff.WriteString(dvalue)
	case "uint32":
		dvalue = fmt.Sprintf("		dd1, _:=strconv.ParseUint(d1,10,64)\n")
		c.buff.WriteString(dvalue)
		nvalue := fmt.Sprintf("		dv1:=uint32(dd1)\n")
		c.buff.WriteString(nvalue)
	case "int32":
		dvalue = fmt.Sprintf("		dd1, _:=strconv.Atoi(d1)\n")
		c.buff.WriteString(dvalue)
		nvalue := fmt.Sprintf("		dv1:=int32(dd1)\n")
		c.buff.WriteString(nvalue)
	case "int64":
		dvalue = fmt.Sprintf("		dv1, _:=strconv.ParseInt(d1,10,64)\n")
		c.buff.WriteString(dvalue)
	case "float64":
		dvalue = fmt.Sprintf("		dv1, _:=strconv.ParseFloat(d1,64)\n")
		c.buff.WriteString(dvalue)
	case "float32":
		dvalue = fmt.Sprintf("		dd1, _:=strconv.ParseFloat(d1,64)\n")
		c.buff.WriteString(dvalue)
		nvalue := fmt.Sprintf("		dv1:=float32(dd1)\n")
		c.buff.WriteString(nvalue)
	case "bool":
		dvalue = fmt.Sprintf("		dv1, _:=strconv.ParseBool(d1)\n")
		c.buff.WriteString(dvalue)
	default:
		have = false
	}
	return
}
