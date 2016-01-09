package main
import (
	"fmt"
	"encoding/json"
	"player"
)
const (
	dataGlossaryNouns = `[
		{"Name": "Troll", "attitude": 0.2},
		{"Name": "Apple", "attitude": 0.3},
		{"Name": "Sword", "attitude": 0.5}
	]`
	dataGlossaryVerbs = `[
		{"Name": "Kill", "attitude": 0.2, "socialStigma": 0.3},
		{"Name": "Getting skill", "attitude": 0.5, "socialStigma": 0.3},
		{"Name": "Eating", "attitude": 0.8, "socialStigma": 0.2}
	]`
	dataAgents = `[
		{"Name": "Bob", "Age": 20}
	]`
	dataEmotions = `[
		{"Name":"Joy", "Value":0},
		{"Name":"Hope", "Value":0},
		{"Name":"Disappointment", "Value":0},
		{"Name":"Distress", "Value":0},
		{"Name":"Fear", "Value":0},
		{"Name":"Relief", "Value":0},
		{"Name":"Shame", "Value":0},
		{"Name":"Anger", "Value":0}
	]`
)

type Verb struct {
	Name         string
	Attitude     float32
	SocialStigma float32
}

type Noun struct {
	Name     string
	Attitude float32
}

type Person struct {
	Name string
	Age  int8
}

/*type emotionApplet struct {
	Agent Agent
	Action Action
	Recipient Recipient
	Probability float64
}*/

/*func createAgent(data []byte) (Person) {
	return Person{}
}*/

func main() {
	fmt.Println(player.EMOTION_JOY | player.EMOTION_RELIEF | player.EMOTION_ANGER)
	fmt.Println((((player.EMOTION_JOY | player.EMOTION_RELIEF | player.EMOTION_ANGER) & player.EMOTION_RELIEF) == player.EMOTION_RELIEF))

	var verbs []Verb

	verbsData := []byte(dataGlossaryVerbs)
	err := json.Unmarshal(verbsData, &verbs)
	if err != nil {
		fmt.Print("Error:", err)
	}
	
	fmt.Println(verbs)

	var nouns []Noun

	nounData := []byte(dataGlossaryNouns)
	err = json.Unmarshal(nounData, &nouns)
	if err != nil {
		fmt.Print("Error:", err)
	}

	fmt.Println(nouns)

	var persons []Person
	agentsData := []byte(dataAgents)
	err = json.Unmarshal(agentsData, &persons)
	fmt.Println(persons)

	var emotions []player.Emotion
	emotionsData := []byte(dataEmotions)
	err = json.Unmarshal(emotionsData, &emotions)
	fmt.Println(emotions)

}