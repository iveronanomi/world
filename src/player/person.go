package player

type Person struct {
	Name					string
	Age						int8
	PsychoemotionalFeatures	PsychoemotionalFeatures
	Emotions				Emotions
}

type Emotions map[Emotion]float32

type PsychoemotionalFeatures struct {
	Openness			float32
	Conscientiousness	float32
	Extraversion		float32
	Pliability			float32
	Neuroticism			float32
}