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
}

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
