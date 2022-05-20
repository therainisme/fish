package petpet

import (
	"github.com/dlclark/regexp2"
	"testing"
)

func TestReg(t *testing.T) {
	reg := regexp2.MustCompile(`(?<=\[CQ:at,qq=)\d*(?=\]\s\u6478)`, 0)

	m, _ := reg.FindStringMatch("[CQ:at,qq=1234567] 摸")
	if m != nil {
		gps := m.Groups()
		atQQ := gps[0].Captures[0].String()
		if atQQ != "1234567" {
			t.Failed()
		}
	} else {
		t.Failed()
	}
}
