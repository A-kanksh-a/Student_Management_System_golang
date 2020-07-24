package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var students = []student{}

	fmt.Println("\n Student Management System : ")
	fmt.Println("\n 1. ADD Student ")
	fmt.Println("\n 2. UPDATE Student")
	fmt.Println("\n 3. DELETE Student")
	fmt.Println("\n 4. GET Student")
	for {
		fmt.Println("       \n ENTER YOUR CHOICE: ")

		text := getString()
		switch text {

		case "1":
			fmt.Println("\n----------------Enter Details: ")
			fmt.Println("\nEnter Roll No")
			id := getString()
			fmt.Println("\nEnter Name")
			name := getString()
			fmt.Println("\nEnter Marks")
			marks := getString()
			newStudent := createStudent(id, name, marks)
			students = append(students, newStudent)
			fmt.Println("\n----------------DETAILS ADDED! ")
		case "2":
			fmt.Println("\n----------------Enter Id to UPDATE")
			id := getString()
			studentToUpdate, index := findByID(students, id)
			fmt.Println("\nEnter Name")
			name := getString()
			fmt.Println("\nEnter Marks")
			marks := getString()
			studentToUpdate.update(id, name, marks)
			students[index] = studentToUpdate
			fmt.Println("\n----------------DETAILS UPDATED! ")
		case "3":
			fmt.Println("\n------------------Enter Id to DELETE")
			id := getString()
			_, index := findByID(students, id)
			students = append(students[:index], students[index+1:]...)
			fmt.Println("----------------ENTRY DELETED! ")
		case "4":
			fmt.Println("\n----------------ALL STUDENT DETAILS: ")
			for index, value := range students {
				fmt.Println(string(index) + ": " + value.id + ", " + value.name + ", " + value.marks)
			}

		default:
			fmt.Println("WRONG INPUT")
		}
	}
}

func findByID(students []student, id string) (s student, index int) {
	for index, student := range students {
		if student.id == id {
			return student, index
		}
	}
	return
}

func getString() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
