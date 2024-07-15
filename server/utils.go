package main

import (
	"encoding/json"
	"os"
)


func readJson(filename string) (content map[string]string, err error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		println("Errorr reading from ", filename, " because ", err)
		return nil, err
	}
	var encryptedLusers map[string]string
	err = json.Unmarshal(file, &encryptedLusers)
	if err != nil {
		println("Error unmarshalling ", filename, " becasue ", err)
		return nil, err
	}
	return encryptedLusers, nil
}

func writeJson(filename string, content map[string]string) (err error) {
	file_data, err := json.Marshal(content)
	if err != nil {
		println("filedata failed to marshall ", content)
		return err
	}
	file, err := os.Open(filename)
	if err != nil {
		println("error opening file", filename, "because", err)
		return err
	}
	defer file.Close()
	_, err = file.Write(file_data)
	if err != nil {
		println("error writing to file ", filename, "because", err)
		return err
	}
	return nil
}


// appends json without overwrite
func appendJson(content map[string]string, newContent map[string]string) (finalContent map[string]string) {
	for key, val := range newContent {
		if _, ok := content[key]; !ok {
			content[key] = val
		}
	}
	return content
}