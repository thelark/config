package xml

import (
	"testing"
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

func TestInit(t *testing.T) {
	c := new(Config)
	err := Init("test_config.xml", c)
	if err != nil {
		t.Error(err)
		return
	}
	if c.Version != "v1.0.1" {
		t.Error("Get value error.")
		return
	}
	t.Log("Init testing success!")
}
