
.PHONEY: all build run usage

TARGET=bolt

all: usage

edit e:
	vi $(TARGET)_sample.go

build b:
	go build -o sample $(TARGET)_sample.go

run r:
	./sample

bolt:
	bolt check sample.db
	bolt stats sample.db

clean:
	rm -f ./sample sample.db

make m:
	vi Makefile

usage:
	@echo ""
	@echo "make [edit|build|run|clean]"
	@echo ""
