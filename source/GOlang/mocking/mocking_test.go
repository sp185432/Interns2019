package main

import "testing"


func TestMystat(t *testing.T) {

		actual := mystat("https://google.com")
		expected := "200 OK"

		if actual != expected {
			t.Fatal("ERROR: Actual and Expected Values are different")
		} else {
			t.Log("SUCCESS 1: Test-cases are passed")
		}


		actual = mystat("http://localhost:8080/blah")
		expected = "404 Not Found"

		if actual != expected {
			t.Fatal("ERROR: Actual and Expected Values are different")
		} else {
			t.Log("SUCCESS 2: Test-cases are passed")
		}
}