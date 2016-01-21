package unit

/* всё то, что способно передвигаться по собственному желанию,
будь то тролль с одной условной извилиной или высокоинтеллектуальная система "дворф".*/
type Unit struct {
	Id   int    `json:"id"`
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
