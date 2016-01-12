package player

type Bitmask uint16
type Emotion Bitmask

const (
	EMOTION_JOY Bitmask = 1
	EMOTION_HOPE Bitmask = 2
	EMOTION_DISAPPOINTMENT Bitmask = 4
	EMOTION_DISTRESS Bitmask = 8
	EMOTION_FEAR Bitmask = 16
	EMOTION_RELIEF Bitmask = 32
	EMOTION_PRIDE Bitmask = 64
	EMOTION_ADMIRATION Bitmask = 128
	EMOTION_SHAME Bitmask = 512
	EMOTION_ANGER Bitmask = 1024
)

type recipient interface{}

type agent interface {
	getAttitude(r recipient) float32
	equal(r recipient) bool
}

type Applet struct {
	Agent       agent
	Action      Verb
	Recipient   recipient
	Possibility float32
}

func equal(i1 interface{}, i2 interface{}) bool {
	return i1 == i2
}

func (p Person) getAttitude(r recipient) float32 {
	if equal(p, r) {
		return 1
	}
	return 0
}

func (p Person) equal(r recipient) bool {
	return equal(p, r)
}

type Verb struct {
	Name         string
	Attitude     float32
	SocialStigma float32
}

func (v Verb) getAttitude(r recipient) float32 {
	return v.Attitude
}

func (v Verb) equal(r recipient) bool {
	return equal(v, r)
}

type Noun struct {
	Name     string
	Attitude float32
}

func (n Noun) getAttitude(r recipient) float32 {
	return 0
}

func (n Noun) equal(r recipient) bool {
	return equal(n, r)
}

type processable interface {
	Process() Bitmask
}

func (a Applet) Process() Bitmask {
	var emotions Bitmask

	if a.Action.SocialStigma > 1 {
		if a.Agent.equal(a.Recipient) {
			emotions = EMOTION_PRIDE
		} else {
			emotions = EMOTION_ADMIRATION
		}
	} else if a.Action.SocialStigma < 1 {
		if a.Agent.equal(a.Recipient) {
			emotions = EMOTION_SHAME
		} else {
			emotions = EMOTION_ANGER
		}
	}

	if (a.Action.Attitude > 0 && a.Agent.getAttitude(a.Recipient) > 0 ||
	a.Action.Attitude < 0 && a.Agent.getAttitude(a.Recipient) < 0) {
		emotions |= EMOTION_JOY//|EMOTION_HOPE|EMOTION_DISAPPOINTMENT
	}

	if (a.Action.Attitude < 0 && a.Agent.getAttitude(a.Recipient) > 0 ||
	a.Action.Attitude > 0 && a.Agent.getAttitude(a.Recipient) < 0) {
		emotions |= EMOTION_DISTRESS//|EMOTION_FEAR|EMOTION_RELIEF
	}

	return emotions
}