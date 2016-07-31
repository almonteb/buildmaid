BINDIR=bin
RELEASEDIR=release
NAME=buildmaid
BINARY=bm

VERSION?=dev
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.buildTime=${BUILD_TIME}"
BUILDCMD=go build -v ${LDFLAGS}

.DEFAULT_GOAL: $(BINARY)
$(BINARY):
	${BUILDCMD} -o ${BINDIR}/${BINARY} .

test:
	go test ./...

run:
	go run *.go

vet:
	go vet ./...

lint:
	golint ./...

fmt:
	go fmt ./...

doc:
	godoc -http=:6060 -index

clean:
	rm -rf ${BINDIR} ${RELEASEDIR}

define generate_release
	$(eval DIR := ${RELEASEDIR}/buildmaid-$(1))
	GOOS=$(1) GOARCH=amd64 ${BUILDCMD} -o ${DIR}/$(2) .
	cp example.yml ${DIR}
	tar -C ${DIR} -cvzf ${DIR}-${VERSION}.tar.gz .
endef

release: release-windows release-linux release-darwin

release-windows:
	$(call generate_release,windows,${BINARY}.exe)

release-linux:
	$(call generate_release,linux,${BINARY})

release-darwin:
	$(call generate_release,darwin,${BINARY})
