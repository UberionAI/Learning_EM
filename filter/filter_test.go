package filter

import (
	"reflect"
	"testing"
)

func TestFilterPositive(t *testing.T) {
	input := []int{1, 2, -1, -5, 4, 0, 4}
	expected := []int{1, 2, 4, 4}

	result := FilterPositive(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FilterPositive(%v) was incorrect, got: %v, want: %v.", input, result, expected)
	}
}

//PS C:\Users\tla\GolandProjects\Learning_EM\filter> go test -v .\...
//=== RUN   TestFilterPositive
//--- PASS: TestFilterPositive (0.00s)
//PASS
//ok      github.com/UberionAI/Learning_EM/filter (cached)
//PS C:\Users\tla\GolandProjects\Learning_EM\filter>
