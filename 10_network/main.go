package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age,string"`
	Nicknames []string `json:"nicknames"`
}

func main() {
	b := []byte(`{"name":"masana", "age":20,"nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames) // masana 0 [a b c]

	v, _ := json.Marshal(p)
	fmt.Println(string(v)) // {"name":"masana","age":"0","nicknames":["a","b","c"]}
}
