package main

import (
	"encoding/json"
	"fmt"
)

// Marshal
// serialize []string to string
func serializeToString() {
	slice := []string{"apple", "peach", "pear"}

	res, _ := json.Marshal(slice)

	// ["apple","peach","pear"] (plain string)
	fmt.Println(string(res))
}

// serialize map to string
func serializeMapToString() {
	data := make(map[string]interface{})

	data["name"] = "dongseon"
	data["sex"] = "male"

	jsonDoc, _ := json.Marshal(data)

	// {"name":"dongseon","sex":"male"} (plain string)
	fmt.Println(string(jsonDoc))
}

// Unmarshal
// unserialize json string to map[string]interface{}
func unserializeJsonStringToMap() {
	var res map[string]interface{}

	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

	if err := json.Unmarshal(byt, &res); err != nil {
		panic(err)
	}

	// map[num:6.13 strs:[a b]]
	fmt.Println(res)
}

// unserialize json string to custom struct
type Config struct {
	Enviroment string
	Version    string
	HostName   string
}

func unserializeJsonStringToStruct() {
	c := new(Config)

	jsonDoc := `
		{
			"Enviroment": "Dev",
			"Version": "1.0",
			"HostNam": ""
		}
	`

	if err := json.Unmarshal([]byte(jsonDoc), &c); err != nil {
		panic("json형태가 올바르지 않습니다")
	}

	fmt.Println(c)
}

type ApiResponse struct {
	Status  string
	Data    map[string]interface{}
	Message string
}

func main() {
	serializeToString()
	serializeMapToString()

	unserializeJsonStringToMap()
	unserializeJsonStringToStruct()
}
