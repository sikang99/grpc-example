# Makefile for go program
.PHONY:	all proto build run test client server kill make usage

PROTOC=/usr/local/bin/protoc
PROGRAM=grpc-example
SERVICE=customer_service

all: usage

ep:
	vi proto/$(SERVICE).proto

epg:
	vi proto/$(SERVICE).pb.go

ec:
	vi client_test/client_test.go

es:
	vi server/server.go

et:
	vi test/load_test.go

proto p:
	$(PROTOC) -I./proto ./proto/$(SERVICE).proto --go_out=plugins=grpc:proto
	@ls -al ./proto

build b:
	go build -o client/client client/client.go
	go build -o server/server server/server.go

rebuild:
	make proto
	make build

run r:
	./server/server &
	sleep 1
	./client/client list

client rc:
	./client/client add stoney 52
	./client/client add mandoo 19
	./client/client add namoo 25
	./client/client list
	./client/client delete 2 
	./client/client get 1 
	./client/client get 2 
	./client/client update 1 younga 48 
	./client/client list
	./client/client list 19

server rs:
	./server/server &

kill k:
	killall server

test t:
	@make server
	cd test && go test -v	
	@make kill

clean:
	rm -f client/client server/server person.db

git-push gpush gu:
	make clean
	git init
	git add *
	git commit -m "more load tests"
	git push -u https://sikang99@github.com/sikang99/$(PROGRAM) master
	#chromium-browser https://github.com/sikang99/$(PROGRAM)

git-status gs:
	git status
	git log --oneline -5


readme md:
	vi README.md

make m:
	vi Makefile

usage:
	@echo ""
	@echo "Makefile for $(PROGRAM), by Stoney Kang, sikang99@gmail.com"
	@echo ""
	@echo "make [proto|build|run|kill|test]"
	@echo "   - proto : compile interface spec"
	@echo "   - build : compile client/server"
	@echo "   - run   : start the server and exec client"
	@echo "   - kill  : stop the server"
	@echo ""
