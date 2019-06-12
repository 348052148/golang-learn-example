package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//writerZip("zipd.zip")
	readerZip("zipd.zip")
}

func readerZip(filename string)  {
	//读取方式有2种
	//1. openReader 2. newReader
	file,_ := os.Open(filename)
	defer file.Close()
	fileInfo ,err := file.Stat()
	reader,err := zip.NewReader(file,fileInfo.Size())
	//reader,err := zip.OpenReader(filename)
	if err != nil {
		panic(err)
	}
	for _,file := range reader.File {
		r,_ := file.Open()
		buf := bufio.NewScanner(r)
		fmt.Println(file.Name)
		for buf.Scan() {
			fmt.Println(buf.Text())
		}
	}
}

func writerZip(filename string)  {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	w := zip.NewWriter(file)
	defer w.Close()
	//
	//filelist := "k.txt"
	//sfile,err := os.Open(filelist)
	//defer sfile.Close()
	//if err!= nil {
	//	panic(err)
	//}
	//fileInfo ,err := sfile.Stat()
	//fileInfo.Size()
	//if err != nil {
	//	panic(err)
	//}
	//fileHeader,_ := zip.FileInfoHeader(fileInfo)
	//fmt.Println(fileHeader)
	//writer,err := w.CreateHeader(fileHeader)
	//if err!=nil {
	//	panic(err)
	//}
	//_,err = io.Copy(writer, sfile)
	//if err!= nil {
	//	panic(err)
	//}
	//2种方式 1.将文件写入，2，创建一个项写入
	writer,_ := w.Create("f.txt")
	writer.Write([]byte("first Item"))
}
