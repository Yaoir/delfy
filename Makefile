# Makefile for delfy

# set BINDIR to the directory in your PATH where you want cont installed
BINDIR=/home/jay/.bin/elf

# set MANDIR to where you want the manual page installed
MANDIR=/usr/share/man/man1

# Release date. For ronn, when making manual page
RELDATE=2019-02-10

delfy: delfy.go
	@go build -o delfy delfy.go

# Manual Page

man: delfy.1.ronn
	@ronn --roff --manual="UNIX Programmer's Manual" --organization="Jay Ts" --date="$(RELDATE)" delfy.1.ronn >/dev/null 2>&1
	@gzip -f delfy.1
	@mv delfy.1.gz man1
	@man -l man1/delfy.1.gz

showman:
	@man -l man1/delfy.1.gz

install:
	@cp delfy $(BINDIR)

installman:
	@cp man1/delfy.1.gz $(MANDIR)

clean:
	@rm -f delfy

backup back bak:
	@cp -a delfy.responses* delfy.1.ronn *.go man1 Makefile push README.md TODO .bak

wc count:
	@wc -l delfy.go
