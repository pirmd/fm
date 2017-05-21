INTRODUCTION
============

`fm` is a command line file browser heavily inspired from [ranger](http://ranger.nongnu.org/).
The main differences that `fm` tries to achieve are:
* stripped down features (e.g. no tabs, no shell execution...)
* implementation of a Virtual File System backend, for example to browse zip or sftp shares (not implemented yet)

If you are looking to a full features command line file manager you most probably want to try [lf](https://github.com/gokcehan/lf)

INSTALLATION
============
Everything should work fine using go standard commands (`build`, `get`, `install`...).
For simplicity, you can just run `make install` if you prefer

I'm trying to keep dependancies number low:
* go (>= 1.8 for sorting suggar that is `sort.Slice`)
- termbox-go
- afero
* go-open
* go-humanize

USAGE
=====
Usage is quite straightforward: `fm [path|url]`
* `path` is a normal path indication being absolute or relative
* `url` is an url-like string (`scheme://[userinfo@]host/path`) to indicate browsing to specific location (like sftp shares or zip file).
It can easily be extended to anything that `afero` supports.

Ìf no `path` or `url` is given, `fm` will start browsing the current folder.


CUSTOMIZATION
=============
I don't feel the need to have a config file system or complicated flag system.
No special customization is proposed up to now, you'll have to deep in the code for that, starting with the `config.go`, which is hopefully commented enough to have you play with customizations.
Once done, run `make install` and you're done.

If you feel like a bit more adventurious, you can dive into the code and:
- `fm.go`: glue between the ui and the underlying file systems. It is responsible for almost all behaviour of the app
- `folder.go`: propose some basic type to facilitate `fm.go`'s life
* `ui.go`: almost anything linked to layout and colors lies here
* `vfs/`: really small interface to `afero`. Main use is to have a factory-like features to adopt the correct `afero.Fs` depending on the url you like to browse
* `tui/`: simple layer above `termbox-go` with basic widgets to conveniently displayed text and lists

CHANGELOG
=========
* _Version 1_ (2017.05.15):
    - basis for simple file browsing
    - connector to `afero` for future extension to browse zip archive or sftp shares

* _Version 2_ (Roadmap):
    - implement zip archive and sftp browsing
    - extend browsing to simple file management (delete, copy/paste, rename)


CONTRIBUTION
============
If you feel like to contribute, just follow github guidelines on [forking](https://help.github.com/articles/fork-a-repo/) then [send a pull request](https://help.github.com/articles/creating-a-pull-request/)

[modeline]: # ( vim: set fenc=utf-8 spell spl=en: )
