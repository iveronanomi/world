package player

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
	Name         string
	Attitude     float32
	SocialStigma float32
}

type Noun struct {
	Name     string
	Attitude float32
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

func (p Person) getAttitude(r IRecipient) float32 {
	return 0
}

func (p Person) equal(r IRecipient) bool {
	cp, ok := r.(Person)
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