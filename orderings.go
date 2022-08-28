/*
@name: Juwuan Turner-Howard
@date: Mar 1, 2021
@course-section: CS 424-02
@description:
	This project creates and displays a grade-book from a user inputted file path and grade scale.
*/

package main

// AscendingLastName sorting order.
var AscendingLastName = func(s1, s2 *Student) bool {
	return s1.LastNameFirst() <= s2.LastNameFirst()
}

// DescendingCourseAverage sorting order.
var DescendingCourseAverage = func(s1, s2 *Student) bool {
	return s1.CourseAverage() >= s2.CourseAverage()
}
