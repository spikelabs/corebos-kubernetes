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

type DatabaseService struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseService) Reset()         { *m = DatabaseService{} }
func (m *DatabaseService) String() string { return proto.CompactTextString(m) }
func (*DatabaseService) ProtoMessage()    {}
func (*DatabaseService) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{5}
}

func (m *DatabaseService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseService.Unmarshal(m, b)
}
func (m *DatabaseService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseService.Marshal(b, m, deterministic)
}
func (m *DatabaseService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseService.Merge(m, src)
}
func (m *DatabaseService) XXX_Size() int {
	return xxx_messageInfo_DatabaseService.Size(m)
}
func (m *DatabaseService) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseService.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseService proto.InternalMessageInfo

func (m *DatabaseService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatabaseService) GetLabel() string {
	if m != nil {
		return m.Label
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
	return fileDescriptor_40204d9320c6ada8, []int{6}
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

type CreateClientDatabaseRequest struct {
	Database             *Database        `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	DatabaseService      *DatabaseService `protobuf:"bytes,2,opt,name=database_service,json=databaseService,proto3" json:"database_service,omitempty"`
	DatabasePvc          *DatabasePvc     `protobuf:"bytes,3,opt,name=database_pvc,json=databasePvc,proto3" json:"database_pvc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateClientDatabaseRequest) Reset()         { *m = CreateClientDatabaseRequest{} }
func (m *CreateClientDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClientDatabaseRequest) ProtoMessage()    {}
func (*CreateClientDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{7}
}

func (m *CreateClientDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClientDatabaseRequest.Unmarshal(m, b)
}
func (m *CreateClientDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClientDatabaseRequest.Marshal(b, m, deterministic)
}
func (m *CreateClientDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClientDatabaseRequest.Merge(m, src)
}
func (m *CreateClientDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateClientDatabaseRequest.Size(m)
}
func (m *CreateClientDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClientDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClientDatabaseRequest proto.InternalMessageInfo

func (m *CreateClientDatabaseRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *CreateClientDatabaseRequest) GetDatabaseService() *DatabaseService {
	if m != nil {
		return m.DatabaseService
	}
	return nil
}

func (m *CreateClientDatabaseRequest) GetDatabasePvc() *DatabasePvc {
	if m != nil {
		return m.DatabasePvc
	}
	return nil
}

type CreateClientDeploymentRequest struct {
	Deployment           *Deployment    `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Service              *Service       `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	DeploymentPvc        *DeploymentPvc `protobuf:"bytes,3,opt,name=deployment_pvc,json=deploymentPvc,proto3" json:"deployment_pvc,omitempty"`
	Ingress              *Ingress       `protobuf:"bytes,4,opt,name=ingress,proto3" json:"ingress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateClientDeploymentRequest) Reset()         { *m = CreateClientDeploymentRequest{} }
func (m *CreateClientDeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClientDeploymentRequest) ProtoMessage()    {}
func (*CreateClientDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{8}
}

func (m *CreateClientDeploymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClientDeploymentRequest.Unmarshal(m, b)
}
func (m *CreateClientDeploymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClientDeploymentRequest.Marshal(b, m, deterministic)
}
func (m *CreateClientDeploymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClientDeploymentRequest.Merge(m, src)
}
func (m *CreateClientDeploymentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateClientDeploymentRequest.Size(m)
}
func (m *CreateClientDeploymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClientDeploymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClientDeploymentRequest proto.InternalMessageInfo

func (m *CreateClientDeploymentRequest) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *CreateClientDeploymentRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *CreateClientDeploymentRequest) GetDeploymentPvc() *DeploymentPvc {
	if m != nil {
		return m.DeploymentPvc
	}
	return nil
}

func (m *CreateClientDeploymentRequest) GetIngress() *Ingress {
	if m != nil {
		return m.Ingress
	}
	return nil
}

type ClientResponse struct {
	Success              int32    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientResponse) Reset()         { *m = ClientResponse{} }
func (m *ClientResponse) String() string { return proto.CompactTextString(m) }
func (*ClientResponse) ProtoMessage()    {}
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40204d9320c6ada8, []int{9}
}

func (m *ClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientResponse.Unmarshal(m, b)
}
func (m *ClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientResponse.Marshal(b, m, deterministic)
}
func (m *ClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientResponse.Merge(m, src)
}
func (m *ClientResponse) XXX_Size() int {
	return xxx_messageInfo_ClientResponse.Size(m)
}
func (m *ClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientResponse proto.InternalMessageInfo

func (m *ClientResponse) GetSuccess() int32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *ClientResponse) GetError() string {
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
	proto.RegisterType((*DatabaseService)(nil), "DatabaseService")
	proto.RegisterType((*DatabasePvc)(nil), "DatabasePvc")
	proto.RegisterType((*CreateClientDatabaseRequest)(nil), "CreateClientDatabaseRequest")
	proto.RegisterType((*CreateClientDeploymentRequest)(nil), "CreateClientDeploymentRequest")
	proto.RegisterType((*ClientResponse)(nil), "ClientResponse")
}

func init() { proto.RegisterFile("kubernetes.proto", fileDescriptor_40204d9320c6ada8) }

var fileDescriptor_40204d9320c6ada8 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xed, 0x8a, 0xd3, 0x40,
	0x14, 0x35, 0xeb, 0xd6, 0xb4, 0x37, 0xbb, 0x6d, 0x19, 0x16, 0x09, 0xab, 0xab, 0xcb, 0x80, 0xb0,
	0x20, 0x8c, 0xd0, 0xc5, 0x5f, 0x45, 0x10, 0xb6, 0x20, 0x22, 0xc2, 0x12, 0x11, 0xfc, 0x57, 0x26,
	0xc9, 0xa5, 0x04, 0xd3, 0x24, 0xce, 0x4c, 0x2a, 0x3e, 0x89, 0x6f, 0xe0, 0x03, 0xf8, 0x1a, 0xbe,
	0x94, 0x64, 0x32, 0x33, 0x4d, 0x96, 0x50, 0xe8, 0xbf, 0xb9, 0x1f, 0xe7, 0xe6, 0x9c, 0x7b, 0x4f,
	0x0b, 0xf3, 0xef, 0x75, 0x8c, 0xa2, 0x40, 0x85, 0x92, 0x55, 0xa2, 0x54, 0x25, 0x8d, 0x00, 0x56,
	0x58, 0xe5, 0xe5, 0xaf, 0x2d, 0x16, 0x8a, 0x5c, 0xc2, 0x58, 0x60, 0x95, 0x67, 0x09, 0x97, 0xa1,
	0x77, 0xed, 0xdd, 0x8c, 0x22, 0x17, 0x13, 0x02, 0xa7, 0x05, 0xdf, 0x62, 0x78, 0x72, 0xed, 0xdd,
	0x4c, 0x22, 0xfd, 0x26, 0x17, 0x30, 0xca, 0x79, 0x8c, 0x79, 0xf8, 0x58, 0x27, 0xdb, 0x80, 0xde,
	0x82, 0xff, 0x05, 0xc5, 0x2e, 0x4b, 0xd0, 0x81, 0xbc, 0x21, 0xd0, 0x49, 0x17, 0xf4, 0x0e, 0xce,
	0xf7, 0x44, 0xee, 0x77, 0xc9, 0x20, 0x34, 0x04, 0x5f, 0xaa, 0x52, 0xf0, 0x8d, 0xa5, 0x61, 0x43,
	0xfa, 0x0d, 0xfc, 0x8f, 0xc5, 0x46, 0xa0, 0x94, 0x83, 0xc0, 0x2b, 0x00, 0x59, 0xc7, 0xeb, 0xb4,
	0xdc, 0xf2, 0xac, 0x30, 0xd8, 0x89, 0xac, 0xe3, 0x95, 0x4e, 0xb4, 0xba, 0x65, 0x59, 0x8b, 0x04,
	0x8d, 0x14, 0x17, 0xd3, 0xdf, 0x1e, 0x8c, 0x57, 0x5c, 0xf1, 0x98, 0xcb, 0x23, 0xf4, 0x90, 0x97,
	0x10, 0xa4, 0xf1, 0xba, 0x96, 0x28, 0x34, 0xa0, 0x9d, 0x0a, 0x69, 0xfc, 0xd5, 0x64, 0x4c, 0x43,
	0xc5, 0xa5, 0xfc, 0x59, 0x8a, 0x34, 0x3c, 0xb5, 0x0d, 0xf7, 0x26, 0x63, 0x1a, 0x52, 0xf3, 0xe9,
	0x70, 0x64, 0x1b, 0x2c, 0x19, 0xba, 0x84, 0x99, 0x7d, 0x1f, 0xbf, 0xef, 0x25, 0x04, 0x16, 0x7c,
	0xfc, 0xb6, 0xff, 0x7a, 0xf0, 0xec, 0x4e, 0x20, 0x57, 0x78, 0x97, 0x67, 0x58, 0x28, 0x3b, 0x29,
	0xc2, 0x1f, 0x35, 0x4a, 0x45, 0x5e, 0xc1, 0xd8, 0xf1, 0x6e, 0x26, 0x06, 0x8b, 0x09, 0x73, 0x3d,
	0xae, 0x44, 0x96, 0x30, 0xb7, 0xef, 0xb5, 0x6c, 0x15, 0xe8, 0x2f, 0x05, 0x8b, 0x39, 0x7b, 0xa0,
	0x2c, 0x9a, 0xa5, 0x0f, 0xa4, 0xbe, 0x81, 0x33, 0x07, 0xae, 0x76, 0x89, 0xde, 0x70, 0xb0, 0x38,
	0x63, 0x1d, 0x55, 0x51, 0x90, 0xee, 0x03, 0xfa, 0xcf, 0x83, 0xab, 0x1e, 0x69, 0x67, 0x37, 0x4b,
	0xfb, 0x35, 0x40, 0xea, 0x92, 0x86, 0x78, 0xc0, 0x3a, 0x7d, 0x9d, 0x32, 0xa1, 0xe0, 0xf7, 0x39,
	0x8f, 0x99, 0xe5, 0x6a, 0x0b, 0xe4, 0x2d, 0x4c, 0xf7, 0x88, 0x0e, 0xcb, 0x29, 0xeb, 0x79, 0x3d,
	0x3a, 0x4f, 0x7b, 0xd6, 0xa7, 0xe0, 0x67, 0xad, 0x99, 0xb5, 0x2d, 0x9a, 0xd1, 0xc6, 0xdc, 0x91,
	0x2d, 0xd0, 0xf7, 0x30, 0x6d, 0x65, 0x44, 0x28, 0xab, 0xb2, 0x90, 0xed, 0xb9, 0xea, 0x24, 0x69,
	0x50, 0xed, 0x6f, 0xd7, 0x86, 0x8d, 0x03, 0x50, 0x88, 0x52, 0x58, 0x07, 0xe8, 0x60, 0xf1, 0xc7,
	0x6b, 0x46, 0xd4, 0x52, 0xa1, 0xf8, 0xcc, 0x0b, 0xbe, 0x41, 0x41, 0x3e, 0xc0, 0xc5, 0xd0, 0x59,
	0xc9, 0x73, 0x76, 0xe0, 0xda, 0x97, 0x33, 0xd6, 0x67, 0x42, 0x1f, 0x91, 0x4f, 0xf0, 0x74, 0x78,
	0xd5, 0xe4, 0x05, 0x3b, 0x78, 0x83, 0x81, 0x61, 0xf1, 0x13, 0xfd, 0x57, 0x75, 0xfb, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0x03, 0x29, 0x76, 0x2b, 0xbe, 0x04, 0x00, 0x00,
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
	CreateClientDatabase(ctx context.Context, in *CreateClientDatabaseRequest, opts ...grpc.CallOption) (*ClientResponse, error)
	CreateClientDeployment(ctx context.Context, in *CreateClientDeploymentRequest, opts ...grpc.CallOption) (*ClientResponse, error)
}

type clusterManagerClient struct {
	cc *grpc.ClientConn
}

func NewClusterManagerClient(cc *grpc.ClientConn) ClusterManagerClient {
	return &clusterManagerClient{cc}
}

func (c *clusterManagerClient) CreateClientDatabase(ctx context.Context, in *CreateClientDatabaseRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/ClusterManager/CreateClientDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterManagerClient) CreateClientDeployment(ctx context.Context, in *CreateClientDeploymentRequest, opts ...grpc.CallOption) (*ClientResponse, error) {
	out := new(ClientResponse)
	err := c.cc.Invoke(ctx, "/ClusterManager/CreateClientDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterManagerServer is the server API for ClusterManager service.
type ClusterManagerServer interface {
	CreateClientDatabase(context.Context, *CreateClientDatabaseRequest) (*ClientResponse, error)
	CreateClientDeployment(context.Context, *CreateClientDeploymentRequest) (*ClientResponse, error)
}

func RegisterClusterManagerServer(s *grpc.Server, srv ClusterManagerServer) {
	s.RegisterService(&_ClusterManager_serviceDesc, srv)
}

func _ClusterManager_CreateClientDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterManagerServer).CreateClientDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClusterManager/CreateClientDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterManagerServer).CreateClientDatabase(ctx, req.(*CreateClientDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterManager_CreateClientDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterManagerServer).CreateClientDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClusterManager/CreateClientDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterManagerServer).CreateClientDeployment(ctx, req.(*CreateClientDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ClusterManager",
	HandlerType: (*ClusterManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClientDatabase",
			Handler:    _ClusterManager_CreateClientDatabase_Handler,
		},
		{
			MethodName: "CreateClientDeployment",
			Handler:    _ClusterManager_CreateClientDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubernetes.proto",
}
