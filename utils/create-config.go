package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"git.ucode.space/Phil/gosh/config"
)

func CreateConfig(cr bool, uri string) {
	config := config.Config{
		Created: cr,
		Server:  uri,
	}

	filepayload, err := json.MarshalIndent(config, "", " ")

	if err != nil {
		log.Fatal(err.Error())
	}

	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("No User home directory found")
	}

	if _, err := os.Stat(home + ("/.config/gosh")); os.IsNotExist(err) {
		err := os.Mkdir(home+("/.config/gosh"), 0777)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	err = ioutil.WriteFile(home+("/.config/gosh/config.json"), filepayload, 0644)

	if err != nil {
		log.Fatal(err.Error())
	}

}
