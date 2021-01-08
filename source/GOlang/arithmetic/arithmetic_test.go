package main

import "testing"

var inputs = []struct {
	n1, n2 int
}{
	{
		n1: 10,
		n2: 20,
	},
	{
		n1: 30,
		n2: 20,
	},
	{
		n1: -45,
		n2: 2,
	},
	{
		n1: -20,
		n2: -10,
	},
}

func TestAdd(t *testing.T) {
	//actual := add(10, 20)
	//expected := 30
	for _, obj := range inputs {
		actual := add(obj.n1, obj.n2)
		expected := obj.n1 + obj.n2
		if actual != expected {
			t.Fatal("ERROR: Actual and Expected Values are different")
		} else {
			t.Log("SUCCESS: Test-cases are passed")
		}
	}
}

func TestSubtract(t *testing.T) {
	//actual := subtract(10, 20)
	//expected := -10

	for _, obj := range inputs {
		actual := subtract(obj.n1, obj.n2)
		expected := obj.n1 - obj.n2
		if actual != expected {
			t.Fatal("ERROR: Actual and Expected Values are different")
		} else {
			t.Log("SUCCESS: Test-cases are passed")
		}
	}
}

func TestMultiply(t *testing.T) {
	//actual := multiply(10, 20)
	//expected := 200

	for _, obj := range inputs {
		actual := multiply(obj.n1, obj.n2)
		expected := obj.n1 * obj.n2
		if actual != expected {
			t.Fatal("ERROR: Actual and Expected Values are different")
		} else {
			t.Log("SUCCESS: Test-cases are passed")
		}
	}
}

func TestDivide(t *testing.T) {
	//actual := divide(10, 20)
	//expected := 0

	for _, obj := range inputs {
		
		if (0 == obj.n2) {
			t.Log("PANIC: divide by zero exception occured")
		} else {
			actual := divide(obj.n1, obj.n2)
			expected := obj.n1 / obj.n2
			if actual != expected {
				t.Fatal("ERROR: Actual and Expected Values are different")
			} else {
				t.Log("SUCCESS: Test-cases are passed")
			}
		}
				
	}
}

func BenchmarkCalculateAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
		for _, obj := range inputs {
			add(obj.n1, obj.n2)			
		}

}
}

func BenchmarkCalculateSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
		for _, obj := range inputs {
			subtract(obj.n1, obj.n2)			
		}

}
}

func BenchmarkCalculateMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
		for _, obj := range inputs {
			multiply(obj.n1, obj.n2)			
		}

}
}

func BenchmarkCalculateDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
		for _, obj := range inputs {
			divide(obj.n1, obj.n2)			
		}

}
}

func BenchmarkCalculate(b *testing.B) {
	BenchmarkCalculateAdd(b)
	BenchmarkCalculateSubtract(b)
	BenchmarkCalculateMultiply(b)
	BenchmarkCalculateDivide(b)
}