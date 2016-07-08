BINDIR=bin
BINARY=bm

VERSION?=dev
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.buildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)
$(BINARY):
	go build -v ${LDFLAGS} -o ${BINDIR}/${BINARY} .

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
	rm -rf ${BINDIR}