package player

import (
	"testing"
	"fmt"
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
				Possibility:0.9,
			},
			expectResult: EMOTION_SHAME | EMOTION_DISTRESS,
		},
		{
			applet : Applet{
				Agent:Person{Name: "Alice"},
				Recipient:Person{Name: "Bob"},
				Action: Verb{Name: "Search", SocialStigma: -0.8, Attitude: -0.5},
				Possibility:0.9,
			},
			expectResult: EMOTION_ANGER | EMOTION_DISTRESS,
		},
	}
	for _, c := range cases {
		fmt.Println(c.applet.Process())
		if c.applet.Process() != c.expectResult {
			t.Fail()
		}
	}

}