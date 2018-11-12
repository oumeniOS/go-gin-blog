package logging

import (
	"fmt"
	"os"
	"log"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath()string  {
	return fmt.Sprintf("%s",LogSavePath)
}

func getLogFileFullPath()string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s",TimeFormat,LogSaveName,LogFileExt)
	return fmt.Sprintf("%s%s",prefixPath,suffixPath)
}

//打开日志文件
func openLogFile(filePath string) *os.File  {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		{
			mkDir()
		}
	case os.IsPermission(err):
		{
			//无权限
			log.Fatalf("Permission :%v", err)
		}
	}

	handle, err := os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_RDWR,0644)
	if err != nil{
		log.Fatalf("Fail to OpenFile :%v",err)
	}
	return handle
}

func mkDir()  {
	dir,_ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(),os.ModePerm)
	if err != nil{
		panic(err)
	}
}











































