/*
@name: Juwuan Turner-Howard
@date: Mar 1, 2021
@course-section: CS 424-02
@description:
	This project creates and displays a grade-book from a user inputted file path and grade scale.
*/

package main

import "sort"

// By acts as a functionality wrapper for sorting orderings. Example: By(ordering).Sort(students)
type By func(s1, s2 *Student) bool

// Sort sorts a slice By the casted function (closure) where the closure is a Less comparison function.
func (by By) Sort(students []Student) {
	gradeBookSorter := &gradeBookSorter{
		students: students,
		order:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(gradeBookSorter)
}

// gradeBookSorter sorting agent for sort.Sort's method containing data slice and sorting closure (Less method).
type gradeBookSorter struct {
	students []Student
	order    By
}

// Len is part of sort.Interface.
func (g *gradeBookSorter) Len() int {
	return len(g.students)
}

// Swap is part of sort.Interface.
func (g *gradeBookSorter) Swap(i, j int) {
	g.students[i], g.students[j] = g.students[j], g.students[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (g *gradeBookSorter) Less(i, j int) bool {
	return g.order(&g.students[i], &g.students[j])
}
