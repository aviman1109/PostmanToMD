package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func license() {
	var lic string
	if _, err := os.Stat("License"); err == nil {
		r, err := ioutil.ReadFile("License")
		check(err)
		lic = string(r)
	} else if os.IsNotExist(err) {
		lic = "no License"
	}
	fmt.Println(lic)
}
func main() {
	// license()
	file := os.Args[1]
	jsonFile, err := os.Open(file)
	check(err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	check(err)
	var jsonCollection *Item
	json.Unmarshal(byteValue, &jsonCollection)

	readJSON(jsonCollection)

}
