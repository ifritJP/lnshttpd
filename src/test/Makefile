LNSC=lnsc

ifneq ($(wildcard make.conf.mk),)
include make.conf.mk
endif

SRCS=
SRCS += $(shell find . -iname '*.lns')


help:
	@echo make build [LNS_OPT=opt]
	@echo make buildDB
	@echo make check
	@echo make update_lns

build: test

test: $(SRCS)
	@echo $(SRCS) | sed 's/ /\n/g' |	\
		$(LNSC) @- save -langGo --main test $(LNS_OPT)
	$(LNSC) test.lns mkmain entry.go
	go build -a -tags gopherlua -o test

