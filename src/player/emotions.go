package player
import "reflect"

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

type Recipient interface{}

type Agent interface {
	getAttitude(r Recipient) float32
	equal(r Recipient) bool
}

type Applet struct {
	Agent       Agent
	Action      Verb
	Recipient   Recipient
	Possibility float32
}

func (p Person) getAttitude(r Recipient) float32 {
	if equal(p, r) {
		return 1
	}
	return 0
}

func (p Person) equal(r Recipient) bool {
	return equal(p, r)
}

type Verb struct {
	Name         string
	Attitude     float32
	SocialStigma float32
}

func (v Verb) getAttitude(r Recipient) float32 {
	return v.Attitude
}

func (v Verb) equal(r Recipient) bool {
	return equal(v, r)
}

type Noun struct {
	Name     string
	Attitude float32
}

func (n Noun) getAttitude(r Recipient) float32 {
	return float32(0)
}

func (n Noun) equal(r Recipient) bool {
	return equal(n, r)
}

type processable interface {
	Process() Bitmask
}

func equal(a interface{}, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func (a Applet) Process() Emotion {
	var emo Emotion

	if a.Action.SocialStigma > 1 {
		if a.Agent.equal(a.Recipient) {
			emo = EMOTION_PRIDE
		} else {
			emo = EMOTION_ADMIRATION
		}
	} else if a.Action.SocialStigma < 1 {
		if a.Agent.equal(a.Recipient) {
			emo = EMOTION_SHAME
		} else {
			emo = EMOTION_ANGER
		}
	}

	if (a.Action.Attitude > 0 && a.Agent.getAttitude(a.Recipient) > 0 ||
	a.Action.Attitude < 0 && a.Agent.getAttitude(a.Recipient) < 0) {
		emo |= EMOTION_JOY//|EMOTION_HOPE|EMOTION_DISAPPOINTMENT
	}

	if (a.Action.Attitude < 0 && a.Agent.getAttitude(a.Recipient) > 0 ||
	a.Action.Attitude > 0 && a.Agent.getAttitude(a.Recipient) < 0) {
		emo |= EMOTION_DISTRESS//|EMOTION_FEAR|EMOTION_RELIEF
	}

	return emo
}