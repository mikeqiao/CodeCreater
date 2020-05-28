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
		c.InitData(v)
		c.Init()
		f := new(file.File)
		path := k
		name := path + "/" + k + ".go"
		f.CreateDir(path)
		f.CreateFile(name)
		f.Write(c.GetBuff())
		f.Close()
	}
}
