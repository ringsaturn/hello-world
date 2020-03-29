package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	JSONStr := `{"a":"c","b":[1,"2",true,[1,2,3,4]]}`
	value := gjson.Parse(JSONStr)
	fmt.Println(value)
	fmt.Println(value.Get("a"))
	fmt.Println(value.Get("b.0"))
	fmt.Println(value.Get("b.1"))
	fmt.Println(value.Get("b.2"))
	fmt.Println(value.Get("b.3"))
}
