# Instrumenting go code

**This is an MVP, doesn't work, DO NOT USE, yet**

This aims to be a single step tool to instrument a program with opentelemetry during the build phase rather than writing
all the specific instrumentation code, this will insert the trace spans into the code as it is being compiled.

For other languages that are more dynamic, automatic instrumentation by adding a library makes integration of
instrumentation exceptionally easy. In go, as a statically compiled language into a non-intermediate vm, it makes it a
little more difficult to insert new code into functions as there is no vtable available to the running code to make
changes to itself.

## Installation

```bash
$ go install github.com/skirmish/instrument/instrument-otel
```

## Compiling an example using it.

Hello World example (using windows, for non windows binaries, drop the .exe)

```bash
$ go build -o hello.exe -a -toolexec instrument-otel ./testdata/hello/.
```

Another build example.

```bash
$ go build -o shello.exe -a -toolexec instrument-otel ./testdata/smallhello/.
```

And another.

```bash
$ go build -o fib.exe -a -toolexec instrument-otel ./testdata/fibonacci/.
```