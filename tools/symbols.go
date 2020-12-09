package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetSymbolKey(name string) string {
	// Open our jsonFile
	jsonFile, err := os.Open("./data/symbols_name.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened symbols_name.json")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Declared an empty map interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(byteValue), &result)

	// Print the data type of result variable
	//fmt.Println(result[name])

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	code := fmt.Sprintf("%v", result[name])
	return code
}

func GetSymbolName(kay string) string {
	// Open our jsonFile
	jsonFile, err := os.Open("./data/symbol_reverse.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened symbols_name.json")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Declared an empty map interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(byteValue), &result)

	// Print the data type of result variable
	//fmt.Println(result[name])

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	code := fmt.Sprintf("%v", result[kay])
	return code
}

func GetSymbolInArray() ([]string, []string) {
	// Open our jsonFile
	jsonFile, err := os.Open("./data/symbols_name.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened symbols_name.json")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Declared an empty map interface
	var result map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(byteValue), &result)

	// Print the data type of result variable
	//fmt.Println(result[name])

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// copy keys to array
	//keys := make([]string, 0, len(result))
	//for k := range result {
	//	keys = append(keys, k)
	//}
	// copy values to array
	values := []string{}
	keys := []string{}
	for k, val := range result {
		values = append(values, val)
		keys = append(keys, k)
	}
	return values, keys
}

// TODO: update symbol list
func DataUpdater() {
	temp := mapReverser()
	saveMapInJson(temp)
}
func mapReverser() map[string]string {
	// Open our jsonFile
	jsonFile, err := os.Open("./data/symbols_name.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened symbols_name.json")

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Declared an empty map interface
	var result map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(byteValue), &result)

	// Print the data type of result variable
	//fmt.Println(result[name])

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	reversedMap := make(map[string]string)

	for key, value := range result {
		reversedMap[value] = key
	}
	return reversedMap
}
func saveMapInJson(obj map[string]string) {
	result := obj
	jsonString, _ := json.Marshal(result)
	err := ioutil.WriteFile("symbol_reverse.json", jsonString, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
