// boltchat
// Copyright (C) 2021  The boltchat Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package errs

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type SyntaxError struct {
	Error error
	Desc  string
}

/*
Emerg displays a message to the user that something has
gone wrong internally, and immediately stops program
execution afterwards.
*/
func Emerg(err error) {
	fmt.Printf(color.HiRedString(
		"An unexpected error has occurred.\nPlease consider creating " +
			"an issue at <https://github.com/bolt-chat/boltchat/issues> " +
			"if this is repetitive behaviour.\n",
	))

	// Immediately stop execution
	panic(err)
}

/*
Syntax tells the user that they've made a syntax error
and cleanly exits the program afterwards.
*/
func Syntax(err SyntaxError) {
	fmt.Printf("Syntax error: %s\n", err.Desc)
	os.Exit(1)
}

/*
Connect informs the user about an error that has occured while
attempting to connect to the server.
*/
func Connect(err error) {
	fmt.Printf("Connection error: %s\n", err.Error())
	os.Exit(1)
}
