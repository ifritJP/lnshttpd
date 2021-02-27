LNSC = lnsc

help:
	@echo make all

bind.go: bind.lns
	$(LNSC) bind.lns save -langGo --package lnshttpd

all:
	$(LNSC) lnsservlet.lns save -langGo --package lnshttpd
