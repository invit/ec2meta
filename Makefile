.PHONY: build install snapshot dist test vet lint fmt run clean
OUT := ec2meta
PKG := github.com/invit/ec2meta
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -o ${OUT} ${PKG}

install:
	CGO_ENABLED=0 GOOS=linux go install -a -v ${PKG}

snapshot:
	goreleaser --snapshot --rm-dist

dist:
	goreleaser --rm-dist

test:
	@go test -v ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

fmt:
	@gofmt -l -w -s ${GO_FILES}

clean:
	-@rm -vf ${OUT}
	-@rm -vrf dist



