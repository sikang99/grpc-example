#
# Makefile for bolt test
#
.PHONEY: all build run usage

TARGET=bolt

usage:
	@echo "usage: make [edit|build|run|clean]"

edit e:
	vi $(TARGET)_sample.go

build b:
	go build -o sample $(TARGET)_sample.go

run r:
	./sample

bolt-install:
	go get github.com/boltdb/bolt

bolt-check:
	bolt check sample.db
	bolt stats sample.db

clean:
	rm -f ./sample sample.db

