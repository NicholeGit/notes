package observer

import (
	"math/rand"
	"testing"
)

func TestObserver(t *testing.T) {

	random := NewRandomNumberGenerator()

	o1 := &DigitObserver{rand.Intn(50)}
	o2 := &DigitObserver{rand.Intn(50)}

	random.AddObserver(o1)
	random.AddObserver(o2)

	result := random.Execute()

	for _, r := range result {
		if len(result) != 2 && r >= 50 {
			t.Errorf("Expect result to equal random int array")
		}
	}
}
