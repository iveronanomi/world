package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"world/src/unit"
)

var verbs []unit.Verb
var nouns []unit.Noun
var units []unit.Unit

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

func getNouns() error {
	var err error
	if len(nouns) == 0 {
		data, err := ioutil.ReadFile("./data/glossary_nouns.json")
		if err == nil {
			nounsData := []byte(data)
			err = json.Unmarshal(nounsData, &nouns)
		}
	}
	return err
}

func getUnits() error {
	var err error
	if len(units) == 0 {
		data, err := ioutil.ReadFile("./data/units.json")
		if err == nil {
			unitsData := []byte(data)
			err = json.Unmarshal(unitsData, &units)
		}
	}
	return err
}

func main() {
	if getVerbs() == nil {
		fmt.Printf("%#v \n\n", verbs)
	}

	if getNouns() == nil {
		fmt.Printf("%#v \n\n", nouns)
	}

	if getUnits() == nil {
		fmt.Printf("%#v \n\n", units)
	}
}
