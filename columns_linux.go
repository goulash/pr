// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package pr

func PrintGrid(columns int, list []string) {
	tw := GetTerminalWidth()
	FprintGrid(os.Stdout, tw, columns, list)
}

func PrintFlex(list []string) {
	tw := GetTerminalWidth()
	FprintFlex(os.Stdout, tw, list)
}
