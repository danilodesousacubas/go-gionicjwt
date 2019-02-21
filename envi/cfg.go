package envi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is struct to model enviroments configs
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Dbname   string `json:"dbname"`
		Password string `json:"password"`
		Sslmode  string `json:"sslmode"`
	}
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
}

func profile(filename string) (Config, error) {
	var config Config

	configFile, err := os.Open(filename)

	defer configFile.Close()

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return config, err
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("File reading file: %v\n", err)
		os.Exit(1)
	}

	json.Unmarshal(file, &config)

	return config, err
}
