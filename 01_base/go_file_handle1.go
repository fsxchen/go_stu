package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
)

func readFromFileByBufio() {
	fileObj, err := os.Open("./go_file_handle1.go")
	if err != nil {
		fmt.Printf("Open file failed, err:%v", err)
		return
	}

	defer fileObj.Close()
	// 创建要给io，用来从文件中读取内容的对象
	reader := bufio.NewReader(fileObj)
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Printf("read line failed, err:%v", err)
		return
	}
	fmt.Println(line)
}

func main() {
	readFromFileByBufio()
}