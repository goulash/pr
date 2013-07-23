// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package pr

func PrintColumns(columns int, text string) {
	tw := GetTerminalWidth()
	FprintColumns(os.Stdout, tw, columns, text)
}

func PrintAutoColumns(text string) {
	tw := GetTerminalWidth()
	FprintAutoColumns(os.Stdout, tw, text)
}

func PrintGrid(columns int, list []string) {
	tw := GetTerminalWidth()
	FprintColumns(os.Stdout, tw, columns, list)
}

func PrintAutoGrid(list []string) {
	tw := GetTerminalWidth()
	FprintColumns(os.Stdout, tw, list)
}
