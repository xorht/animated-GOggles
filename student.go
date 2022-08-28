/*
@name: Juwuan Turner-Howard
@date: Mar 1, 2021
@course-section: CS 424-02
@description:
	This project creates and displays a grade-book from a user inputted file path and grade scale.
*/

package main

// Student holds general student information and methods.
type Student struct {
	firstName, lastName                         string
	testGrades, homeworkGrades                  []float64
	testAverage, homeworkAverage, courseAverage float64
}

// UpdateAverages updates the Student's test, homework, and course averages given a GradeScale.
func (s *Student) UpdateAverages(scale GradeScale) {
	s.testAverage = AverageOf(s.Tests())
	s.homeworkAverage = AverageOf(s.Homeworks())
	s.courseAverage = CalculateStudentAverage(scale, *s)
}

// NewStudent creates Student object in a parametrized fashion.
func NewStudent(firstName string, lastName string, testGrades []float64, homeworkGrades []float64) Student {
	return Student{firstName: firstName, lastName: lastName, testGrades: testGrades, homeworkGrades: homeworkGrades}
}

// HomeworkAverage returns the Student's homework average.
func (s Student) HomeworkAverage() float64 {
	return s.homeworkAverage
}

// TestAverage returns the Student's test average.
func (s Student) TestAverage() float64 {
	return s.testAverage
}

// FullName returns the Student's first and last names spaced.
func (s Student) FullName() string {
	return s.firstName + " " + s.lastName
}

// LastNameFirst returns the Student's last name delimited by a comma followed by first name.
func (s Student) LastNameFirst() string {
	return s.lastName + ", " + s.firstName
}

// CourseAverage returns Student's course average.
func (s Student) CourseAverage() float64 {
	return s.courseAverage
}

// Tests returns a copy of the Student's test grades values.
func (s Student) Tests() []float64 {
	return s.testGrades
}

// Homeworks returns copy of the Student's homework grade values.
func (s Student) Homeworks() []float64 {
	return s.homeworkGrades
}
