// MIT License

// Copyright (c) 2020 Uday Hiwarale

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"github.com/thatisuday/commando"
)

func main() {

	// configure commando
	commando.
		SetExecutableName("tree").
		SetVersion("1.0.1").
		SetDescription("This tool lists the contents of a directory in tree-like format.\nIt can also display information about files and folders like size, permission and ownership.")

	// configure the root command
	commando.
		Register(nil).
		AddArgument("dir", "local directory path", "./").                                                   // default `./`
		AddFlag("level,l", "level of depth to travel", commando.Int, 1).                                    // default `1`
		AddFlag("size", "display size of the each file", commando.Bool, nil).                               // default `false`
		AddFlag("mode", "display mode of the each file", commando.Bool, nil).                               // default `false`
		AddFlag("no-color", "ignore colored output", commando.Bool, nil).                                   // default `true`
		AddFlag("ignore", "ignore directories (separated by comma)", commando.String, ".git,node_modules"). // default `.git,node_modules`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			dir := args["dir"].Value

			// call `list` function
			list(false, dir, flags)
		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("displays detailed information of a directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		AddArgument("dir", "local directory path", "./").
		AddFlag("level,l", "level of depth to travel", commando.Int, nil).
		AddFlag("no-color", "ignore colored output", commando.Bool, nil).                                   // default `true`
		AddFlag("ignore", "ignore directories (separated by comma)", commando.String, ".git,node_modules"). // default `.git,node_modules`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			dir := args["dir"].Value

			// call `list` function
			list(true, dir, flags)
		})

	// parse command-line arguments
	commando.Parse(nil)

}
