package comparevalue

import (
	"testing"
)

func TestFooer(t *testing.T) {

	want := []int{1, 100, 69, 50, 90}

	if want[1] < 100 && want[1] > 50 {

		if want[1]%2 == 0 {
			t.Log("PASS: It's closer to 100 and it's even!")
		}
	}

	if want[2] < 100 && want[2] > 50 {

		if want[2]%2 != 0 {
			t.Log("PASS: It's closer to 100")
		}
	}

	if want[3] == 50 {
		t.Log("PASS: It's 50!")
	}
	if want[0] < 50 {
		t.Log("PASS: It's closer to 0")
	}

}
