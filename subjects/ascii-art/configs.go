package asciiart

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configs struct {
	FileTheme string `json: "FileTheme"`
	Color     string `json: "Color"`
}

// Set GConfigs From File
// Creating Configs File if not exists
func ReadConfigs() error {
	fileName := "configs.json"
	dataInBytes, err := os.ReadFile(fileName)

	//  Creating "configs.json" IF NOT EXISTS
	if err != nil {
		return SaveConfigs()
	}

	// Set Configs From File
	err = json.Unmarshal(dataInBytes, GConfigs)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// Save GConfigs to File
// Creating Configs File if not exists
func SaveConfigs() error {
	fileName := "configs.json"
	dataInBytes, err := json.Marshal(GConfigs)
	if err != nil {
		return err
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(dataInBytes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// func SaveConfigsFromParams() error {

// }
