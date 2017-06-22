package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

func main() {

	//go run .\json.go .\file.json fileOut.json
	if len(os.Args) != 3 {
		log.Fatal("Usage json fileIn fileOut")
	}
	fmt.Println(os.Args)

	fmt.Println("JSON BEGIN")

	jsonFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error opening JSON file ", err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Error reading JSON file ", err)
	}

	var post Post
	json.Unmarshal(jsonData, &post)

	fmt.Println(post)

	jsonOut, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		log.Fatal("Error marshalling JSON ", err)
	}

	jsonFileDest, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal("Error creating JSON destination file ", err)
	}
	defer jsonFileDest.Close()

	dest := io.MultiWriter(os.Stdout, jsonFileDest)

	if _, err := io.Copy(dest, bytes.NewReader(jsonOut)); err != nil {
		log.Fatal("Error writing JSON file ", err)
	}

	fmt.Println("JSON END")

}
