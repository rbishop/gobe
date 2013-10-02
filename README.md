gobe
====

gobe (`go build` + exec) is a command line script for Linux users that are sick of this error when trying to use `go run`:

`fork/exec /tmp/go-build952472105/command-line-arguments/_obj/exe/gobe: permission denied`

The work around for the above error on Linux is to re-mount your `/tmp` directory and give it exec privileges. But c'mon, we're programmers. We're lazier than that!

gobe just simply does a `go build` followed by executing the compiled program and forwarding any arguments onto the program. After the program has run gobe will remove the compiled file for you.

Install
-------

Use go's great built in package getter, `go get`. The following will get the source and install gobe to your `$GOPATH` directory:

`go get github.com/rjbishop/gobe`

Use
---

Just give gobe the filename and any arguments the program will take. Example:

`gobe hello.go World`

`=> Hello World`
