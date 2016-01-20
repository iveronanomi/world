package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"world/src/player"
)

var verbs []player.Verb

func getVerbs() error {
	var err error
	if len(verbs) == 0 {
		data, err := ioutil.ReadFile("./data/glossary_verbs.json")
		if err == nil {
			verbsData := []byte(data)
			err = json.Unmarshal(verbsData, &verbs)
		}
	}
	return err
}

func main() {
	if getVerbs() == nil {
		fmt.Printf("%#v \n", verbs)
	}
}
