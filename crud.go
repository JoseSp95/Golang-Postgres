package main

import (
	//"log"
	"fmt"
	"log"
)

func createStudent()  {
	st := Student{
		Name:   "Victor",
		Age:    19,
		Active: true,
	}
	fmt.Println(st)
	err := Create(st)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creado exitosamente")
}

func getAllStudents(){
	listStudents, err := GetAll()
	if err != nil{
		panic(err)
	}

	for _, student := range listStudents{
		fmt.Println(student.Name)
	}

}

func getStudent(id int) {
	student, err := GetByID(id)
	if err != nil{
		fmt.Println("Error")
		log.Fatal(err)
	}
	fmt.Println(student)
}

func UpdateStudent()  {
	student := Student{
		ID:2,
		Name:"Sony",
		Age:30,
		Active:true}


	Update(student)
}


func main() {
	UpdateStudent()
}
