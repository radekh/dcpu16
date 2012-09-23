#!Makefile
# $Id$
# Copyright (c) 2012 Radek Hnilica

SOURCE_DIR := .
DESTDIR := ../html

HTMLS = $(subst .md,.html,$(shell ls *.md))

all: $(HTMLS)

#	rsync $(HTMLS) *.css $(DESTDIR)

# Generic rules
%.html: %.md
	multimarkdown $< >$@
