package main
import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"player"
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
//	fmt.Println(player.EMOTION_JOY | player.EMOTION_RELIEF | player.EMOTION_ANGER)
//	fmt.Println((((player.EMOTION_JOY | player.EMOTION_RELIEF | player.EMOTION_ANGER) & player.EMOTION_RELIEF) == player.EMOTION_RELIEF))
	if getVerbs() == nil {
		fmt.Println(verbs)
	}

}