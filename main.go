package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	chartPath := os.Getenv("CHART_PATH")
	imageTag := os.Getenv("IMAGE_TAG")

	data, err := ioutil.ReadFile(chartPath)
	if err != nil {
		log.Fatalf("Error reading chart: %v", err)
	}

	var valuesData map[string]interface{}
	err = yaml.Unmarshal(data, &valuesData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Ensure that "image" key exists and is a map
	imageData, ok := valuesData["image"].(map[string]interface{})
	if !ok {
		log.Fatalf("'image' key is missing or is not a map")
	}

	imageData["tag"] = imageTag

	modifiedData, err := yaml.Marshal(valuesData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(modifiedData))
	err = ioutil.WriteFile(chartPath, modifiedData, 0644)
	if err != nil {
		log.Fatalf("Error writing modified data to file: %v", err)
	}
}
