include $(GOROOT)/src/Make.$(GOARCH)

TARG=db
CGOFILES=mysql.go
GOFILES=api.go

CGO_CFLAGS=$(shell mysql_config --include)
CGO_LDFLAGS=$(shell mysql_config --libs)

CLEANFILES+=example

include $(GOROOT)/src/Make.pkg

all: example

%: install %.go
	$(QUOTED_GOBIN)/$(GC) $*.go
	$(QUOTED_GOBIN)/$(LD) -o $@ $*.$O

