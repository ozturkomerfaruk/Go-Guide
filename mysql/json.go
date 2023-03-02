package main

import (
	"encoding/json"
	"fmt"
)

func maind() {
	jsonStr := `
	{
		"data": {
			"object" : "card",
			"id" : "card_1231",
			"first_name" : "Omer",
			"last_name" : "Ozturk",
			"balance" : "54.32"			
		}
	}
	`
	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}

	fmt.Println(m)

	//------------------
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
