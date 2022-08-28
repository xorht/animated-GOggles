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
	"fmt"
	"os"
	"strings"
	"time"
)

// Entry-point for Grade-Book program.
func main() {
	fmt.Println()
	fmt.Print(ProgramOverview)
	fmt.Print(FileExample)
	fmt.Println()
	time.Sleep(3 * time.Second)
	for run := true; run; { // run program while run is true
		run = Run()
	}
	ExitOutput()
}

// Run prompts for and reads a file path and test weight from the console, displaying a sorted GradeBook in tabular
// form with ReportGradeBook.
func Run() bool {
	fmt.Print(FilenamePrompt)
	var path string               // relative or absolute path of file
	fmt.Scanf("%s", &path) // read path

	// Attempt to open file for fail-fast approach
	var inputFile, err = os.OpenFile(path, os.O_RDONLY, 0644)
	for err != nil { // unhappy path
		fmt.Print(err)         // display the error
		if !continueOption() { // if user does not want to continue, exit
			return false
		}
		fmt.Print(FilenamePrompt)
		fmt.Scanf("%s", &path)                                // read path
		inputFile, err = os.OpenFile(path, os.O_RDONLY, 0644) // try to open path again
	}
	defer inputFile.Close() // ensure file is closed before function exits

	fmt.Printf("\nSuccessfully loaded grade book from: %s\n", path)
	students := ReadStudentsFrom(inputFile) // Read students

	fmt.Print(GradeScalePrompt)
	println()
	fmt.Print(TestWeightPrompt)
	var testWeight float64
	fmt.Scanf("%f", &testWeight) // read in test weight percentage
	homeworkWeight := 100 - testWeight  // set homework to remaining percentage of 100

	gradeBook := GradeBook{
		students:   students,
		scale:      GradeScale{testWeight: testWeight, homeworkWeight: homeworkWeight},
		sourcePath: path,
	}

	gradeBook.UpdateAverages()                         // update grade book to have most recent calculated averages
	gradeBook.UpdateStudentOrdering(AscendingLastName) // order by last name

	fmt.Println()
	ReportGradeBook(gradeBook) // report grade book

	return continueOption()
}

// continueOption prompts the user to determine program continuation.
func continueOption() bool {
	fmt.Print("\n\nWould you like to continue[y/n]? ")
	var choice string
	fmt.Scanf("%s", &choice)
	return strings.ToLower(choice) == "y"
}
