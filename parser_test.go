package jsonparser_test

import (
	"testing"
	"github.com/Junbong/static-json-parser"
	"fmt"
	"io/ioutil"
	"reflect"
)

func TestNew(t *testing.T) {
	b, e := ioutil.ReadFile("config.json")
	if e != nil {
		t.Errorf("Configuration read error.. %s", e)
	}

	if p := jsonparser.New(b); p != nil {
		fmt.Println("Instance created")
	} else {
		t.Error("Instance not created")
	}
}

func TestParser_Marshal(t *testing.T) {
	// TODO
}

func TestParser_Unmarshal(t *testing.T) {
	config := `
		{
		  "name": "text",
		  "age": "uint",
		  "height": "float32",
		  "weight": "float32",
		  "score": "int",
		  "graduated": "bool"
		}
	`

	input := `
		{
		  "name": "Lloyd",
		  "age": 30,
		  "height": 184.53,
		  "weight": 86.5423,
		  "score": -47,
		  "graduated": "true",
		  "hobby": "Listening Music",
		  "languages": [ "KR", "EN", "JP", "CN" ]
		}
	`

	p := jsonparser.New([]byte(config))
	res, err := p.Unmarshal([]byte(input))

	fmt.Println(res)

	if err != nil {
		t.Error(err)
	} else if res["name"] != "Lloyd" {
		t.Error("Name not matched", res["name"], "with", "'Lloyd'")
	} else if res["age"] != uint(30) {
		t.Error("Age not matched", res["age"], "with", 30)
	} else if res["height"] != float32(184.53) {
		t.Error("Height not matched", res["height"], "with", 184.53)
	} else if res["weight"] != float32(86.5423) {
		t.Error("Weight not matched", res["weight"], "with", 86.5423)
	} else if res["score"] != int(-47) {
		t.Error("Score not matched", res["score"], "with", -47)
	} else if res["graduated"] != true {
		t.Error("Graduation not matched", res["graduated"], "with", true)
	} else if res["hobby"] != "Listening Music" {
		t.Error("Hobby not matched", res["name"], "with", "'Listening Music'")
	}
	langs := []interface{} { "KR", "EN", "JP", "CN" }
	s := reflect.ValueOf(res["languages"])
	for i:=0; i<s.Len(); i++ {
		if s.Index(i).Interface().(string) != langs[i] {
			t.Error("Array not matched", s.Index(i).Interface().(string), "with", langs[i])
		}
	}
}
