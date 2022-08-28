/*
@name: Juwuan Turner-Howard
@date: Mar 1, 2021
@course-section: CS 424-02
@description:
	This project creates and displays a grade-book from a user inputted file path and grade scale.
*/

package main

// CalculateStudentAverage calculates the Student's average based on the grade scale.
func CalculateStudentAverage(scale GradeScale, student Student) float64 {
	return (scale.TestWeight()*student.TestAverage() + scale.HomeworkWeight()*student.HomeworkAverage()) / 100
}

// CalculateStudentAverage calculates the GradeBook's average based on Student averages.
func CalculateGradeBookAverage(gradeBook GradeBook) float64 {
	var sum float64 = 0
	students := gradeBook.Students()
	for _, student := range students {
		sum += student.CourseAverage()
	}
	return sum / float64(len(students))
}

// averageOf calculates the average of a float64 slice.
func AverageOf(values []float64) float64 {
	sum := float64(0)
	for _, val := range values {
		sum += val
	}
	return sum / float64(len(values))
}
