LNSC = lnsc

ifneq ($(wildcard make.conf.mk),)
include make.conf.mk
endif

help:
	@echo make all

# lns/httpd/bind.go: lns/httpd/bind.lns
# 	$(LNSC) lns/httpd/bind.lns save -langGo --package httpd
lns/httpd/lnsservlet.go: lns/httpd/lnsservlet.lns
	$(LNSC) lns/httpd/lnsservlet.lns save -langGo --package httpd


all: lns/httpd/lnsservlet.go
