package player

import (
	"testing"
)

func TestProcess(t *testing.T) {
	cases := []struct {
		applet       Applet
		expectResult Emotion
	}{
		{
			applet: Applet{
				Agent:       Person{Name: "Alice"},
				Recipient:   Person{Name: "Alice"},
				Action:      Verb{Name: "Kill", SocialStigma: -0.8, Attitude: -0.5},
				Possibility: 0.9,
			},
			expectResult: EMOTION_SHAME | EMOTION_DISTRESS,
		},
		{
			applet: Applet{
				Agent:       Person{Name: "Alice"},
				Recipient:   Person{Name: "Bob"},
				Action:      Verb{Name: "Search", SocialStigma: -0.8, Attitude: -0.5},
				Possibility: 0.9,
			},
			expectResult: EMOTION_ANGER | EMOTION_DISTRESS,
		},
	}
	for _, c := range cases {
		if p := c.applet.Process(); p != c.expectResult {
			t.Fail()
			t.Logf("{%d} Error: Emotion expected: %d, getted: %d\n", p, c.expectResult)
		} else {
			t.Logf("{%d} Succcess: Emotion expected: %d, getted: %d\n", p, c.expectResult)
		}
	}

}
