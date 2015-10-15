# Test Drive Go

The purpose of this test drive is to understand the setup and capabilities of the [Go programming language](http://golang.org).  Before wasting your time, take a look at the resources used to fuel this test drive.

* [The Go Programming Language Documentation](http://golang.org/doc)
* [Go by Example](https://gobyexample.com)

## Setup

### Installation

To easily install, manage, and work with multiple Go versions, install [Go Version Manager](https://github.com/moovweb/gvm) with these [instructions](https://github.com/moovweb/gvm#installing).

Prepend to your `~/.bash_profile`:

    [[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"

Reload your profile:

    $ source ~/.bash_profile

Install latest Go version:

    $ gvm install go1.4

Set as default:

    $ gvm use go1.4 --default

### Linting

To lint your code, install [Golint](https://github.com/golang/lint#installation).

    $ go get -u github.com/golang/lint/golint

Append to your `~/.vimrc` to print style mistakes after a save:

```
" https://github.com/golang/lint#vim
set rtp+=$GOPATH/src/github.com/golang/lint/misc/vim
autocmd BufWritePost,FileWritePost *.go execute 'Lint' | cwindow
```

### Documentation

To view your program documentation, install [GoDoc](http://godoc.org/golang.org/x/tools/cmd/godoc).

    $ go get -u golang.org/x/tools/cmd/godoc

Refer to the [Godoc: documenting Go code](http://blog.golang.org/godoc-documenting-go-code) blog post for some background and comment convention.

### Other Tools

For other awesome tools, check out [Go Tools](https://github.com/golang/tools).

## Capabilities

### Run from source

Write a [program](src/testdrive/main.go) and use [run](https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program):

    $ go run src/testdrive/main.go

### Run from binary

Compile your source code using [build](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies):

    $ go build -o bin/testdrive src/testdrive/main.go

Run:

    $ ./bin/testdrive

### Example Packages

* [Basics](src/testdrive/basics/basics.go) - Importing packages, variables, functions, etc.