// Copyright (c) 2013, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

//
// Package pr implements functions for pretty printing of information.
//
// Example:
//
//  import (
//  	"bufio"
//  	"flag"
//  	"fmt"
//  	"io"
//  	"os"
//  	"strings"
//
//  	"github.com/cassava/pr"
//  )
//
//  func main() {
//  	reader := bufio.NewReader(os.Stdin)
//  	buffer := make([]string, 0, 32)
//
//  	for {
//  		line, err := reader.ReadString('\n')
//  		if err != nil {
//  			if err != io.EOF {
//  				fmt.Errorf("error: %s\n", err)
//  			}
//  			break
//  		}
//
//  		buffer = append(buffer, strings.TrimSpace(line))
//  	}
//
//  	width := flag.Int("width", -1, "width of the terminal")
//  	flag.Parse()
//
//  	pr.Printc(buffer, *width)
//  }
package pr
