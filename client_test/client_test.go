package client

import (
	"io"
	"testing"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var address string = "127.0.0.1:11111"

func TestCustomerService(t *testing.T) {
	conn, err := grpc.Dial(address)
	if err != nil {
		t.Errorf("connect error %v\n", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	//var age int
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
}
