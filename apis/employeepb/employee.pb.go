// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: apis/employeepb/employee.proto

package employeepb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId uint32 `protobuf:"varint,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
}

func (x *EmployeeRequest) Reset() {
	*x = EmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_employeepb_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeRequest) ProtoMessage() {}

func (x *EmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_employeepb_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeRequest.ProtoReflect.Descriptor instead.
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return file_apis_employeepb_employee_proto_rawDescGZIP(), []int{0}
}

func (x *EmployeeRequest) GetEmployeeId() uint32 {
	if x != nil {
		return x.EmployeeId
	}
	return 0
}

type EmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId uint32 `protobuf:"varint,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	FullName   string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone      string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Address    string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Gender     string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	JobTitle   string `protobuf:"bytes,7,opt,name=job_title,json=jobTitle,proto3" json:"job_title,omitempty"`
}

func (x *EmployeeResponse) Reset() {
	*x = EmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_employeepb_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeResponse) ProtoMessage() {}

func (x *EmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_employeepb_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeResponse.ProtoReflect.Descriptor instead.
func (*EmployeeResponse) Descriptor() ([]byte, []int) {
	return file_apis_employeepb_employee_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeResponse) GetEmployeeId() uint32 {
	if x != nil {
		return x.EmployeeId
	}
	return 0
}

func (x *EmployeeResponse) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *EmployeeResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmployeeResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *EmployeeResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EmployeeResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *EmployeeResponse) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

type ListEmployeesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListEmployeesRequest) Reset() {
	*x = ListEmployeesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_employeepb_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmployeesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmployeesRequest) ProtoMessage() {}

func (x *ListEmployeesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_employeepb_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmployeesRequest.ProtoReflect.Descriptor instead.
func (*ListEmployeesRequest) Descriptor() ([]byte, []int) {
	return file_apis_employeepb_employee_proto_rawDescGZIP(), []int{2}
}

func (x *ListEmployeesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListEmployeesRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_apis_employeepb_employee_proto protoreflect.FileDescriptor

var file_apis_employeepb_employee_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x70,
	0x62, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x32, 0x0a, 0x0f, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x22, 0xcb,
	0x01, 0x0a, 0x10, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x44, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x32, 0xad, 0x01, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x19, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x73, 0x12, 0x1e, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_employeepb_employee_proto_rawDescOnce sync.Once
	file_apis_employeepb_employee_proto_rawDescData = file_apis_employeepb_employee_proto_rawDesc
)

func file_apis_employeepb_employee_proto_rawDescGZIP() []byte {
	file_apis_employeepb_employee_proto_rawDescOnce.Do(func() {
		file_apis_employeepb_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_employeepb_employee_proto_rawDescData)
	})
	return file_apis_employeepb_employee_proto_rawDescData
}

var file_apis_employeepb_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apis_employeepb_employee_proto_goTypes = []interface{}{
	(*EmployeeRequest)(nil),      // 0: employee.EmployeeRequest
	(*EmployeeResponse)(nil),     // 1: employee.EmployeeResponse
	(*ListEmployeesRequest)(nil), // 2: employee.ListEmployeesRequest
}
var file_apis_employeepb_employee_proto_depIdxs = []int32{
	0, // 0: employee.EmployeeService.GetEmployee:input_type -> employee.EmployeeRequest
	2, // 1: employee.EmployeeService.GetListEmployees:input_type -> employee.ListEmployeesRequest
	1, // 2: employee.EmployeeService.GetEmployee:output_type -> employee.EmployeeResponse
	1, // 3: employee.EmployeeService.GetListEmployees:output_type -> employee.EmployeeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_employeepb_employee_proto_init() }
func file_apis_employeepb_employee_proto_init() {
	if File_apis_employeepb_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_employeepb_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_employeepb_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apis_employeepb_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEmployeesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_employeepb_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_employeepb_employee_proto_goTypes,
		DependencyIndexes: file_apis_employeepb_employee_proto_depIdxs,
		MessageInfos:      file_apis_employeepb_employee_proto_msgTypes,
	}.Build()
	File_apis_employeepb_employee_proto = out.File
	file_apis_employeepb_employee_proto_rawDesc = nil
	file_apis_employeepb_employee_proto_goTypes = nil
	file_apis_employeepb_employee_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	GetEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	GetListEmployees(ctx context.Context, in *ListEmployeesRequest, opts ...grpc.CallOption) (EmployeeService_GetListEmployeesClient, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/employee.EmployeeService/GetEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetListEmployees(ctx context.Context, in *ListEmployeesRequest, opts ...grpc.CallOption) (EmployeeService_GetListEmployeesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EmployeeService_serviceDesc.Streams[0], "/employee.EmployeeService/GetListEmployees", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceGetListEmployeesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmployeeService_GetListEmployeesClient interface {
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceGetListEmployeesClient struct {
	grpc.ClientStream
}

func (x *employeeServiceGetListEmployeesClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
type EmployeeServiceServer interface {
	GetEmployee(context.Context, *EmployeeRequest) (*EmployeeResponse, error)
	GetListEmployees(*ListEmployeesRequest, EmployeeService_GetListEmployeesServer) error
}

// UnimplementedEmployeeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (*UnimplementedEmployeeServiceServer) GetEmployee(context.Context, *EmployeeRequest) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}
func (*UnimplementedEmployeeServiceServer) GetListEmployees(*ListEmployeesRequest, EmployeeService_GetListEmployeesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetListEmployees not implemented")
}

func RegisterEmployeeServiceServer(s *grpc.Server, srv EmployeeServiceServer) {
	s.RegisterService(&_EmployeeService_serviceDesc, srv)
}

func _EmployeeService_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.EmployeeService/GetEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetListEmployees_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEmployeesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeeServiceServer).GetListEmployees(m, &employeeServiceGetListEmployeesServer{stream})
}

type EmployeeService_GetListEmployeesServer interface {
	Send(*EmployeeResponse) error
	grpc.ServerStream
}

type employeeServiceGetListEmployeesServer struct {
	grpc.ServerStream
}

func (x *employeeServiceGetListEmployeesServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _EmployeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "employee.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployee",
			Handler:    _EmployeeService_GetEmployee_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetListEmployees",
			Handler:       _EmployeeService_GetListEmployees_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apis/employeepb/employee.proto",
}
