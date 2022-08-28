/*
@name: Juwuan Turner-Howard
@date: Mar 1, 2021
@course-section: CS 424-02
@description:
	This project creates and displays a grade-book from a user inputted file path and grade scale.

@notes:
	Student Data Schema in File:
		<firstname> <lastname>
		<list of test grades>
		<list of homework grades>
*/

package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// ReadStudentsFrom reads Student information from an io.Reader.
func ReadStudentsFrom(reader io.Reader) (students []Student) {
	students = make([]Student, 0) // slice of students
	// Set scanner and scan while there is more to scan
	for scanner := bufio.NewScanner(reader); scanner.Scan(); {
		students = append(students, readStudent(scanner))
	}
	return
}

// readStudent reads section of student information including name, test grades, and homework grades.
func readStudent(scanner *bufio.Scanner) Student {
	fullName := strings.Split(scanner.Text(), " ")  // read in student full name
	firstName, lastName := fullName[0], fullName[1] // unpack student's first and last name
	scanner.Scan()                                  // move to next line in file
	testGrades := readGrades(*scanner)              // read test grades
	scanner.Scan()                                  // move to next line in file
	homeworkGrades := readGrades(*scanner)          // read homework grades
	return NewStudent(firstName, lastName, testGrades, homeworkGrades)
}

// readGrades reads row of grades into a float64 slice.
func readGrades(scanner bufio.Scanner) (grades []float64) {
	gradesInput := strings.Split(scanner.Text(), " ")
	grades = make([]float64, len(gradesInput)) // initialize slice with tight bounded size
	for i := range gradesInput {
		grades[i], _ = strconv.ParseFloat(gradesInput[i], 64)
	}
	return
}
