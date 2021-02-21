// Parse arguments.

package main

import "strings"

//
// Examples:
//  C:\Go\pkg\tool\windows_amd64\compile.exe
// 		-o $WORK\b026\_pkg_.a
//		-trimpath $WORK\b026=>
//		-p internal/poll
//		-std
//		-buildid JttlkUAjQvxzpPTPE4ew/JttlkUAjQvxzpPTPE4ew
//		-goversion go1.14.2
//		-D
//		-importcfg $WORK\b026\importcfg
//		-pack
//		-c=4
//		C:\Go\src\internal\poll\errno_windows.go
//		C:\Go\src\internal\poll\fd.go C:\Go\src\internal\poll\fd_fsync_windows.go
//		C:\Go\src\internal\poll\fd_mutex.go
//		C:\Go\src\internal\poll\fd_poll_runtime.go
//		C:\Go\src\internal\poll\fd_posix.go
//		C:\Go\src\internal\poll\fd_windows.go
//		C:\Go\src\internal\poll\hook_windows.go C:\Go\src\internal\poll\sendfile_windows.go
//		C:\Go\src\internal\poll\sockopt.go
//		C:\Go\src\internal\poll\sockopt_windows.go
//		C:\Go\src\internal\poll\sockoptip.go
//
// C:\Go\pkg\tool\windows_amd64\compile.exe
//		-o $WORK\b001\_pkg_.a
//		-trimpath $WORK\b001=>
//		-p main
//		-complete
//		-buildid ZhbgFFYdodKdwrtFK8AS/ZhbgFFYdodKdwrtFK8AS -goversion go1.14.2
//		-D _/D_/dev/go/src/github.com/skirmish/instrument/testdata/hello
//		-importcfg $WORK\b001\importcfg
//		-pack
//		-c=4
//		testdata\hello\main.go

type Build struct {
	program string
	args    []string
	files   []string
}

// isOpt determines if the current item is a flag or not.
func isOption(this string, next string) (int, bool) {
	switch {
	case this[0] == '-' && strings.Contains(this[1:], "="): // -flag=val
		//kv := strings.SplitN(this[1:],"=",2)
		return 1, true
		break
	case this[0] == '-' && ((len(next) > 0 && next[0] != '-') || len(next) == 0): // -flag value
		return 2, true
		break
	case this[0] == '-': // -flag
		return 1, true
		break
	default:
		return 1, false
		break
	}
	return 1, false
}

// Parse the arguments into a Build object which we can then work with.
func parseArgs(args []string) (*Build, error) {
	b := &Build{}

	b.program = args[0]

	// Go through all the arguments, skipping the first as it is the program to invoke
	i := 1
	for i < len(args)-1 {
		inc, isOpt := isOption(args[i], args[i+1])
		if isOpt {
			if inc == 1 {
				b.args = append(b.args, args[i])
			} else {
				b.args = append(b.args, args[i]+" "+args[i+1])
			}
		} else {
			b.files = append(b.files, args[i])
		}
		i += inc
	}
	// Do the last argument
	if i < len(args) {
		_, isOpt := isOption(args[i], "")
		if isOpt {
			b.args = append(b.args, args[i])
		} else {
			b.files = append(b.files, args[i])
		}
	}

	return b, nil
}
