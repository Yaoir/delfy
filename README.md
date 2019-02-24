## Overview

**delfy** is a simple divination program.

## Introduction

## Compiling and Installing

**delfy** is written in Go. To compile it, you need to have Go installed.

To compile:

```
$ go build delfy.go
```
or if you have GNU **make** installed:
```
$ make
```

To install the manual page, copy the file **man1/delfy.1.gz** to the directory where your manual pages are located. On Linux, this is typically **/usr/share/man/man1**.

To install **delfy** program using **make**, edit **Makefile** to set BINDIR appropriately, then run

```
$ make install
```

To install the manual page using **make**, edit **Makefile** to set MANDIR appropriately, then run

```
$ sudo make installman
```

## Manual Page

A copy of the manual page is included here for convenience. To display it better, install it on your system and use the `man` command to view it.
