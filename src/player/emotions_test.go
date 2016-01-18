package player

import (
	"testing"
)

func TestProcess (t *testing.T) {
	cases := []struct {
		applet			Applet
		expectResult	Emotion
	}{
		{
			applet : Applet{
				Agent:Person{Name: "Alice"},
				Recipient:Person{Name: "Alice"},
				Action: Verb{Name: "Kill", SocialStigma: -0.8, Attitude: -0.5},
				Possibility:float32(0.9),
			},
			expectResult: EMOTION_JOY,
		},
		{
			applet : Applet{
				Agent:Person{Name: "Alice"},
				Recipient:Person{Name: "Bob"},
				Action: Verb{Name: "Kill", SocialStigma: -0.8, Attitude: -0.5},
				Possibility:float32(0.9),
			},
			expectResult: EMOTION_ADMIRATION|EMOTION_HOPE,
		},
	}
	for _, c := range cases  {
		if c.applet.Process() != c.expectResult {
			t.Fail()
		}
	}

}