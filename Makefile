# Makefile for go program
.PHONY:	all proto build run test client server kill make usage

PROTOC=protoc
PROGRAM=grpc-example
SERVICE=customer_service

usage:
	@echo "Makefile for $(PROGRAM), by Stoney Kang, sikang99@gmail.com"
	@echo "make [proto|build|run|kill|test]"
	@echo "   - proto : compile interface spec"
	@echo "   - build : compile client/server"
	@echo "   - run   : start the server and exec client"
	@echo "   - kill  : stop the server"

ep:
	vi proto/$(SERVICE).proto
epg:
	vi proto/$(SERVICE).pb.go
ec:
	vi client/client.go
es:
	vi server/server.go
ed:
	vi proto/$(SERVICE)_db.go
et:
	vi test/load_test.go

proto p:
	$(PROTOC) -I./proto ./proto/$(SERVICE).proto --go_out=plugins=grpc:proto
	@ls -al ./proto

build b:
	go build -o client/client client/client.go
	go build -o server/server server/*.go

rebuild:
	make proto
	make build

run r:
	@echo "> make (run) [all|client|server]"

run-all ra:
	./server/server &
	sleep 1
	./client/client list

run-client rc:
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

run-server rs:
	./server/server &

kill k:
	killall server

test t:
	@make server
	cd test && go test -v	
	@make kill

db-check dc:
	bolt check person.db

clean:
	cd bolt_test && make clean
	rm -f client/client server/server person.db

git-push gpush gu:
	make clean
	git add * .gitignore 
	git commit -m "modify README.md"
	git push

git-store gs:
	git config credential.helper store


