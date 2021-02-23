

install:
	go install -v github.com/skirmish/instrument/instrument-otel

test: install testhello testfib testhelloweb


testhello:
	go build -o hello -a -toolexec instrument-otel ./testdata/hello/.

testfib:
	go build -o fib -a -toolexec instrument-otel ./testdata/fibonacci/.

testhelloweb:
	go build -o helloweb -a -toolexec instrument-otel ./testdata/simplehelloweb/.

clean:
	rm -f hello
	rm -f fib
	rm -f helloweb

