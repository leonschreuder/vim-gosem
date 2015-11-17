vim-gosem
================================================================================

This is a semantic highlighter for go sourcecode files in vim. It uses
[gosem][] to do the analysing. [gosem][] uses go's built in go-parser, so it is
blasing fast, even on large files.

![Screenshot with solarized highlighting](img/screen.png)

Screenshot with [solarized](http://ethanschoonover.com/solarized) color scheme.

Installation
================================================================================

You need go installed to make all of this work, but since this is a tool for go
development this might not be a problem for you.

gosem
--------------------------------------------------------------------------------

To get the syntaxhighlighting, you first need [gosem][]. Just `go get` it:

    go get -u https://github.com/meonlol/gosem

This should just work. Ask aunt google to help you with your environment if it
won't. You can test gosem by navigating to some go source file on the
commandline and running `gosem -f yourfile.go`. If it ouputs something like the
below, your good to go.

    someFieldToPrint someOtherField|8,16,stringToPrint a b|18,21,alocalhost


vim-gosem
--------------------------------------------------------------------------------

You know the drill! Copy-pase, away!

*  [Pathogen][]
    * `git clone https://github.com/meonlol/vim-gosem.git ~/.vim/bundle/vim-gosem`
*  [vim-plug][]
    * `Plug 'meonlol/vim-gosem'`
*  [NeoBundle][]
    * `NeoBundle 'meonlol/vim-gosem'`
*  [Vundle][]
    * `Plugin 'meonlol/vim-gosem'`

Then put this in your vimrc...

    " make vim-gosem your highlighter once you open a go source file, and refresh on save
    autocmd BufRead,BufNewFile *.go       setlocal syntax=vim-gosem
    autocmd BufWritePost *.go     silent! setlocal syntax=vim-gosem

... `source %` it, and open a go source file.


[gosem]: https://github.com/meonlol/gosem
[Pathogen]: https://github.com/tpope/vim-pathogen
[vim-plug]: https://github.com/junegunn/vim-plug
[NeoBundle]: https://github.com/Shougo/neobundle.vim
[Vundle]: https://github.com/gmarik/vundle
