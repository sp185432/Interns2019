package arithmetic

import "testing"

/*
------BASIC TEST-------
func TestAddition(t *testing.T) {

	actual := Addition(10, 20)
	expected := 10 + 20
	if actual != expected {
		t.Fatal("Different value is achieved ")
		t.Fatalf("Actual: %d Expected: %d", actual, expected)
	} else {
		t.Log("The result is obained as expected!")
		t.Logf("Actual: %d Expected: %d", actual, expected)
	}
}*/

var testcases = []struct {
	num1, num2 int
}{
	{num1: 12, num2: 0},
	{num1: 22, num2: 03},
	{num1: 15, num2: 13},
}

//------TABLE TEST FOR ADDITION--------
func TestAddition(t *testing.T) {
	for _, tc := range testcases {
		actual := Addition(tc.num1, tc.num2)
		expected := tc.num1 + tc.num2
		if actual != expected {
			t.Fatal("\nDifferent value is achieved ")
			t.Fatalf("\nActual: %d Expected: %d num1: %d num2: %d", actual, expected, tc.num1, tc.num2)
		} else {
			t.Log("\nThe result is obained as expected!")
			t.Logf("\nActual: %d Expected: %d num1: %d num2: %d", actual, expected, tc.num1, tc.num2)
		}

	}
}

//------TABLE TEST FOR SUBTRACTION--------
func TestSubtraction(t *testing.T) {
	for _, tc := range testcases {
		actual := Subtraction(tc.num1, tc.num2)
		expected := tc.num1 - tc.num2
		if actual != expected {
			t.Fatal("\nDifferent value is achieved ")
			t.Fatalf("\nActual: %d Expected: %d num1: %d num2: %d", actual, expected, tc.num1, tc.num2)
		} else {
			t.Log("\nThe result is obained as expected!")
			t.Logf("\nActual: %d Expected: %d num1: %d num2: %d", actual, expected, tc.num1, tc.num2)
		}

	}
}

//------TABLE TEST FOR MULTIPLICATION----------
func TestMultiplication(t *testing.T) {
	for _, tc := range testcases {
		actual := Multiplication(tc.num1, tc.num2)
		expected := tc.num1 * tc.num2
		if actual != expected {
			t.Fatal("\nDifferent value is achieved ")
			t.Fatalf("\nActual: %d Expected: %d num1: %d num2: %d", actual, expected, tc.num1, tc.num2)
		} else {
			t.Log("\nThe result is obained as expected!")
			t.Logf("\nActual: %d Expected: %d num1: %d num2: %d", actual, expected, tc.num1, tc.num2)
		}

	}
}

//------TABLE TEST FOR DIVISION---------
func TestDivision(t *testing.T) {
	for _, tc := range testcases {

		if tc.num2 == 0 {
			t.Log("\nDivide by zero is not possible. Denominator should be a non zero value")
		} else {
			actual := Division(tc.num1, tc.num2)
			expected := tc.num1 / tc.num2
			if actual != expected {
				t.Fatal("\nDifferent value is achieved ")
				t.Fatalf("\nActual: %d Expected: %d num1: %d num2: %d", actual, expected, tc.num1, tc.num2)
			} else {
				t.Log("\nThe result is obained as expected!")
				t.Logf("\nActual: %d Expected: %d num1: %d num2: %d", actual, expected, tc.num1, tc.num2)
			}
		}

	}
}

//------BENCHMARKING------
func BenchmarkAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Addition(12, 15)
	}
}

func BenchmarkSubtraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtraction(7, 3)
	}
}

func BenchmarkMultiplication(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplication(12, 7)
	}
}

func BenchmarkDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(12, 2)
	}
}
