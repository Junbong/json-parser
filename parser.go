package jsonparser

import (
	"encoding/json"
	"fmt"
	"strings"
	"errors"
	"reflect"
	"runtime"
	"strconv"
)

type Parser struct {
	TypeMap		map[string]string
}

func New(config []byte) *Parser {
	ps := &Parser{
		TypeMap:make(map[string]string),
	}

	// Parse config
	if config != nil {
		var res map[string]interface{}
		json.Unmarshal(config, &res)

		for k, v := range res {
			ps.TypeMap[k] = v.(string)
		}
		fmt.Println("TypeMap", ps.TypeMap)
	}

	return ps
}

func (p Parser) Marshal(v interface{}) ([]byte, error) {
	// TODO
	return json.Marshal(v)
}

func (p Parser) Unmarshal(data []byte) (res map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	res = make(map[string]interface{})
	m := make(map[string]interface{})

	// Use decoder for decode natively
	dec := json.NewDecoder(strings.NewReader(string(data)))
	dec.UseNumber()
	if e := dec.Decode(&m); e != nil {
		return nil, e
	}

	for k, v := range m {
		if t, contains := p.TypeMap[k]; contains {
			if tv, e := castValue(v, t); e == nil {
				res[k] = tv
			} else {
				return nil, e
			}
		} else {
			res[k] = v
		}
	}
	return
}

func castValue(v interface{}, t string) (interface{}, error) {
	switch t {
	case "string", "text":
		return v.(string), nil
	case "bool", "boolean":
		if tv, e := parseBool(v, t); e == nil {
			return tv, nil
		} else {
			return nil, e
		}
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		if tv, e := parseInt(v, t); e == nil {
			return tv, nil
		} else {
			return nil, e
		}
	case "float", "float32", "float64":
		if tv, e := parseFloat(v, t); e == nil {
			return tv, nil
		} else {
			return nil, e
		}
	default:
		return nil, errors.New(fmt.Sprintf("Type '%s' is not supported", t))
	}
}

func parseBool(v interface{}, _ string) (interface{}, error) {
	tp := reflect.TypeOf(v)
	if tp.Kind() == reflect.Bool {
		return v.(bool), nil
	} else if tp.Kind() == reflect.String {
		return strconv.ParseBool(v.(string))
	} else {
		return nil, errors.New("Unconvertable type")
	}
}

func parseInt(v interface{}, t string) (interface{}, error) {
	tp := reflect.TypeOf(v)
	if tp.String() != "json.Number" {
		return nil, errors.New("Is not a type of JSON number")
	}
	tv := v.(json.Number)

	if x64, err := tv.Int64(); err == nil {
		switch t {
		case "int":
			return int(x64), nil
		case "int8":
			return int8(x64), nil
		case "int16":
			return int16(x64), nil
		case "int32":
			return int32(x64), nil
		case "int64":
			return int64(x64), nil
		case "uint":
			return uint(x64), nil
		case "uint8":
			return uint8(x64), nil
		case "uint16":
			return uint16(x64), nil
		case "uint32":
			return uint32(x64), nil
		case "uint64":
			return uint64(x64), nil
		default:
			return uint64(x64), nil
		}
	} else {
		return nil, err
	}
}

func parseFloat(v interface{}, t string) (interface{}, error) {
	tp := reflect.TypeOf(v)
	if tp.String() != "json.Number" {
		return nil, errors.New("Is not a type of JSON number")
	}
	tv := v.(json.Number)

	if f64, err := tv.Float64(); err == nil {
		switch t {
		case "float", "float32":
			return float32(f64), nil
		case "float64":
			return f64, nil
		default:
			return f64, nil
		}
	} else {
		return nil, err
	}
}
