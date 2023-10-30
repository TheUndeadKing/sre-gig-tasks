package generatevalue

import (
	"testing"
)

func TestFooer(t *testing.T) {

	result := GenerateValue()

	got := result
	want := 100

	if got <= want {
		t.Log("PASS")
	} else {
		t.Errorf("want %d, got %d", want, got)
	}

}
