package main

import (
	"fmt"
	"student/student"
	"time"
)

func checkforerrors(err string) {
	if err == "" {
		fmt.Println("Created Student successfully")
	} else {
		fmt.Println(err)
	}
}
func main() {
	// Create some students
	student1, err1 := student.NewStudent("Gloria", "Ramirez", time.Date(1996, 5, 15, 0, 0, 0, 0, time.UTC), []float64{8.5, 8.7, 9.0}, 2018, 2022)
	checkforerrors(err1)

	student2, err2 := student.NewStudent("Manny", "Delgado", time.Date(2002, 06, 19, 0, 0, 0, 0, time.UTC), []float64{7.5, 7.8, 8.0}, 2017, 2021)
	checkforerrors(err2)

	student3, err3 := student.NewStudent("Phil", "Dunphy", time.Date(1998, 11, 15, 0, 0, 0, 0, time.UTC), []float64{7.5, 8.0, 8.2}, 2016, 2020)
	checkforerrors(err3)

	fmt.Println("--------------------------------------------------------------------------------------------")

	student1.PrintDetails()
	student2.PrintDetails()
	student3.PrintDetails()

	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("READ ALL STUDENTS\n\n")
	allStudents := student.ReadAllStudents()
	fmt.Println("All Students:")
	for _, s := range allStudents {
		// fmt.Printf("%+v\n", s)
		s.PrintDetails()
	}

	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("READ STUDENT BY ROLL NO\n\n")

	rollNoToRead := 2
	studentFound, err := student.ReadStudentByRollNo(rollNoToRead)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("Student with Roll No", rollNoToRead, ":")
		studentFound.PrintDetails()
	}

	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("UPDATE STUDENT BY ROLL NO\n\n")

	rollNoToUpdate := 1
	studentToBeUpdated, err := student.ReadStudentByRollNo(rollNoToUpdate)

	if err == "" {
		err = studentToBeUpdated.UpdateStudentByParameter("firstname", "Sonia")
		if err != "" {
			fmt.Println(err)

		}

		err = studentToBeUpdated.UpdateStudentByParameter("yearOfPassing", 2025)
		if err != "" {
			fmt.Println(err)

		}

		fmt.Println("Updated Student", rollNoToUpdate)
		studentToBeUpdated.PrintDetails()

	} else {
		fmt.Println(err)
	}

	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("DELETE STUDENT BY ROLL NO\n\n")

	rollNoToDelete := 3
	err = student.DeleteStudentByRollNo(rollNoToDelete)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted Student with Roll No ", rollNoToDelete)
	}

	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("READ ALL STUDENTS AGAIN\n\n")

	allStudents = student.ReadAllStudents()
	for _, s := range allStudents {
		s.PrintDetails()
	}
}
