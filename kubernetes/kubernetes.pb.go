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
	SubDomain            string   `protobuf:"bytes,2,opt,name=sub_domain,json=subDomain,proto3" json:"sub_domain,omitempty"`
	Resource             string   `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
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

func (m *Ingress) GetSubDomain() string {
	if m != nil {
		return m.SubDomain
	}
	return ""
}

func (m *Ingress) GetResource() string {
	if m != nil {
		return m.Resource
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
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0x36, 0xf5, 0x7a, 0x69, 0x27, 0xbd, 0x22, 0x6b, 0x85, 0x70, 0x20, 0x1e, 0x0b, 0xc2, 0x81,
	0x10, 0x21, 0x87, 0x4f, 0x72, 0x4f, 0x2d, 0x82, 0x0f, 0xe2, 0xb1, 0x22, 0xf8, 0x56, 0x36, 0xc9,
	0x50, 0x82, 0x69, 0x36, 0xee, 0x6c, 0x2a, 0xfe, 0x12, 0xf1, 0xdf, 0x4a, 0x37, 0xbb, 0x69, 0x03,
	0x79, 0xe9, 0xdb, 0x7e, 0x33, 0xf3, 0xcd, 0xce, 0x7c, 0xdf, 0x2e, 0xbc, 0xf8, 0xd9, 0x66, 0xa8,
	0x6b, 0x34, 0x48, 0x49, 0xa3, 0x95, 0x51, 0x5c, 0x00, 0x6c, 0xb0, 0xa9, 0xd4, 0x9f, 0x3d, 0xd6,
	0x86, 0xdd, 0xc2, 0x4c, 0x63, 0x53, 0x95, 0xb9, 0xa4, 0x38, 0xb8, 0x0b, 0xee, 0xa7, 0xa2, 0xc7,
	0x8c, 0xc1, 0x55, 0x2d, 0xf7, 0x18, 0x4f, 0xee, 0x82, 0xfb, 0xb9, 0xb0, 0x67, 0xb6, 0x82, 0x69,
	0x25, 0x33, 0xac, 0xe2, 0xe7, 0x36, 0xd8, 0x01, 0xfe, 0x00, 0xe1, 0x37, 0xd4, 0x87, 0x32, 0xc7,
	0x9e, 0x14, 0x8c, 0x91, 0x26, 0xe7, 0xa4, 0x47, 0xb8, 0x39, 0x0d, 0xf2, 0x74, 0xc8, 0x47, 0xa9,
	0x31, 0x84, 0x64, 0x94, 0x96, 0x3b, 0x3f, 0x86, 0x87, 0xfc, 0x07, 0x84, 0x9f, 0xeb, 0x9d, 0x46,
	0xa2, 0x51, 0xe2, 0x6b, 0x00, 0x6a, 0xb3, 0x6d, 0xa1, 0xf6, 0xb2, 0xac, 0x1d, 0x77, 0x4e, 0x6d,
	0xb6, 0xb1, 0x81, 0x6e, 0x6f, 0x52, 0xad, 0xce, 0xd1, 0xad, 0xd2, 0x63, 0xfe, 0x37, 0x80, 0xd9,
	0x46, 0x1a, 0x99, 0x49, 0xba, 0x60, 0x1f, 0xf6, 0x06, 0xa2, 0x22, 0xdb, 0xb6, 0x84, 0xda, 0x12,
	0xba, 0xae, 0x50, 0x64, 0xdf, 0x5d, 0xc4, 0x15, 0x34, 0x92, 0xe8, 0xb7, 0xd2, 0x45, 0x7c, 0xe5,
	0x0b, 0x9e, 0x5c, 0xc4, 0x15, 0x14, 0xee, 0xea, 0x78, 0xea, 0x0b, 0xfc, 0x30, 0xfc, 0x23, 0x44,
	0xfe, 0x7c, 0xb9, 0x60, 0xff, 0x26, 0xf0, 0x72, 0xad, 0x51, 0x1a, 0x5c, 0x57, 0x25, 0xd6, 0x46,
	0xe0, 0xaf, 0x16, 0xc9, 0xb0, 0x77, 0x00, 0x45, 0xef, 0x83, 0xed, 0x15, 0xa5, 0x51, 0x72, 0xb2,
	0x46, 0x9c, 0xa5, 0x19, 0x87, 0x90, 0x3a, 0xa7, 0x6d, 0xfb, 0x28, 0x9d, 0x25, 0xce, 0x79, 0xe1,
	0x13, 0xec, 0x03, 0x2c, 0x4f, 0x8c, 0x6d, 0x73, 0xc8, 0xad, 0x16, 0x51, 0xba, 0x4c, 0x06, 0x7e,
	0x8b, 0x9b, 0x62, 0x60, 0x3f, 0x87, 0xb0, 0xec, 0x0c, 0xb5, 0xd2, 0x1c, 0x5b, 0x3b, 0x83, 0x85,
	0x4f, 0xb0, 0xb7, 0x30, 0x1b, 0xc8, 0x13, 0xa5, 0xf3, 0xc4, 0x2b, 0x22, 0xfa, 0x14, 0x7b, 0x0f,
	0x0b, 0x7f, 0xb6, 0xf7, 0x5f, 0xdb, 0xd2, 0x45, 0x72, 0x26, 0x9e, 0x88, 0x8a, 0x13, 0xe0, 0x9f,
	0x60, 0x35, 0x94, 0x86, 0x1a, 0x55, 0x53, 0xa7, 0x66, 0x9b, 0xe7, 0xc7, 0x99, 0xba, 0xdf, 0xe1,
	0xe1, 0xf1, 0x0d, 0xa0, 0xd6, 0x4a, 0xfb, 0x37, 0x60, 0x41, 0xfa, 0x15, 0x96, 0xeb, 0xaa, 0x25,
	0x83, 0xfa, 0x8b, 0xac, 0xe5, 0x0e, 0x35, 0x7b, 0x84, 0xc5, 0x79, 0x67, 0xb6, 0x4a, 0x46, 0x3c,
	0xb8, 0x7d, 0x95, 0x8c, 0x5d, 0xcf, 0x9f, 0x65, 0xd7, 0xf6, 0xd3, 0x3e, 0xfc, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x89, 0xc6, 0x05, 0xc0, 0xc8, 0x03, 0x00, 0x00,
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
