package file

import (
	"bytes"
	"fmt"
	"os"
)

type File struct {
	F *os.File
}

//创建文件
func (f *File) CreateFile(path string) bool {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	//	defer file.Close()
	if err != nil {
		return false
	}
	f.F = file
	return true
}

//创建目录
func (f *File) CreateDir(path string) bool {
	if f.IsDirOrFileExist(path) == false {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true

}

//判断文件 或 目录是否存在
func (f *File) IsDirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)

}

// 判断给定文件名是否是一个目录
// 如果文件名存在并且为目录则返回 true。如果 filename 是一个相对路径，则按照当前工作目录检查其相对路径。
func (f *File) IsDir(filename string) bool {
	return f.isFileOrDir(filename, true)
}

// 判断给定文件名是否为一个正常的文件
// 如果文件存在且为正常的文件则返回 true
func (f *File) IsFile(filename string) bool {

	return f.isFileOrDir(filename, false)

}

// 判断是文件还是目录，根据decideDir为true表示判断是否为目录；否则判断是否为文件
func (f *File) isFileOrDir(filename string, decideDir bool) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	isDir := fileInfo.IsDir()
	if decideDir {
		return isDir
	}
	return !isDir
}

func (f *File) Write(b *bytes.Buffer) {
	if nil != f.F && nil != b {
		n, err := f.F.Write(b.Bytes())
		fmt.Println(n, err)
	}

}

func (f *File) Close() {
	if nil != f.F {
		f.F.Close()
	}

}
