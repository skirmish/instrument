package main

import (
	"reflect"
	"testing"
)

var argTests = []struct {
	name     string
	inArgs   []string
	outBuild *Build
}{
	{
		name: "first",
		inArgs: []string{
			"C:\\Go\\pkg\\tool\\windows_amd64\\compile.exe",
			"-o", "$WORK\\b026\\_pkg_.a",
			"-trimpath", "$WORK\\b026=>",
			"-p", "internal/poll",
			"-std",
			"-buildid", "JttlkUAjQvxzpPTPE4ew/JttlkUAjQvxzpPTPE4ew",
			"-goversion", "go1.14.2",
			"-D", "",
			"-importcfg", "$WORK\\b026\\importcfg",
			"-pack",
			"-c=4",
			"C:\\Go\\src\\internal\\poll\\errno_windows.go",
			"C:\\Go\\src\\internal\\poll\\fd.go",
			"C:\\Go\\src\\internal\\poll\\fd_fsync_windows.go",
			"C:\\Go\\src\\internal\\poll\\fd_mutex.go",
			"C:\\Go\\src\\internal\\poll\\fd_poll_runtime.go",
			"C:\\Go\\src\\internal\\poll\\fd_posix.go",
			"C:\\Go\\src\\internal\\poll\\fd_windows.go",
			"C:\\Go\\src\\internal\\poll\\hook_windows.go",
			"C:\\Go\\src\\internal\\poll\\sendfile_windows.go",
			"C:\\Go\\src\\internal\\poll\\sockopt.go",
			"C:\\Go\\src\\internal\\poll\\sockopt_windows.go",
			"C:\\Go\\src\\internal\\poll\\sockoptip.go",
		},
		outBuild: &Build{
			program: "C:\\Go\\pkg\\tool\\windows_amd64\\compile.exe",
			pkg:     "internal/poll",
			stdlib:  true,
			output:  "$WORK\\b026\\_pkg_.a",
			args: []string{
				"-o $WORK\\b026\\_pkg_.a",
				"-trimpath $WORK\\b026=>",
				"-p internal/poll",
				"-std",
				"-buildid JttlkUAjQvxzpPTPE4ew/JttlkUAjQvxzpPTPE4ew",
				"-goversion go1.14.2",
				"-D ",
				"-importcfg $WORK\\b026\\importcfg",
				"-pack",
				"-c=4",
			},
			files: []string{
				"C:\\Go\\src\\internal\\poll\\errno_windows.go",
				"C:\\Go\\src\\internal\\poll\\fd.go",
				"C:\\Go\\src\\internal\\poll\\fd_fsync_windows.go",
				"C:\\Go\\src\\internal\\poll\\fd_mutex.go",
				"C:\\Go\\src\\internal\\poll\\fd_poll_runtime.go",
				"C:\\Go\\src\\internal\\poll\\fd_posix.go",
				"C:\\Go\\src\\internal\\poll\\fd_windows.go",
				"C:\\Go\\src\\internal\\poll\\hook_windows.go",
				"C:\\Go\\src\\internal\\poll\\sendfile_windows.go",
				"C:\\Go\\src\\internal\\poll\\sockopt.go",
				"C:\\Go\\src\\internal\\poll\\sockopt_windows.go",
				"C:\\Go\\src\\internal\\poll\\sockoptip.go",
			},
		},
	},
	{
		name: "fibmain",
		inArgs: []string{
			"C:\\Go\\pkg\\tool\\windows_amd64\\compile.exe",
			"-o", "$WORK\\b001\\_pkg_.a",
			"-trimpath", "$WORK\\b001=>",
			"-p", "main",
			"-complete",
			"-buildid", "07gIkXLEJooNJ0O9y49j/07gIkXLEJooNJ0O9y49j",
			"-goversion", "go1.14.2",
			"-D", "_/D_/dev/go/src/github.com/skirmish/instrument/testdata/fibonacci",
			"-importcfg", "$WORK\\b001\\importcfg",
			"-pack",
			"-c=4",
			"testdata\\fibonacci\\main.go",
		},
		outBuild: &Build{
			program: "C:\\Go\\pkg\\tool\\windows_amd64\\compile.exe",
			output:  "$WORK\\b001\\_pkg_.a",
			pkg:     "main",
			args: []string{
				"-o $WORK\\b001\\_pkg_.a",
				"-trimpath $WORK\\b001=>",
				"-p main",
				"-complete",
				"-buildid 07gIkXLEJooNJ0O9y49j/07gIkXLEJooNJ0O9y49j",
				"-goversion go1.14.2",
				"-D _/D_/dev/go/src/github.com/skirmish/instrument/testdata/fibonacci",
				"-importcfg $WORK\\b001\\importcfg",
				"-pack",
				"-c=4",
			},
			files: []string{
				"testdata\\fibonacci\\main.go",
			},
			stdlib: false,
		},
	},
}

//
// Examples:
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

func Test_Args1(t *testing.T) {
	for _, tc := range argTests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseArgs(tc.inArgs)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.outBuild, got) {
				t.Fatalf("\n expected: %#v\n got: %#v", tc.outBuild, got)
			}
		})
	}
}
