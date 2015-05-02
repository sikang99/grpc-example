package main

import (
	"fmt"
	"io"
	"strconv"

	pb "../proto"
	"github.com/mattn/sc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var address string = "127.0.0.1:11111"

func list(age int) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	req := &pb.RequestType{}
	if age > 0 {
		ps := &pb.Person{
			Age: int32(age),
		}
		req.Person = ps
	}

	stream, err := client.ListPersons(context.Background(), req)
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
		fmt.Printf("List: %v\n", person)
	}
	return nil
}

func purge(age int) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	req := &pb.RequestType{}
	if age > 0 {
		ps := &pb.Person{
			Age: int32(age),
		}
		req.Person = ps
	}

	stream, err := client.PurgePersons(context.Background(), req)
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
		fmt.Printf("Purge: %v\n", person)
	}

	return nil
}

func add(name string, age int) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Name: name,
		Age:  int32(age),
	}

	res, err := client.AddPerson(context.Background(), person)
	fmt.Printf("Add: %v\n", res.Person)
	return err
}

func get(id int) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Id: int32(id),
	}

	res, err := client.GetPerson(context.Background(), person)
	fmt.Printf("Get (%d): %v\n", id, res.Person)
	return err
}

func delete(id int) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Id: int32(id),
	}

	res, err := client.DeletePerson(context.Background(), person)
	fmt.Printf("Delete (%d): %v\n", id, res.Person)
	return err
}

func update(id int, name string, age int) error {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Id:   int32(id),
		Name: name,
		Age:  int32(age),
	}

	res, err := client.UpdatePerson(context.Background(), person)
	fmt.Printf("Update (%d): %v\n", id, res.Person)
	return err
}

// client main function
func main() {
	(&sc.Cmds{
		{
			Name: "list",
			Desc: "list [age]: list person(s), optinally with age",
			Run: func(c *sc.C, args []string) error {
				if len(args) == 1 {
					age, _ := strconv.Atoi(args[0])
					return list(age)
				} else {
					return list(0)
				}
			},
		},
		{
			Name: "purge",
			Desc: "purge [age]: purge person(s), optinally with age",
			Run: func(c *sc.C, args []string) error {
				if len(args) == 1 {
					age, _ := strconv.Atoi(args[0])
					return purge(age)
				} else {
					return purge(0)
				}
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
			Name: "get",
			Desc: "get [id]: get person",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 1 {
					return sc.UsageError
				}
				id, _ := strconv.Atoi(args[0])
				return get(id)
			},
		},
		{
			Name: "update",
			Desc: "delete [id] [name] [age]: update person",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 3 {
					return sc.UsageError
				}
				id, _ := strconv.Atoi(args[0])
				name := args[1]
				age, err := strconv.Atoi(args[2])
				if err != nil {
					return err
				}
				return update(id, name, age)
			},
		},
		{
			Name: "delete",
			Desc: "delete [id]: delete person",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 1 {
					return sc.UsageError
				}
				id, _ := strconv.Atoi(args[0])
				return delete(id)
			},
		},
	}).Run(&sc.C{
		Desc: "guthub.com/sikang99/grpc-example",
	})
}
