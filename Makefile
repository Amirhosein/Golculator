export APP=golculator

export LDFLAGS="-w -s"

all: build install

build:
	cd ./cmd/golculator/ && go build -ldflags $(LDFLAGS)

install:
	cd ./cmd/golculator/ && go install -ldflags $(LDFLAGS)