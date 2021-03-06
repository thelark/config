package yaml

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
	"reflect"
)

var data map[interface{}]interface{}

func Init(filename string) error {
	data = make(map[interface{}]interface{})
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal([]byte(fileData), data); err != nil {
		return err
	}
	return nil
}

func Get(key string) interface{} {
	var retData interface{}
	searchKeys := strings.Split(key, ".")
	tempData := data
	for index := 0; index < len(searchKeys); index++ {
		if reflect.TypeOf(tempData[searchKeys[index]]) == reflect.TypeOf(data) {
			tempData = tempData[searchKeys[index]].(map[interface{}]interface{})
		} else {
			retData = tempData[searchKeys[index]]
		}
	}
	if retData != nil {
		return retData
	} else {
		return ""
	}
}
