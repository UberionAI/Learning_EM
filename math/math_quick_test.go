package math

import (
	"testing"
	"testing/quick"
)

func TestAbsProperties(t *testing.T) {
	config := &quick.Config{
		MaxCount: 500,
		Rand:     nil,
	}

	f1 := func(x int) bool {
		return Abs(x) >= 0
	}

	if err := quick.Check(f1, config); err != nil {
		t.Errorf("Abs(x) must be >= 0: %v", err)
	}

	f2 := func(x int) bool {
		return Abs(x) == Abs(-x)
	}

	if err := quick.Check(f2, config); err != nil {
		t.Errorf("Abs(-x) must be = Abs(x): %v", err)
	}

	f3 := func(x int) bool {
		return Abs(Abs(x)) == Abs(x)
	}

	if err := quick.Check(f3, config); err != nil {
		t.Errorf("x * Abs(x) must be >= 0: %v", err)
	}
}

//PS C:\Users\tla\GolandProjects\Learning_EM\math> go test -v -quickchecks=1000 ./...
//=== RUN   TestAbsProperties
//--- PASS: TestAbsProperties (0.00s)
//PASS
//ok      github.com/UberionAI/Learning_EM/math   0.123s
//PS C:\Users\tla\GolandProjects\Learning_EM\math>
