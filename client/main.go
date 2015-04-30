package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/mattn/sc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//pb "github.com/sikang99/grpc-example/proto"
	pb "../proto"
)

func add(name string, age int) error {
	conn, err := grpc.Dial("127.0.0.1:11111")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Name: name,
		Age:  int32(age),
	}
	_, err = client.AddPerson(context.Background(), person)
	return err
}

func delete(name string) error {
	conn, err := grpc.Dial("127.0.0.1:11111")
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Name: name,
	}
	_, err = client.DeletePerson(context.Background(), person)
	return err
}

func list() error {
	conn, err := grpc.Dial("127.0.0.1:11111")
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewCustomerServiceClient(conn)

	req := &pb.RequestType{
		Op: "match",
	}
	stream, err := client.ListPerson(context.Background(), req)
	//stream, err := client.ListPerson(context.Background(), new(pb.RequestType))
	if err != nil {
		return err
	}
	for {
		person, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(person)
	}
	return nil
}

func main() {
	(&sc.Cmds{
		{
			Name: "list",
			Desc: "list: listing person",
			Run: func(c *sc.C, args []string) error {
				return list()
			},
		},
		{
			Name: "add",
			Desc: "add [name] [age]: add person",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 2 {
					return sc.UsageError
				}
				name := args[0]
				age, err := strconv.Atoi(args[1])
				if err != nil {
					return err
				}
				return add(name, age)
			},
		},
		{
			Name: "delete",
			Desc: "delete [name]: delete person",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 1 {
					return sc.UsageError
				}
				name := args[0]
				return delete(name)
			},
		},
	}).Run(&sc.C{
		Desc: "grpc example",
	})
}
