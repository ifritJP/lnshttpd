LNSC = lnsc

help:
	@echo make all

lns/httpd/bind.go: lns/httpd/bind.lns
	$(LNSC) lns/httpd/bind.lns save -langGo --package httpd

all:
	$(LNSC) lns/httpd/lnsservlet.lns save -langGo --package httpd
