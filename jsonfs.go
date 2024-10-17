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
			relativePath, err := filepath.Rel(folderPath, path)
			if err != nil {
				return err
			}
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			result.Files = append(result.Files, File{
				Name:    relativePath,
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

func Unmarshal(jsonStr, destPath string) error {
	var result Result
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	folderPath := filepath.Join(destPath, result.Folder)
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating folder: %w", err)
	}

	for _, file := range result.Files {
		filePath := filepath.Join(folderPath, file.Name)
		dir := filepath.Dir(filePath)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
		err := ioutil.WriteFile(filePath, []byte(file.Content), os.ModePerm)
		if err != nil {
			return fmt.Errorf("error writing file %s: %w", filePath, err)
		}
	}

	return nil
}
