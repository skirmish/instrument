// Parse arguments.

package main

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

func parseArgs(args []string) (*Build, error) {
	return nil, nil
}
