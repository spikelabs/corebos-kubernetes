// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kubernetes.proto

package kubernetes

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Deployment struct {
	Replicas             int32    `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Label                string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{0}
}

func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *Deployment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Deployment) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type Service struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{1}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type DeploymentPvc struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Storage              string   `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeploymentPvc) Reset()         { *m = DeploymentPvc{} }
func (m *DeploymentPvc) String() string { return proto.CompactTextString(m) }
func (*DeploymentPvc) ProtoMessage()    {}
func (*DeploymentPvc) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{2}
}

func (m *DeploymentPvc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentPvc.Unmarshal(m, b)
}
func (m *DeploymentPvc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentPvc.Marshal(b, m, deterministic)
}
func (m *DeploymentPvc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentPvc.Merge(m, src)
}
func (m *DeploymentPvc) XXX_Size() int {
	return xxx_messageInfo_DeploymentPvc.Size(m)
}
func (m *DeploymentPvc) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentPvc.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentPvc proto.InternalMessageInfo

func (m *DeploymentPvc) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeploymentPvc) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type Ingress struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ingress) Reset()         { *m = Ingress{} }
func (m *Ingress) String() string { return proto.CompactTextString(m) }
func (*Ingress) ProtoMessage()    {}
func (*Ingress) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{3}
}

func (m *Ingress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ingress.Unmarshal(m, b)
}
func (m *Ingress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ingress.Marshal(b, m, deterministic)
}
func (m *Ingress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingress.Merge(m, src)
}
func (m *Ingress) XXX_Size() int {
	return xxx_messageInfo_Ingress.Size(m)
}
func (m *Ingress) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingress.DiscardUnknown(m)
}

var xxx_messageInfo_Ingress proto.InternalMessageInfo

func (m *Ingress) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Database struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	DbUsername           string   `protobuf:"bytes,3,opt,name=db_username,json=dbUsername,proto3" json:"db_username,omitempty"`
	DbPassword           string   `protobuf:"bytes,4,opt,name=db_password,json=dbPassword,proto3" json:"db_password,omitempty"`
	DbDatabase           string   `protobuf:"bytes,5,opt,name=db_database,json=dbDatabase,proto3" json:"db_database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Database) Reset()         { *m = Database{} }
func (m *Database) String() string { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()    {}
func (*Database) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{4}
}

func (m *Database) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Database.Unmarshal(m, b)
}
func (m *Database) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Database.Marshal(b, m, deterministic)
}
func (m *Database) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Database.Merge(m, src)
}
func (m *Database) XXX_Size() int {
	return xxx_messageInfo_Database.Size(m)
}
func (m *Database) XXX_DiscardUnknown() {
	xxx_messageInfo_Database.DiscardUnknown(m)
}

var xxx_messageInfo_Database proto.InternalMessageInfo

func (m *Database) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Database) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Database) GetDbUsername() string {
	if m != nil {
		return m.DbUsername
	}
	return ""
}

func (m *Database) GetDbPassword() string {
	if m != nil {
		return m.DbPassword
	}
	return ""
}

func (m *Database) GetDbDatabase() string {
	if m != nil {
		return m.DbDatabase
	}
	return ""
}

type DatabasePvc struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Storage              string   `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabasePvc) Reset()         { *m = DatabasePvc{} }
func (m *DatabasePvc) String() string { return proto.CompactTextString(m) }
func (*DatabasePvc) ProtoMessage()    {}
func (*DatabasePvc) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{5}
}

func (m *DatabasePvc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabasePvc.Unmarshal(m, b)
}
func (m *DatabasePvc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabasePvc.Marshal(b, m, deterministic)
}
func (m *DatabasePvc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabasePvc.Merge(m, src)
}
func (m *DatabasePvc) XXX_Size() int {
	return xxx_messageInfo_DatabasePvc.Size(m)
}
func (m *DatabasePvc) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabasePvc.DiscardUnknown(m)
}

var xxx_messageInfo_DatabasePvc proto.InternalMessageInfo

func (m *DatabasePvc) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatabasePvc) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type CreateClientRequest struct {
	Deployment           *Deployment    `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Service              *Service       `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	DeploymentPvc        *DeploymentPvc `protobuf:"bytes,3,opt,name=deployment_pvc,json=deploymentPvc,proto3" json:"deployment_pvc,omitempty"`
	Ingress              *Ingress       `protobuf:"bytes,4,opt,name=ingress,proto3" json:"ingress,omitempty"`
	Database             *Database      `protobuf:"bytes,5,opt,name=database,proto3" json:"database,omitempty"`
	DatabasePvc          *DatabasePvc   `protobuf:"bytes,6,opt,name=database_pvc,json=databasePvc,proto3" json:"database_pvc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateClientRequest) Reset()         { *m = CreateClientRequest{} }
func (m *CreateClientRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClientRequest) ProtoMessage()    {}
func (*CreateClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{6}
}

func (m *CreateClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClientRequest.Unmarshal(m, b)
}
func (m *CreateClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClientRequest.Marshal(b, m, deterministic)
}
func (m *CreateClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClientRequest.Merge(m, src)
}
func (m *CreateClientRequest) XXX_Size() int {
	return xxx_messageInfo_CreateClientRequest.Size(m)
}
func (m *CreateClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClientRequest proto.InternalMessageInfo

func (m *CreateClientRequest) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *CreateClientRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *CreateClientRequest) GetDeploymentPvc() *DeploymentPvc {
	if m != nil {
		return m.DeploymentPvc
	}
	return nil
}

func (m *CreateClientRequest) GetIngress() *Ingress {
	if m != nil {
		return m.Ingress
	}
	return nil
}

func (m *CreateClientRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *CreateClientRequest) GetDatabasePvc() *DatabasePvc {
	if m != nil {
		return m.DatabasePvc
	}
	return nil
}

type CreateClientResponse struct {
	Success              int32    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateClientResponse) Reset()         { *m = CreateClientResponse{} }
func (m *CreateClientResponse) String() string { return proto.CompactTextString(m) }
func (*CreateClientResponse) ProtoMessage()    {}
func (*CreateClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{7}
}

func (m *CreateClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClientResponse.Unmarshal(m, b)
}
func (m *CreateClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClientResponse.Marshal(b, m, deterministic)
}
func (m *CreateClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClientResponse.Merge(m, src)
}
func (m *CreateClientResponse) XXX_Size() int {
	return xxx_messageInfo_CreateClientResponse.Size(m)
}
func (m *CreateClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClientResponse proto.InternalMessageInfo

func (m *CreateClientResponse) GetSuccess() int32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *CreateClientResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Deployment)(nil), "Deployment")
	proto.RegisterType((*Service)(nil), "Service")
	proto.RegisterType((*DeploymentPvc)(nil), "DeploymentPvc")
	proto.RegisterType((*Ingress)(nil), "Ingress")
	proto.RegisterType((*Database)(nil), "Database")
	proto.RegisterType((*DatabasePvc)(nil), "DatabasePvc")
	proto.RegisterType((*CreateClientRequest)(nil), "CreateClientRequest")
	proto.RegisterType((*CreateClientResponse)(nil), "CreateClientResponse")
}

func init() { proto.RegisterFile("kubernetes.proto", fileDescriptor_40204d9320c6ada8) }

var fileDescriptor_40204d9320c6ada8 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x8b, 0xdb, 0x30,
	0x10, 0xad, 0xd3, 0xcd, 0xda, 0x3b, 0xca, 0x86, 0xa2, 0xa6, 0x60, 0x16, 0x4a, 0x17, 0x41, 0x61,
	0xa1, 0xa0, 0x82, 0x43, 0x4f, 0x25, 0xa7, 0x84, 0x42, 0x0f, 0xa5, 0x41, 0xa5, 0xe7, 0x20, 0xdb,
	0x43, 0x30, 0x75, 0x6c, 0x57, 0x92, 0x53, 0xfa, 0x4b, 0x4a, 0xff, 0xed, 0x12, 0xd9, 0xf2, 0x07,
	0xf8, 0x92, 0x9b, 0xde, 0xcc, 0xbc, 0x99, 0xf1, 0x7b, 0x63, 0x78, 0xf5, 0xab, 0x8e, 0x51, 0x15,
	0x68, 0x50, 0xf3, 0x4a, 0x95, 0xa6, 0x64, 0x02, 0x60, 0x87, 0x55, 0x5e, 0xfe, 0x3d, 0x61, 0x61,
	0xe8, 0x03, 0x04, 0x0a, 0xab, 0x3c, 0x4b, 0xa4, 0x0e, 0xbd, 0x47, 0xef, 0x69, 0x2e, 0x3a, 0x4c,
	0x29, 0xdc, 0x14, 0xf2, 0x84, 0xe1, 0xec, 0xd1, 0x7b, 0xba, 0x13, 0xf6, 0x4d, 0x57, 0x30, 0xcf,
	0x65, 0x8c, 0x79, 0xf8, 0xd2, 0x06, 0x1b, 0xc0, 0xd6, 0xe0, 0xff, 0x40, 0x75, 0xce, 0x12, 0xec,
	0x48, 0xde, 0x14, 0x69, 0x36, 0x24, 0x6d, 0xe0, 0xbe, 0x5f, 0x64, 0x7f, 0x4e, 0x26, 0xa9, 0x21,
	0xf8, 0xda, 0x94, 0x4a, 0x1e, 0xdd, 0x1a, 0x0e, 0xb2, 0xb7, 0xe0, 0x7f, 0x2d, 0x8e, 0x0a, 0xb5,
	0x9e, 0x22, 0xb2, 0x7f, 0x1e, 0x04, 0x3b, 0x69, 0x64, 0x2c, 0xf5, 0x15, 0x4b, 0xd1, 0x77, 0x40,
	0xd2, 0xf8, 0x50, 0x6b, 0x54, 0x96, 0xd0, 0x7c, 0x25, 0xa4, 0xf1, 0xcf, 0x36, 0xd2, 0x16, 0x54,
	0x52, 0xeb, 0x3f, 0xa5, 0x4a, 0xc3, 0x1b, 0x57, 0xb0, 0x6f, 0x23, 0x6d, 0x41, 0xda, 0x8e, 0x0e,
	0xe7, 0xae, 0xc0, 0x2d, 0xc3, 0x3e, 0x03, 0x71, 0xef, 0xeb, 0xbf, 0xfa, 0xff, 0x0c, 0x5e, 0x6f,
	0x15, 0x4a, 0x83, 0xdb, 0x3c, 0xc3, 0xc2, 0x08, 0xfc, 0x5d, 0xa3, 0x36, 0xf4, 0x03, 0x40, 0xda,
	0x89, 0x69, 0x7b, 0x91, 0x88, 0xf0, 0x5e, 0x5f, 0x31, 0x48, 0x53, 0x06, 0xbe, 0x6e, 0xec, 0xb2,
	0xed, 0x49, 0x14, 0xf0, 0xd6, 0x3e, 0xe1, 0x12, 0xf4, 0x13, 0x2c, 0x7b, 0xc6, 0xa1, 0x3a, 0x27,
	0x56, 0x0b, 0x12, 0x2d, 0xf9, 0xc8, 0x34, 0x71, 0x9f, 0x8e, 0x3c, 0x64, 0xe0, 0x67, 0x8d, 0x2b,
	0x56, 0x9a, 0x4b, 0xeb, 0xd6, 0x25, 0xe1, 0x12, 0xf4, 0x3d, 0x04, 0x23, 0x79, 0x48, 0x74, 0xc7,
	0x9d, 0x22, 0xa2, 0x4b, 0xd1, 0x8f, 0xb0, 0x70, 0x6f, 0x3b, 0xff, 0xd6, 0x96, 0x2e, 0xf8, 0x40,
	0x3c, 0x41, 0xd2, 0x1e, 0xb0, 0x2f, 0xb0, 0x1a, 0x4b, 0xa3, 0xab, 0xb2, 0xd0, 0x8d, 0x9a, 0x75,
	0x92, 0x5c, 0x76, 0x6a, 0x4e, 0xdc, 0xc1, 0xcb, 0x0d, 0xa0, 0x52, 0xa5, 0x72, 0x37, 0x60, 0x41,
	0xf4, 0x1d, 0x96, 0xdb, 0xbc, 0xd6, 0x06, 0xd5, 0x37, 0x59, 0xc8, 0x23, 0x2a, 0xba, 0x81, 0xc5,
	0xb0, 0x33, 0x5d, 0xf1, 0x09, 0x0f, 0x1e, 0xde, 0xf0, 0xa9, 0xf1, 0xec, 0x45, 0x7c, 0x6b, 0xff,
	0xbc, 0xf5, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xc5, 0xc1, 0x80, 0x8d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterManagerClient is the client API for ClusterManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterManagerClient interface {
	CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error)
}

type clusterManagerClient struct {
	cc *grpc.ClientConn
}

func NewClusterManagerClient(cc *grpc.ClientConn) ClusterManagerClient {
	return &clusterManagerClient{cc}
}

func (c *clusterManagerClient) CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error) {
	out := new(CreateClientResponse)
	err := c.cc.Invoke(ctx, "/ClusterManager/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterManagerServer is the server API for ClusterManager service.
type ClusterManagerServer interface {
	CreateClient(context.Context, *CreateClientRequest) (*CreateClientResponse, error)
}

func RegisterClusterManagerServer(s *grpc.Server, srv ClusterManagerServer) {
	s.RegisterService(&_ClusterManager_serviceDesc, srv)
}

func _ClusterManager_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterManagerServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClusterManager/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterManagerServer).CreateClient(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ClusterManager",
	HandlerType: (*ClusterManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ClusterManager_CreateClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubernetes.proto",
}
