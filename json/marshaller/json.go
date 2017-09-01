package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jbbarquero/gofirst/json/post"
)

func main() {
	fmt.Println("Unmarshalling JSON file")

	filename := os.Args[1]

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening JSON file %s: %s\n", filename, err)
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file %s: %s\n", filename, err)
	}

	fmt.Println(string(jsonData))
	var post post.Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
