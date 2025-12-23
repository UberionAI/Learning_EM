package math

import "testing"

//Задание №1
//func TestSum(t *testing.T) {
//	result := Sum(1, 2)
//	if result != 3 {
//		t.Errorf("Sum(1, 2) = %d; want 3", result)
//	}
//}

//PS C:\Users\tla\GolandProjects\Learning_EM\math> go test -v
//=== RUN   TestSum
//--- PASS: TestSum (0.00s)
//PASS
//ok      github.com/UberionAI/Learning_EM/math   0.128s

// Задание №2
func TestSumOnlyNatural(t *testing.T) {
	result := SumOnlyNatural(1, -2)
	if result == 0 {
		t.Errorf("SumOnlyNatural failed")
	}
}

//PS C:\Users\tla\GolandProjects\Learning_EM\math> go test -v
//=== RUN   TestSumOnlyNatural
//--- PASS: TestSumOnlyNatural (0.00s)
//PASS
//ok      github.com/UberionAI/Learning_EM/math   0.128s
//PS C:\Users\tla\GolandProjects\Learning_EM\math>
