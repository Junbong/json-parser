# Strict JSON Parser

Example:
```go
json := `
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

// Type configs
config := `
  {
    "name": "text",
    "age": "uint",
    "height": "float",
    "weight": "float32",
    "score": "int",
    "graduated": "bool"
  }
`

p := sparser.New([]byte(config))

// Parser parses given JSON string to map[string]interface{}
res, err := p.Unmarshal([]byte(json))

for k, v := range res {
  fmt.Println(k, v, reflect.TypeOf(v).String())
}
```

Result:
```
weight 86.5423 float32
score -47 int
graduated true bool
hobby Listening Music string
languages [KR EN JP CN] []interface {}
name Lloyd string
age 30 uint
height 184.53 float32
```
