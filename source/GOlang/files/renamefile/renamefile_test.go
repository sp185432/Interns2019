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

	actual := my_rename("file1.txt","file2.txt")
	
		for _, ob := range op_test {	
			if ob.name != actual {
				t.Log("ERROR: Test-case failed")
			} else {
				t.Log("SUCCESS: Test-case passed")
			}
		}
}

func BenchmarkFileFile_Exist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file_exist("file2.txt")
	}
}

// func BenchmarkFileMy_Rename(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		my_rename("file1.txt","")
// 	}
// }

func BenchmarkFile(b *testing.B) {
	BenchmarkFileFile_Exist(b)
	//BenchmarkFileMy_Rename(b)
}