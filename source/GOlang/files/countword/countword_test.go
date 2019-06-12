package main

import "testing"

var op_test = []struct {
	name string
}{

	{
		name: "yes",
	},
	{
		name: "no",
	},
	{
		name: "abc",
	},
	{
		name: "exist",
	},
}

var count_test = []struct {
	num int
}{

	{
		num: 1,
	},
	{
		num: 20,
	},
	{
		num: 0,
	},
	{
		num: 4,
	},
}

func TestFile_Exist(t *testing.T) {

	actual := file_exist("file1.txt")
	
		for _, ob := range op_test {	
			if ob.name != actual {
				t.Log("ERROR: Test-case failed")
			} else {
				t.Log("SUCCESS: Test-case passed")
			}
		}
}
func TestMy_Rename(t *testing.T) {

	actual := readFile("file1.txt")
	
		for _, ob := range count_test {	
			if ob.num != actual {
				t.Log("ERROR: Test-case failed")
			} else {
				t.Log("SUCCESS: Test-case passed")
			}
		}
}

func BenchmarkFileFile_Exist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file_exist("file1.txt")
	}
}

func BenchmarkFileReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readFile("file1.txt")
	}
}

func BenchmarkFile(b *testing.B) {
	BenchmarkFileFile_Exist(b)
	BenchmarkFileReadFile(b)
}