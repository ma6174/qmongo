build:
	go build -v

test: 
	./qmongo ./examples/crud.ql

all: build
