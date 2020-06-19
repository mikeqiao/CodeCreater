package class

import "strings"

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
