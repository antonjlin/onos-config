// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/northbound/proto/diags.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ChangesRequest is a message for specifying GetChanges query parameters.
type ChangesRequest struct {
	ChangeIds            []string `protobuf:"bytes,1,rep,name=changeIds,proto3" json:"changeIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangesRequest) Reset()         { *m = ChangesRequest{} }
func (m *ChangesRequest) String() string { return proto.CompactTextString(m) }
func (*ChangesRequest) ProtoMessage()    {}
func (*ChangesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{0}
}

func (m *ChangesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangesRequest.Unmarshal(m, b)
}
func (m *ChangesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangesRequest.Marshal(b, m, deterministic)
}
func (m *ChangesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangesRequest.Merge(m, src)
}
func (m *ChangesRequest) XXX_Size() int {
	return xxx_messageInfo_ChangesRequest.Size(m)
}
func (m *ChangesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangesRequest proto.InternalMessageInfo

func (m *ChangesRequest) GetChangeIds() []string {
	if m != nil {
		return m.ChangeIds
	}
	return nil
}

// ConfigRequest is a message for specifying GetConfigurations query parameters.
type ConfigRequest struct {
	DeviceIds            []string `protobuf:"bytes,1,rep,name=deviceIds,proto3" json:"deviceIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{1}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetDeviceIds() []string {
	if m != nil {
		return m.DeviceIds
	}
	return nil
}

// ChangeValue is an individual Path/Value combination in a Change
type ChangeValue struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Removed              bool     `protobuf:"varint,3,opt,name=removed,proto3" json:"removed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeValue) Reset()         { *m = ChangeValue{} }
func (m *ChangeValue) String() string { return proto.CompactTextString(m) }
func (*ChangeValue) ProtoMessage()    {}
func (*ChangeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{2}
}

func (m *ChangeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeValue.Unmarshal(m, b)
}
func (m *ChangeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeValue.Marshal(b, m, deterministic)
}
func (m *ChangeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeValue.Merge(m, src)
}
func (m *ChangeValue) XXX_Size() int {
	return xxx_messageInfo_ChangeValue.Size(m)
}
func (m *ChangeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeValue.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeValue proto.InternalMessageInfo

func (m *ChangeValue) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChangeValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ChangeValue) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

// Change is a descriptor of a submitted configuration change targeted at a single device.
type Change struct {
	Time                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Desc                 string               `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Changevalues         []*ChangeValue       `protobuf:"bytes,4,rep,name=changevalues,proto3" json:"changevalues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{3}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Change) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Change) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Change) GetChangevalues() []*ChangeValue {
	if m != nil {
		return m.Changevalues
	}
	return nil
}

// Change is a descriptor of a submitted configuration change targeted at a single device.
type Configuration struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Deviceid             string               `protobuf:"bytes,2,opt,name=deviceid,proto3" json:"deviceid,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	User                 string               `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Desc                 string               `protobuf:"bytes,6,opt,name=desc,proto3" json:"desc,omitempty"`
	ChangeIDs            []string             `protobuf:"bytes,7,rep,name=changeIDs,proto3" json:"changeIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7617167392ea9d7, []int{4}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Configuration) GetDeviceid() string {
	if m != nil {
		return m.Deviceid
	}
	return ""
}

func (m *Configuration) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Configuration) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *Configuration) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Configuration) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Configuration) GetChangeIDs() []string {
	if m != nil {
		return m.ChangeIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangesRequest)(nil), "proto.ChangesRequest")
	proto.RegisterType((*ConfigRequest)(nil), "proto.ConfigRequest")
	proto.RegisterType((*ChangeValue)(nil), "proto.ChangeValue")
	proto.RegisterType((*Change)(nil), "proto.Change")
	proto.RegisterType((*Configuration)(nil), "proto.Configuration")
}

func init() { proto.RegisterFile("pkg/northbound/proto/diags.proto", fileDescriptor_c7617167392ea9d7) }

var fileDescriptor_c7617167392ea9d7 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x8e, 0x9b, 0x30,
	0x10, 0x95, 0x13, 0x92, 0x2c, 0x43, 0x77, 0xa5, 0x8e, 0xb6, 0x92, 0x85, 0x2a, 0x15, 0x71, 0xe2,
	0x52, 0x58, 0xa5, 0xab, 0xde, 0xab, 0x8d, 0x54, 0xf5, 0x58, 0xab, 0xea, 0xdd, 0xc1, 0x5e, 0x82,
	0xba, 0xc1, 0x14, 0x9b, 0x7c, 0x42, 0xbf, 0xa0, 0x3f, 0xda, 0x3f, 0xa8, 0x6c, 0x03, 0x81, 0x4b,
	0x7b, 0x62, 0xe6, 0xf1, 0x66, 0x3c, 0xef, 0xe9, 0x41, 0xd2, 0xfe, 0xa8, 0x8a, 0x46, 0x75, 0xe6,
	0x74, 0x54, 0x7d, 0x23, 0x8a, 0xb6, 0x53, 0x46, 0x15, 0xa2, 0xe6, 0x95, 0xce, 0x5d, 0x8d, 0x1b,
	0xf7, 0x89, 0xdf, 0x55, 0x4a, 0x55, 0x2f, 0xd2, 0x13, 0x8e, 0xfd, 0x73, 0x61, 0xea, 0xb3, 0xd4,
	0x86, 0x9f, 0x5b, 0xcf, 0x4b, 0x73, 0xb8, 0x7b, 0x3a, 0xf1, 0xa6, 0x92, 0x9a, 0xc9, 0x9f, 0xbd,
	0xd4, 0x06, 0xdf, 0x42, 0x58, 0x3a, 0xe4, 0x8b, 0xd0, 0x94, 0x24, 0xeb, 0x2c, 0x64, 0x57, 0x20,
	0x7d, 0x0f, 0xb7, 0x4f, 0xaa, 0x79, 0xae, 0xab, 0x19, 0x5d, 0xc8, 0x4b, 0x5d, 0xce, 0xe9, 0x13,
	0x90, 0x7e, 0x85, 0xc8, 0xaf, 0xff, 0xce, 0x5f, 0x7a, 0x89, 0x08, 0x41, 0xcb, 0xcd, 0x89, 0x92,
	0x84, 0x64, 0x21, 0x73, 0x35, 0xde, 0xc3, 0xe6, 0x62, 0x7f, 0xd2, 0x95, 0x03, 0x7d, 0x83, 0x14,
	0x76, 0x9d, 0x3c, 0xab, 0x8b, 0x14, 0x74, 0x9d, 0x90, 0xec, 0x86, 0x8d, 0x6d, 0xfa, 0x9b, 0xc0,
	0xd6, 0xef, 0xc4, 0x1c, 0x02, 0xab, 0xc7, 0xad, 0x8b, 0xf6, 0x71, 0xee, 0xc5, 0xe6, 0xa3, 0xd8,
	0xfc, 0xdb, 0x28, 0x96, 0x39, 0x1e, 0xde, 0xc1, 0xaa, 0x16, 0xc3, 0x3b, 0xab, 0x5a, 0xd8, 0x73,
	0x84, 0xd4, 0xa5, 0x7b, 0x21, 0x64, 0xae, 0xc6, 0x8f, 0xf0, 0xca, 0xab, 0x75, 0x77, 0x68, 0x1a,
	0x24, 0xeb, 0x2c, 0xda, 0xa3, 0x5f, 0x9a, 0xcf, 0xc4, 0xb0, 0x05, 0x2f, 0xfd, 0x43, 0x46, 0x67,
	0xfa, 0x8e, 0x9b, 0x5a, 0x35, 0x76, 0x7b, 0xc3, 0x87, 0xeb, 0x42, 0xe6, 0x6a, 0x8c, 0xe1, 0xc6,
	0x9b, 0x33, 0xdd, 0x31, 0xf5, 0xf8, 0x08, 0xbb, 0xb2, 0x93, 0xdc, 0x0c, 0x92, 0xff, 0x2d, 0x68,
	0xa4, 0xda, 0xa9, 0xbe, 0x15, 0x6e, 0x2a, 0xf8, 0xff, 0xd4, 0x40, 0xb5, 0xb7, 0xf5, 0x5a, 0x76,
	0x74, 0xe3, 0x6f, 0xb3, 0xf5, 0xe4, 0xc6, 0x76, 0xe6, 0xc6, 0x35, 0x0c, 0x07, 0x4d, 0x77, 0x8b,
	0x30, 0x1c, 0xf4, 0xfe, 0x17, 0x81, 0xc8, 0x6b, 0x3e, 0xd8, 0xe8, 0xe1, 0x23, 0xc0, 0x67, 0x69,
	0x86, 0x3c, 0xe1, 0x9b, 0x85, 0x67, 0x63, 0xbe, 0xe2, 0xdb, 0x05, 0xfc, 0x40, 0xf0, 0x13, 0xbc,
	0xb6, 0x53, 0x73, 0xef, 0x34, 0xde, 0x8f, 0xac, 0x79, 0xd8, 0xe2, 0x25, 0x3a, 0x90, 0x1f, 0xc8,
	0x71, 0xeb, 0xe0, 0x0f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x53, 0x00, 0x6d, 0x30, 0x18, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigDiagsClient is the client API for ConfigDiags service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigDiagsClient interface {
	// GetChanges returns a stream of submitted changes objects.
	GetChanges(ctx context.Context, in *ChangesRequest, opts ...grpc.CallOption) (ConfigDiags_GetChangesClient, error)
	// GetConfigurations returns a stream of submitted configurations aimed at individual devices.
	GetConfigurations(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (ConfigDiags_GetConfigurationsClient, error)
}

type configDiagsClient struct {
	cc *grpc.ClientConn
}

func NewConfigDiagsClient(cc *grpc.ClientConn) ConfigDiagsClient {
	return &configDiagsClient{cc}
}

func (c *configDiagsClient) GetChanges(ctx context.Context, in *ChangesRequest, opts ...grpc.CallOption) (ConfigDiags_GetChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConfigDiags_serviceDesc.Streams[0], "/proto.ConfigDiags/GetChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &configDiagsGetChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfigDiags_GetChangesClient interface {
	Recv() (*Change, error)
	grpc.ClientStream
}

type configDiagsGetChangesClient struct {
	grpc.ClientStream
}

func (x *configDiagsGetChangesClient) Recv() (*Change, error) {
	m := new(Change)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *configDiagsClient) GetConfigurations(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (ConfigDiags_GetConfigurationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConfigDiags_serviceDesc.Streams[1], "/proto.ConfigDiags/GetConfigurations", opts...)
	if err != nil {
		return nil, err
	}
	x := &configDiagsGetConfigurationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfigDiags_GetConfigurationsClient interface {
	Recv() (*Configuration, error)
	grpc.ClientStream
}

type configDiagsGetConfigurationsClient struct {
	grpc.ClientStream
}

func (x *configDiagsGetConfigurationsClient) Recv() (*Configuration, error) {
	m := new(Configuration)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConfigDiagsServer is the server API for ConfigDiags service.
type ConfigDiagsServer interface {
	// GetChanges returns a stream of submitted changes objects.
	GetChanges(*ChangesRequest, ConfigDiags_GetChangesServer) error
	// GetConfigurations returns a stream of submitted configurations aimed at individual devices.
	GetConfigurations(*ConfigRequest, ConfigDiags_GetConfigurationsServer) error
}

// UnimplementedConfigDiagsServer can be embedded to have forward compatible implementations.
type UnimplementedConfigDiagsServer struct {
}

func (*UnimplementedConfigDiagsServer) GetChanges(req *ChangesRequest, srv ConfigDiags_GetChangesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChanges not implemented")
}
func (*UnimplementedConfigDiagsServer) GetConfigurations(req *ConfigRequest, srv ConfigDiags_GetConfigurationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConfigurations not implemented")
}

func RegisterConfigDiagsServer(s *grpc.Server, srv ConfigDiagsServer) {
	s.RegisterService(&_ConfigDiags_serviceDesc, srv)
}

func _ConfigDiags_GetChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChangesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigDiagsServer).GetChanges(m, &configDiagsGetChangesServer{stream})
}

type ConfigDiags_GetChangesServer interface {
	Send(*Change) error
	grpc.ServerStream
}

type configDiagsGetChangesServer struct {
	grpc.ServerStream
}

func (x *configDiagsGetChangesServer) Send(m *Change) error {
	return x.ServerStream.SendMsg(m)
}

func _ConfigDiags_GetConfigurations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfigRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigDiagsServer).GetConfigurations(m, &configDiagsGetConfigurationsServer{stream})
}

type ConfigDiags_GetConfigurationsServer interface {
	Send(*Configuration) error
	grpc.ServerStream
}

type configDiagsGetConfigurationsServer struct {
	grpc.ServerStream
}

func (x *configDiagsGetConfigurationsServer) Send(m *Configuration) error {
	return x.ServerStream.SendMsg(m)
}

var _ConfigDiags_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ConfigDiags",
	HandlerType: (*ConfigDiagsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetChanges",
			Handler:       _ConfigDiags_GetChanges_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetConfigurations",
			Handler:       _ConfigDiags_GetConfigurations_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/northbound/proto/diags.proto",
}
