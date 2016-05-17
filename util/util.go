package util

import (
	"io"
	"os"
)

func Check(e error) {
	if e != nil {
		//panic(e)
		errStr := e.Error() + "\n"
		WriteInfoFile(errStr, "log")
	}
}

func CheckFileIsExist(fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func OpenVideoInfoFile(fileName string) (f *os.File, err error) {
	if CheckFileIsExist(fileName) { //如果文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(fileName) //创建文件
	}
	return
}

func WriteInfoFile(videoInfo string, fileName string) error {
	f, err := OpenVideoInfoFile(fileName)
	defer f.Close()
	Check(err)
	_, err = io.WriteString(f, videoInfo) //写入文件(字符串)
	Check(err)
	return err
}
