package main

import "testing"

var op_test = []struct {
	name string
}{

	{
		name: "read",
	},
	{
		name: "update",
	},
	{
		name: "abc",
	},
	{
		name: "exist",
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
func TestUpdate(t *testing.T) {

	actual := update("file1.txt")
	
		for _, ob := range op_test {	
			if ob.name != actual {
				t.Log("ERROR: Test-case failed")
			} else {
				t.Log("SUCCESS: Test-case passed")
			}
		}
}

func TestRead(t *testing.T) {

	actual := read("file1.txt")
	
		for _, ob := range op_test {	
			if ob.name != actual {
				t.Log("ERROR: Test-case failed")
			} else {
				t.Log("SUCCESS: Test-cases passed")
			}
		}
}

func BenchmarkFileFile_Exist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file_exist("file1.txt")
	}
}
func BenchmarkFileUpdate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		update("file1.txt")
	}
}
func BenchmarkFileRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read("file1.txt")
	}
}

// func BenchmarkFileMy_Rename(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		my_rename("file1.txt","")
// 	}
// }

func BenchmarkFile(b *testing.B) {
	BenchmarkFileFile_Exist(b)
	BenchmarkFileUpdate(b)
	BenchmarkFileRead(b)
	//BenchmarkFileMy_Rename(b)
}