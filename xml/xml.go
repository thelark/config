package xml

import (
	"io/ioutil"
	"encoding/xml"
)

func Init(filename string, r interface{}) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal(bytes, r); err != nil {
		return err
	}
	return nil
}
