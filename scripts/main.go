package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// var list, f []string
	folderList := []string{}
	path := "../containers"
	folders, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	} else if (len(folders)) == 0 {
		log.Println("empty source folder")
	}

	for i := 0; i < len(folders); i++ {

		if folders[i].IsDir() {
			// fmt.Println(folders[i].Name())
			folderList = append(folderList, folders[i].Name())
		}
	}
	// fmt.Println(folderList)
	output, err := json.Marshal(folderList)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(output))
}
