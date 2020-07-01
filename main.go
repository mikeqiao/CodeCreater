// codecreater project main.go
package main

import (
	"flag"
	"fmt"

	"github.com/mikeqiao/codecreater/read"
)

var Path string
var Type string

func init() {
	flag.StringVar(&Path, "path", "default", "common")
	flag.StringVar(&Type, "type", "default", "data")
}
func main() {
	fmt.Println("Hello World!")
	/*

	 */
	flag.Parse() //暂停获取参数
	if "data" == Type {
		read.InitData()
		read.CreateData(Path)
	}
	if "service" == Type {
		read.InitService()
		read.CreateService(Path)
	}
}
