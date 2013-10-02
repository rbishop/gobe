gobe
===

gobe (go build + execute) is a simple command line script for Linux users that are sick of this error:
`fork/exec /tmp/go-build952472105/command-line-arguments/_obj/exe/gobe: permission denied`

gobe just simply does a `go build` followed by executing the program and forwarding any arguments onto the program. After the program has run gobe will remove the compiled file for you.
