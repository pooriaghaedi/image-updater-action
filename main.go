package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	chartPath := os.Getenv("CHART_PATH")
	data, err := ioutil.ReadFile(chartPath)
	if err != nil {
		fmt.Println("Chart Not found", err)
	}

	var valuesData map[string]interface{}

	err = yaml.Unmarshal(data, &valuesData)
	if err != nil {
		fmt.Println("JSON ERROR", err)
	}
	valuesData["image"].(map[string]interface{})["tag"] = os.Getenv("IMAGE_TAG")

	modifiedData, err := yaml.Marshal(valuesData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(modifiedData))
	err = ioutil.WriteFile(chartPath, modifiedData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
