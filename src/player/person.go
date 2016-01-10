package player

type Person struct {
	Name                    string
	Age                     int8
	PsychoemotionalFeatures PsychoemotionalFeatures
}

type PsychoemotionalFeatures struct {
	Openness          float32
	Conscientiousness float32
	Extraversion      float32
	Pliability        float32
	Neuroticism       float32
}
