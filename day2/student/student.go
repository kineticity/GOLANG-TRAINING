package student

import (
	"time"
	"fmt"
)

var rollNumberCounter = 0 //roll number which increments for each new Student

var students = []*Student{}

type Student struct {
	RollNo                 int       
	Firstname              string
	Lastname               string
	Fullname               string    
	DOB                    time.Time
	Age                    int       
	SemesterCGPAArray      []float64
	FinalCGPA              float64   
	SemesterGrades         []string  
	FinalGrade             string    
	YearOfEnrollment       int
	YearOfPassing          int
	NumberOfYearsToGraduate int       
}

//CREATE
func NewStudent(firstname, lastname string, dob time.Time, semesterCGPAArray []float64, yearOfEnrollment int, yearOfPassing int) (*Student, string) {
	err:=validateStudentInput(firstname, lastname, dob, semesterCGPAArray, yearOfEnrollment, yearOfPassing)
	if  err!= "" {
		return nil, err
	}

	rollNumberCounter++ 

	student := &Student{
		RollNo:                 rollNumberCounter,
		Firstname:              firstname,
		Lastname:               lastname,
		Fullname:               getFullname(firstname, lastname),
		DOB:                    dob,
		Age:                    calculateAge(dob),
		SemesterCGPAArray:      semesterCGPAArray,
		FinalCGPA:              calculateFinalCGPA(semesterCGPAArray),
		SemesterGrades:         calculateSemesterGrades(semesterCGPAArray),
		FinalGrade:             getFinalGrade(calculateFinalCGPA(semesterCGPAArray)),
		YearOfEnrollment:       yearOfEnrollment,
		YearOfPassing:          yearOfPassing,
		NumberOfYearsToGraduate: calculateNumberOfYearsToGraduate(yearOfPassing, yearOfEnrollment),
	}

	students = append(students, student)

	return student, ""
}

func validateStudentInput(firstname, lastname string, dob time.Time, semesterCGPAArray []float64, yearOfEnrollment, yearOfPassing int) string {
	if firstname == "" || lastname == "" {
		return "first name and last name cannot be empty"
	}

	if dob.After(time.Now()) {
		return "invalid date of birth"
	}

	for _, cgpa := range semesterCGPAArray {
		if cgpa < 0.0 || cgpa > 10.0 {
			return "CGPA must be between 0 and 10"
		}
	}

	if yearOfEnrollment==0 || yearOfEnrollment > time.Now().Year() {
		return "invalid year of enrollment"
	}

	if yearOfPassing == 0 || yearOfPassing < yearOfEnrollment {
		return "invalid year of passing"
	}

	return ""
}


//READ ALL
func ReadAllStudents() []*Student {
	return students
}


//READ BY ROLLNO
func ReadStudentByRollNo(rollNo int) (*Student, string) {
	for _, student := range students {
		if student.RollNo == rollNo {
			return student, ""
		}
	}
	return nil, "This student does not exist"
}

//UPDATE BY ROLLNO
func (s *Student) UpdateStudent(firstname, lastname string, dob time.Time, semesterCGPAArray []float64, yearOfEnrollment, yearOfPassing int) string {

	err:=validateStudentInput(firstname, lastname, dob, semesterCGPAArray, yearOfEnrollment, yearOfPassing)
	if  err!= "" {
		return err
	}

	s.Firstname = firstname
	s.Lastname = lastname
	s.Fullname = getFullname(firstname, lastname)
	s.DOB = dob
	s.Age = calculateAge(dob)
	s.SemesterCGPAArray = semesterCGPAArray
	s.FinalCGPA = calculateFinalCGPA(semesterCGPAArray)
	s.SemesterGrades = calculateSemesterGrades(semesterCGPAArray)
	s.FinalGrade = getFinalGrade(s.FinalCGPA)
	s.YearOfEnrollment = yearOfEnrollment
	s.YearOfPassing = yearOfPassing
	s.NumberOfYearsToGraduate = calculateNumberOfYearsToGraduate(yearOfPassing, yearOfEnrollment)

	return ""
}


//DELETE BY ROLLNO
func DeleteStudentByRollNo(rollNo int) string {
	for i, student := range students {
		if student.RollNo == rollNo {
			students = append(students[:i], students[i+1:]...) 
			return ""
		}
	}
	return "This student does not exist"
}

func getFullname(firstname, lastname string) string {
	return firstname + " " + lastname
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.Before(dob.AddDate(age, 0, 0)) {
		age--
	}
	return age
}

func calculateFinalCGPA(semesterCGPAArray []float64) float64 {
	total := 0.0
	for _, cgpa := range semesterCGPAArray {
		total += cgpa
	}
	if len(semesterCGPAArray) == 0 {
		return 0
	}
	return total / float64(len(semesterCGPAArray))
}

func calculateSemesterGrades(semesterCGPAArray []float64) []string {
	grades := []string{}
	for _, cgpa := range semesterCGPAArray {
		grades = append(grades, cgpaToGrade(cgpa))
	}
	return grades
}

func getFinalGrade(finalCGPA float64) string {
	return cgpaToGrade(finalCGPA)
}

func calculateNumberOfYearsToGraduate(yearOfPassing, yearOfEnrollment int) int {
	if yearOfPassing > 0 {
		return yearOfPassing - yearOfEnrollment
	}
	return 0
}

func cgpaToGrade(cgpa float64) string {
	if cgpa >= 9.0 {
		return "A"
	} else if cgpa >= 8.0 {
		return "B"
	} else if cgpa >= 7.0 {
		return "C"
	} else if cgpa >= 6.0 {
		return "D"
	}
	return "F"
}

func (s *Student) PrintDetails() {
	if s==nil{
		fmt.Println("Student object was not created")
	} else{
		fmt.Println("Roll No:", s.RollNo)
		fmt.Println("First Name:", s.Firstname)
		fmt.Println("Last Name:", s.Lastname)
		fmt.Println("Full Name:", s.Fullname)
		fmt.Println("Date of Birth:", s.DOB.Format("2006-01-02")) 
		fmt.Println("Age:", s.Age)
		fmt.Println("Semester CGPA Array:", s.SemesterCGPAArray)
		fmt.Println("Final CGPA:", s.FinalCGPA)
		fmt.Println("Semester Grades:", s.SemesterGrades)
		fmt.Println("Final Grade:", s.FinalGrade)
		fmt.Println("Year of Enrollment:", s.YearOfEnrollment)
		fmt.Println("Year of Passing:", s.YearOfPassing)
		fmt.Println("Number of Years to Graduate:", s.NumberOfYearsToGraduate)
		fmt.Println() 
	}
	
}
