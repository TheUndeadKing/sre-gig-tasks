package comparevalue

import (
	generatevalue "tasktwo/generateValue"
	"testing"
)

func TestFooer(t *testing.T) {

	value := generatevalue.GenerateValue()

	if value < 100 && value > 50 { // Modify the conditional in the code you previously wrote to check not only if the  number is higher than 50, but also if it's even. If it's even and higher than 50, print,"It's closer to 100, and it's even!"

		if value%2 == 0 {
			t.Log("It's closer to 100 and it's even!")
		} else {
			t.Log("It's closer to 100")
		}
	} else if value == 50 { //  Modify the previous code to print "It's 50!" if the random number is 50
		t.Log("It's 50!")
	} else if value < 50 {
		t.Log("It's closer to 0")
	}
}
