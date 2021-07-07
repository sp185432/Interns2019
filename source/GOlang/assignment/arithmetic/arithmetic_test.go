package arithmetic

import (
	"fmt"
	"testing"
)

var testcases = []struct {
	num1, num2 int
}{
	{num1: 10, num2: 15},
	{num1: 12, num2: 0},
	{num1: 22, num2: 03},
	{num1: 15, num2: 13},
}

func TestAdd(t *testing.T) {
	for _, tc := range testcases {

		actualAdd := Addition(tc.num1, tc.num2)
		expectedAdd := tc.num1 + tc.num2
		if actualAdd != expectedAdd {
			t.Fatal("\n------ADDITION------Different value is achieved ")
			t.Fatalf("\nActual: %d Expected: %d num1: %d num2: %d", actualAdd, expectedAdd, tc.num1, tc.num2)
		} else {
			t.Log("\n------ADDITION-----The result is obained as expected!")
			t.Logf("\nActual: %d Expected: %d num1: %d num2: %d", actualAdd, expectedAdd, tc.num1, tc.num2)
		}

		actualSub := Subtraction(tc.num1, tc.num2)
		expectedSub := tc.num1 - tc.num2
		if actualSub != expectedSub {
			t.Fatal("\n-----SUBTRACTION------Different value is achieved ")
			t.Fatalf("\nActual: %d Expected: %d num1: %d num2: %d", actualSub, expectedSub, tc.num1, tc.num2)
		} else {
			t.Log("\n-----SUBTRACTION------The result is obained as expected!")
			t.Logf("\nActual: %d Expected: %d num1: %d num2: %d", actualSub, expectedSub, tc.num1, tc.num2)
		}

		actualMul := Multiplication(tc.num1, tc.num2)
		expectedMul := tc.num1 * tc.num2
		if actualMul != expectedMul {
			t.Fatal("\n------MULTIPLICATION-----Different value is achieved ")
			t.Fatalf("\nActual: %d Expected: %d num1: %d num2: %d", actualMul, expectedMul, tc.num1, tc.num2)
		} else {
			t.Log("\n-----MULTIPLICATION------The result is obained as expected!")
			t.Logf("\nActual: %d Expected: %d num1: %d num2: %d", actualMul, expectedMul, tc.num1, tc.num2)
		}

		if tc.num2 == 0 {
			t.Log("\nDivide by zero is not possible. Denominator should be a non zero value")
		} else {
			actualDiv := Division(tc.num1, tc.num2)
			expectedDiv := tc.num1 / tc.num2
			if actualDiv != expectedDiv {
				t.Fatal("\n-----DIVISION-----Different value is achieved ")
				t.Fatalf("\nActual: %d Expected: %d num1: %d num2: %d", actualDiv, expectedDiv, tc.num1, tc.num2)
			} else {
				t.Log("\n-----DIVISION----The result is obained as expected!")
				t.Logf("\nActual: %d Expected: %d num1: %d num2: %d", actualDiv, expectedDiv, tc.num1, tc.num2)
			}
		}
	}
}

var funcs = []struct {
	name string
	f    func(int, int) int
}{
	{"Add", Addition},
	{"Sub", Subtraction},
	{"Mul", Multiplication},
	//{"Div", Division},
}

/*func Benchmark(b *testing.B) {
	for _, f := range funcs {
		//FOR MULTIPLE SIZES TO SEE HOW IT RUNS
		for n := 1; n <= 16; n *= 2 {
			b.Run(fmt.Sprintf("%s %d", f.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, tc := range testcases {
						f.f(tc.num1, tc.num2)
					}
				}
			})
		}

	}
}*/

//BENCHMARK PER FUNCTION
func BenchmarkArithmetic(b *testing.B) {
	for _, f := range funcs {
		b.Run(fmt.Sprintf("%s ", f.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tc := range testcases {
					f.f(tc.num1, tc.num2)
				}
			}
		})

	}
}
