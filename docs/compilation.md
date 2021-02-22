# Compilation

How does this thing do its job.

You `go install` the thing. Then on compilation to auto-instrument you use the `-toolexec` option to specify the
instrumenter. It then rewrites the go code and passes the re-written code to the compiler, which then compiles it and
gives you an executable with instrumentation of all the things.

Ideally, you shouldn't have to do anything to the code to help the instrumenter, but there will have to be some method
of configuring what not to instrument and helper comments to collect metrics and logs as well as advice the instrumenter
on the right thing to do.

The biggest change to the instrumented code is `context` will be added as the first parameter to any function
invocation, if one doesn't exist. It should be able to co-exist with any `context` that already exists.

## How

In the post;
<https://xdg.me/rewriting-go-with-ast-transformation/>
The author gives an example.

https://play.golang.org/p/-mS9Vglbdi



