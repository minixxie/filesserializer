package main

import (
	"fmt"
	"github.com/minixxie/jsonfs"
)

func main() {
	projectFolderPath := "/tmp/old/nodejs-project"

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
