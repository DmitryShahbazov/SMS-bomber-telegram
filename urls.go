package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Urls struct which contains an array of urls
// type Urls struct {
// 	Urls []URLData `json:"users"`
// }

// // URLData struct which contains a name, type and list of url data
// type URLData struct {
// 	Name   string `json:"name"`
// 	Method string `json:"method"`
// 	Link   string `json:"link"`
// 	Data   Data   `json:"data"`
// }

// // Data struct which contains a list of data
// type Data struct {
// 	Phone      string `json:"phone"`
// 	Identifier string `json:"identifier"`
// }

// JSONPath path to JSON urls
var JSONPath = "urls_data.json"

func parseJSON() []map[string]interface{} {
	jsonFile, err := os.Open(JSONPath)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result

}
