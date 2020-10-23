// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alert/alert.proto

package alert

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

// Event is inspired by Google Analytics events
// https://developers.google.com/analytics/devguides/collection/analyticsjs/events
type Event struct {
	// id is not required for inserts
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Category             string            `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Action               string            `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Label                string            `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Value                uint64            `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_975149e03a08f257, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Event) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Event) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Event) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Event) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ReportEventRequest struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportEventRequest) Reset()         { *m = ReportEventRequest{} }
func (m *ReportEventRequest) String() string { return proto.CompactTextString(m) }
func (*ReportEventRequest) ProtoMessage()    {}
func (*ReportEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_975149e03a08f257, []int{1}
}

func (m *ReportEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportEventRequest.Unmarshal(m, b)
}
func (m *ReportEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportEventRequest.Marshal(b, m, deterministic)
}
func (m *ReportEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportEventRequest.Merge(m, src)
}
func (m *ReportEventRequest) XXX_Size() int {
	return xxx_messageInfo_ReportEventRequest.Size(m)
}
func (m *ReportEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportEventRequest proto.InternalMessageInfo

func (m *ReportEventRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type ReportEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportEventResponse) Reset()         { *m = ReportEventResponse{} }
func (m *ReportEventResponse) String() string { return proto.CompactTextString(m) }
func (*ReportEventResponse) ProtoMessage()    {}
func (*ReportEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_975149e03a08f257, []int{2}
}

func (m *ReportEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportEventResponse.Unmarshal(m, b)
}
func (m *ReportEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportEventResponse.Marshal(b, m, deterministic)
}
func (m *ReportEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportEventResponse.Merge(m, src)
}
func (m *ReportEventResponse) XXX_Size() int {
	return xxx_messageInfo_ReportEventResponse.Size(m)
}
func (m *ReportEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportEventResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Event)(nil), "alert.Event")
	proto.RegisterMapType((map[string]string)(nil), "alert.Event.MetadataEntry")
	proto.RegisterType((*ReportEventRequest)(nil), "alert.ReportEventRequest")
	proto.RegisterType((*ReportEventResponse)(nil), "alert.ReportEventResponse")
}

func init() { proto.RegisterFile("alert/alert.proto", fileDescriptor_975149e03a08f257) }

var fileDescriptor_975149e03a08f257 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0xda, 0x2d, 0x75, 0xaa, 0xa2, 0xeb, 0x1f, 0xd6, 0x9c, 0x4a, 0x4e, 0x05, 0x31,
	0x81, 0x14, 0xa4, 0xd8, 0x93, 0x42, 0xbd, 0x89, 0xb0, 0x47, 0x6f, 0x9b, 0x64, 0xa8, 0xc1, 0x24,
	0x1b, 0x37, 0x9b, 0x40, 0xbe, 0xb3, 0x1f, 0x42, 0xb2, 0x1b, 0x6a, 0x8b, 0x5e, 0x86, 0xfd, 0xbd,
	0x37, 0x3b, 0xf3, 0x60, 0xe0, 0x42, 0xe4, 0xa8, 0x74, 0x68, 0x6a, 0x50, 0x29, 0xa9, 0x25, 0x25,
	0x06, 0xfc, 0x6f, 0x07, 0xc8, 0xa6, 0xc5, 0x52, 0xd3, 0x33, 0x70, 0xb3, 0x94, 0x39, 0x73, 0x67,
	0x71, 0xcc, 0xdd, 0x2c, 0xa5, 0x1e, 0x4c, 0x13, 0xa1, 0x71, 0x2b, 0x55, 0xc7, 0x5c, 0xa3, 0xee,
	0x98, 0xde, 0xc0, 0x44, 0x24, 0x3a, 0x93, 0x25, 0x1b, 0x19, 0x67, 0x20, 0x7a, 0x05, 0x24, 0x17,
	0x31, 0xe6, 0x6c, 0x6c, 0x64, 0x0b, 0xbd, 0xda, 0x8a, 0xbc, 0x41, 0x46, 0xe6, 0xce, 0x62, 0xcc,
	0x2d, 0xd0, 0x07, 0x98, 0x16, 0xa8, 0x45, 0x2a, 0xb4, 0x60, 0x93, 0xf9, 0x68, 0x31, 0x8b, 0xbc,
	0xc0, 0x06, 0x34, 0x79, 0x82, 0xd7, 0xc1, 0xdc, 0x94, 0x5a, 0x75, 0x7c, 0xd7, 0xeb, 0xad, 0xe1,
	0xf4, 0xc0, 0xa2, 0xe7, 0x30, 0xfa, 0xc4, 0x6e, 0x48, 0xde, 0x3f, 0x7f, 0x17, 0xda, 0xdc, 0x16,
	0x1e, 0xdd, 0x95, 0xe3, 0xaf, 0x80, 0x72, 0xac, 0xa4, 0xd2, 0x66, 0x07, 0xc7, 0xaf, 0x06, 0x6b,
	0x4d, 0x7d, 0x20, 0xd8, 0xb3, 0x99, 0x31, 0x8b, 0x4e, 0xf6, 0x73, 0x70, 0x6b, 0xf9, 0xd7, 0x70,
	0x79, 0xf0, 0xb3, 0xae, 0x64, 0x59, 0x63, 0xf4, 0x06, 0xe4, 0xa9, 0x6f, 0xa6, 0x2f, 0x30, 0xdb,
	0xf3, 0xe9, 0xed, 0x30, 0xe3, 0xef, 0x36, 0xcf, 0xfb, 0xcf, 0xb2, 0xe3, 0xfc, 0xa3, 0xe7, 0xfb,
	0xf7, 0xbb, 0x6d, 0xa6, 0x3f, 0x9a, 0x38, 0x48, 0x64, 0x11, 0x16, 0x59, 0xa2, 0xe4, 0x50, 0xdb,
	0x65, 0x68, 0x4e, 0x67, 0xcf, 0xb8, 0x36, 0x35, 0x9e, 0x18, 0x69, 0xf9, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xab, 0x72, 0x88, 0x2f, 0xe2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlertClient is the client API for Alert service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlertClient interface {
	// ReportEvent does event ingestions.
	ReportEvent(ctx context.Context, in *ReportEventRequest, opts ...grpc.CallOption) (*ReportEventResponse, error)
}

type alertClient struct {
	cc *grpc.ClientConn
}

func NewAlertClient(cc *grpc.ClientConn) AlertClient {
	return &alertClient{cc}
}

func (c *alertClient) ReportEvent(ctx context.Context, in *ReportEventRequest, opts ...grpc.CallOption) (*ReportEventResponse, error) {
	out := new(ReportEventResponse)
	err := c.cc.Invoke(ctx, "/alert.Alert/ReportEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertServer is the server API for Alert service.
type AlertServer interface {
	// ReportEvent does event ingestions.
	ReportEvent(context.Context, *ReportEventRequest) (*ReportEventResponse, error)
}

func RegisterAlertServer(s *grpc.Server, srv AlertServer) {
	s.RegisterService(&_Alert_serviceDesc, srv)
}

func _Alert_ReportEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServer).ReportEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alert/ReportEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServer).ReportEvent(ctx, req.(*ReportEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alert_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alert.Alert",
	HandlerType: (*AlertServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportEvent",
			Handler:    _Alert_ReportEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alert/alert.proto",
}
