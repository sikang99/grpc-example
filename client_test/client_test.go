package client

import (
	"io"
	"testing"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var address string = "127.0.0.1:11111"

func TestBaseForCustomerService(t *testing.T) {
	// for List
	conn, err := grpc.Dial(address)
	if err != nil {
		t.Errorf("connect error %v\n", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	req := &pb.RequestType{}

	stream, err := client.ListPersons(context.Background(), req)
	if err != nil {
		t.Errorf("request error %v\n", err)
	}
	for {
		person, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Errorf("recv error %v\n", err)
		}
		t.Logf("List: %v\n", person)
	}

	// test functions of CRUD: Add, Get, Update, and Delete

	// for Add
	person := &pb.Person{
		Name: "Stoney",
		Age:  52,
	}

	res, err := client.AddPerson(context.Background(), person)
	t.Logf("Add: %v\n", res.Man)

	// for Get
	person.Id = res.Man.Id

	res, err = client.GetPerson(context.Background(), person)
	t.Logf("Get: %v\n", res.Man)

	// for Update
	person.Name = "Minwoo"
	person.Age = 19

	res, err = client.UpdatePerson(context.Background(), person)
	t.Logf("Update: %v\n", res.Man)

	// for Delete
	res, err = client.DeletePerson(context.Background(), person)
	t.Logf("Delete: %v\n", res.Man)

	// check if the res is nil
	res, err = client.GetPerson(context.Background(), person)
	t.Logf("Get: %v\n", res.Man)
	if res.Man != nil {
		t.Errorf("result error %v\n", res.Man)
	}
}

func TestLoadForCustomerService(t *testing.T) {

}
