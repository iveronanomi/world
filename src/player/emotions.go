package player

type Bitmask uint16

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

type Emotion struct {
	Name  string
	Value Bitmask
}