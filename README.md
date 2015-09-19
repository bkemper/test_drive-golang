# Test Drive Go

The purpose of this test drive is to understand the setup and capabilities of [Go](http://golang.org) the language.

## Setup

Follow these [instructions](https://github.com/moovweb/gvm#installing) and install [Go Version Manager](https://github.com/moovweb/gvm).

Prepend to your `~/.bash_profile`:

    [[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"

Reload your profile:

    $ source ~/.bash_profile

Install Go:

    $ gvm install go1.4  --default
    
Set as default:

    $ gvm use go1.4 --default

Most importantly, install [Golint](https://github.com/golang/lint#installation).

Append to your `~/.vimrc` to lint after a save:

```bash
" https://github.com/golang/lint#vim
set rtp+=$GOPATH/src/github.com/golang/lint/misc/vim
autocmd BufWritePost,FileWritePost *.go execute 'Lint' | cwindow
```