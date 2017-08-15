package json

import "testing"

func TestInit(t *testing.T) {
	err := Init("test_config.json")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Init testing success!")
	}
}
func TestGet(t *testing.T) {
	Init("test_config.json")
	if Get("server.port").(string) == ":1213" {
		t.Log("Get testing success!")
	} else {
		t.Error("Get testing failed")
	}
}
