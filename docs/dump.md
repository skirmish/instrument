## Dump

This is a dump of a run provided purely for observational purposes.

```bash
D:\dev\go\src\github.com\skirmish\instrument>go install github.com/skirmish/instrument/instrument-otel
```

```bash
D:\dev\go\src\github.com\skirmish\instrument>go build -o fib.exe -a -toolexec instrument-otel ./testdata/fibonacci/.
# internal/syscall/windows/sysdll
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  internal/syscall/windows/sysdll
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000116010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b028\_pkg_.a -trimpath $WORK\b028=> -p internal/syscall/windows/sysdll -std -complete -buildid KfHvX7d14D9jQrKvPHoA/KfHvX7d14D9jQrKvPHoA -goversion go1.14.2 -D  -importcfg $WORK\
b028\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\internal\syscall\windows\sysdll\sysdll.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b028\_pkg_.a -trimpath $WORK\b028=> -p internal/syscall/windows/sysdll -std -complete -buildid KfHvX7d14D9jQrKvPHoA/KfHvX7d14D9jQrKvPHoA -goversion go1.14.2 -D  -importcfg $WORK\b028\importcfg -pack -c=4 C:\Go\src\internal\syscall\windows\sysdll\sysdll.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b028\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b028=>" "-p" "internal/syscall/windows/sysdll" "-std" "-complete" "-buildid" "KfHvX7d14D9jQrKvPHoA/KfHvX7d14D9jQrKvPHoA" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b028\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\syscall\\windows\\sysdll\\sysdll.go"`
```

```bash
# runtime/internal/sys
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  runtime/internal/sys
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc00015e010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b011\_pkg_.a -trimpath $WORK\b011=> -p runtime/internal/sys -std -+ -complete -buildid Op1CdGbTIuE-wwfb1N6k/Op1CdGbTIuE-wwfb1N6k -goversion go1.14.2 -D  -importcfg $WORK\b011\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\runtime\internal\sys\arch.go C:\Go\src\runtime\internal\sys\arch_amd64.go C:\Go\src\runtime\internal\sys\intrinsics.go C:\Go\src\runtime\internal\sys\intrinsics_common.go C:\Go\src\runtime\internal\sys\stubs.go C:\Go\src\runtime\internal\sys\sys.go C:\Go\src\runtime\internal\sys\zgoarch_amd64.go C:\Go\src\runtime\internal\sys\zgoos_windows.go C:\Go\src\runtime\internal\sys\zversion.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b011\_pkg_.a -trimpath $WORK\b011=> -p runtime/internal/sys -std -+ -complete -buildid Op1CdGbTIuE-wwfb1N6k/Op1CdGbTIuE-wwfb1N6k -goversion go1.14.2 -D  -importcfg $WORK\b011\importcfg -pack -c=4 C:\Go\src\runtime\internal\sys\arch.go C:\Go\src\runtime\internal\sys\arch_amd64.go C:\Go\src\runtime\internal\sys\intrinsics.go C:\Go\src\runtime\internal\sys\intrinsics_common.go C:\Go\src\runtime\internal\sys\stubs.go C:\Go\src\runtime\internal\sys\sys.go C:\Go\src\runtime\internal\sys\zgoarch_amd64.go C:\Go\src\runtime\internal\sys\zgoos_windows.go C:\Go\src\runtime\internal\sys\zversion.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b011\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b011=>" "-p" "runtime/internal/sys" "-std" "-+" "-complete" "-buildid" "Op1CdGbTIuE-wwfb1N6k/Op1CdGbTIuE-wwfb1N6k" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b011\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\runtime\\internal\\sys\\arch.go" "C:\\Go\\src\\runtime\\internal\\sys\\arch_amd64.go" "C:\\Go\\src\\runtime\\internal\\sys\\intrinsics.go" "C:\\Go\\src\\runtime\\internal\\sys\\intrinsics_common.go" "C:\\Go\\src\\runtime\\internal\\sys\\stubs.go" "C:\\Go\\src\\runtime\\internal\\sys\\sys.go" "C:\\Go\\src\\runtime\\internal\\sys\\zgoarch_amd64.go" "C:\\Go\\src\\runtime\\internal\\sys\\zgoos_windows.go" "C:\\Go\\src\\runtime\\internal\\sys\\zversion.go"`
```

```bash
# internal/race
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  internal/race
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000116010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b019\_pkg_.a -trimpath $WORK\b019=> -p internal/race -std -complete -buildid Y9V0TbwEiA_yxVC5J3Af/Y9V0TbwEiA_yxVC5J3Af -goversion go1.14.2 -D  -importcfg $WORK\b019\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\internal\race\doc.go C:\Go\src\internal\race\norace.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b019\_pkg_.a -trimpath $WORK\b019=> -p internal/race -std -complete -buildid Y9V0TbwEiA_yxVC5J3Af/Y9V0TbwEiA_yxVC5J3Af -goversiongo1.14.2 -D  -importcfg $WORK\b019\importcfg -pack -c=4 C:\Go\src\internal\race\doc.go C:\Go\src\internal\race\norace.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b019\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b019=>" "-p" "internal/race" "-std" "-complete" "-buildid" "Y9V0TbwEiA_yxVC5J3Af/Y9V0TbwEiA_yxVC5J3Af" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b019\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\race\\doc.go" "C:\\Go\\src\\internal\\race\\norace.go"`
```

```bash
# math/bits
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  math/bits
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc00015c010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b015\_pkg_.a -trimpath $WORK\b015=> -p math/bits -std -complete -buildid PNMKhvgwzskiwR4h_-ED/PNMKhvgwzskiwR4h_-ED -goversion go1.14.2 -D  -importcfg $WORK\b015\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\math\bits\bits.go C:\Go\src\math\bits\bits_errors.go C:\Go\src\math\bits\bits_tables.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b015\_pkg_.a -trimpath $WORK\b015=> -p math/bits -std -complete -buildid PNMKhvgwzskiwR4h_-ED/PNMKhvgwzskiwR4h_-ED -goversion go1.14.2 -D  -importcfg $WORK\b015\importcfg -pack -c=4 C:\Go\src\math\bits\bits.go C:\Go\src\math\bits\bits_errors.go C:\Go\src\math\bits\bits_tables.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b015\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b015=>" "-p" "math/bits" "-std" "-complete" "-buildid" "PNMKhvgwzskiwR4h_-ED/PNMKhvgwzskiwR4h_-ED" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b015\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\math\\bits\\bits.go" "C:\\Go\\src\\math\\bits\\bits_errors.go" "C:\\Go\\src\\math\\bits\\bits_tables.go"`
```

```bash
# unicode/utf16
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  unicode/utf16
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000098010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b030\_pkg_.a -trimpath $WORK\b030=> -p unicode/utf16 -std -complete -buildid nCg9Zd-RVE8Oehnw0Xd4/nCg9Zd-RVE8Oehnw0Xd4 -goversion go1.14.2 -D  -importcfg $WORK\b030\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\unicode\utf16\utf16.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b030\_pkg_.a -trimpath $WORK\b030=> -p unicode/utf16 -std -complete -buildid nCg9Zd-RVE8Oehnw0Xd4/nCg9Zd-RVE8Oehnw0Xd4 -goversiongo1.14.2 -D  -importcfg $WORK\b030\importcfg -pack -c=4 C:\Go\src\unicode\utf16\utf16.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b030\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b030=>" "-p" "unicode/utf16" "-std" "-complete" "-buildid" "nCg9Zd-RVE8Oehnw0Xd4/nCg9Zd-RVE8Oehnw0Xd4" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b030\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\unicode\\utf16\\utf16.go"`
```

```bash
# unicode/utf8
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  unicode/utf8
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000098010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b017\_pkg_.a -trimpath $WORK\b017=> -p unicode/utf8 -std -complete -buildid BsMfO-Z2dYTX3y0jz9Yz/BsMfO-Z2dYTX3y0jz9Yz -goversion go1.14.2 -D  -importcfg $WORK\b017\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\unicode\utf8\utf8.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b017\_pkg_.a -trimpath $WORK\b017=> -p unicode/utf8 -std -complete -buildid BsMfO-Z2dYTX3y0jz9Yz/BsMfO-Z2dYTX3y0jz9Yz -goversion go1.14.2 -D  -importcfg $WORK\b017\importcfg -pack -c=4 C:\Go\src\unicode\utf8\utf8.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b017\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b017=>" "-p" "unicode/utf8" "-std" "-complete" "-buildid" "BsMfO-Z2dYTX3y0jz9Yz/BsMfO-Z2dYTX3y0jz9Yz" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b017\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\unicode\\utf8\\utf8.go"`
```

```bash
# runtime/internal/math
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  runtime/internal/math
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000118010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b010\_pkg_.a -trimpath $WORK\b010=> -p runtime/internal/math -std -+ -complete -buildid Y1il53jR879KNSveIYWn/Y1il53jR879KNSveIYWn -goversion go1.14.2 -D  -importcfg $WORK\b010\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\runtime\internal\math\math.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b010\_pkg_.a -trimpath $WORK\b010=> -p runtime/internal/math -std -+ -complete -buildid Y1il53jR879KNSveIYWn/Y1il53jR879KNSveIYWn-goversion go1.14.2 -D  -importcfg $WORK\b010\importcfg -pack -c=4 C:\Go\src\runtime\internal\math\math.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b010\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b010=>" "-p" "runtime/internal/math" "-std" "-+" "-complete" "-buildid" "Y1il53jR879KNSveIYWn/Y1il53jR879KNSveIYWn" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b010\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\runtime\\internal\\math\\math.go"`
```

```bash
# sync/atomic
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:10 Package:
instrumenter: 2021/02/21 12:10:10 stdlib:  false
instrumenter: 2021/02/21 12:10:10 output:  0xc0000da010
instrumenter: 2021/02/21 12:10:10 args:  [-trimpath $WORK\b020=> -I $WORK\b020\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b020\symabis]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\sync\atomic\asm.s]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b020=> -I $WORK\b020\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b020\symabis C:\Go\src\sync\atomic\asm.s]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020\\symabis" "C:\\Go\\src\\sync\\atomic\\asm.s"`
```

```bash
# sync/atomic
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  sync/atomic
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000096010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b020\_pkg_.a -trimpath $WORK\b020=> -p sync/atomic -std -buildid MJ6qnH8ziDUpAaXAHKsd/MJ6qnH8ziDUpAaXAHKsd -goversion go1.14.2 -symabis $WORK\b020\symabis -D  -importcfg $WORK\b020\importcfg -pack -asmhdr $WORK\b020\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\sync\atomic\doc.go C:\Go\src\sync\atomic\value.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b020\_pkg_.a -trimpath $WORK\b020=> -p sync/atomic -std -buildid MJ6qnH8ziDUpAaXAHKsd/MJ6qnH8ziDUpAaXAHKsd -goversion go1.14.2 -symabis $WORK\b020\symabis -D  -importcfg $WORK\b020\importcfg -pack -asmhdr $WORK\b020\go_asm.h -c=4 C:\Go\src\sync\atomic\doc.go C:\Go\src\sync\atomic\value.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020=>" "-p" "sync/atomic" "-std" "-buildid" "MJ6qnH8ziDUpAaXAHKsd/MJ6qnH8ziDUpAaXAHKsd" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020\\go_asm.h" "-c=4" "C:\\Go\\src\\sync\\atomic\\doc.go" "C:\\Go\\src\\sync\\atomic\\value.go"`
```

```bash
# sync/atomic
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:10 Package:
instrumenter: 2021/02/21 12:10:10 stdlib:  false
instrumenter: 2021/02/21 12:10:10 output:  0xc000096010
instrumenter: 2021/02/21 12:10:10 args:  [-trimpath $WORK\b020=> -I $WORK\b020\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b020\asm.o]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\sync\atomic\asm.s]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b020=> -I $WORK\b020\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b020\asm.o C:\Go\src\sync\atomic\asm.s]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b020\\asm.o" "C:\\Go\\src\\sync\\atomic\\asm.s"`
```

```bash
# unicode
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  unicode
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000096010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b021\_pkg_.a -trimpath $WORK\b021=> -p unicode -std -complete -buildid S0YkY-HO-1O3-36O1NG8/S0YkY-HO-1O3-36O1NG8 -goversion go1.14.2 -D  -importcfg $WORK\b021\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\unicode\casetables.go C:\Go\src\unicode\digit.go C:\Go\src\unicode\graphic.go C:\Go\src\unicode\letter.go C:\Go\src\unicode\tables.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b021\_pkg_.a -trimpath $WORK\b021=> -p unicode -std -complete -buildid S0YkY-HO-1O3-36O1NG8/S0YkY-HO-1O3-36O1NG8 -goversion go1.14.2 -D  -importcfg $WORK\b021\importcfg -pack -c=4 C:\Go\src\unicode\casetables.go C:\Go\src\unicode\digit.go C:\Go\src\unicode\graphic.go C:\Go\src\unicode\letter.go C:\Go\src\unicode\tables.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b021\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b021=>" "-p" "unicode" "-std" "-complete" "-buildid" "S0YkY-HO-1O3-36O1NG8/S0YkY-HO-1O3-36O1NG8" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b021\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\unicode\\casetables.go" "C:\\Go\\src\\unicode\\digit.go" "C:\\Go\\src\\unicode\\graphic.go" "C:\\Go\\src\\unicode\\letter.go" "C:\\Go\\src\\unicode\\tables.go"`
```

```bash
# internal/cpu
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:10 Package:
instrumenter: 2021/02/21 12:10:10 stdlib:  false
instrumenter: 2021/02/21 12:10:10 output:  0xc000096010
instrumenter: 2021/02/21 12:10:10 args:  [-trimpath $WORK\b007=> -I $WORK\b007\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b007\symabis]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\internal\cpu\cpu_x86.s]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b007=> -I $WORK\b007\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b007\symabis C:\Go\src\internal\cpu\cpu_x86.s]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007\\symabis" "C:\\Go\\src\\internal\\cpu\\cpu_x86.s"`
```

```bash
# internal/cpu
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  internal/cpu
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000096010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b007\_pkg_.a -trimpath $WORK\b007=> -p internal/cpu -std -+ -buildid uc9DHwK9tFXI11CMLuCT/uc9DHwK9tFXI11CMLuCT -goversion go1.14.2 -symabis $WORK\b007\symabis -D  -importcfg $WORK\b007\importcfg -pack -asmhdr $WORK\b007\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\internal\cpu\cpu.go C:\Go\src\internal\cpu\cpu_amd64.go C:\Go\src\internal\cpu\cpu_x86.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b007\_pkg_.a -trimpath $WORK\b007=> -p internal/cpu -std -+ -buildid uc9DHwK9tFXI11CMLuCT/uc9DHwK9tFXI11CMLuCT -goversion go1.14.2 -symabis $WORK\b007\symabis -D  -importcfg $WORK\b007\importcfg -pack -asmhdr $WORK\b007\go_asm.h -c=4 C:\Go\src\internal\cpu\cpu.go C:\Go\src\internal\cpu\cpu_amd64.go C:\Go\src\internal\cpu\cpu_x86.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007=>" "-p" "internal/cpu" "-std" "-+" "-buildid" "uc9DHwK9tFXI11CMLuCT/uc9DHwK9tFXI11CMLuCT" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007\\go_asm.h" "-c=4" "C:\\Go\\src\\internal\\cpu\\cpu.go" "C:\\Go\\src\\internal\\cpu\\cpu_amd64.go" "C:\\Go\\src\\internal\\cpu\\cpu_x86.go"`
```

```bash
# internal/cpu
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:10 Package:
instrumenter: 2021/02/21 12:10:10 stdlib:  false
instrumenter: 2021/02/21 12:10:10 output:  0xc0000de010
instrumenter: 2021/02/21 12:10:10 args:  [-trimpath $WORK\b007=> -I $WORK\b007\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b007\cpu_x86.o]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\internal\cpu\cpu_x86.s]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b007=> -I $WORK\b007\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b007\cpu_x86.o C:\Go\src\internal\cpu\cpu_x86.s]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b007\\cpu_x86.o" "C:\\Go\\src\\internal\\cpu\\cpu_x86.s"`
```

```bash
# runtime/internal/atomic
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:10 Package:
instrumenter: 2021/02/21 12:10:10 stdlib:  false
instrumenter: 2021/02/21 12:10:10 output:  0xc000094010
instrumenter: 2021/02/21 12:10:10 args:  [-trimpath $WORK\b009=> -I $WORK\b009\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b009\symabis]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\runtime\internal\atomic\asm_amd64.s]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b009=> -I $WORK\b009\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b009\symabis C:\Go\src\runtime\internal\atomic\asm_amd64.s]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009\\symabis" "C:\\Go\\src\\runtime\\internal\\atomic\\asm_amd64.s"`
```

```bash
# runtime/internal/atomic
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:10 Package:  runtime/internal/atomic
instrumenter: 2021/02/21 12:10:10 stdlib:  true
instrumenter: 2021/02/21 12:10:10 output:  0xc000096010
instrumenter: 2021/02/21 12:10:10 args:  [-o $WORK\b009\_pkg_.a -trimpath $WORK\b009=> -p runtime/internal/atomic -std -+ -buildid LXeORvlfeXo63DDEArrz/LXeORvlfeXo63DDEArrz -goversion go1.14.2 -symabis $WORK\b009\symabis -D  -importcfg $WORK\b009\importcfg -pack -asmhdr $WORK\b009\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\runtime\internal\atomic\atomic_amd64.go C:\Go\src\runtime\internal\atomic\stubs.go]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b009\_pkg_.a -trimpath $WORK\b009=> -p runtime/internal/atomic -std -+ -buildid LXeORvlfeXo63DDEArrz/LXeORvlfeXo63DDEArrz -goversion go1.14.2 -symabis $WORK\b009\symabis -D  -importcfg $WORK\b009\importcfg -pack -asmhdr $WORK\b009\go_asm.h -c=4 C:\Go\src\runtime\internal\atomic\atomic_amd64.go C:\Go\src\runtime\internal\atomic\stubs.go]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009=>" "-p" "runtime/internal/atomic" "-std" "-+" "-buildid" "LXeORvlfeXo63DDEArrz/LXeORvlfeXo63DDEArrz" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009\\go_asm.h" "-c=4" "C:\\Go\\src\\runtime\\internal\\atomic\\atomic_amd64.go" "C:\\Go\\src\\runtime\\internal\\atomic\\stubs.go"`
```

```bash
# runtime/internal/atomic
instrumenter: 2021/02/21 12:10:10 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:10 Package:
instrumenter: 2021/02/21 12:10:10 stdlib:  false
instrumenter: 2021/02/21 12:10:10 output:  0xc000116010
instrumenter: 2021/02/21 12:10:10 args:  [-trimpath $WORK\b009=> -I $WORK\b009\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b009\asm_amd64.o]
instrumenter: 2021/02/21 12:10:10 files:  [C:\Go\src\runtime\internal\atomic\asm_amd64.s]
instrumenter: 2021/02/21 12:10:10 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b009=> -I $WORK\b009\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b009\asm_amd64.o C:\Go\src\runtime\internal\atomic\asm_amd64.s]
instrumenter: 2021/02/21 12:10:10 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b009\\asm_amd64.o" "C:\\Go\\src\\runtime\\internal\\atomic\\asm_amd64.s"`

# internal/testlog
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:11 Package:  internal/testlog
instrumenter: 2021/02/21 12:10:11 stdlib:  true
instrumenter: 2021/02/21 12:10:11 output:  0xc0000e0010
instrumenter: 2021/02/21 12:10:11 args:  [-o $WORK\b034\_pkg_.a -trimpath $WORK\b034=> -p internal/testlog -std -complete -buildid ZTx8HOxpm-br7fjQieQr/ZTx8HOxpm-br7fjQieQr -goversion go1.14.2 -D  -importcfg $WORK\b034\importcfg-pack -c=4]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\internal\testlog\log.go]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b034\_pkg_.a -trimpath $WORK\b034=> -p internal/testlog -std -complete -buildid ZTx8HOxpm-br7fjQieQr/ZTx8HOxpm-br7fjQieQr -goversion go1.14.2 -D  -importcfg $WORK\b034\importcfg -pack -c=4 C:\Go\src\internal\testlog\log.go]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b034\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b034=>" "-p" "internal/testlog" "-std" "-complete" "-buildid" "ZTx8HOxpm-br7fjQieQr/ZTx8HOxpm-br7fjQieQr" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b034\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\testlog\\log.go"`
```

```bash
# internal/bytealg
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b006\symabis]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\internal\bytealg\compare_amd64.s C:\Go\src\internal\bytealg\count_amd64.s C:\Go\src\internal\bytealg\equal_amd64.s C:\Go\src\internal\bytealg\index_amd64.s C:\Go\src\internal\bytealg\indexbyte_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b006\symabis C:\Go\src\internal\bytealg\compare_amd64.s C:\Go\src\internal\bytealg\count_amd64.s C:\Go\src\internal\bytealg\equal_amd64.s C:\Go\src\internal\bytealg\index_amd64.s C:\Go\src\internal\bytealg\indexbyte_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\symabis" "C:\\Go\\src\\internal\\bytealg\\compare_amd64.s" "C:\\Go\\src\\internal\\bytealg\\count_amd64.s" "C:\\Go\\src\\internal\\bytealg\\equal_amd64.s" "C:\\Go\\src\\internal\\bytealg\\index_amd64.s" "C:\\Go\\src\\internal\\bytealg\\indexbyte_amd64.s"`
```

```bash
# internal/bytealg
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:11 Package:  internal/bytealg
instrumenter: 2021/02/21 12:10:11 stdlib:  true
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-o $WORK\b006\_pkg_.a -trimpath $WORK\b006=> -p internal/bytealg -std -+ -buildid 4QmI2kV8R9Dbmvgs2IAh/4QmI2kV8R9Dbmvgs2IAh -goversion go1.14.2 -symabis $WORK\b006\symabis -D  -importcfg$WORK\b006\importcfg -pack -asmhdr $WORK\b006\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\internal\bytealg\bytealg.go C:\Go\src\internal\bytealg\compare_native.go C:\Go\src\internal\bytealg\count_native.go C:\Go\src\internal\bytealg\equal_generic.go C:\Go\src\internal\bytealg\equal_native.go C:\Go\src\internal\bytealg\index_amd64.go C:\Go\src\internal\bytealg\index_native.go C:\Go\src\internal\bytealg\indexbyte_native.go]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b006\_pkg_.a -trimpath $WORK\b006=> -p internal/bytealg -std -+ -buildid 4QmI2kV8R9Dbmvgs2IAh/4QmI2kV8R9Dbmvgs2IAh -goversion go1.14.2 -symabis $WORK\b006\symabis -D  -importcfg $WORK\b006\importcfg -pack -asmhdr $WORK\b006\go_asm.h -c=4 C:\Go\src\internal\bytealg\bytealg.go C:\Go\src\internal\bytealg\compare_native.go C:\Go\src\internal\bytealg\count_native.go C:\Go\src\internal\bytealg\equal_generic.go C:\Go\src\internal\bytealg\equal_native.go C:\Go\src\internal\bytealg\index_amd64.go C:\Go\src\internal\bytealg\index_native.go C:\Go\src\internal\bytealg\indexbyte_native.go]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006=>" "-p" "internal/bytealg" "-std" "-+" "-buildid" "4QmI2kV8R9Dbmvgs2IAh/4QmI2kV8R9Dbmvgs2IAh" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\go_asm.h" "-c=4" "C:\\Go\\src\\internal\\bytealg\\bytealg.go" "C:\\Go\\src\\internal\\bytealg\\compare_native.go" "C:\\Go\\src\\internal\\bytealg\\count_native.go" "C:\\Go\\src\\internal\\bytealg\\equal_generic.go" "C:\\Go\\src\\internal\\bytealg\\equal_native.go" "C:\\Go\\src\\internal\\bytealg\\index_amd64.go" "C:\\Go\\src\\internal\\bytealg\\index_native.go" "C:\\Go\\src\\internal\\bytealg\\indexbyte_native.go"`
```

```bash
# internal/bytealg
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc00015c010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\compare_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\internal\bytealg\compare_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\compare_amd64.o C:\Go\src\internal\bytealg\compare_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\compare_amd64.o" "C:\\Go\\src\\internal\\bytealg\\compare_amd64.s"`
```

```bash
# internal/bytealg
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\count_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\internal\bytealg\count_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\count_amd64.o C:\Go\src\internal\bytealg\count_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\count_amd64.o" "C:\\Go\\src\\internal\\bytealg\\count_amd64.s"`
```

```bash
# internal/bytealg
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\equal_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\internal\bytealg\equal_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\equal_amd64.o C:\Go\src\internal\bytealg\equal_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\equal_amd64.o" "C:\\Go\\src\\internal\\bytealg\\equal_amd64.s"`
```

```bash
# internal/bytealg
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\index_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\internal\bytealg\index_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\index_amd64.o C:\Go\src\internal\bytealg\index_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\index_amd64.o" "C:\\Go\\src\\internal\\bytealg\\index_amd64.s"`

# internal/bytealg
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000114010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\indexbyte_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\internal\bytealg\indexbyte_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b006=> -I $WORK\b006\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b006\indexbyte_amd64.o C:\Go\src\internal\bytealg\indexbyte_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b006\\indexbyte_amd64.o" "C:\\Go\\src\\internal\\bytealg\\indexbyte_amd64.s"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b014\symabis]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\dim_amd64.s C:\Go\src\math\exp_amd64.s C:\Go\src\math\floor_amd64.s C:\Go\src\math\hypot_amd64.s C:\Go\src\math\log_amd64.s C:\Go\src\math\sqrt_amd64.s C:\Go\src\math\stubs_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b014\symabis C:\Go\src\math\dim_amd64.s C:\Go\src\math\exp_amd64.s C:\Go\src\math\floor_amd64.s C:\Go\src\math\hypot_amd64.s C:\Go\src\math\log_amd64.s C:\Go\src\math\sqrt_amd64.s C:\Go\src\math\stubs_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\symabis" "C:\\Go\\src\\math\\dim_amd64.s" "C:\\Go\\src\\math\\exp_amd64.s" "C:\\Go\\src\\math\\floor_amd64.s" "C:\\Go\\src\\math\\hypot_amd64.s" "C:\\Go\\src\\math\\log_amd64.s" "C:\\Go\\src\\math\\sqrt_amd64.s" "C:\\Go\\src\\math\\stubs_amd64.s"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:11 Package:  math
instrumenter: 2021/02/21 12:10:11 stdlib:  true
instrumenter: 2021/02/21 12:10:11 output:  0xc000122010
instrumenter: 2021/02/21 12:10:11 args:  [-o $WORK\b014\_pkg_.a -trimpath $WORK\b014=> -p math -std -buildid v42vL_QuEpACRlz19hhK/v42vL_QuEpACRlz19hhK -goversion go1.14.2 -symabis $WORK\b014\symabis -D  -importcfg $WORK\b014\importcfg -pack -asmhdr $WORK\b014\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\abs.go C:\Go\src\math\acosh.go C:\Go\src\math\asin.go C:\Go\src\math\asinh.go C:\Go\src\math\atan.go C:\Go\src\math\atan2.go C:\Go\src\math\atanh.go C:\Go\src\math\bits.go C:\Go\src\math\cbrt.go C:\Go\src\math\const.go C:\Go\src\math\copysign.go C:\Go\src\math\dim.go C:\Go\src\math\erf.go C:\Go\src\math\erfinv.go C:\Go\src\math\exp.go C:\Go\src\math\exp_asm.go C:\Go\src\math\expm1.go C:\Go\src\math\floor.go C:\Go\src\math\fma.go C:\Go\src\math\frexp.go C:\Go\src\math\gamma.go C:\Go\src\math\hypot.go C:\Go\src\math\j0.go C:\Go\src\math\j1.go C:\Go\src\math\jn.go C:\Go\src\math\ldexp.go C:\Go\src\math\lgamma.go C:\Go\src\math\log.go C:\Go\src\math\log10.go C:\Go\src\math\log1p.go C:\Go\src\math\logb.go C:\Go\src\math\mod.go C:\Go\src\math\modf.go C:\Go\src\math\nextafter.go C:\Go\src\math\pow.go C:\Go\src\math\pow10.go C:\Go\src\math\remainder.go C:\Go\src\math\signbit.go C:\Go\src\math\sin.go C:\Go\src\math\sincos.go C:\Go\src\math\sinh.go C:\Go\src\math\sqrt.go C:\Go\src\math\tan.go C:\Go\src\math\tanh.go C:\Go\src\math\trig_reduce.go C:\Go\src\math\unsafe.go]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b014\_pkg_.a -trimpath $WORK\b014=> -p math -std -buildid v42vL_QuEpACRlz19hhK/v42vL_QuEpACRlz19hhK -goversion go1.14.2 -symabis $WORK\b014\symabis -D  -importcfg $WORK\b014\importcfg -pack -asmhdr $WORK\b014\go_asm.h -c=4 C:\Go\src\math\abs.go C:\Go\src\math\acosh.go C:\Go\src\math\asin.go C:\Go\src\math\asinh.go C:\Go\src\math\atan.go C:\Go\src\math\atan2.go C:\Go\src\math\atanh.go C:\Go\src\math\bits.go C:\Go\src\math\cbrt.go C:\Go\src\math\const.go C:\Go\src\math\copysign.go C:\Go\src\math\dim.go C:\Go\src\math\erf.go C:\Go\src\math\erfinv.go C:\Go\src\math\exp.go C:\Go\src\math\exp_asm.go C:\Go\src\math\expm1.go C:\Go\src\math\floor.go C:\Go\src\math\fma.go C:\Go\src\math\frexp.go C:\Go\src\math\gamma.go C:\Go\src\math\hypot.go C:\Go\src\math\j0.go C:\Go\src\math\j1.go C:\Go\src\math\jn.go C:\Go\src\math\ldexp.go C:\Go\src\math\lgamma.go C:\Go\src\math\log.go C:\Go\src\math\log10.go C:\Go\src\math\log1p.go C:\Go\src\math\logb.go C:\Go\src\math\mod.go C:\Go\src\math\modf.go C:\Go\src\math\nextafter.go C:\Go\src\math\pow.go C:\Go\src\math\pow10.go C:\Go\src\math\remainder.go C:\Go\src\math\signbit.go C:\Go\src\math\sin.go C:\Go\src\math\sincos.go C:\Go\src\math\sinh.go C:\Go\src\math\sqrt.go C:\Go\src\math\tan.go C:\Go\src\math\tanh.go C:\Go\src\math\trig_reduce.go C:\Go\src\math\unsafe.go]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-p" "math" "-std" "-buildid" "v42vL_QuEpACRlz19hhK/v42vL_QuEpACRlz19hhK" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\go_asm.h" "-c=4" "C:\\Go\\src\\math\\abs.go" "C:\\Go\\src\\math\\acosh.go" "C:\\Go\\src\\math\\asin.go" "C:\\Go\\src\\math\\asinh.go" "C:\\Go\\src\\math\\atan.go" "C:\\Go\\src\\math\\atan2.go" "C:\\Go\\src\\math\\atanh.go" "C:\\Go\\src\\math\\bits.go" "C:\\Go\\src\\math\\cbrt.go" "C:\\Go\\src\\math\\const.go" "C:\\Go\\src\\math\\copysign.go" "C:\\Go\\src\\math\\dim.go" "C:\\Go\\src\\math\\erf.go" "C:\\Go\\src\\math\\erfinv.go" "C:\\Go\\src\\math\\exp.go" "C:\\Go\\src\\math\\exp_asm.go" "C:\\Go\\src\\math\\expm1.go" "C:\\Go\\src\\math\\floor.go" "C:\\Go\\src\\math\\fma.go" "C:\\Go\\src\\math\\frexp.go" "C:\\Go\\src\\math\\gamma.go" "C:\\Go\\src\\math\\hypot.go" "C:\\Go\\src\\math\\j0.go" "C:\\Go\\src\\math\\j1.go" "C:\\Go\\src\\math\\jn.go" "C:\\Go\\src\\math\\ldexp.go" "C:\\Go\\src\\math\\lgamma.go" "C:\\Go\\src\\math\\log.go" "C:\\Go\\src\\math\\log10.go" "C:\\Go\\src\\math\\log1p.go" "C:\\Go\\src\\math\\logb.go" "C:\\Go\\src\\math\\mod.go" "C:\\Go\\src\\math\\modf.go" "C:\\Go\\src\\math\\nextafter.go" "C:\\Go\\src\\math\\pow.go" "C:\\Go\\src\\math\\pow10.go" "C:\\Go\\src\\math\\remainder.go" "C:\\Go\\src\\math\\signbit.go" "C:\\Go\\src\\math\\sin.go" "C:\\Go\\src\\math\\sincos.go" "C:\\Go\\src\\math\\sinh.go" "C:\\Go\\src\\math\\sqrt.go" "C:\\Go\\src\\math\\tan.go" "C:\\Go\\src\\math\\tanh.go" "C:\\Go\\src\\math\\trig_reduce.go" "C:\\Go\\src\\math\\unsafe.go"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\dim_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\dim_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\dim_amd64.o C:\Go\src\math\dim_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\dim_amd64.o" "C:\\Go\\src\\math\\dim_amd64.s"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\exp_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\exp_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\exp_amd64.o C:\Go\src\math\exp_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\exp_amd64.o" "C:\\Go\\src\\math\\exp_amd64.s"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\floor_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\floor_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\floor_amd64.o C:\Go\src\math\floor_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\floor_amd64.o" "C:\\Go\\src\\math\\floor_amd64.s"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\hypot_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\hypot_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\hypot_amd64.o C:\Go\src\math\hypot_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\hypot_amd64.o" "C:\\Go\\src\\math\\hypot_amd64.s"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\log_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\log_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\log_amd64.o C:\Go\src\math\log_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\log_amd64.o" "C:\\Go\\src\\math\\log_amd64.s"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\sqrt_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\sqrt_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\sqrt_amd64.o C:\Go\src\math\sqrt_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\sqrt_amd64.o" "C:\\Go\\src\\math\\sqrt_amd64.s"`
```

```bash
# math
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\stubs_amd64.o]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\math\stubs_amd64.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b014=> -I $WORK\b014\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b014\stubs_amd64.o C:\Go\src\math\stubs_amd64.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b014\\stubs_amd64.o" "C:\\Go\\src\\math\\stubs_amd64.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:11 Package:
instrumenter: 2021/02/21 12:10:11 stdlib:  false
instrumenter: 2021/02/21 12:10:11 output:  0xc000116010
instrumenter: 2021/02/21 12:10:11 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b005\symabis]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\runtime\asm.s C:\Go\src\runtime\asm_amd64.s C:\Go\src\runtime\duff_amd64.s C:\Go\src\runtime\memclr_amd64.s C:\Go\src\runtime\memmove_amd64.s C:\Go\src\runtime\preempt_amd64.sC:\Go\src\runtime\rt0_windows_amd64.s C:\Go\src\runtime\sys_windows_amd64.s C:\Go\src\runtime\zcallback_windows.s]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b005\symabis C:\Go\src\runtime\asm.s C:\Go\src\runtime\asm_amd64.s C:\Go\src\runtime\duff_amd64.s C:\Go\src\runtime\memclr_amd64.s C:\Go\src\runtime\memmove_amd64.s C:\Go\src\runtime\preempt_amd64.s C:\Go\src\runtime\rt0_windows_amd64.s C:\Go\src\runtime\sys_windows_amd64.s C:\Go\src\runtime\zcallback_windows.s]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\symabis" "C:\\Go\\src\\runtime\\asm.s" "C:\\Go\\src\\runtime\\asm_amd64.s" "C:\\Go\\src\\runtime\\duff_amd64.s" "C:\\Go\\src\\runtime\\memclr_amd64.s" "C:\\Go\\src\\runtime\\memmove_amd64.s" "C:\\Go\\src\\runtime\\preempt_amd64.s" "C:\\Go\\src\\runtime\\rt0_windows_amd64.s" "C:\\Go\\src\\runtime\\sys_windows_amd64.s" "C:\\Go\\src\\runtime\\zcallback_windows.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:11 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:11 Package:  runtime
instrumenter: 2021/02/21 12:10:11 stdlib:  true
instrumenter: 2021/02/21 12:10:11 output:  0xc00011c010
instrumenter: 2021/02/21 12:10:11 args:  [-o $WORK\b005\_pkg_.a -trimpath $WORK\b005=> -p runtime -std -+ -buildid iTvJR38i3qSocelUSHI8/iTvJR38i3qSocelUSHI8 -goversion go1.14.2 -symabis $WORK\b005\symabis -D  -importcfg $WORK\b005\importcfg -pack -asmhdr $WORK\b005\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:11 files:  [C:\Go\src\runtime\alg.go C:\Go\src\runtime\atomic_pointer.go C:\Go\src\runtime\auxv_none.go C:\Go\src\runtime\cgo.go C:\Go\src\runtime\cgocall.go C:\Go\src\runtime\cgocallback.go C:\Go\src\runtime\cgocheck.go C:\Go\src\runtime\chan.go C:\Go\src\runtime\checkptr.go C:\Go\src\runtime\compiler.go C:\Go\src\runtime\complex.go C:\Go\src\runtime\cpuflags.go C:\Go\src\runtime\cpuflags_amd64.go C:\Go\src\runtime\cpuprof.go C:\Go\src\runtime\cputicks.go C:\Go\src\runtime\debug.go C:\Go\src\runtime\debugcall.go C:\Go\src\runtime\debuglog.go C:\Go\src\runtime\debuglog_off.go C:\Go\src\runtime\defs_windows_amd64.go C:\Go\src\runtime\env_posix.go C:\Go\src\runtime\error.go C:\Go\src\runtime\extern.go C:\Go\src\runtime\fastlog2.go C:\Go\src\runtime\fastlog2table.go C:\Go\src\runtime\float.go C:\Go\src\runtime\hash64.go C:\Go\src\runtime\heapdump.go C:\Go\src\runtime\iface.go C:\Go\src\runtime\lfstack.go C:\Go\src\runtime\lfstack_64bit.go C:\Go\src\runtime\lock_sema.go C:\Go\src\runtime\malloc.go C:\Go\src\runtime\map.go C:\Go\src\runtime\map_fast32.go C:\Go\src\runtime\map_fast64.go C:\Go\src\runtime\map_faststr.go C:\Go\src\runtime\mbarrier.go C:\Go\src\runtime\mbitmap.go C:\Go\src\runtime\mcache.go C:\Go\src\runtime\mcentral.go C:\Go\src\runtime\mem_windows.go C:\Go\src\runtime\mfinal.go C:\Go\src\runtime\mfixalloc.go C:\Go\src\runtime\mgc.go C:\Go\src\runtime\mgcmark.go C:\Go\src\runtime\mgcscavenge.go C:\Go\src\runtime\mgcstack.go C:\Go\src\runtime\mgcsweep.go C:\Go\src\runtime\mgcsweepbuf.go C:\Go\src\runtime\mgcwork.go C:\Go\src\runtime\mheap.go C:\Go\src\runtime\mpagealloc.go C:\Go\src\runtime\mpagealloc_64bit.go C:\Go\src\runtime\mpagecache.go C:\Go\src\runtime\mpallocbits.go C:\Go\src\runtime\mprof.go C:\Go\src\runtime\mranges.go C:\Go\src\runtime\msan0.go C:\Go\src\runtime\msize.go C:\Go\src\runtime\mstats.go C:\Go\src\runtime\mwbbuf.go C:\Go\src\runtime\netpoll.go C:\Go\src\runtime\netpoll_windows.go C:\Go\src\runtime\os_nonopenbsd.go C:\Go\src\runtime\os_windows.go C:\Go\src\runtime\panic.go C:\Go\src\runtime\plugin.go C:\Go\src\runtime\preempt.go C:\Go\src\runtime\print.go C:\Go\src\runtime\proc.go C:\Go\src\runtime\profbuf.go C:\Go\src\runtime\proflabel.go C:\Go\src\runtime\race0.go C:\Go\src\runtime\rdebug.go C:\Go\src\runtime\runtime.go C:\Go\src\runtime\runtime1.go C:\Go\src\runtime\runtime2.go C:\Go\src\runtime\rwmutex.go C:\Go\src\runtime\select.go C:\Go\src\runtime\sema.go C:\Go\src\runtime\signal_windows.go C:\Go\src\runtime\sigqueue.go C:\Go\src\runtime\sigqueue_note.go C:\Go\src\runtime\sizeclasses.go C:\Go\src\runtime\slice.go C:\Go\src\runtime\softfloat64.go C:\Go\src\runtime\stack.go C:\Go\src\runtime\string.go C:\Go\src\runtime\stubs.go C:\Go\src\runtime\stubs3.go C:\Go\src\runtime\stubs_amd64.go C:\Go\src\runtime\stubs_nonlinux.go C:\Go\src\runtime\symtab.go C:\Go\src\runtime\sys_nonppc64x.go C:\Go\src\runtime\sys_x86.go C:\Go\src\runtime\syscall_windows.go C:\Go\src\runtime\time.go C:\Go\src\runtime\time_nofake.go C:\Go\src\runtime\timeasm.go C:\Go\src\runtime\trace.go C:\Go\src\runtime\traceback.go C:\Go\src\runtime\type.go C:\Go\src\runtime\typekind.go C:\Go\src\runtime\utf8.go C:\Go\src\runtime\vdso_in_none.go C:\Go\src\runtime\write_err.go C:\Go\src\runtime\zcallback_windows.go]
instrumenter: 2021/02/21 12:10:11 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b005\_pkg_.a -trimpath $WORK\b005=> -p runtime -std -+ -buildid iTvJR38i3qSocelUSHI8/iTvJR38i3qSocelUSHI8 -goversion go1.14.2 -symabis $WORK\b005\symabis -D  -importcfg $WORK\b005\importcfg -pack -asmhdr $WORK\b005\go_asm.h -c=4 C:\Go\src\runtime\alg.go C:\Go\src\runtime\atomic_pointer.go C:\Go\src\runtime\auxv_none.go C:\Go\src\runtime\cgo.go C:\Go\src\runtime\cgocall.go C:\Go\src\runtime\cgocallback.go C:\Go\src\runtime\cgocheck.go C:\Go\src\runtime\chan.go C:\Go\src\runtime\checkptr.go C:\Go\src\runtime\compiler.go C:\Go\src\runtime\complex.go C:\Go\src\runtime\cpuflags.go C:\Go\src\runtime\cpuflags_amd64.go C:\Go\src\runtime\cpuprof.go C:\Go\src\runtime\cputicks.go C:\Go\src\runtime\debug.go C:\Go\src\runtime\debugcall.go C:\Go\src\runtime\debuglog.go C:\Go\src\runtime\debuglog_off.go C:\Go\src\runtime\defs_windows_amd64.go C:\Go\src\runtime\env_posix.go C:\Go\src\runtime\error.go C:\Go\src\runtime\extern.go C:\Go\src\runtime\fastlog2.go C:\Go\src\runtime\fastlog2table.go C:\Go\src\runtime\float.go C:\Go\src\runtime\hash64.goC:\Go\src\runtime\heapdump.go C:\Go\src\runtime\iface.go C:\Go\src\runtime\lfstack.go C:\Go\src\runtime\lfstack_64bit.go C:\Go\src\runtime\lock_sema.go C:\Go\src\runtime\malloc.go C:\Go\src\runtime\map.go C:\Go\src\runtime\map_fast32.go C:\Go\src\runtime\map_fast64.go C:\Go\src\runtime\map_faststr.go C:\Go\src\runtime\mbarrier.go C:\Go\src\runtime\mbitmap.go C:\Go\src\runtime\mcache.go C:\Go\src\runtime\mcentral.go C:\Go\src\runtime\mem_windows.go C:\Go\src\runtime\mfinal.go C:\Go\src\runtime\mfixalloc.go C:\Go\src\runtime\mgc.go C:\Go\src\runtime\mgcmark.go C:\Go\src\runtime\mgcscavenge.go C:\Go\src\runtime\mgcstack.go C:\Go\src\runtime\mgcsweep.go C:\Go\src\runtime\mgcsweepbuf.go C:\Go\src\runtime\mgcwork.go C:\Go\src\runtime\mheap.go C:\Go\src\runtime\mpagealloc.go C:\Go\src\runtime\mpagealloc_64bit.go C:\Go\src\runtime\mpagecache.go C:\Go\src\runtime\mpallocbits.go C:\Go\src\runtime\mprof.go C:\Go\src\runtime\mranges.go C:\Go\src\runtime\msan0.go C:\Go\src\runtime\msize.go C:\Go\src\runtime\mstats.go C:\Go\src\runtime\mwbbuf.go C:\Go\src\runtime\netpoll.go C:\Go\src\runtime\netpoll_windows.go C:\Go\src\runtime\os_nonopenbsd.go C:\Go\src\runtime\os_windows.go C:\Go\src\runtime\panic.go C:\Go\src\runtime\plugin.go C:\Go\src\runtime\preempt.go C:\Go\src\runtime\print.go C:\Go\src\runtime\proc.go C:\Go\src\runtime\profbuf.go C:\Go\src\runtime\proflabel.go C:\Go\src\runtime\race0.go C:\Go\src\runtime\rdebug.go C:\Go\src\runtime\runtime.go C:\Go\src\runtime\runtime1.go C:\Go\src\runtime\runtime2.go C:\Go\src\runtime\rwmutex.go C:\Go\src\runtime\select.go C:\Go\src\runtime\sema.go C:\Go\src\runtime\signal_windows.go C:\Go\src\runtime\sigqueue.go C:\Go\src\runtime\sigqueue_note.go C:\Go\src\runtime\sizeclasses.go C:\Go\src\runtime\slice.go C:\Go\src\runtime\softfloat64.go C:\Go\src\runtime\stack.go C:\Go\src\runtime\string.go C:\Go\src\runtime\stubs.go C:\Go\src\runtime\stubs3.go C:\Go\src\runtime\stubs_amd64.go C:\Go\src\runtime\stubs_nonlinux.go C:\Go\src\runtime\symtab.go C:\Go\src\runtime\sys_nonppc64x.go C:\Go\src\runtime\sys_x86.go C:\Go\src\runtime\syscall_windows.go C:\Go\src\runtime\time.go C:\Go\src\runtime\time_nofake.go C:\Go\src\runtime\timeasm.go C:\Go\src\runtime\trace.go C:\Go\src\runtime\traceback.go C:\Go\src\runtime\type.go C:\Go\src\runtime\typekind.go C:\Go\src\runtime\utf8.go C:\Go\src\runtime\vdso_in_none.go C:\Go\src\runtime\write_err.go C:\Go\src\runtime\zcallback_windows.go]
instrumenter: 2021/02/21 12:10:11 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-p" "runtime" "-std" "-+" "-buildid" "iTvJR38i3qSocelUSHI8/iTvJR38i3qSocelUSHI8" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\go_asm.h" "-c=4" "C:\\Go\\src\\runtime\\alg.go" "C:\\Go\\src\\runtime\\atomic_pointer.go" "C:\\Go\\src\\runtime\\auxv_none.go" "C:\\Go\\src\\runtime\\cgo.go" "C:\\Go\\src\\runtime\\cgocall.go" "C:\\Go\\src\\runtime\\cgocallback.go" "C:\\Go\\src\\runtime\\cgocheck.go" "C:\\Go\\src\\runtime\\chan.go" "C:\\Go\\src\\runtime\\checkptr.go" "C:\\Go\\src\\runtime\\compiler.go" "C:\\Go\\src\\runtime\\complex.go" "C:\\Go\\src\\runtime\\cpuflags.go" "C:\\Go\\src\\runtime\\cpuflags_amd64.go" "C:\\Go\\src\\runtime\\cpuprof.go" "C:\\Go\\src\\runtime\\cputicks.go" "C:\\Go\\src\\runtime\\debug.go" "C:\\Go\\src\\runtime\\debugcall.go" "C:\\Go\\src\\runtime\\debuglog.go" "C:\\Go\\src\\runtime\\debuglog_off.go" "C:\\Go\\src\\runtime\\defs_windows_amd64.go" "C:\\Go\\src\\runtime\\env_posix.go" "C:\\Go\\src\\runtime\\error.go" "C:\\Go\\src\\runtime\\extern.go" "C:\\Go\\src\\runtime\\fastlog2.go" "C:\\Go\\src\\runtime\\fastlog2table.go" "C:\\Go\\src\\runtime\\float.go" "C:\\Go\\src\\runtime\\hash64.go" "C:\\Go\\src\\runtime\\heapdump.go" "C:\\Go\\src\\runtime\\iface.go" "C:\\Go\\src\\runtime\\lfstack.go" "C:\\Go\\src\\runtime\\lfstack_64bit.go" "C:\\Go\\src\\runtime\\lock_sema.go" "C:\\Go\\src\\runtime\\malloc.go" "C:\\Go\\src\\runtime\\map.go" "C:\\Go\\src\\runtime\\map_fast32.go" "C:\\Go\\src\\runtime\\map_fast64.go" "C:\\Go\\src\\runtime\\map_faststr.go" "C:\\Go\\src\\runtime\\mbarrier.go" "C:\\Go\\src\\runtime\\mbitmap.go" "C:\\Go\\src\\runtime\\mcache.go" "C:\\Go\\src\\runtime\\mcentral.go" "C:\\Go\\src\\runtime\\mem_windows.go" "C:\\Go\\src\\runtime\\mfinal.go" "C:\\Go\\src\\runtime\\mfixalloc.go" "C:\\Go\\src\\runtime\\mgc.go" "C:\\Go\\src\\runtime\\mgcmark.go" "C:\\Go\\src\\runtime\\mgcscavenge.go" "C:\\Go\\src\\runtime\\mgcstack.go" "C:\\Go\\src\\runtime\\mgcsweep.go" "C:\\Go\\src\\runtime\\mgcsweepbuf.go" "C:\\Go\\src\\runtime\\mgcwork.go" "C:\\Go\\src\\runtime\\mheap.go" "C:\\Go\\src\\runtime\\mpagealloc.go" "C:\\Go\\src\\runtime\\mpagealloc_64bit.go" "C:\\Go\\src\\runtime\\mpagecache.go" "C:\\Go\\src\\runtime\\mpallocbits.go" "C:\\Go\\src\\runtime\\mprof.go" "C:\\Go\\src\\runtime\\mranges.go" "C:\\Go\\src\\runtime\\msan0.go" "C:\\Go\\src\\runtime\\msize.go" "C:\\Go\\src\\runtime\\mstats.go" "C:\\Go\\src\\runtime\\mwbbuf.go" "C:\\Go\\src\\runtime\\netpoll.go" "C:\\Go\\src\\runtime\\netpoll_windows.go" "C:\\Go\\src\\runtime\\os_nonopenbsd.go" "C:\\Go\\src\\runtime\\os_windows.go" "C:\\Go\\src\\runtime\\panic.go" "C:\\Go\\src\\runtime\\plugin.go" "C:\\Go\\src\\runtime\\preempt.go" "C:\\Go\\src\\runtime\\print.go" "C:\\Go\\src\\runtime\\proc.go" "C:\\Go\\src\\runtime\\profbuf.go" "C:\\Go\\src\\runtime\\proflabel.go" "C:\\Go\\src\\runtime\\race0.go" "C:\\Go\\src\\runtime\\rdebug.go" "C:\\Go\\src\\runtime\\runtime.go" "C:\\Go\\src\\runtime\\runtime1.go" "C:\\Go\\src\\runtime\\runtime2.go" "C:\\Go\\src\\runtime\\rwmutex.go" "C:\\Go\\src\\runtime\\select.go" "C:\\Go\\src\\runtime\\sema.go" "C:\\Go\\src\\runtime\\signal_windows.go" "C:\\Go\\src\\runtime\\sigqueue.go" "C:\\Go\\src\\runtime\\sigqueue_note.go" "C:\\Go\\src\\runtime\\sizeclasses.go" "C:\\Go\\src\\runtime\\slice.go" "C:\\Go\\src\\runtime\\softfloat64.go" "C:\\Go\\src\\runtime\\stack.go" "C:\\Go\\src\\runtime\\string.go" "C:\\Go\\src\\runtime\\stubs.go" "C:\\Go\\src\\runtime\\stubs3.go" "C:\\Go\\src\\runtime\\stubs_amd64.go" "C:\\Go\\src\\runtime\\stubs_nonlinux.go" "C:\\Go\\src\\runtime\\symtab.go" "C:\\Go\\src\\runtime\\sys_nonppc64x.go" "C:\\Go\\src\\runtime\\sys_x86.go" "C:\\Go\\src\\runtime\\syscall_windows.go" "C:\\Go\\src\\runtime\\time.go" "C:\\Go\\src\\runtime\\time_nofake.go" "C:\\Go\\src\\runtime\\timeasm.go" "C:\\Go\\src\\runtime\\trace.go" "C:\\Go\\src\\runtime\\traceback.go" "C:\\Go\\src\\runtime\\type.go" "C:\\Go\\src\\runtime\\typekind.go" "C:\\Go\\src\\runtime\\utf8.go" "C:\\Go\\src\\runtime\\vdso_in_none.go" "C:\\Go\\src\\runtime\\write_err.go" "C:\\Go\\src\\runtime\\zcallback_windows.go"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc0000dc010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\asm.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\asm.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\asm.o C:\Go\src\runtime\asm.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\asm.o" "C:\\Go\\src\\runtime\\asm.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc0000de010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\asm_amd64.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\asm_amd64.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\asm_amd64.o C:\Go\src\runtime\asm_amd64.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\asm_amd64.o" "C:\\Go\\src\\runtime\\asm_amd64.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc0000de010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\duff_amd64.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\duff_amd64.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\duff_amd64.o C:\Go\src\runtime\duff_amd64.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\duff_amd64.o" "C:\\Go\\src\\runtime\\duff_amd64.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc000116010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\memclr_amd64.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\memclr_amd64.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\memclr_amd64.o C:\Go\src\runtime\memclr_amd64.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\memclr_amd64.o" "C:\\Go\\src\\runtime\\memclr_amd64.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc000116010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\memmove_amd64.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\memmove_amd64.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\memmove_amd64.o C:\Go\src\runtime\memmove_amd64.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\memmove_amd64.o" "C:\\Go\\src\\runtime\\memmove_amd64.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc0000de010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\preempt_amd64.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\preempt_amd64.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\preempt_amd64.o C:\Go\src\runtime\preempt_amd64.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\preempt_amd64.o" "C:\\Go\\src\\runtime\\preempt_amd64.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc000116010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\rt0_windows_amd64.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\rt0_windows_amd64.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\rt0_windows_amd64.o C:\Go\src\runtime\rt0_windows_amd64.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\rt0_windows_amd64.o" "C:\\Go\\src\\runtime\\rt0_windows_amd64.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc000116010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\sys_windows_amd64.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\sys_windows_amd64.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\sys_windows_amd64.o C:\Go\src\runtime\sys_windows_amd64.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\sys_windows_amd64.o" "C:\\Go\\src\\runtime\\sys_windows_amd64.s"`
```

```bash
# runtime
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc000116010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\zcallback_windows.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\runtime\zcallback_windows.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b005=> -I $WORK\b005\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b005\zcallback_windows.o C:\Go\src\runtime\zcallback_windows.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b005\\zcallback_windows.o" "C:\\Go\\src\\runtime\\zcallback_windows.s"`

# sync
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:12 Package:  sync
instrumenter: 2021/02/21 12:10:12 stdlib:  true
instrumenter: 2021/02/21 12:10:12 output:  0xc000116010
instrumenter: 2021/02/21 12:10:12 args:  [-o $WORK\b018\_pkg_.a -trimpath $WORK\b018=> -p sync -std -buildid w6gD1-B21NztnfaOy6F6/w6gD1-B21NztnfaOy6F6 -goversion go1.14.2 -D  -importcfg $WORK\b018\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\sync\cond.go C:\Go\src\sync\map.go C:\Go\src\sync\mutex.go C:\Go\src\sync\once.go C:\Go\src\sync\pool.go C:\Go\src\sync\poolqueue.go C:\Go\src\sync\runtime.go C:\Go\src\sync\rwmutex.go C:\Go\src\sync\waitgroup.go]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b018\_pkg_.a -trimpath $WORK\b018=> -p sync -std -buildid w6gD1-B21NztnfaOy6F6/w6gD1-B21NztnfaOy6F6 -goversion go1.14.2 -D  -importcfg $WORK\b018\importcfg -pack -c=4 C:\Go\src\sync\cond.go C:\Go\src\sync\map.go C:\Go\src\sync\mutex.go C:\Go\src\sync\once.go C:\Go\src\sync\pool.go C:\Go\src\sync\poolqueue.go C:\Go\src\sync\runtime.go C:\Go\src\sync\rwmutex.go C:\Go\src\sync\waitgroup.go]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b018\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b018=>" "-p" "sync" "-std" "-buildid" "w6gD1-B21NztnfaOy6F6/w6gD1-B21NztnfaOy6F6" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b018\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\sync\\cond.go" "C:\\Go\\src\\sync\\map.go" "C:\\Go\\src\\sync\\mutex.go" "C:\\Go\\src\\sync\\once.go" "C:\\Go\\src\\sync\\pool.go" "C:\\Go\\src\\sync\\poolqueue.go" "C:\\Go\\src\\sync\\runtime.go" "C:\\Go\\src\\sync\\rwmutex.go" "C:\\Go\\src\\sync\\waitgroup.go"`
```

```bash
# internal/reflectlite
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc0000dc010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b004=> -I $WORK\b004\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b004\symabis]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\internal\reflectlite\asm.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b004=> -I $WORK\b004\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b004\symabis C:\Go\src\internal\reflectlite\asm.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004\\symabis" "C:\\Go\\src\\internal\\reflectlite\\asm.s"`
```

```bash
# internal/reflectlite
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:12 Package:  internal/reflectlite
instrumenter: 2021/02/21 12:10:12 stdlib:  true
instrumenter: 2021/02/21 12:10:12 output:  0xc0000de010
instrumenter: 2021/02/21 12:10:12 args:  [-o $WORK\b004\_pkg_.a -trimpath $WORK\b004=> -p internal/reflectlite -std -buildid Y8UGe87RKPo-HFcvp3Xu/Y8UGe87RKPo-HFcvp3Xu -goversion go1.14.2 -symabis $WORK\b004\symabis -D  -importcfg $WORK\b004\importcfg -pack -asmhdr $WORK\b004\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\internal\reflectlite\swapper.go C:\Go\src\internal\reflectlite\type.go C:\Go\src\internal\reflectlite\value.go]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b004\_pkg_.a -trimpath $WORK\b004=> -p internal/reflectlite -std -buildid Y8UGe87RKPo-HFcvp3Xu/Y8UGe87RKPo-HFcvp3Xu -goversion go1.14.2 -symabis $WORK\b004\symabis -D  -importcfg $WORK\b004\importcfg -pack -asmhdr $WORK\b004\go_asm.h -c=4 C:\Go\src\internal\reflectlite\swapper.go C:\Go\src\internal\reflectlite\type.go C:\Go\src\internal\reflectlite\value.go]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004=>" "-p" "internal/reflectlite" "-std" "-buildid" "Y8UGe87RKPo-HFcvp3Xu/Y8UGe87RKPo-HFcvp3Xu" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004\\go_asm.h" "-c=4" "C:\\Go\\src\\internal\\reflectlite\\swapper.go" "C:\\Go\\src\\internal\\reflectlite\\type.go" "C:\\Go\\src\\internal\\reflectlite\\value.go"`
```

```bash
# internal/reflectlite
instrumenter: 2021/02/21 12:10:12 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:12 Package:
instrumenter: 2021/02/21 12:10:12 stdlib:  false
instrumenter: 2021/02/21 12:10:12 output:  0xc0000dc010
instrumenter: 2021/02/21 12:10:12 args:  [-trimpath $WORK\b004=> -I $WORK\b004\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b004\asm.o]
instrumenter: 2021/02/21 12:10:12 files:  [C:\Go\src\internal\reflectlite\asm.s]
instrumenter: 2021/02/21 12:10:12 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b004=> -I $WORK\b004\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b004\asm.o C:\Go\src\internal\reflectlite\asm.s]
instrumenter: 2021/02/21 12:10:12 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b004\\asm.o" "C:\\Go\\src\\internal\\reflectlite\\asm.s"`
```

```bash
# errors
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  errors
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc000118010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b003\_pkg_.a -trimpath $WORK\b003=> -p errors -std -complete -buildid 0PRcBv1RBUdAFD9I9MPP/0PRcBv1RBUdAFD9I9MPP -goversion go1.14.2 -D  -importcfg $WORK\b003\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\errors\errors.go C:\Go\src\errors\wrap.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b003\_pkg_.a -trimpath $WORK\b003=> -p errors -std -complete -buildid 0PRcBv1RBUdAFD9I9MPP/0PRcBv1RBUdAFD9I9MPP -goversion go1.14.2 -D  -importcfg $WORK\b003\importcfg -pack -c=4 C:\Go\src\errors\errors.go C:\Go\src\errors\wrap.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b003\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b003=>" "-p" "errors" "-std" "-complete" "-buildid" "0PRcBv1RBUdAFD9I9MPP/0PRcBv1RBUdAFD9I9MPP" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b003\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\errors\\errors.go" "C:\\Go\\src\\errors\\wrap.go"`
```

```bash
# sort
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  sort
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc0000dc010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b022\_pkg_.a -trimpath $WORK\b022=> -p sort -std -complete -buildid n9q1KFzZAomUd9OGsGK_/n9q1KFzZAomUd9OGsGK_ -goversion go1.14.2 -D  -importcfg $WORK\b022\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\sort\search.go C:\Go\src\sort\slice.go C:\Go\src\sort\slice_go113.go C:\Go\src\sort\sort.go C:\Go\src\sort\zfuncversion.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b022\_pkg_.a -trimpath $WORK\b022=> -p sort -std -complete -buildid n9q1KFzZAomUd9OGsGK_/n9q1KFzZAomUd9OGsGK_ -goversion go1.14.2-D  -importcfg $WORK\b022\importcfg -pack -c=4 C:\Go\src\sort\search.go C:\Go\src\sort\slice.go C:\Go\src\sort\slice_go113.go C:\Go\src\sort\sort.go C:\Go\src\sort\zfuncversion.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b022\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b022=>" "-p" "sort" "-std" "-complete" "-buildid" "n9q1KFzZAomUd9OGsGK_/n9q1KFzZAomUd9OGsGK_" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b022\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\sort\\search.go" "C:\\Go\\src\\sort\\slice.go" "C:\\Go\\src\\sort\\slice_go113.go" "C:\\Go\\src\\sort\\sort.go" "C:\\Go\\src\\sort\\zfuncversion.go"`
```

```bash
# internal/oserror
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  internal/oserror
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc000118010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b025\_pkg_.a -trimpath $WORK\b025=> -p internal/oserror -std -complete -buildid M-5Y1LWkfTFlKEy6fxQN/M-5Y1LWkfTFlKEy6fxQN -goversion go1.14.2 -D  -importcfg $WORK\b025\importcfg-pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\internal\oserror\errors.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b025\_pkg_.a -trimpath $WORK\b025=> -p internal/oserror -std -complete -buildid M-5Y1LWkfTFlKEy6fxQN/M-5Y1LWkfTFlKEy6fxQN -goversion go1.14.2 -D  -importcfg $WORK\b025\importcfg -pack -c=4 C:\Go\src\internal\oserror\errors.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b025\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b025=>" "-p" "internal/oserror" "-std" "-complete" "-buildid" "M-5Y1LWkfTFlKEy6fxQN/M-5Y1LWkfTFlKEy6fxQN" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b025\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\oserror\\errors.go"`
```

```bash
# io
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  io
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc000118010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b023\_pkg_.a -trimpath $WORK\b023=> -p io -std -complete -buildid Yifk7lvCDGVN4vQdz_hn/Yifk7lvCDGVN4vQdz_hn -goversion go1.14.2 -D  -importcfg $WORK\b023\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\io\io.go C:\Go\src\io\multi.go C:\Go\src\io\pipe.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b023\_pkg_.a -trimpath $WORK\b023=> -p io -std -complete -buildid Yifk7lvCDGVN4vQdz_hn/Yifk7lvCDGVN4vQdz_hn -goversion go1.14.2 -D  -importcfg $WORK\b023\importcfg -pack -c=4 C:\Go\src\io\io.go C:\Go\src\io\multi.go C:\Go\src\io\pipe.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b023\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b023=>" "-p" "io" "-std" "-complete" "-buildid" "Yifk7lvCDGVN4vQdz_hn/Yifk7lvCDGVN4vQdz_hn" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b023\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\io\\io.go" "C:\\Go\\src\\io\\multi.go" "C:\\Go\\src\\io\\pipe.go"`
```

```bash
# strconv
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  strconv
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc0000de010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b016\_pkg_.a -trimpath $WORK\b016=> -p strconv -std -complete -buildid dXWPha3CawVjNZTJz3_C/dXWPha3CawVjNZTJz3_C -goversion go1.14.2 -D  -importcfg $WORK\b016\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\strconv\atob.go C:\Go\src\strconv\atof.go C:\Go\src\strconv\atoi.go C:\Go\src\strconv\decimal.go C:\Go\src\strconv\doc.go C:\Go\src\strconv\extfloat.go C:\Go\src\strconv\ftoa.go C:\Go\src\strconv\isprint.go C:\Go\src\strconv\itoa.go C:\Go\src\strconv\quote.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b016\_pkg_.a -trimpath $WORK\b016=> -p strconv -std -complete -buildid dXWPha3CawVjNZTJz3_C/dXWPha3CawVjNZTJz3_C -goversion go1.14.2 -D  -importcfg $WORK\b016\importcfg -pack -c=4 C:\Go\src\strconv\atob.go C:\Go\src\strconv\atof.go C:\Go\src\strconv\atoi.go C:\Go\src\strconv\decimal.go C:\Go\src\strconv\doc.go C:\Go\src\strconv\extfloat.go C:\Go\src\strconv\ftoa.go C:\Go\src\strconv\isprint.go C:\Go\src\strconv\itoa.go C:\Go\src\strconv\quote.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b016\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b016=>" "-p" "strconv" "-std" "-complete" "-buildid" "dXWPha3CawVjNZTJz3_C/dXWPha3CawVjNZTJz3_C" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b016\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\strconv\\atob.go" "C:\\Go\\src\\strconv\\atof.go" "C:\\Go\\src\\strconv\\atoi.go" "C:\\Go\\src\\strconv\\decimal.go" "C:\\Go\\src\\strconv\\doc.go" "C:\\Go\\src\\strconv\\extfloat.go" "C:\\Go\\src\\strconv\\ftoa.go" "C:\\Go\\src\\strconv\\isprint.go" "C:\\Go\\src\\strconv\\itoa.go" "C:\\Go\\src\\strconv\\quote.go"`
```

```bash
# syscall
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:13 Package:
instrumenter: 2021/02/21 12:10:13 stdlib:  false
instrumenter: 2021/02/21 12:10:13 output:  0xc000096010
instrumenter: 2021/02/21 12:10:13 args:  [-trimpath $WORK\b029=> -I $WORK\b029\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b029\symabis]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\syscall\asm_windows.s]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b029=> -I $WORK\b029\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b029\symabis C:\Go\src\syscall\asm_windows.s]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029\\symabis" "C:\\Go\\src\\syscall\\asm_windows.s"`
```

```bash
# syscall
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  syscall
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc00011a010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b029\_pkg_.a -trimpath $WORK\b029=> -p syscall -std -buildid fx4XtyzozfnIrS3P0jmX/fx4XtyzozfnIrS3P0jmX -goversion go1.14.2 -symabis $WORK\b029\symabis -D  -importcfg $WORK\b029\importcfg -pack -asmhdr $WORK\b029\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\syscall\dll_windows.go C:\Go\src\syscall\endian_little.go C:\Go\src\syscall\env_windows.go C:\Go\src\syscall\exec_windows.go C:\Go\src\syscall\msan0.go C:\Go\src\syscall\net.go C:\Go\src\syscall\security_windows.go C:\Go\src\syscall\str.go C:\Go\src\syscall\syscall.go C:\Go\src\syscall\syscall_windows.go C:\Go\src\syscall\syscall_windows_amd64.go C:\Go\src\syscall\time_nofake.go C:\Go\src\syscall\types_windows.go C:\Go\src\syscall\types_windows_amd64.go C:\Go\src\syscall\zerrors_windows.go C:\Go\src\syscall\zerrors_windows_amd64.go C:\Go\src\syscall\zsyscall_windows.go C:\Go\src\syscall\zsysnum_windows_amd64.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b029\_pkg_.a -trimpath $WORK\b029=> -p syscall -std -buildid fx4XtyzozfnIrS3P0jmX/fx4XtyzozfnIrS3P0jmX -goversion go1.14.2 -symabis $WORK\b029\symabis -D  -importcfg $WORK\b029\importcfg -pack -asmhdr $WORK\b029\go_asm.h -c=4 C:\Go\src\syscall\dll_windows.go C:\Go\src\syscall\endian_little.go C:\Go\src\syscall\env_windows.go C:\Go\src\syscall\exec_windows.go C:\Go\src\syscall\msan0.go C:\Go\src\syscall\net.go C:\Go\src\syscall\security_windows.go C:\Go\src\syscall\str.go C:\Go\src\syscall\syscall.go C:\Go\src\syscall\syscall_windows.go C:\Go\src\syscall\syscall_windows_amd64.go C:\Go\src\syscall\time_nofake.go C:\Go\src\syscall\types_windows.go C:\Go\src\syscall\types_windows_amd64.go C:\Go\src\syscall\zerrors_windows.go C:\Go\src\syscall\zerrors_windows_amd64.go C:\Go\src\syscall\zsyscall_windows.go C:\Go\src\syscall\zsysnum_windows_amd64.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029=>" "-p" "syscall" "-std" "-buildid" "fx4XtyzozfnIrS3P0jmX/fx4XtyzozfnIrS3P0jmX" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029\\go_asm.h" "-c=4" "C:\\Go\\src\\syscall\\dll_windows.go" "C:\\Go\\src\\syscall\\endian_little.go" "C:\\Go\\src\\syscall\\env_windows.go" "C:\\Go\\src\\syscall\\exec_windows.go" "C:\\Go\\src\\syscall\\msan0.go" "C:\\Go\\src\\syscall\\net.go" "C:\\Go\\src\\syscall\\security_windows.go" "C:\\Go\\src\\syscall\\str.go" "C:\\Go\\src\\syscall\\syscall.go" "C:\\Go\\src\\syscall\\syscall_windows.go" "C:\\Go\\src\\syscall\\syscall_windows_amd64.go" "C:\\Go\\src\\syscall\\time_nofake.go" "C:\\Go\\src\\syscall\\types_windows.go" "C:\\Go\\src\\syscall\\types_windows_amd64.go" "C:\\Go\\src\\syscall\\zerrors_windows.go" "C:\\Go\\src\\syscall\\zerrors_windows_amd64.go" "C:\\Go\\src\\syscall\\zsyscall_windows.go" "C:\\Go\\src\\syscall\\zsysnum_windows_amd64.go"`
```

```bash
# syscall
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:13 Package:
instrumenter: 2021/02/21 12:10:13 stdlib:  false
instrumenter: 2021/02/21 12:10:13 output:  0xc000116010
instrumenter: 2021/02/21 12:10:13 args:  [-trimpath $WORK\b029=> -I $WORK\b029\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b029\asm_windows.o]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\syscall\asm_windows.s]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b029=> -I $WORK\b029\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b029\asm_windows.o C:\Go\src\syscall\asm_windows.s]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b029\\asm_windows.o" "C:\\Go\\src\\syscall\\asm_windows.s"`
```

```bash
# internal/syscall/windows/registry
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  internal/syscall/windows/registry
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc00015c010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b032\_pkg_.a -trimpath $WORK\b032=> -p internal/syscall/windows/registry -std -complete -buildid 7mrTpu24KSkR2H12Fczm/7mrTpu24KSkR2H12Fczm -goversion go1.14.2 -D  -importcfg $WORK\b032\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\internal\syscall\windows\registry\key.go C:\Go\src\internal\syscall\windows\registry\syscall.go C:\Go\src\internal\syscall\windows\registry\value.go C:\Go\src\internal\syscall\windows\registry\zsyscall_windows.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b032\_pkg_.a -trimpath $WORK\b032=> -p internal/syscall/windows/registry -std -complete -buildid 7mrTpu24KSkR2H12Fczm/7mrTpu24KSkR2H12Fczm -goversion go1.14.2 -D  -importcfg $WORK\b032\importcfg -pack -c=4 C:\Go\src\internal\syscall\windows\registry\key.go C:\Go\src\internal\syscall\windows\registry\syscall.go C:\Go\src\internal\syscall\windows\registry\value.go C:\Go\src\internal\syscall\windows\registry\zsyscall_windows.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b032\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b032=>" "-p" "internal/syscall/windows/registry" "-std" "-complete" "-buildid" "7mrTpu24KSkR2H12Fczm/7mrTpu24KSkR2H12Fczm" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b032\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\syscall\\windows\\registry\\key.go" "C:\\Go\\src\\internal\\syscall\\windows\\registry\\syscall.go" "C:\\Go\\src\\internal\\syscall\\windows\\registry\\value.go" "C:\\Go\\src\\internal\\syscall\\windows\\registry\\zsyscall_windows.go"`
```

```bash
# internal/syscall/windows
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  internal/syscall/windows
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc00011a010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b027\_pkg_.a -trimpath $WORK\b027=> -p internal/syscall/windows -std -complete -buildid 78egX7aDM9EBW0k9ZtI6/78egX7aDM9EBW0k9ZtI6 -goversion go1.14.2 -D  -importcfg $WORK\b027\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\internal\syscall\windows\psapi_windows.go C:\Go\src\internal\syscall\windows\reparse_windows.go C:\Go\src\internal\syscall\windows\security_windows.go C:\Go\src\internal\syscall\windows\symlink_windows.go C:\Go\src\internal\syscall\windows\syscall_windows.go C:\Go\src\internal\syscall\windows\zsyscall_windows.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b027\_pkg_.a -trimpath $WORK\b027=> -p internal/syscall/windows -std -complete -buildid 78egX7aDM9EBW0k9ZtI6/78egX7aDM9EBW0k9ZtI6-goversion go1.14.2 -D  -importcfg $WORK\b027\importcfg -pack -c=4 C:\Go\src\internal\syscall\windows\psapi_windows.go C:\Go\src\internal\syscall\windows\reparse_windows.go C:\Go\src\internal\syscall\windows\security_windows.go C:\Go\src\internal\syscall\windows\symlink_windows.go C:\Go\src\internal\syscall\windows\syscall_windows.go C:\Go\src\internal\syscall\windows\zsyscall_windows.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b027\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b027=>" "-p" "internal/syscall/windows" "-std" "-complete" "-buildid" "78egX7aDM9EBW0k9ZtI6/78egX7aDM9EBW0k9ZtI6" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b027\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\syscall\\windows\\psapi_windows.go" "C:\\Go\\src\\internal\\syscall\\windows\\reparse_windows.go" "C:\\Go\\src\\internal\\syscall\\windows\\security_windows.go" "C:\\Go\\src\\internal\\syscall\\windows\\symlink_windows.go" "C:\\Go\\src\\internal\\syscall\\windows\\syscall_windows.go" "C:\\Go\\src\\internal\\syscall\\windows\\zsyscall_windows.go"`
```

```bash
# reflect
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:13 Package:
instrumenter: 2021/02/21 12:10:13 stdlib:  false
instrumenter: 2021/02/21 12:10:13 output:  0xc000116010
instrumenter: 2021/02/21 12:10:13 args:  [-trimpath $WORK\b013=> -I $WORK\b013\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b013\symabis]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\reflect\asm_amd64.s]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b013=> -I $WORK\b013\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -gensymabis -o $WORK\b013\symabis C:\Go\src\reflect\asm_amd64.s]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-gensymabis" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013\\symabis" "C:\\Go\\src\\reflect\\asm_amd64.s"`
```

```bash
# reflect
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  reflect
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc000116010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b013\_pkg_.a -trimpath $WORK\b013=> -p reflect -std -buildid HB08pl4YmDeyeWR4HtgU/HB08pl4YmDeyeWR4HtgU -goversion go1.14.2 -symabis $WORK\b013\symabis -D  -importcfg $WORK\b013\importcfg -pack -asmhdr $WORK\b013\go_asm.h -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\reflect\deepequal.go C:\Go\src\reflect\makefunc.go C:\Go\src\reflect\swapper.go C:\Go\src\reflect\type.go C:\Go\src\reflect\value.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b013\_pkg_.a -trimpath $WORK\b013=> -p reflect -std -buildid HB08pl4YmDeyeWR4HtgU/HB08pl4YmDeyeWR4HtgU -goversion go1.14.2 -symabis $WORK\b013\symabis -D  -importcfg $WORK\b013\importcfg -pack -asmhdr $WORK\b013\go_asm.h -c=4 C:\Go\src\reflect\deepequal.go C:\Go\src\reflect\makefunc.go C:\Go\src\reflect\swapper.go C:\Go\src\reflect\type.go C:\Go\src\reflect\value.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013=>" "-p" "reflect" "-std" "-buildid" "HB08pl4YmDeyeWR4HtgU/HB08pl4YmDeyeWR4HtgU" "-goversion" "go1.14.2" "-symabis" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013\\symabis" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013\\importcfg" "-pack" "-asmhdr" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013\\go_asm.h" "-c=4" "C:\\Go\\src\\reflect\\deepequal.go" "C:\\Go\\src\\reflect\\makefunc.go" "C:\\Go\\src\\reflect\\swapper.go" "C:\\Go\\src\\reflect\\type.go" "C:\\Go\\src\\reflect\\value.go"`
```

```bash
# reflect
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\asm.exe
instrumenter: 2021/02/21 12:10:13 Package:
instrumenter: 2021/02/21 12:10:13 stdlib:  false
instrumenter: 2021/02/21 12:10:13 output:  0xc000116010
instrumenter: 2021/02/21 12:10:13 args:  [-trimpath $WORK\b013=> -I $WORK\b013\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b013\asm_amd64.o]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\reflect\asm_amd64.s]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\asm.exe -trimpath $WORK\b013=> -I $WORK\b013\ -I C:\Go\pkg\include -D GOOS_windows -D GOARCH_amd64 -o $WORK\b013\asm_amd64.o C:\Go\src\reflect\asm_amd64.s]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\asm.exe "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013=>" "-I" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013\\" "-I" "C:\\Go\\pkg\\include" "-D" "GOOS_windows" "-D" "GOARCH_amd64" "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b013\\asm_amd64.o" "C:\\Go\\src\\reflect\\asm_amd64.s"`
```

```bash
# internal/syscall/execenv
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  internal/syscall/execenv
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc000116010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b033\_pkg_.a -trimpath $WORK\b033=> -p internal/syscall/execenv -std -complete -buildid 4JCJAW3mWDXKe_tY2wSZ/4JCJAW3mWDXKe_tY2wSZ -goversion go1.14.2 -D  -importcfg $WORK\b033\im
portcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\internal\syscall\execenv\execenv_windows.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b033\_pkg_.a -trimpath $WORK\b033=> -p internal/syscall/execenv -std -complete -buildid 4JCJAW3mWDXKe_tY2wSZ/4JCJAW3mWDXKe_tY2wSZ
-goversion go1.14.2 -D  -importcfg $WORK\b033\importcfg -pack -c=4 C:\Go\src\internal\syscall\execenv\execenv_windows.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b033\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b033=>" "-p" "internal/syscall/execenv" "-std" "-complete" "-buildid" "4JCJAW3mWDXKe_tY2wSZ/4JCJAW3mWDXKe_tY2wSZ" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b033\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\syscall\\execenv\\execenv_windows.go"`
```

```bash
# internal/fmtsort
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  internal/fmtsort
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc000118010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b012\_pkg_.a -trimpath $WORK\b012=> -p internal/fmtsort -std -complete -buildid P2yyJEJr0ZfmGV7l14Ld/P2yyJEJr0ZfmGV7l14Ld -goversion go1.14.2 -D  -importcfg $WORK\b012\importcfg-pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\internal\fmtsort\sort.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b012\_pkg_.a -trimpath $WORK\b012=> -p internal/fmtsort -std -complete -buildid P2yyJEJr0ZfmGV7l14Ld/P2yyJEJr0ZfmGV7l14Ld -goversion go1.14.2 -D  -importcfg $WORK\b012\importcfg -pack -c=4 C:\Go\src\internal\fmtsort\sort.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b012\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b012=>" "-p" "internal/fmtsort" "-std" "-complete" "-buildid" "P2yyJEJr0ZfmGV7l14Ld/P2yyJEJr0ZfmGV7l14Ld" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b012\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\fmtsort\\sort.go"`
```

```bash
# time
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  time
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc000116010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b031\_pkg_.a -trimpath $WORK\b031=> -p time -std -buildid yKnSIvj03T6uD38AQI8q/yKnSIvj03T6uD38AQI8q -goversion go1.14.2 -D  -importcfg $WORK\b031\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\time\format.go C:\Go\src\time\sleep.go C:\Go\src\time\sys_windows.go C:\Go\src\time\tick.go C:\Go\src\time\time.go C:\Go\src\time\zoneinfo.go C:\Go\src\time\zoneinfo_abbrs_windows.go C:\Go\src\time\zoneinfo_read.go C:\Go\src\time\zoneinfo_windows.go]instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b031\_pkg_.a -trimpath $WORK\b031=> -p time -std -buildid yKnSIvj03T6uD38AQI8q/yKnSIvj03T6uD38AQI8q -goversion go1.14.2 -D  -importcfg $WORK\b031\importcfg -pack -c=4 C:\Go\src\time\format.go C:\Go\src\time\sleep.go C:\Go\src\time\sys_windows.go C:\Go\src\time\tick.go C:\Go\src\time\time.go C:\Go\src\time\zoneinfo.go C:\Go\src\time\zoneinfo_abbrs_windows.go C:\Go\src\time\zoneinfo_read.go C:\Go\src\time\zoneinfo_windows.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b031\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b031=>" "-p" "time" "-std" "-buildid" "yKnSIvj03T6uD38AQI8q/yKnSIvj03T6uD38AQI8q" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b031\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\time\\format.go" "C:\\Go\\src\\time\\sleep.go" "C:\\Go\\src\\time\\sys_windows.go" "C:\\Go\\src\\time\\tick.go" "C:\\Go\\src\\time\\time.go" "C:\\Go\\src\\time\\zoneinfo.go" "C:\\Go\\src\\time\\zoneinfo_abbrs_windows.go" "C:\\Go\\src\\time\\zoneinfo_read.go" "C:\\Go\\src\\time\\zoneinfo_windows.go"`
```

```bash
# internal/poll
instrumenter: 2021/02/21 12:10:13 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:13 Package:  internal/poll
instrumenter: 2021/02/21 12:10:13 stdlib:  true
instrumenter: 2021/02/21 12:10:13 output:  0xc0000e2010
instrumenter: 2021/02/21 12:10:13 args:  [-o $WORK\b026\_pkg_.a -trimpath $WORK\b026=> -p internal/poll -std -buildid JttlkUAjQvxzpPTPE4ew/JttlkUAjQvxzpPTPE4ew -goversion go1.14.2 -D  -importcfg $WORK\b026\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:13 files:  [C:\Go\src\internal\poll\errno_windows.go C:\Go\src\internal\poll\fd.go C:\Go\src\internal\poll\fd_fsync_windows.go C:\Go\src\internal\poll\fd_mutex.go C:\Go\src\internal\poll\fd_poll_runtime.go C:\Go\src\internal\poll\fd_posix.go C:\Go\src\internal\poll\fd_windows.go C:\Go\src\internal\poll\hook_windows.go C:\Go\src\internal\poll\sendfile_windows.go C:\Go\src\internal\poll\sockopt.go C:\Go\src\internal\poll\sockopt_windows.go C:\Go\src\internal\poll\sockoptip.go]
instrumenter: 2021/02/21 12:10:13 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b026\_pkg_.a -trimpath $WORK\b026=> -p internal/poll -std -buildid JttlkUAjQvxzpPTPE4ew/JttlkUAjQvxzpPTPE4ew -goversion go1.14.2 -D  -importcfg $WORK\b026\importcfg -pack -c=4 C:\Go\src\internal\poll\errno_windows.go C:\Go\src\internal\poll\fd.go C:\Go\src\internal\poll\fd_fsync_windows.go C:\Go\src\internal\poll\fd_mutex.go C:\Go\src\internal\poll\fd_poll_runtime.go C:\Go\src\internal\poll\fd_posix.go C:\Go\src\internal\poll\fd_windows.go C:\Go\src\internal\poll\hook_windows.go C:\Go\src\internal\poll\sendfile_windows.go C:\Go\src\internal\poll\sockopt.go C:\Go\src\internal\poll\sockopt_windows.go C:\Go\src\internal\poll\sockoptip.go]
instrumenter: 2021/02/21 12:10:13 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b026\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b026=>" "-p" "internal/poll" "-std" "-buildid" "JttlkUAjQvxzpPTPE4ew/JttlkUAjQvxzpPTPE4ew" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b026\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\internal\\poll\\errno_windows.go" "C:\\Go\\src\\internal\\poll\\fd.go" "C:\\Go\\src\\internal\\poll\\fd_fsync_windows.go" "C:\\Go\\src\\internal\\poll\\fd_mutex.go" "C:\\Go\\src\\internal\\poll\\fd_poll_runtime.go" "C:\\Go\\src\\internal\\poll\\fd_posix.go" "C:\\Go\\src\\internal\\poll\\fd_windows.go" "C:\\Go\\src\\internal\\poll\\hook_windows.go" "C:\\Go\\src\\internal\\poll\\sendfile_windows.go" "C:\\Go\\src\\internal\\poll\\sockopt.go" "C:\\Go\\src\\internal\\poll\\sockopt_windows.go" "C:\\Go\\src\\internal\\poll\\sockoptip.go"`
```

```bash
# os
instrumenter: 2021/02/21 12:10:14 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:14 Package:  os
instrumenter: 2021/02/21 12:10:14 stdlib:  true
instrumenter: 2021/02/21 12:10:14 output:  0xc00011a010
instrumenter: 2021/02/21 12:10:14 args:  [-o $WORK\b024\_pkg_.a -trimpath $WORK\b024=> -p os -std -buildid 0HV8V0lK_XyoHssRipsi/0HV8V0lK_XyoHssRipsi -goversion go1.14.2 -D  -importcfg $WORK\b024\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:14 files:  [C:\Go\src\os\dir.go C:\Go\src\os\dir_windows.go C:\Go\src\os\env.go C:\Go\src\os\error.go C:\Go\src\os\error_errno.go C:\Go\src\os\error_posix.go C:\Go\src\os\exec.go C:\Go\src\os\exec_posix.go C:\Go\src\os\exec_windows.go C:\Go\src\os\executable.go C:\Go\src\os\executable_windows.go C:\Go\src\os\file.go C:\Go\src\os\file_posix.go C:\Go\src\os\file_windows.go C:\Go\src\os\getwd.go C:\Go\src\os\path.go C:\Go\src\os\path_windows.go C:\Go\src\os\proc.go C:\Go\src\os\rawconn.go C:\Go\src\os\removeall_noat.go C:\Go\src\os\stat.go C:\Go\src\os\stat_windows.go C:\Go\src\os\sticky_notbsd.go C:\Go\src\os\str.go C:\Go\src\os\sys.go C:\Go\src\os\sys_windows.go C:\Go\src\os\types.go C:\Go\src\os\types_windows.go]
instrumenter: 2021/02/21 12:10:14 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b024\_pkg_.a -trimpath $WORK\b024=> -p os -std -buildid 0HV8V0lK_XyoHssRipsi/0HV8V0lK_XyoHssRipsi -goversion go1.14.2 -D  -importcfg $WORK\b024\importcfg -pack -c=4 C:\Go\src\os\dir.go C:\Go\src\os\dir_windows.go C:\Go\src\os\env.go C:\Go\src\os\error.go C:\Go\src\os\error_errno.go C:\Go\src\os\error_posix.go C:\Go\src\os\exec.go C:\Go\src\os\exec_posix.goC:\Go\src\os\exec_windows.go C:\Go\src\os\executable.go C:\Go\src\os\executable_windows.go C:\Go\src\os\file.go C:\Go\src\os\file_posix.go C:\Go\src\os\file_windows.go C:\Go\src\os\getwd.go C:\Go\src\os\path.go C:\Go\src\os\path_windows.go C:\Go\src\os\proc.go C:\Go\src\os\rawconn.go C:\Go\src\os\removeall_noat.go C:\Go\src\os\stat.go C:\Go\src\os\stat_windows.go C:\Go\src\os\sticky_notbsd.go C:\Go\src\os\str.go C:\Go\src\os\sys.go C:\Go\src\os\sys_windows.go C:\Go\src\os\types.go C:\Go\src\os\types_windows.go]
instrumenter: 2021/02/21 12:10:14 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b024\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b024=>" "-p" "os" "-std" "-buildid" "0HV8V0lK_XyoHssRipsi/0HV8V0lK_XyoHssRipsi" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b024\\importcfg" "-pack" "-c=4" "C:\\Go\\src\\os\\dir.go" "C:\\Go\\src\\os\\dir_windows.go" "C:\\Go\\src\\os\\env.go" "C:\\Go\\src\\os\\error.go" "C:\\Go\\src\\os\\error_errno.go" "C:\\Go\\src\\os\\error_posix.go" "C:\\Go\\src\\os\\exec.go" "C:\\Go\\src\\os\\exec_posix.go" "C:\\Go\\src\\os\\exec_windows.go" "C:\\Go\\src\\os\\executable.go" "C:\\Go\\src\\os\\executable_windows.go" "C:\\Go\\src\\os\\file.go" "C:\\Go\\src\\os\\file_posix.go" "C:\\Go\\src\\os\\file_windows.go" "C:\\Go\\src\\os\\getwd.go" "C:\\Go\\src\\os\\path.go" "C:\\Go\\src\\os\\path_windows.go" "C:\\Go\\src\\os\\proc.go" "C:\\Go\\src\\os\\rawconn.go" "C:\\Go\\src\\os\\removeall_noat.go" "C:\\Go\\src\\os\\stat.go" "C:\\Go\\src\\os\\stat_windows.go" "C:\\Go\\src\\os\\sticky_notbsd.go" "C:\\Go\\src\\os\\str.go" "C:\\Go\\src\\os\\sys.go" "C:\\Go\\src\\os\\sys_windows.go" "C:\\Go\\src\\os\\types.go" "C:\\Go\\src\\os\\types_windows.go"`
```

```bash
# fmt
instrumenter: 2021/02/21 12:10:14 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:14 Package:  fmt
instrumenter: 2021/02/21 12:10:14 stdlib:  true
instrumenter: 2021/02/21 12:10:14 output:  0xc0000de010
instrumenter: 2021/02/21 12:10:14 args:  [-o $WORK\b002\_pkg_.a -trimpath $WORK\b002=> -p fmt -std -complete -buildid oT7FK11pC__XgTDvOw8t/oT7FK11pC__XgTDvOw8t -goversion go1.14.2 -D  -importcfg $WORK\b002\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:14 files:  [C:\Go\src\fmt\doc.go C:\Go\src\fmt\errors.go C:\Go\src\fmt\format.go C:\Go\src\fmt\print.go C:\Go\src\fmt\scan.go]
instrumenter: 2021/02/21 12:10:14 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b002\_pkg_.a -trimpath $WORK\b002=> -p fmt -std -complete -buildid oT7FK11pC__XgTDvOw8t/oT7FK11pC__XgTDvOw8t -goversion go1.14.2 -D  -importcfg $WORK\b002\importcfg -pack -c=4 C:\Go\src\fmt\doc.go C:\Go\src\fmt\errors.go C:\Go\src\fmt\format.go C:\Go\src\fmt\print.go C:\Go\src\fmt\scan.go]
instrumenter: 2021/02/21 12:10:14 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b002\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b002=>" "-p" "fmt" "-std" "-complete" "-buildid" "oT7FK11pC__XgTDvOw8t/oT7FK11pC__XgTDvOw8t" "-goversion" "go1.14.2" "-D" "" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b002\\importcfg""-pack" "-c=4" "C:\\Go\\src\\fmt\\doc.go" "C:\\Go\\src\\fmt\\errors.go" "C:\\Go\\src\\fmt\\format.go" "C:\\Go\\src\\fmt\\print.go" "C:\\Go\\src\\fmt\\scan.go"`
```

```bash
# _/D_/dev/go/src/github.com/skirmish/instrument/testdata/fibonacci
instrumenter: 2021/02/21 12:10:14 Program:  C:\Go\pkg\tool\windows_amd64\compile.exe
instrumenter: 2021/02/21 12:10:14 Package:  main
instrumenter: 2021/02/21 12:10:14 stdlib:  false
instrumenter: 2021/02/21 12:10:14 output:  0xc000116010
instrumenter: 2021/02/21 12:10:14 args:  [-o $WORK\b001\_pkg_.a -trimpath $WORK\b001=> -p main -complete -buildid 07gIkXLEJooNJ0O9y49j/07gIkXLEJooNJ0O9y49j -goversion go1.14.2 -D _/D_/dev/go/src/github.com/skirmish/instrument/testdata/fibonacci -importcfg $WORK\b001\importcfg -pack -c=4]
instrumenter: 2021/02/21 12:10:14 files:  [D:\dev\go\src\github.com\skirmish\instrument\testdata\fibonacci\main.go]
instrumenter: 2021/02/21 12:10:14 incoming args [C:\Go\pkg\tool\windows_amd64\compile.exe -o $WORK\b001\_pkg_.a -trimpath $WORK\b001=> -p main -complete -buildid 07gIkXLEJooNJ0O9y49j/07gIkXLEJooNJ0O9y49j -goversion go1.14.2 -D _/D_/dev/go/src/github.com/skirmish/instrument/testdata/fibonacci -importcfg $WORK\b001\importcfg -pack -c=4 testdata\fibonacci\main.go]
instrumenter: 2021/02/21 12:10:14 executing command `C:\Go\pkg\tool\windows_amd64\compile.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b001\\_pkg_.a" "-trimpath" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b001=>" "-p" "main" "-complete" "-buildid" "07gIkXLEJooNJ0O9y49j/07gIkXLEJooNJ0O9y49j" "-goversion" "go1.14.2" "-D" "_/D_/dev/go/src/github.com/skirmish/instrument/testdata/fibonacci" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b001\\importcfg" "-pack" "-c=4" "D:\\dev\\go\\src\\github.com\\skirmish\\instrument\\testdata\\fibonacci\\main.go"`
```

```bash
instrumenter: 2021/02/21 12:10:14 Program:  C:\Go\pkg\tool\windows_amd64\link.exe
instrumenter: 2021/02/21 12:10:14 Package:
instrumenter: 2021/02/21 12:10:14 stdlib:  false
instrumenter: 2021/02/21 12:10:14 output:  0xc000116010
instrumenter: 2021/02/21 12:10:14 args:  [-o $WORK\b001\exe\a.out.exe -importcfg $WORK\b001\importcfg.link -buildmode=exe -buildid=z0fEn4ig47ItyZemdzLr/07gIkXLEJooNJ0O9y49j/o1dCexGIUjdXqSZatCzk/z0fEn4ig47ItyZemdzLr -extld=gcc]
instrumenter: 2021/02/21 12:10:14 files:  [C:\Users\adam\AppData\Local\Temp\go-build902137514\b001\_pkg_.a]
instrumenter: 2021/02/21 12:10:14 incoming args [C:\Go\pkg\tool\windows_amd64\link.exe -o $WORK\b001\exe\a.out.exe -importcfg $WORK\b001\importcfg.link -buildmode=exe -buildid=z0fEn4ig47ItyZemdzLr/07gIkXLEJooNJ0O9y49j/o1dCexGIUjdXqSZatCzk/z0fEn4ig47ItyZemdzLr -extld=gcc $WORK\b001\_pkg_.a]
instrumenter: 2021/02/21 12:10:14 executing command `C:\Go\pkg\tool\windows_amd64\link.exe "-o" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b001\\exe\\a.out.exe" "-importcfg" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b001\\importcfg.link" "-buildmode=exe" "-buildid=z0fEn4ig47ItyZemdzLr/07gIkXLEJooNJ0O9y49j/o1dCexGIUjdXqSZatCzk/z0fEn4ig47ItyZemdzLr" "-extld=gcc" "C:\\Users\\adam\\AppData\\Local\\Temp\\go-build902137514\\b001\\_pkg_.a"`
```