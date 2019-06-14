package creates

import (
	"testing"
)

func TestNewSingleton(t *testing.T) {
	s := NewSingleton()
	s["s"] = "a"
	a := NewSingleton()
	if a["s"] != "a" {
		t.Errorf("%s", "!=")
	}
	t.Log("==")
}