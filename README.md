# Static http/https server

Simple gRPC example

## Install

    $ go get github.com/sikang99/grpc-example

## Usage

refer Makefile to know how to use
	
	$ make 
	make [proto|build|run|kill]

compile IDL proto of gRPC
	
	$ make proto

build client and server programs

	$ make build

run and test the service
	
	$ make run

## Options


## References

- [gRPC-JSON Proxy](http://yugui.jp/articles/889)
- [Protocol Buffers を利用した RPC、gRPC を golang から試してみ](http://mattn.kaoriya.net/software/lang/go/20150227144125.htm) 
- [mattn/grpc-example](https://github.com/mattn/grpc-example)


## License

MIT

