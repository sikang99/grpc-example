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
	sync.Mutex
}

func NewCustomerService() *customerService {
	cs := customerService{
		customers: make(map[int]*pb.Person),
		id:        1,
	}

	return &cs
}

func (cs *customerService) ListPersons(p *pb.RequestType, stream pb.CustomerService_ListPersonsServer) error {
	log.Printf("list (%v)\n", p.Man)
	cs.Lock()
	defer cs.Unlock()

	// check the condition
	var age int
	if p.Man != nil {
		age = int(p.Man.Age)
	}

	for _, ps := range cs.customers {
		// conditional listing
		if age > 0 && int(ps.Age) != age {
			continue
		}
		if err := stream.Send(ps); err != nil {
			return err
		}
	}
	return nil
}

func (cs *customerService) GetPerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	log.Printf("get (%d)\n", p.Id)
	cs.Lock()
	defer cs.Unlock()

	resp := new(pb.ResponseType)

	ps, ok := cs.customers[int(p.Id)]
	if ok {
		resp.Man = ps
	}

	log.Printf("%v\n", resp.Man)
	return resp, nil
}

func (cs *customerService) AddPerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	log.Printf("add (%d)\n", cs.id)
	cs.Lock()
	defer cs.Unlock()

	p.Id = int32(cs.id)
	cs.customers[int(p.Id)] = p
	cs.id++

	resp := new(pb.ResponseType)

	ps, ok := cs.customers[int(p.Id)]
	if ok {
		resp.Man = ps
	}

	log.Printf("%v\n", resp.Man)
	return resp, nil
}

func (cs *customerService) DeletePerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	log.Printf("delete (%d)\n", p.Id)
	cs.Lock()
	defer cs.Unlock()

	resp := new(pb.ResponseType)

	ps, ok := cs.customers[int(p.Id)]
	if ok {
		resp.Man = ps
		delete(cs.customers, int(p.Id))
	}

	log.Printf("%v\n", resp.Man)
	return resp, nil
}

func (cs *customerService) UpdatePerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {
	log.Printf("update (%d)\n", p.Id)
	cs.Lock()
	defer cs.Unlock()

	resp := new(pb.ResponseType)

	_, ok := cs.customers[int(p.Id)]
	if ok {
		cs.customers[int(p.Id)] = p
		ps, ok := cs.customers[int(p.Id)]
		if ok {
			resp.Man = ps
		}
	}

	log.Printf("%v\n", resp.Man)
	return resp, nil
}

// server function
func main() {
	lis, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterCustomerServiceServer(server, NewCustomerService())
	server.Serve(lis)
}
