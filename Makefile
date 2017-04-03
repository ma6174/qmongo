build:
	go build -v

test: 
	./qmongo ./examples/curd.ql

all: build
