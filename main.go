// codecreater project main.go
package main

import (
	"fmt"

	"github.com/mikeqiao/codecreater/class"
	"github.com/mikeqiao/codecreater/file"
	"github.com/mikeqiao/codecreater/read"
)

func main() {
	fmt.Println("Hello World!")
	/*

	 */
	read.Init()
	Test()
}

func Test() {
	for k, v := range read.Dmap {

		c := class.NewClass(k)
		c.Lock = false
		c.InitData(v)
		c.Init()
		c.ManagerInit()
		f := new(file.File)
		path := k
		name := path + "/" + k + ".go"
		name2 := path + "/" + "manager.go"

		f.CreateDir(path)

		f.CreateFile(name)
		f.Write(c.GetBuff())
		f.Close()
		f2 := new(file.File)
		f2.CreateFile(name2)
		f2.Write(c.GetManagerBuff())
		f2.Close()
	}
}
