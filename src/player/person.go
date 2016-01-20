package player

type Person struct {
	Name string
	Age  int8
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
