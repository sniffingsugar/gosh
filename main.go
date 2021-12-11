package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"git.ucode.space/Phil/gosh/config"
	"git.ucode.space/Phil/gosh/utils"
)

type SUrl struct {
	CLI  bool   `json:"cli"`
	Surl string `json:"surl"`
}

type RUrl struct {
	URL string `json:"URL"`
}

func main() {

	utils.Check_Conf()

	if len(os.Args) > 1 {
		surl := SUrl{CLI: true, Surl: os.Args[1]}
		b, err := json.Marshal(surl)
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}
		resp, err := client.Post(config.GConfig.Server, "application/json", bytes.NewBuffer(b))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var eurl RUrl

		err = json.Unmarshal(body, &eurl)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(eurl.URL)

	} else {
		utils.Info()
	}

}
