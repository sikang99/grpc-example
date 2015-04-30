package main

import (
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//pb "github.com/sikang99/grpc-example/proto"
	pb "../proto"
)

type customerService struct {
	customers []*pb.Person
	m         sync.Mutex
}

func (cs *customerService) ListPerson(p *pb.RequestType, stream pb.CustomerService_ListPersonServer) error {
	log.Println("list")
	cs.m.Lock()
	defer cs.m.Unlock()
	for _, p := range cs.customers {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

func (cs *customerService) AddPerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	log.Println("add")
	cs.m.Lock()
	defer cs.m.Unlock()
	cs.customers = append(cs.customers, p)
	return new(pb.ResponseType), nil
}

func (cs *customerService) DeletePerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	log.Println("delete")
	cs.m.Lock()
	defer cs.m.Unlock()
	//cs.customers = delete(cs.customers, p)
	return new(pb.ResponseType), nil
}

func main() {
	lis, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterCustomerServiceServer(server, new(customerService))
	server.Serve(lis)
}
