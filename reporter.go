/*
@name: Juwuan Turner-Howard
@date: Mar 1, 2021
@course-section: CS 424-02
@description:
	This project creates and displays a grade-book from a user inputted file path and grade scale.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

const FilenamePrompt = "Please enter file name: "
const TestWeightPrompt = "Please enter test weight: "
const ProgramOverview = "Welcome to the grade-book calculator test program." +
	"\nThe software will read students from an input data file." +
	"\nWhen prompted, please enter the absolute or directory relative path to file with students.\n"
const FileExample = "Student File Data Schema:" +
	"\n\t<firstname> <lastname>" +
	"\n\t<list of test grades>" +
	"\n\t<list of homework grades>" +
	"\n\tNote: No newline should be at the end of the file.\n"
const GradeScalePrompt = "Please enter the gradebook's scale information.\n" +
	"Note, there are only two weight groups at this time.\n"

// ExitOutput outputs exit message.
func ExitOutput() {
	fmt.Println("Exiting...")
}

// ReportOverview reports the GradeBook's student count, data source, test and homework weight, and class average.
func ReportOverview(gradeBook GradeBook) {
	fmt.Printf(
		" %s ---%d STUDENTS FOUND IN FILE %s\n",
		"GRADE REPORT",
		len(gradeBook.Students()),
		gradeBook.SourcePath(),
	)
	decimalLabelFormat := " %-21s: %.1f\n" // format each average output
	fmt.Printf(decimalLabelFormat, "TEST WEIGHT", gradeBook.TestWeight())
	fmt.Printf(decimalLabelFormat, "HOMEWORK WEIGHT", gradeBook.HomeworkWeight())
	fmt.Printf(decimalLabelFormat, "CLASS AVERAGE", gradeBook.Average())
}

// ReportTable reports GradeBook student information in a table form to the console.
func ReportTable(gradeBook GradeBook, tableName string) {
	// configure row format
	columnFormat := "%-20s"
	lastColumnFormat := "%-30s" // ensure text has room in last column for proper formatting.
	rowFormat := fmt.Sprintf(
		"| %s: %s| %s| %s| %s|", columnFormat, columnFormat, columnFormat, columnFormat, lastColumnFormat)

	// configure header output
	headerFormat := fmt.Sprintf(
		"| %s| %s| %s| %s| %s|", columnFormat, columnFormat, columnFormat, columnFormat, lastColumnFormat)
	headerOutput := fmt.Sprintf(
		headerFormat, "NAME", "TEST AVERAGE", "HOMEWORK AVERAGE", "COURSE AVERAGE", "STATUS")

	// configure table name format
	tableNameFormat := "| %-" + strconv.Itoa(len(headerOutput)-3) + "s|\n"

	// make sure line length is sufficient
	line := strings.Repeat("-", len(headerOutput))

	// output table name and header
	fmt.Println(line)
	fmt.Printf(tableNameFormat, tableName)
	fmt.Println(line)
	fmt.Println(headerOutput)
	fmt.Println(line)

	// setup assigment column format
	const (
		decimalFormat           = "%5.1f" // width five, right-justified, with one trailing value
		assignmentAmountFormat  = " (%d)"
		assignmentAverageFormat = decimalFormat + assignmentAmountFormat
	)

	students := gradeBook.Students()
	gradeBookAverage := gradeBook.Average()
	var testAverage, homeworkAverage, courseAverage float64 // reusable variables

	// print student rows
	for _, student := range students {
		// setup test average column output
		testLen := len(student.Tests())
		testAverage = student.TestAverage()
		testAverageColOutput := fmt.Sprintf(assignmentAverageFormat, testAverage, testLen)

		// setup homework average column output
		homeworkLen := len(student.Homeworks())
		homeworkAverage = student.HomeworkAverage()
		homeworkAverageColOutput := fmt.Sprintf(assignmentAverageFormat, homeworkAverage, homeworkLen)

		// setup course average row output
		courseAverage = student.CourseAverage()
		courseAverageColOutput := fmt.Sprintf(decimalFormat, student.CourseAverage())

		// setup course average status
		var courseAverageStatus = ""
		if courseAverage > gradeBookAverage {
			courseAverageStatus = "  ** above class average **"
		}
		courseStatusColOutput := fmt.Sprintf("%s", courseAverageStatus)

		rowOutput := fmt.Sprintf(
			rowFormat /* Name | Test Average | Homework Average | Course Average | Course Status */,
			student.LastNameFirst(),
			testAverageColOutput,
			homeworkAverageColOutput,
			courseAverageColOutput,
			courseStatusColOutput)

		fmt.Println(rowOutput)
	}
	fmt.Println(line) // table bottom border
}

// ReportGradeBook default report format of GradeBook.
func ReportGradeBook(gradeBook GradeBook) {
	ReportOverview(gradeBook)
	ReportTable(gradeBook, "GRADE-BOOK")
}
