## Overview

**delfy** is a simple divination tool.

Its operation is based on completely unscientific principles unless you are very open minded and understand quantum mechanics really well. https://www.livescience.com/26444-quantum-mechanics-physicists-poll.html

Neither **delfy** nor its documentation are meant to be viewed, interpreted, or used in a serious way. If you do, it will probably stop working immediately.

## Back Story

My first experience with computerized divination was around 1981 when a friend showed me the Unix 7th ed. ching(1) program. I was a student at a highly-regarded, science-oriented college, and I had been brought up in a conservative environment with almost no exposure to metaphysical or shamanic concepts. I found the program to actually be helpful, which I found intriguing.

ching(6) manual page: http://www.cfcl.com/ching/man/

ching(6) source code at GitHub: https://github.com/floren/ching

Later in life, I learned more about Taoism and other philosophies, religions, teachings, and cultures that use (or have used) various divination practices. There were and are many! Divination has been used for probably many thousands of years as an aid to making decisions and clarifying various aspects of one's life. Have you ever flipped a coin to decide something?

## Introduction

**delfy** works in a manner simpler than many traditional systems for divination, while being more complicated than flipping a coin.

### Using **delfy**

The user is prompted with **If**, with the expectation that the rest of an _if clause_ will be supplied. **delfy** responds with a _then clause_ to complete the sentence. Here's an example:

```
$ delfy
--
If I get up early tomorrow morning
then look at each other again and smile.
--
If 
```

**delfy** will continue to prompt for more input in case further clarification is desired. The user may exit the program simply by typing the Enter (Return) key or a Control-D immediately after the prompt. A Control-C quits the program at any time.

Interpretation of **delfy**'s output is up to the user. Here are some suggestions that may help:

Everyone has at least some intuition. Although you may not get the best results, you do not need to be a trained or professional "psychic" or "shaman".

Relax and **stay open to your feelings**. Allow the "meaning" to come to you in the first instant you see the response. It's more about how you personally feel and react, rather than the formal meaning of the sentences.

**Have a question clearly in mind.**

The combination of a strongly- and sincerely-held question or intensely-felt personal issue, along with the openness to receive, are perhaps the two most significant determinators of the quality of any intuitive reading. Conversely, if you ask frivolous questions, don't blame the oracle if you get a lame or nonsensical result. Or in **delfy**'s own words,

    If I ask you stupid questions
    Then remember that thing that exists that you forgot about?

Remember:

* Ask sincerely.
* Receive openly.

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

## Credits

**delfy** is a cheap workalike of a program called **laroec** written around 1997 by Ben Osheroff. Ben and his girlfriend encountered a similar program in operation in an exhibit at the Exploratorium in San Francisco, and wrote **laroec** to mimic it.

**laroec** has a curses-based full-screen interface and a "Comes from" mode. I hardly ever used either, so I did not implement them in **delfy**.

The `delfy.responses` data file included with **delfy** is a slightly-edited copy of the one included with **laroec**. Thanks go to Ben for giving me permission to include it. Creating my own would have taken more time than I'm willing to spend on this right now.

## Manual Page

A copy of the **delfy** manual page is included here for convenience. To display it better, install it on your system and use the `man` command to view it.

```
DELFY(6)                           UNIX Programmer's Manual                           DELFY(6)

NAME
       delfy - oracle

SYNOPSIS
       delfy

DESCRIPTION
       delfy is a divination tool that may be used to amplify or clarify intuition.

       delfy prints an If prompt and waits for the user to fill in an if clause of a sentence,
       followed by a newline. delfy then fills in the then clause to complete a  sentence.  As
       an  example,  if  the  user  were to type "I think divination is nonsense", delfy might
       respond as follows:

           If I think divination is nonsense
           then roses are red, violets are blue.

       To end the program, type an Enter (Return) key or a Control-D at the If prompt.

FILES
       delphy searches for its data file delfy.responses in the following directories, in  the
       order listed:

           /usr/share/games/delfy
           /usr/share/delfy
           /usr/local/share/delfy
           the current working directory
           $HOME/.delfy

SEE ALSO
       If you want to learn more about divination
       then gather an immense collection of inexplicable bits of news.

       fortune(1), ching(1) from Seventh Edition UNIX

DIAGNOSTICS
       If you want error messages
       then it would be bad.  Real bad.  Quite ironic, too.

BUGS
       If you don´t like delfy
       then the camp will break up.

AUTHOR
       Jay Ts (http://jayts.com)

Jay Ts                                   February 2019                                DELFY(6)
```
