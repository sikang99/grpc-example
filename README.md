# Static http/https server

Simple gRPC service example

## Install

    $ go get github.com/sikang99/grpc-example

## Usage

refer Makefile if you know how to use.
	
	$ make 
	Makefile for grpc-example, by Stoney Kang, sikang99@gmail.com

	make [proto|build|run|kill|test]
   	- proto : compile interface spec
   	- build : compile client/server
   	- run   : start the server and exec client
   	- kill  : stop the server


compile IDL proto of gRPC.
	
	$ make proto

build client and server programs

	$ make build

run and test the service
	
	$ make run
	$ make test

## History

- 2015/05/01 : list support search with condition optionally
- 2015/04/30 : start to code with mattn/grpc-example


## References

- [gRPC-JSON Proxy](http://yugui.jp/articles/889)
- [Protocol Buffers を利用した RPC、gRPC を golang から試してみ](http://mattn.kaoriya.net/software/lang/go/20150227144125.htm) 
- [mattn/grpc-example](https://github.com/mattn/grpc-example)


## License

MIT

