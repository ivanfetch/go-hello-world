# Hello World - A Git Repo vs. Go Package Naming Test

This repo is to experiment with naming a Github repo and Go package differently.

* The Github repository `go-hello-world` is more descriptive (let's pretend that's the case for this example), than `hello` which is what I would like to use for my Go package name.
* Binary installation using Go is also descriptive - that you will end up with a `hello-world` binary: `go install github.com/ivanfetch/go-hello-world/cmd/hello-world`
* Although I have to import `"github.com/ivanfetch/go-hello-world"`, I can still refer to my desired Go package name of `hello`, for example when calling `hello.Print()`.
* If others import this code for use in their own project, is the different `import` (go-hello-world) and package prefix (hello) a bad experience?
* I could name my Github repo `ivanfetch/hello` which would match the Go package name, but it's nice for the repo name to be more descriptive than a GO package should be.