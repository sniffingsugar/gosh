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
)

type SUrl struct {
	CLI  bool   `json:"cli"`
	Surl string `json:"surl"`
}

type RUrl struct {
	URL string `json:"URL"`
}

func main() {

	if len(os.Args) > 1 {
		surl := SUrl{CLI: true, Surl: os.Args[1]}
		b, err := json.Marshal(surl)
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}
		resp, err := client.Post(config.URL+"/", "application/json", bytes.NewBuffer(b))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var eurl RUrl

		json.Unmarshal(body, &eurl)

		fmt.Println(eurl.URL)

	} else {
		fmt.Println("------ GLY.ONE CLI /v0.0.1 ------")
		fmt.Println("--- Usage:")
		fmt.Println("gly <url>  // Shorten a URL")
		fmt.Println("---------------------------------")
	}

}
