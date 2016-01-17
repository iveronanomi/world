package main
import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"world/src/player"
)

var verbs []player.Verb

func getVerbs() (error) {
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
		fmt.Println(verbs)
	}

	testApplet := player.Applet{
		Agent:    player.Person{Name:"Dave"},
		Action: player.Verb{Name:"Kill", SocialStigma:-1, Attitude:-0.9},
		Recipient: player.Person{Name:"Dave"},
		Possibility: float32(0.9)}

	fmt.Println(testApplet.Process())
}