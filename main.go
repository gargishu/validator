package main

import (
	// "fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func validate(input string) bool {

	filePath := "modids.json"
	var modids [] string
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
        return false
    }
    err = json.Unmarshal(file, &modids)
    if err != nil {
        return false
    }
	var inputValue = strings.ToUpper(input)

	for _, value := range modids {
			if value == inputValue {
				fmt.Println("hello yes")
				return true
			}
		}
		fmt.Println("hello no")
	return false

}
