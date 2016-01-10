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

type Verb struct {
	Name         string
	Attitude     float32
	SocialStigma float32
}

type Noun struct {
	Name     string
	Attitude float32
}

type Processable interface {
	process() Bitmask
}

type Recipient interface {
	get()
	getAttitude() float32
}

type Applet struct {
	Agent       Person
	Action      Verb
	Recipient   Recipient
	Possibility float32
}

func (p Person) get() Person {
	return p
}

func (p Person) getAttitude(p2 Person) float32 {
	return 0
}

func (v Verb) get() Verb {
	return v
}

func (v Verb) getAttitude(p2 Person) float32 {
	return v.Attitude
}

func (n *Noun) get() Noun {
	return n
}

func (n *Noun) getAttitude(p2 Person) float32 {
	return 0
}

func (a Applet) process() (Bitmask, error) {
	var emotions Bitmask
	var err error

	if a.Action.SocialStigma > 1 {
		if a.Agent == a.Recipient.get() {
			emotions = EMOTION_PRIDE
		} else {
			emotions = EMOTION_ADMIRATION
		}
	} else if a.Action.SocialStigma < 1 {
		if a.Agent == a.Recipient.get() {
			emotions = EMOTION_SHAME
		} else {
			emotions = EMOTION_ANGER
		}
	}

	if (a.Action.Attitude > 0 && a.Recipient.getAttitude(a.Agent) > 0 ||
	a.Action.Attitude < 0 && a.Recipient.getAttitude(a.Agent) < 0) {
		emotions |= EMOTION_JOY
	}
	if (a.Action.Attitude < 0 && a.Recipient.getAttitude(a.Agent) > 0 ||
	a.Action.Attitude > 0 && a.Recipient.getAttitude(a.Agent) < 0) {
		emotions |= EMOTION_ANGER
	}
	return emotions, err
}