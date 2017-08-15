package main

import (
	"thelark.cn/config/json"
	"thelark.cn/config/yaml"
	"fmt"
)

func main() {
	var err error
	err = json.Init("./json/test_config.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json.Get("server.key"))
	err = yaml.Init("./yaml/test_config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yaml.Get("server.key"))
}
