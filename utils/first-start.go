package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"git.ucode.space/Phil/gosh/config"
)

func Check_Conf() {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("No User home directory found")
	}

	jsonFile, err := os.Open(home + "/.config/gosh/config.json")

	if err != nil {
		var uri string
		fmt.Println("----- gosh CLI -----")
		fmt.Println("Welcome to gosh config generator")
		fmt.Print("Please enter your Server URI: ")
		_, err := fmt.Scanln(&uri)

		if err != nil {
			log.Fatal(err)
		}

		if !regexp.MustCompile(`^(http|https)://`).MatchString(uri) {
			fmt.Println("Invalid URI | Must start with http:// or https://")
			os.Exit(1)
		}

		CreateConfig(true, uri)

		fmt.Println("Config file created, you can now run gosh")
		os.Exit(0)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &config.GConfig)

	if err != nil {
		log.Fatal("No valid Json!")
	}
}
