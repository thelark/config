package json

import (
	"io/ioutil"
	"encoding/json"
	"strings"
	"reflect"
)

var data map[string]interface{}

func Init(filename string) error {
	data = make(map[string]interface{})
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
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
			tempData = tempData[searchKeys[index]].(map[string]interface{})
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
