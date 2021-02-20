# Instrumenting go code

**This is an MVP, doesn't work, DO NOT USE, yet**

This aims to be a single step tool to instrument a program with opentelemetry during the build phase rather than writing
all the specific instrumentation code, this will insert the trace spans into the code as it is being compiled.

## Installation

```bash
$ go install github.com/skirmish/instrument/instrument
```

## Compiling an example using it.

Hello World example

```bash
$ go build -o hello.exe -a -toolexec instrument ./testdata/hello/.
```

Another build example.

```bash
$ go build -o shello.exe -a -toolexec instrument ./testdata/smallhello/.
```

```bash
$ go build -o fib.exe -a -toolexec instrument ./testdata/fibonacci/.
```