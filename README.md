Title:		Dcpu16 Readme
Author:		Radek Hnilica

dcpu16
======


> Emulator, documentation, and other stuff for 0x10c computer DCPU-16

This project is my experimental emulator and tools for DCPU-16.
I use this while slowly learnig Go language. It has no other value.


Documentation
-------------

Documentation is in directory doc and is written using multimarkdown.
For transforming into html I use [fletcher/peg-multimarkdown](fletcher/peg-multimarkdown).
How it is used see in the `doc/Makefile`.

Go source
---------

Go source files are in `go` directory.  It has structure of workspace.
Script `setenv` is setting the environment so it is recognized as workspace.
Use it this way

    $ cd go
    $ . setenv


Structure of source
-------------------
