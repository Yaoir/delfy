## Overview

**delfy** is a simple divination tool.

It's operation is based on completely unscientific principles unless you are open minded and understand quantum mechanics really well.

The contents of this README are not meant to be taken in a serious way.

## Back Story

My first experience with computerized divination was around 1981 when a friend showed me the Unix 7th ed. ching(1) program. I was a student at a highly-regarded, science-oriented college, and I had been brought up in a conservative environment with almost no exposure to metaphysical or shamanic concepts. I found the program interesting to use, and actually helpful.

ching(6) manual page: http://www.cfcl.com/ching/man/
ching(6) source code at GitHub: https://github.com/floren/ching

Later in life, I learned more about Taoism and other philosophies, religions, teachings, and cultures that use (or have used) various divination practices. There were and are many! Divination has been used for probably thousands of years as an aid to making decisions and clarifying various aspects of one's life. Have you ever flipped a coin to decide something?

## Introduction

**delfy** works in a manner simpler than many traditional systems for divination, while being more complicated than flipping a coin.

### Using **delfy**

The user is prompted with `If`, with the expectation that the rest of an _if clause_ will be supplied. It responds with a _then clause_ to complete the sentence. Here's an example:

```
$ delfy
--
If I get up early tomorrow morning
Then look at each other again and smile.
--
If 
```

**delfy** will continue to prompt for more input. The user may exit the program simply by typing the Enter/Return key or Control-D immediately after the prompt. Control-C quits the program at any time.

Interpretation of **delfy**'s output is up to the user. Here are some suggestions to help:

Good readers use their own intuitive abilities to interpret cards, tea leaves, or whatever
You can do readings on your own using your own intuition.
Relax, stay open to your feelings. Allow the "meaning" to come to you in the first instant you see the response.
It's more about how you personally feel and react, rathan than the formal meaning of the sentences.

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
