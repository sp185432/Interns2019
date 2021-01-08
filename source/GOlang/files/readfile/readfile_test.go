package main

import "testing"

var names = []struct {
	name string
}{

	{
		name: "",
	},
	{
		name: ".txt",
	},
	{
		name: "file1.txt",
	},
	{
		name: "file1",
	},
	
}

var op_test = []struct {
	name string
}{

	{
		name: "read",
	},
	{
		name: "create",
	},
	{
		name: "",
	},
	{
		name: "write",
	},
	{
		name: "nil",
	},
	
}

func TestReadName(t *testing.T) {

	for _, ob := range names {
		name := "file1.txt"

		if ob.name != name {
			t.Log("ERROR: Wrong file name")
		} else {
			t.Log("SUCCESS: Test-cases are passed")
		}
	}
}

func TestCreateFile(t *testing.T) {
	actual := createFile("file5.txt")
	
		for _, ob := range op_test {	
			if ob.name != actual {
				t.Log("ERROR: Test-case failed")
			} else {
				t.Log("SUCCESS: Test-case passed")
			}
		}
}

func TestWriteFile(t *testing.T) {
	actual := writeFile("file5.txt")
	
		for _, ob := range op_test {	
			if ob.name != actual {
				t.Log("ERROR: Test-case failed")
			} else {
				t.Log("SUCCESS: Test-case passed")
			}
		}
}

func TestReadFile(t *testing.T) {
	actual := readFile("file5.txt")
	
		for _, ob := range op_test {	
			if ob.name != actual {
				t.Log("ERROR: Test-case failed")
			} else {
				t.Log("SUCCESS: Test-case passed")
			}
		}
}

func BenchmarkFileReadName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readName()
	}
}

func BenchmarkFileCreateFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createFile("file5.txt")
	}
}

func BenchmarkFileWriteFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeFile("file5.txt")
	}
}

func BenchmarkFileReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readFile("file5.txt")
	}
}
func BenchmarkFile(b *testing.B) {
	BenchmarkFileReadName(b)
	BenchmarkFileCreateFile(b)
	BenchmarkFileWriteFile(b)
	BenchmarkFileReadFile(b)
}