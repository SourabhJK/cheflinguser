package config

import (
	"encoding/json"
	"fmt"
	"os"
	"flag"
)

var File = flag.String("f", "config.json", "config file location")
var configGlb Config

func ReadConfig() {

	flag.Parse()
	file, err := os.Open(*File)
	if err != nil {
		fmt.Println("file error:", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	erro := decoder.Decode(&configGlb)
	if erro != nil {
		fmt.Println("error in decoding properties.json:", erro)
		os.Exit(1)
	}

	fmt.Println("configGlb:", configGlb)
	
}

func GetConfig() Config{
	return configGlb
}