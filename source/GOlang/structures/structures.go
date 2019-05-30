package main

import "fmt"

type Student struct {
	roll_no, class_no int
}

type Person struct {
	id, age int
	name    string
	stud    Student
}

func main() {
	//person_obj := Person{1,20,"blah",Student{101,2}}
	//person_obj := Person{id:1,name:"blah",age:20,stud:Student{roll_no:101,class_no:2}}


	var person_obj Person
	num := 0
	fmt.Println("Enter number of persons: ")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Println("enter id, name, age of person-", i, ": ")
		fmt.Scanln(&person_obj.id, &person_obj.name, &person_obj.age)
	}

	for i := 0; i < num; i++ {
		fmt.Println("Details of person-", i, ": ")
		fmt.Println(person_obj.id, person_obj.name, person_obj.age)
	}

	for i := 0; i < num; i++ {
		fmt.Println("enter roll no., class no. of student-", i, ": ")
		fmt.Scanln(&person_obj.stud.roll_no, &person_obj.stud.class_no)
	}

	for i := 0; i < num; i++ {
		fmt.Println("Details of student-", i, ": ")
		fmt.Println(person_obj.stud.roll_no, person_obj.stud.class_no)
	}
}
