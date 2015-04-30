.PHONY:	all proto build run client server kill make usage

PROGRAM=grpc-example
PROTOC=/usr/local/bin/protoc

all: usage

ep:
	vi proto/customer_service.proto

ec:
	vi client/main.go

es:
	vi server/main.go

proto:
	$(PROTOC) -I./proto ./proto/customer_service.proto --go_out=plugins=grpc:proto
	@ls -al ./proto


build b:
	go build -o client/client client/main.go
	go build -o server/server server/main.go

run r:
	./server/server &
	sleep 1
	./client/client list

client:
	./client/client add stoney 52
	./client/client add mandoo 19
	./client/client list

server:
	./server/server &

kill k:
	killall server

git-push gpush gu:
	git init
	git add *
	git commit -m "write readme"
	git push -u https://sikang99@github.com/sikang99/$(PROGRAM) master
	#chromium-browser https://github.com/sikang99/$(PROGRAM)

readme md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "make [proto|build|run|kill]"
	@echo ""
