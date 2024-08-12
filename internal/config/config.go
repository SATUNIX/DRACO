package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Command struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Arguments   []string `json:"arguments"`
	Paths       []string `json:"paths,omitempty"`
}

type Config struct {
	Menus        map[string][]Command `json:"menus"`
	ColorPalette struct {
		CursorColor   string `json:"cursorColor"`
		SelectedColor string `json:"selectedColor"`
		TextColor     string `json:"textColor"`
		BackgroundColor string `json:"backgroundColor"` // Added this if needed
	} `json:"colorPalette"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading config file: %v\n", err)
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Printf("Error unmarshaling config file: %v\n", err)
		return config, err
	}

	log.Println("Config loaded successfully.")
	return config, err
}
