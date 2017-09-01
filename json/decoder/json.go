package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	decoder := json.NewDecoder(jsonFile)
	for {
		var post post.Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error decoding JSON file %s: %s\n", filename, err)
		}
		fmt.Println(post)
	}
}
