package player

type Person struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
	PsychoEmotionalFeatures
	Emotions
}

type Emotions map[Emotion]float32

type PsychoEmotionalFeatures struct {
	Openness          float32
	Conscientiousness float32
	Extroversion      float32
	Pliability        float32
	Neuroticism       float32
}
