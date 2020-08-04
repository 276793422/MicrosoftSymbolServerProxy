package main

import (
	"encoding/json"
	"io/ioutil"
)


//定义配置文件解析后的结构
type PathConfig struct {
	Path      string
	Exec      string
}

type Config struct {
	Ip      string
	Port    string
	Root    string
	Log		string
	Path  []PathConfig
}

type zJson struct {
}

func NewJson() *zJson {
	return &zJson{}
}

func (jst *zJson) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func LoadConfig(filename string) Config{
	v := Config{}
	JsonParse := NewJson()
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load(filename, &v)
	return  v
}