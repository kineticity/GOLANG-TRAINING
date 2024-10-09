package student

import (
	"fmt"
	"time"
)

var rollNumberCounter = 0 //roll number which increments for each new Student

var students = []*Student{}

type Student struct {
	RollNo                  int
	Firstname               string
	Lastname                string
	fullname                string
	DOB                     time.Time
	age                     int
	SemesterCGPAArray       []float64
	finalCGPA               float64
	semesterGrades          []string
	finalGrade              string
	YearOfEnrollment        int
	YearOfPassing           int
	numberOfYearsToGraduate int
}

// CREATE
func NewStudent(firstname, lastname string, dob time.Time, semesterCGPAArray []float64, yearOfEnrollment int, yearOfPassing int) (*Student, string) {
	err := validateStudentInput(firstname, lastname, dob, semesterCGPAArray, yearOfEnrollment, yearOfPassing)
	if err != "" {
		return nil, err
	}

	rollNumberCounter++

	student := &Student{
		RollNo:                  rollNumberCounter,
		Firstname:               firstname,
		Lastname:                lastname,
		fullname:                getFullname(firstname, lastname),
		DOB:                     dob,
		age:                     calculateAge(dob),
		SemesterCGPAArray:       semesterCGPAArray,
		finalCGPA:               calculateFinalCGPA(semesterCGPAArray),
		semesterGrades:          calculateSemesterGrades(semesterCGPAArray),
		finalGrade:              getFinalGrade(calculateFinalCGPA(semesterCGPAArray)),
		YearOfEnrollment:        yearOfEnrollment,
		YearOfPassing:           yearOfPassing,
		numberOfYearsToGraduate: calculateNumberOfYearsToGraduate(yearOfPassing, yearOfEnrollment),
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

	if yearOfEnrollment == 0 || yearOfEnrollment > time.Now().Year() {
		return "invalid year of enrollment"
	}

	if yearOfPassing == 0 || yearOfPassing < yearOfEnrollment {
		return "invalid year of passing"
	}

	return ""
}

// READ ALL
func ReadAllStudents() []*Student {
	return students
}

// READ BY ROLLNO
func ReadStudentByRollNo(rollNo int) (*Student, string) {
	for _, student := range students {
		if student.RollNo == rollNo {
			return student, ""
		}
	}
	return nil, "This student does not exist"
}



//VALIDATION FNS
func validname(name interface{}) string{
	namevalue, ok := name.(string)
	if !ok {
		return "Invalid value type for firstname"
	}

	if namevalue!=""{
		return ""
	} else{
		return "Firstname or Lastname cannot be empty"
	}
}

func validDOB(dob interface{}) string {
	dobValue, ok := dob.(time.Time)
	if !ok {
		return "Invalid value type for date of birth"
	}
	if dobValue.After(time.Now()) {
		return "Date of birth cannot be in the future"
	}
	return ""
}

func validSemesterCGPAArray(cgpaArray interface{}) string {
	cgpaArrayValue, ok := cgpaArray.([]float64)
	if !ok {
		return "Invalid value type for CGPA array"
	}
	for _, cgpa := range cgpaArrayValue {
		if cgpa < 0 || cgpa > 10 {
			return "CGPA values must be between 0 and 10"
		}
	}
	return ""
}

func validYearOfEnrollment(year interface{}) string {
	yearValue, ok := year.(int)
	if !ok {
		return "Invalid value type for year of enrollment"
	}
	currentYear := time.Now().Year()
	if yearValue < 1900 || yearValue > currentYear {
		return "Year of enrollment must be between 1900 and the current year"
	}
	return ""
}

func validYearOfPassing(year interface{}, yearOfEnrollment int) string {
	yearValue, ok := year.(int)
	if !ok {
		return "Invalid value type for year of passing"
	}
	currentYear := time.Now().Year()
	if yearValue < yearOfEnrollment || yearValue > currentYear {
		return "Year of passing must be after the year of enrollment and not in the future"
	}
	return ""
}



//UPDATE BY PARAMETER
func (s *Student) UpdateStudentByParameter(parameter string, newValue interface{}) string {

	switch parameter {
	case "firstname":

		err:=validname(newValue)
		if err != "" {
			return err
		}
		firstname,_:=newValue.(string)
		s.Firstname = firstname
		s.fullname = getFullname(s.Firstname, s.Lastname)

	case "lastname":

		err:=validname(newValue)
		if err != "" {
			return err
		}
		lastname, _ := newValue.(string)

		s.Lastname = lastname
		s.fullname = getFullname(s.Firstname, s.Lastname)

	case "dob":
		err:=validDOB(newValue)
		if err != "" {
			return err
		}
	
		dob, _ := newValue.(time.Time)

		s.DOB = dob
		s.age = calculateAge(dob)

	case "semesterCGPAArray":
		err:=validSemesterCGPAArray(newValue)
		if err != "" {
			return err
		}
		
		semesterCGPAArray, _ := newValue.([]float64)
		s.SemesterCGPAArray = semesterCGPAArray
		s.finalCGPA = calculateFinalCGPA(semesterCGPAArray)
		s.semesterGrades = calculateSemesterGrades(semesterCGPAArray)
		s.finalGrade = getFinalGrade(s.finalCGPA)

	case "yearOfEnrollment":
		err:=validYearOfEnrollment(newValue)
		if err != "" {
			return err
		}
		yearOfEnrollment, _ := newValue.(int)

		s.YearOfEnrollment = yearOfEnrollment
		s.numberOfYearsToGraduate = calculateNumberOfYearsToGraduate(s.YearOfPassing, s.YearOfEnrollment)

	case "yearOfPassing":
		err:=validYearOfPassing(newValue,s.YearOfEnrollment)
		if err != "" {
			return err
		}
		yearOfPassing, _ := newValue.(int)
		s.YearOfPassing = yearOfPassing
		s.numberOfYearsToGraduate = calculateNumberOfYearsToGraduate(s.YearOfPassing, s.YearOfEnrollment)

	default:
		return "Invalid parameter"
	}

	return ""
}


// DELETE BY ROLLNO

func DeleteStudentByRollNo(rollNo int) string {
	for i, student := range students {
		if student.RollNo == rollNo {
			students = append(students[:i], students[i+1:]...)
			student=nil //make deleted student object nil
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
	if s == nil {
		fmt.Println("Student object was not created")
	} else {
		fmt.Println("Roll No:", s.RollNo)
		fmt.Println("First Name:", s.Firstname)
		fmt.Println("Last Name:", s.Lastname)
		fmt.Println("Full Name:", s.fullname)
		fmt.Println("Date of Birth:", s.DOB.Format("2006-01-02"))
		fmt.Println("Age:", s.age)
		fmt.Println("Semester CGPA Array:", s.SemesterCGPAArray)
		fmt.Println("Final CGPA:", s.finalCGPA)
		fmt.Println("Semester Grades:", s.semesterGrades)
		fmt.Println("Final Grade:", s.finalGrade)
		fmt.Println("Year of Enrollment:", s.YearOfEnrollment)
		fmt.Println("Year of Passing:", s.YearOfPassing)
		fmt.Println("Number of Years to Graduate:", s.numberOfYearsToGraduate)
		fmt.Println()
	}

}
