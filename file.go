package main

import (
	"io/ioutil"
	"os"
	"speedControl/config"
	"strconv"
)

// 获取数据库中转速
func getCurrent() int  {
	currentStr := readDb()
	current,_ :=strconv.Atoi(currentStr)
	return current
}

//设置转速到数据库
func setCurrent(persent int)  {
	str := strconv.Itoa(persent)
	writeDb(str)
}

//读文件内容
func readDb() string  {
	b, err := ioutil.ReadFile("t.db") // just pass the file name
	if err != nil {
		config.Log.Error(err)
	}
	str := string(b) // convert content to a 'string'
	return str
}
//写文件内容
func writeDb(str string)  {
	f,err := os.OpenFile("t.db",os.O_WRONLY|os.O_CREATE,0666)
	defer f.Close()
	if err != nil{
		config.Log.Error(err.Error())
	}
	_,err = f.Write([]byte(str))
	if err != nil{
		config.Log.Error(err.Error())
	}
}
