package main

import (
	"log"
	"net"
	"sync"

	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type customerService struct {
	customers map[int]*pb.Person
	id        int // current unique id number
	m         sync.Mutex
}

func NewCustomerService() *customerService {
	cs := customerService{
		customers: make(map[int]*pb.Person),
	}

	return &cs
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
	cs.id++
	p.Id = int32(cs.id)
	cs.customers[int(p.Id)] = p
	return new(pb.ResponseType), nil
}

func (cs *customerService) DeletePerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	log.Println("delete")
	cs.m.Lock()
	defer cs.m.Unlock()
	delete(cs.customers, int(p.Id))
	return new(pb.ResponseType), nil
}

func main() {
	lis, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterCustomerServiceServer(server, NewCustomerService())
	server.Serve(lis)
}
