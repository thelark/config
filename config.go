package main

import (
	"thelark.cn/config/xml"
	"fmt"
)

type Database struct {
	Addr     string `xml:"addr"`
	Password string `xml:"password"`
	DB       int `xml:"db"`
}
type Session struct {
	Time     string `xml:"time"`
	Open     bool `xml:"open"`
	Database Database `xml:"database"`
}
type Server struct {
	Port    string `xml:"port"`
	Session Session `xml:"session"`
}
type Config struct {
	Version      string `xml:"version"`
	Environments string `xml:"environments"`
	Server       Server `xml:"server"`
}

func main() {
	var err error

	c := new(Config)
	err = xml.Init("./xml/test_config.xml", c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

}
