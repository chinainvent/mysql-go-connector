include $(GOROOT)/src/Make.$(GOARCH)

TARG=db
CGOFILES=mysql.go
GOFILES=api.go

CGO_CFLAGS=-I/usr/include/mysql
CGO_LDFLAGS=-rdynamic -L/usr/lib/mysql -lmysqlclient_r -lz -lpthread -lcrypt -lnsl -lm -lpthread -lssl -lcrypto

CLEANFILES+=example

include $(GOROOT)/src/Make.pkg

all: example

%: install %.go
	$(QUOTED_GOBIN)/$(GC) $*.go
	$(QUOTED_GOBIN)/$(LD) -o $@ $*.$O

