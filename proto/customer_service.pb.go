// Code generated by protoc-gen-go.
// source: customer_service.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	customer_service.proto

It has these top-level messages:
	ResponseType
	RequestType
	Person
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type ResponseType struct {
}

func (m *ResponseType) Reset()         { *m = ResponseType{} }
func (m *ResponseType) String() string { return proto1.CompactTextString(m) }
func (*ResponseType) ProtoMessage()    {}

type RequestType struct {
	Op string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
}

func (m *RequestType) Reset()         { *m = RequestType{} }
func (m *RequestType) String() string { return proto1.CompactTextString(m) }
func (*RequestType) ProtoMessage()    {}

type Person struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto1.CompactTextString(m) }
func (*Person) ProtoMessage()    {}

func init() {
}

// Client API for CustomerService service

type CustomerServiceClient interface {
	ListPerson(ctx context.Context, in *RequestType, opts ...grpc.CallOption) (CustomerService_ListPersonClient, error)
	AddPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*ResponseType, error)
	DeletePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*ResponseType, error)
}

type customerServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerServiceClient(cc *grpc.ClientConn) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) ListPerson(ctx context.Context, in *RequestType, opts ...grpc.CallOption) (CustomerService_ListPersonClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CustomerService_serviceDesc.Streams[0], c.cc, "/proto.CustomerService/ListPerson", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerServiceListPersonClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CustomerService_ListPersonClient interface {
	Recv() (*Person, error)
	grpc.ClientStream
}

type customerServiceListPersonClient struct {
	grpc.ClientStream
}

func (x *customerServiceListPersonClient) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerServiceClient) AddPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*ResponseType, error) {
	out := new(ResponseType)
	err := grpc.Invoke(ctx, "/proto.CustomerService/AddPerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) DeletePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*ResponseType, error) {
	out := new(ResponseType)
	err := grpc.Invoke(ctx, "/proto.CustomerService/DeletePerson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomerService service

type CustomerServiceServer interface {
	ListPerson(*RequestType, CustomerService_ListPersonServer) error
	AddPerson(context.Context, *Person) (*ResponseType, error)
	DeletePerson(context.Context, *Person) (*ResponseType, error)
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_ListPerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestType)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServiceServer).ListPerson(m, &customerServiceListPersonServer{stream})
}

type CustomerService_ListPersonServer interface {
	Send(*Person) error
	grpc.ServerStream
}

type customerServiceListPersonServer struct {
	grpc.ServerStream
}

func (x *customerServiceListPersonServer) Send(m *Person) error {
	return x.ServerStream.SendMsg(m)
}

func _CustomerService_AddPerson_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Person)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CustomerServiceServer).AddPerson(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CustomerService_DeletePerson_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Person)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CustomerServiceServer).DeletePerson(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPerson",
			Handler:    _CustomerService_AddPerson_Handler,
		},
		{
			MethodName: "DeletePerson",
			Handler:    _CustomerService_DeletePerson_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPerson",
			Handler:       _CustomerService_ListPerson_Handler,
			ServerStreams: true,
		},
	},
}
