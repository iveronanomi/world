package unit

//import "math"

type Bitmask uint16
type Emotion Bitmask

const (
	EMOTION_JOY Emotion = 1 << iota
	EMOTION_HOPE
	EMOTION_DISAPPOINTMENT
	EMOTION_DISTRESS
	EMOTION_FEAR
	EMOTION_RELIEF
	EMOTION_PRIDE
	EMOTION_ADMIRATION
	EMOTION_SHAME
	EMOTION_ANGER
)

type IRecipient interface{}

type IAgent interface {
	getAttitude(r IRecipient) float32
	equal(r IRecipient) bool
}

type Verb struct {
	Name         string  `json:"name"`
	Attitude     float32 `json:"attitude"`
	SocialStigma float32 `json:"social_stigma"`
}

type Noun struct {
	Name     string  `json:"name"`
	Attitude float32 `json:"attitude"`
}

type IProcessable interface {
	Process() Bitmask
}

type Applet struct {
	Agent       IAgent
	Recipient   IRecipient
	Action      Verb
	Possibility float32
}

func (p Unit) getAttitude(r IRecipient) float32 {
	return float32(1)
}

func (p Unit) equal(r IRecipient) bool {
	cp, ok := r.(Unit)
	if ok {
		return cp.Name == p.Name
	}
	return false
}

func (v Verb) getAttitude(r IRecipient) float32 {
	return v.Attitude
}

func (n Noun) getAttitude(r IRecipient) float32 {
	return float32(0)
}

func (n Noun) equal(r IRecipient) bool {
	cp, ok := r.(Noun)
	if ok {
		return cp.Name == n.Name
	}
	return false
}

//func (a Applet) GetEmotionIntensity(e Emotion) float64 {
//	if ((EMOTION_JOY|EMOTION_DISAPPOINTMENT|EMOTION_DISTRESS|EMOTION_RELIEF)&e == e) {
//		return (math.Abs(float64(a.Action)) + math.Abs(float64(a.Recipient)))/2
//	}
//	if ((EMOTION_HOPE|EMOTION_FEAR)&e == e) {
//		return (math.Abs(float64(a.Action)) + math.Abs(float64(a.Recipient)))/2 * a.Possibility
//	}
//	if ((EMOTION_PRIDE|EMOTION_ADMIRATION|EMOTION_SHAME|EMOTION_ANGER)&e == e) {
//		return math.Abs(a.Action.SocialStigma)
//	}
//	return float64(0)
//}

func (a Applet) Process() Emotion {
	var e Emotion

	if a.Action.SocialStigma > 1 {
		if a.Agent.equal(a.Recipient) {
			e = EMOTION_PRIDE
		} else {
			e = EMOTION_ADMIRATION
		}
	} else if a.Action.SocialStigma < 1 {
		if a.Agent.equal(a.Recipient) {
			e = EMOTION_SHAME
		} else {
			e = EMOTION_ANGER
		}
	}

	if a.Action.Attitude > 0 && a.Agent.getAttitude(a.Recipient) > 0 ||
		a.Action.Attitude < 0 && a.Agent.getAttitude(a.Recipient) < 0 {
		e |= EMOTION_JOY //|EMOTION_HOPE|EMOTION_DISAPPOINTMENT
	}

	if a.Action.Attitude < 0 && a.Agent.getAttitude(a.Recipient) > 0 ||
		a.Action.Attitude > 0 && a.Agent.getAttitude(a.Recipient) < 0 {
		e |= EMOTION_DISTRESS //|EMOTION_FEAR|EMOTION_RELIEF
	}

	return e
}
