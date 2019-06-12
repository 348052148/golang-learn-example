package main

import (
	"archive/tar"
	"bufio"
	"fmt"
	"io"
	"os"
)

func readerTar(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := tar.NewReader(file)
	for  {
		if header, ok := reader.Next(); ok == nil {
			fmt.Println(header.Name)
			if header.Name != "targo/i.php" && header.Name != "k.txt" {
				continue
			}
			buf := bufio.NewScanner(reader)
			for buf.Scan() {
				fmt.Printf(" %s, \n", buf.Text())
			}
		} else {
			fmt.Println(ok)
			break
		}

	}
}

func writerTar(filename, files string)  {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	w := tar.NewWriter(file)
	fileInfo, err := os.Stat(files)
	if err != nil {
		panic(err)
	}
	header, err := tar.FileInfoHeader(fileInfo, files)
	if ok :=w.WriteHeader(header); ok != nil {
		panic(ok)

	} else {
		fl, _ := os.Open(files)
		_, er := io.Copy(w, fl)
		if er != nil {
			panic(er)
		}
		defer fl.Close()
	}

	defer file.Close()

}

//tar
// reader,writer
// fileInfo -> fileHeader
// header -> os.FileInfoHeader

func main() {
	//writerTar("z1.tar", "k.txt")
	readerTar("z1.tar")
}
