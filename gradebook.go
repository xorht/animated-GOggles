/*
@name: Juwuan Turner-Howard
@date: Mar 1, 2021
@course-section: CS 424-02
@description:
	This project creates and displays a grade-book from a user inputted file path and grade scale.
*/
package main

// GradeScale contains grade scale information, including test and homework weights.
type GradeScale struct {
	testWeight, homeworkWeight float64
}

// TestWeight returns the test weight of the GradeScale.
func (g GradeScale) TestWeight() float64 {
	return g.testWeight
}

// HomeworkWeight returns the homework weight of the GradeScale.
func (g GradeScale) HomeworkWeight() float64 {
	return g.homeworkWeight
}

// GradeBook holds a collection of Students, a GradeScale, and its overall average.
// The path of the source file with the students is kept as a record.
type GradeBook struct {
	students   []Student
	scale      GradeScale
	average    float64
	sourcePath string
}

// Students returns value slice of GradeBook's Students.
func (g GradeBook) Students() []Student {
	return g.students
}

// TestWeight returns the test weight percentage of the GradeBook.
func (g GradeBook) TestWeight() float64 {
	return g.scale.testWeight
}

// HomeworkWeight returns the homeworkWeight percentage of the GradeBook.
func (g GradeBook) HomeworkWeight() float64 {
	return g.scale.homeworkWeight
}

// Average returns the course average of the GradeBook.
func (g GradeBook) Average() float64 {
	return g.average
}

// SourcePath returns the data source of GradeBook.
func (g *GradeBook) SourcePath() string {
	return g.sourcePath
}

// UpdateAverages updates the GradeBook's and Students' averages.
// Calls UpdateStudentAverages and UpdateGradeBookAverage.
func (g *GradeBook) UpdateAverages() {
	g.UpdateStudentAverages()
	g.UpdateGradeBookAverage()
}

// UpdateGradeBookAverage updates the overall average of the GradeBook's course average.
func (g *GradeBook) UpdateGradeBookAverage() {
	g.average = CalculateGradeBookAverage(*g)
}

// UpdateStudentAverages updates all Student averages in the GradeBook.
func (g *GradeBook) UpdateStudentAverages() {
	students := g.students
	for i := range students {
		students[i].UpdateAverages(g.scale)
	}
}

// UpdateStudentOrdering updates the ordering of students in the GradeBook by the given order function.
func (g *GradeBook) UpdateStudentOrdering(order func(s1, s2 *Student) bool) {
	By(order).Sort(g.Students())
}

// AddStudent adds given student to the GradeBook.
func (g *GradeBook) AddStudent(student Student) {
	g.students = append(g.students, student)
}
