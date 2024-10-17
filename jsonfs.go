package jsonfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Result struct {
	Folder string `json:"folder"`
	Files  []File `json:"files"`
}

func Marshal(folderPath string) (string, error) {
	var result Result
	result.Folder = filepath.Base(folderPath)

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			result.Files = append(result.Files, File{
				Name:    info.Name(),
				Content: string(content),
			})
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error reading folder: %w", err)
	}

	resultBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %w", err)
	}

	return string(resultBytes), nil
}
