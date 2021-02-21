// Parse arguments.

package main

import "strings"

// Build is the output of parsing the arguments coming through.
type Build struct {
	program string
	output  string // output flag -o
	pkg     string // package flag -p
	args    []string
	files   []string
	stdlib  bool // stdlibrary?
}

// isOpt determines if the current item is a flag or not.
// This list of arguments can be one of 4 things;
//		-key=value
//		-key value
//		-key
//		filepath
// We need to decipher which is what.
func isOption(this string, next string) (int, bool, string, string) {
	switch {
	case this[0] == '-' && strings.Contains(this[1:], "="): // -flag=val
		kv := strings.SplitN(this[1:], "=", 2)
		return 1, true, kv[0], kv[1]
		break
	case this[0] == '-' && ((len(next) > 0 && next[0] != '-') || len(next) == 0): // -flag value
		return 2, true, this[1:], next
		break
	case this[0] == '-': // -flag
		return 1, true, this[1:], ""
		break
	default:
		return 1, false, "", this
		break
	}
	return 1, false, "", ""
}

// Parse the arguments into a Build object which we can then work with.
func parseArgs(args []string) (*Build, error) {
	b := &Build{}

	b.program = args[0]

	// Go through all the arguments, skipping the first as it is the program to invoke
	i := 1
	for i < len(args)-1 {
		inc, isOpt, k, v := isOption(args[i], args[i+1])
		if isOpt {
			if inc == 1 {
				b.args = append(b.args, args[i])
				switch k {
				case "std":
					b.stdlib = true
					break
				}
			} else {
				switch k {
				case "p": // package
					b.pkg = v
					break
				case "o": // output
					b.output = v
					break
				}
				b.args = append(b.args, args[i]+" "+args[i+1])
			}
		} else {
			b.files = append(b.files, args[i])
		}
		i += inc
	}
	// Do the last argument
	if i < len(args) {
		_, isOpt, _, _ := isOption(args[i], "") // can ignore the k,v portion as it is the last one.
		if isOpt {
			b.args = append(b.args, args[i])
		} else {
			b.files = append(b.files, args[i])
		}
	}

	return b, nil
}
