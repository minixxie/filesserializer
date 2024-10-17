package main

import (
	"fmt"
	"jsonfs"
)

func main() {
	projectFolderPath := "/tmp/nodejs-project"

	folderJSONStr, err := jsonfs.Marshal(projectFolderPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(folderJSONStr)

	err = jsonfs.Unmarshal(folderJSONStr, "/tmp/new/")
	if err != nil {
		fmt.Println(err)
		return
	}
}
