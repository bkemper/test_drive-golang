我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

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
