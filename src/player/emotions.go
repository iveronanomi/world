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

func (p Person) getAttitude(r recipient) float32 {
	if p == r {
		return 1
	}
	return 0
}

func (p Person) equal(r recipient) bool {
	return true
}

type Verb struct {
	Name         string
	Attitude     float32
	SocialStigma float32
}

func (v Verb) getAttitude(r recipient) float32 {
	return v.Attitude
}

func (p Verb) equal(r recipient) bool {
	return true
}

type Noun struct {
	Name     string
	Attitude float32
}

func (n Noun) getAttitude(r recipient) float32 {
	return 0
}

func (p Noun) equal(r recipient) bool {
	return true
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
		emotions |= EMOTION_JOY
	}

	if (a.Action.Attitude < 0 && a.Agent.getAttitude(a.Recipient) > 0 ||
	a.Action.Attitude > 0 && a.Agent.getAttitude(a.Recipient) < 0) {
		emotions |= EMOTION_ANGER
	}

	return emotions
}