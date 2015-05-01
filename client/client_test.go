package client_test

import (
	"fmt"
	"io"
	"testing"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var address string = "127.0.0.1:11111"

func TestListPersons(t *testing.T) {
	conn, err := grpc.Dial(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	var age int
	req := &pb.RequestType{}

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
