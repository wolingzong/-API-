package main
/*

#cgo CFLAGS : -I../c
#cgo LDFLAGS: -L../c -lvideo_so
#include "video_so.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)
import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const UploadPath string ="../upload/"

func Hello(res http.ResponseWriter, req *http.Request) {

	// 根据字段名获取表单文件
	formFile, header, err := req.FormFile("cardFile")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		return
	}
	defer formFile.Close()

	// 创建保存文件
	destFile, err := os.Create(UploadPath + header.Filename)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return
	}

	name := C.CString(UploadPath + header.Filename)
	result := C.exeFFmpegCmd(name)
	outVideo := C.GoString(result)
	C.free(unsafe.Pointer(name))
	C.free(unsafe.Pointer(result))

	fmt.Println(outVideo)


	res.Header().Set("Content-Type", "image/png")
	res.Header().Set("Content-Disposition",fmt.Sprintf("inline; filename=\"%s\"",header.Filename))

	file, err := ioutil.ReadFile(outVideo)
	if err != nil {
		fmt.Fprintf(res,"查无此图片")
		return
	}
	res.Write(file)


	//fmt.Fprintln(res, getHelloMsg(header.Filename))

}

func getHelloMsg(filename string) string {
	log.Println("getHelloMsg ", filename)

	//fmt.Println("20*30=", C.test_so_func(20, 30))
	//fmt.Println("hello world")




	//name := C.CString("World")
	//defer C.free(unsafe.Pointer(name))
	//
	//C.exeFFmpegCmd(name)




	return filename

}

func main() {





	http.HandleFunc("/upload_card", Hello) //注册URI路径与相应的处理函数
	fmt.Println("Start listening...")

	err := http.ListenAndServe(":80", nil) // 监听9090端口，就跟javaweb中tomcat用的8080差不多一个意思吧
	if err != nil {
		panic(err)
		log.Fatal("ListenAndServe: ", err)
	}

}