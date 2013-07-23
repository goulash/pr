// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Naming
//
// The issue at the moment is deciding what names to give the functions
// and the future functions. So here I will try to list the functions
// I anticipate in this package concerning columns:
//
// The functions are prefixed with 'F' when with a Writer.
//
// PrintColumns(Writer, int, int, string)
//     Print a string into a defined set of columns
//
// PrintAutoColumns(Writer, int, string)
//     Arrange a string into as many columns as possible
//
// PrintGrid(Writer, int, int, []string)
//	   Print a list of strings into the specified number of columns
//
// PrintAutoGrid(Writer, int, []string)
//     Print a list of strings into an optimal grid formation

package pr

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

var columnPadding uint = 2

func GetColumnPadding() uint {
	return columnPadding
}

func SetColumnPadding(uint val) {
	columnPadding = val
}

// FprintGrid prints the items in the given list in as many columns as
// makes sense, given the horizontal space available.
//
// It will not print items in more columns than necessary: the minimum
// number of columns is used to attain the minimum row count.
func FprintAutoGrid(w Writer, hspace uint, list []string) {
	n := len(list)
	rc := runes(list)
	span := columns(rc, columnPadding, hspace)
	cols := len(span)

	if cols <= 1 {
		for _, s := range list {
			fmt.Println(s)
		}
		return
	}

	g := newGridFromCols(n, cols)
	for i := range g.IterRows() {
		if i.Ok {
			fmt.Print(list[i.Idx])
			fmt.Print(strings.Repeat(" ", span[i.Col]-rc[i.Idx]))
		}

		// At the last column, print a newline
		if i.Col == cols-1 {
			fmt.Print("\n")
		}
	}
	if n%cols != 0 {
		fmt.Print("\n")
	}
}

// columns returns the amount of columns that can fit in space, taken the list
// and the padding into consideration.
//
// If the return value is nil, then only one column is supported (if at all),
// otherwise it contains the length of each column, including the padding.
func columns(list []int, padding, hspace int) []int {
	if hspace <= 0 {
		return nil
	}

	n := len(list)
	rows := n
	cols := []int(nil)

trial:
	for c := 2; c <= n; c++ {
		g := mappings.NewGridFromCols(n, c)

		// Continue if we don't reduce rows with this many columns.
		if g.Rows() >= rows {
			continue
		}

		// Get the maximum widths of the individual columns.
		span := make([]int, c)
		for i := range g.IterRows() {
			if i.Ok && list[i.Idx] > span[i.Col] {
				span[i.Col] = list[i.Idx]
			}
		}

		// Have we reached the limit yet?
		padspace := padding * (len(cols) - 1)
		if sum(span)+padspace > hspace {
			break trial
		}

		// Update our data.
		rows = g.Rows()
		cols = span
	}

	// Add the padding.
	for i := 0; i < len(cols)-1; i++ {
		cols[i] += padding
	}
	return cols
}

// runes returns the rune counts of each of the strings in the list.
func runes(list []string) []int {
	retval := make([]int, len(list))
	for i, s := range list {
		retval[i] = utf8.RuneCountInString(s)
	}
	return retval
}

// avg computes the average over the uints in the list.
func avg(list []int) int {
	return sum(list) / len(list)
}

// sum computes the total over the uints in the list.
func sum(list []int) int {
	var total int
	for _, i := range list {
		total += i
	}
	return total
}
