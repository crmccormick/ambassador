// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/auth/v3/external_auth.proto

package envoy_service_auth_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	rpc "istio.io/gogo-genproto/googleapis/google/rpc"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CheckRequest struct {
	// The request attributes.
	Attributes           *AttributeContext `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a53ffd03e1a7630f, []int{0}
}
func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return m.Size()
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetAttributes() *AttributeContext {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// HTTP attributes for a denied response.
type DeniedHttpResponse struct {
	// This field allows the authorization service to send a HTTP response status
	// code to the downstream client other than 403 (Forbidden).
	Status *v3.HttpStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// This field allows the authorization service to send HTTP response headers
	// to the downstream client.
	Headers []*v31.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	// This field allows the authorization service to send a response body data
	// to the downstream client.
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeniedHttpResponse) Reset()         { *m = DeniedHttpResponse{} }
func (m *DeniedHttpResponse) String() string { return proto.CompactTextString(m) }
func (*DeniedHttpResponse) ProtoMessage()    {}
func (*DeniedHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a53ffd03e1a7630f, []int{1}
}
func (m *DeniedHttpResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeniedHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeniedHttpResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeniedHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeniedHttpResponse.Merge(m, src)
}
func (m *DeniedHttpResponse) XXX_Size() int {
	return m.Size()
}
func (m *DeniedHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeniedHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeniedHttpResponse proto.InternalMessageInfo

func (m *DeniedHttpResponse) GetStatus() *v3.HttpStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DeniedHttpResponse) GetHeaders() []*v31.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DeniedHttpResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// HTTP attributes for an ok response.
type OkHttpResponse struct {
	// HTTP entity headers in addition to the original request headers. This allows the authorization
	// service to append, to add or to override headers from the original request before
	// dispatching it to the upstream. By setting `append` field to `true` in the `HeaderValueOption`,
	// the filter will append the correspondent header value to the matched request header. Note that
	// by Leaving `append` as false, the filter will either add a new header, or override an existing
	// one if there is a match.
	Headers              []*v31.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *OkHttpResponse) Reset()         { *m = OkHttpResponse{} }
func (m *OkHttpResponse) String() string { return proto.CompactTextString(m) }
func (*OkHttpResponse) ProtoMessage()    {}
func (*OkHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a53ffd03e1a7630f, []int{2}
}
func (m *OkHttpResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OkHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OkHttpResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OkHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkHttpResponse.Merge(m, src)
}
func (m *OkHttpResponse) XXX_Size() int {
	return m.Size()
}
func (m *OkHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OkHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OkHttpResponse proto.InternalMessageInfo

func (m *OkHttpResponse) GetHeaders() []*v31.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Intended for gRPC and Network Authorization servers `only`.
type CheckResponse struct {
	// Status `OK` allows the request. Any other status indicates the request should be denied.
	Status *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// An message that contains HTTP response attributes. This message is
	// used when the authorization service needs to send custom responses to the
	// downstream client or, to modify/add request headers being dispatched to the upstream.
	//
	// Types that are valid to be assigned to HttpResponse:
	//	*CheckResponse_DeniedResponse
	//	*CheckResponse_OkResponse
	HttpResponse         isCheckResponse_HttpResponse `protobuf_oneof:"http_response"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a53ffd03e1a7630f, []int{3}
}
func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return m.Size()
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

type isCheckResponse_HttpResponse interface {
	isCheckResponse_HttpResponse()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CheckResponse_DeniedResponse struct {
	DeniedResponse *DeniedHttpResponse `protobuf:"bytes,2,opt,name=denied_response,json=deniedResponse,proto3,oneof" json:"denied_response,omitempty"`
}
type CheckResponse_OkResponse struct {
	OkResponse *OkHttpResponse `protobuf:"bytes,3,opt,name=ok_response,json=okResponse,proto3,oneof" json:"ok_response,omitempty"`
}

func (*CheckResponse_DeniedResponse) isCheckResponse_HttpResponse() {}
func (*CheckResponse_OkResponse) isCheckResponse_HttpResponse()     {}

func (m *CheckResponse) GetHttpResponse() isCheckResponse_HttpResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *CheckResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CheckResponse) GetDeniedResponse() *DeniedHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_DeniedResponse); ok {
		return x.DeniedResponse
	}
	return nil
}

func (m *CheckResponse) GetOkResponse() *OkHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_OkResponse); ok {
		return x.OkResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CheckResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CheckResponse_DeniedResponse)(nil),
		(*CheckResponse_OkResponse)(nil),
	}
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "envoy.service.auth.v3.CheckRequest")
	proto.RegisterType((*DeniedHttpResponse)(nil), "envoy.service.auth.v3.DeniedHttpResponse")
	proto.RegisterType((*OkHttpResponse)(nil), "envoy.service.auth.v3.OkHttpResponse")
	proto.RegisterType((*CheckResponse)(nil), "envoy.service.auth.v3.CheckResponse")
}

func init() {
	proto.RegisterFile("envoy/service/auth/v3/external_auth.proto", fileDescriptor_a53ffd03e1a7630f)
}

var fileDescriptor_a53ffd03e1a7630f = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x8e, 0x12, 0x4f,
	0x18, 0xdc, 0x86, 0xdf, 0xee, 0x4f, 0x7b, 0x65, 0x57, 0x27, 0x31, 0x22, 0x07, 0x44, 0x60, 0x15,
	0x30, 0xce, 0x24, 0x70, 0xc3, 0x13, 0xac, 0x46, 0x3c, 0xed, 0x66, 0xdc, 0x78, 0x25, 0xcd, 0xcc,
	0x27, 0x4c, 0x20, 0xdd, 0x63, 0x77, 0xcf, 0x04, 0x3c, 0x19, 0x4f, 0x1b, 0x9f, 0xc0, 0xf8, 0x3e,
	0x26, 0x1e, 0xbd, 0x7a, 0x33, 0x3c, 0x86, 0x27, 0xd3, 0x7f, 0x40, 0x70, 0x87, 0xbd, 0x78, 0x9b,
	0xd0, 0x55, 0xf5, 0x55, 0x55, 0xf7, 0x07, 0x6e, 0x02, 0x4d, 0xd9, 0xc2, 0x13, 0xc0, 0xd3, 0x28,
	0x00, 0x8f, 0x24, 0x72, 0xe2, 0xa5, 0x1d, 0x0f, 0xe6, 0x12, 0x38, 0x25, 0xb3, 0xa1, 0xfa, 0xc1,
	0x8d, 0x39, 0x93, 0xcc, 0xb9, 0xab, 0xa1, 0xae, 0x85, 0xba, 0xfa, 0x24, 0xed, 0x94, 0x1e, 0x18,
	0x85, 0x80, 0xd1, 0xb7, 0xd1, 0xd8, 0x0b, 0x18, 0x07, 0x25, 0x30, 0x22, 0x02, 0x0c, 0xaf, 0xf4,
	0x34, 0x7b, 0x04, 0x91, 0x92, 0x47, 0xa3, 0x44, 0xc2, 0x30, 0x60, 0x54, 0xc2, 0x5c, 0x5a, 0xb8,
	0xd5, 0x93, 0x8b, 0x58, 0x0b, 0x4d, 0xa4, 0x8c, 0x87, 0x42, 0x12, 0x99, 0x08, 0x0b, 0xb8, 0x37,
	0x66, 0x6c, 0x3c, 0x03, 0x8f, 0xc7, 0x81, 0xb7, 0x75, 0xf0, 0x30, 0x09, 0x63, 0xe2, 0x11, 0x4a,
	0x99, 0x24, 0x32, 0x62, 0x54, 0x78, 0x29, 0x70, 0x11, 0x31, 0x1a, 0xd1, 0xf1, 0x8a, 0x9b, 0x92,
	0x59, 0x14, 0x12, 0x09, 0xde, 0xea, 0xc3, 0x1c, 0x54, 0x3f, 0x22, 0x7c, 0xeb, 0x74, 0x02, 0xc1,
	0xd4, 0x87, 0x77, 0x09, 0x08, 0xe9, 0xbc, 0xc4, 0x78, 0xed, 0x50, 0x14, 0x51, 0x05, 0x35, 0x0e,
	0xdb, 0x8f, 0xdd, 0xcc, 0x0a, 0xdc, 0xde, 0x0a, 0x78, 0x6a, 0x92, 0xf8, 0x1b, 0xd4, 0x6e, 0xf3,
	0xcb, 0xd7, 0xcb, 0x72, 0x1d, 0x57, 0xb3, 0xa8, 0x6d, 0x77, 0x73, 0x66, 0xf5, 0x07, 0xc2, 0xce,
	0x73, 0xa0, 0x11, 0x84, 0x03, 0x29, 0x63, 0x1f, 0x44, 0xcc, 0xa8, 0x00, 0xe7, 0x19, 0x3e, 0x30,
	0x39, 0xad, 0x8d, 0xfb, 0xd6, 0x86, 0xaa, 0x48, 0x8d, 0x57, 0xe0, 0xd7, 0x1a, 0xd0, 0xbf, 0xf1,
	0xab, 0xbf, 0xff, 0x09, 0xe5, 0x6e, 0x23, 0xdf, 0x52, 0x9c, 0x1e, 0xfe, 0x7f, 0x02, 0x24, 0x04,
	0x2e, 0x8a, 0xb9, 0x4a, 0x7e, 0x23, 0x84, 0xb9, 0x30, 0x57, 0x5d, 0x98, 0x16, 0xd1, 0xa0, 0x37,
	0x64, 0x96, 0xc0, 0x59, 0xac, 0xda, 0xf3, 0x57, 0x3c, 0xc7, 0xc1, 0xff, 0x8d, 0x58, 0xb8, 0x28,
	0xe6, 0x2b, 0xa8, 0x71, 0xd3, 0xd7, 0xdf, 0x5d, 0x4f, 0xa5, 0x6a, 0xe1, 0x46, 0x76, 0xaa, 0xab,
	0x21, 0xaa, 0x1f, 0x10, 0x3e, 0x3a, 0x9b, 0x6e, 0xe5, 0xfa, 0x77, 0x6b, 0xdd, 0x27, 0xca, 0xc6,
	0x23, 0x5c, 0xcf, 0xb6, 0xb1, 0x3d, 0xaf, 0xfa, 0x39, 0x87, 0x0b, 0xb6, 0x6f, 0xeb, 0xa0, 0xf5,
	0x57, 0xb3, 0x8e, 0x6b, 0xde, 0x96, 0xcb, 0xe3, 0xc0, 0x35, 0x95, 0xae, 0x8b, 0xbc, 0xc0, 0xc7,
	0xa1, 0x8e, 0x35, 0xe4, 0x96, 0x5e, 0xcc, 0x69, 0x52, 0x73, 0xc7, 0xab, 0xb8, 0x5a, 0xc2, 0x60,
	0xcf, 0x3f, 0x32, 0x1a, 0x6b, 0x07, 0x03, 0x7c, 0xc8, 0xa6, 0x7f, 0x14, 0xf3, 0x5a, 0xf1, 0x64,
	0x87, 0xe2, 0x76, 0x9e, 0xc1, 0x9e, 0x8f, 0xd9, 0x3a, 0x4b, 0xb7, 0xa5, 0xaa, 0x38, 0xc1, 0xb5,
	0x6b, 0xdf, 0x99, 0xc1, 0xf6, 0x8f, 0x71, 0x41, 0xef, 0xd5, 0x6a, 0x6e, 0x1b, 0x70, 0xa1, 0x97,
	0xc8, 0x09, 0xe3, 0xd1, 0x7b, 0xbd, 0x3a, 0xce, 0x05, 0xde, 0xd7, 0x14, 0xa7, 0xb6, 0xc3, 0xcb,
	0xe6, 0xc3, 0x2d, 0xd5, 0xaf, 0x07, 0xd9, 0xfe, 0xf7, 0xfa, 0xaf, 0xbe, 0x2d, 0xcb, 0xe8, 0xfb,
	0xb2, 0x8c, 0x7e, 0x2e, 0xcb, 0x08, 0xd7, 0x22, 0x66, 0x78, 0x31, 0x67, 0xf3, 0x45, 0xb6, 0x44,
	0xff, 0xce, 0x0b, 0xfb, 0x57, 0xa4, 0xfc, 0x9d, 0xab, 0x5d, 0x3d, 0x47, 0x97, 0x08, 0x8d, 0x0e,
	0xf4, 0xde, 0x76, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x43, 0x8e, 0x91, 0x2e, 0xc1, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.auth.v3.Authorization/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

// UnimplementedAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (*UnimplementedAuthorizationServer) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v3.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v3.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v3/external_auth.proto",
}

func (m *CheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Attributes != nil {
		{
			size, err := m.Attributes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExternalAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeniedHttpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeniedHttpResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeniedHttpResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintExternalAuth(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Headers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExternalAuth(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Status != nil {
		{
			size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExternalAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OkHttpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OkHttpResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OkHttpResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Headers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExternalAuth(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *CheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.HttpResponse != nil {
		{
			size := m.HttpResponse.Size()
			i -= size
			if _, err := m.HttpResponse.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Status != nil {
		{
			size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExternalAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CheckResponse_DeniedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckResponse_DeniedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DeniedResponse != nil {
		{
			size, err := m.DeniedResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExternalAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *CheckResponse_OkResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckResponse_OkResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OkResponse != nil {
		{
			size, err := m.OkResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExternalAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func encodeVarintExternalAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovExternalAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CheckRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Attributes != nil {
		l = m.Attributes.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeniedHttpResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovExternalAuth(uint64(l))
		}
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OkHttpResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovExternalAuth(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CheckResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	if m.HttpResponse != nil {
		n += m.HttpResponse.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CheckResponse_DeniedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeniedResponse != nil {
		l = m.DeniedResponse.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}
func (m *CheckResponse_OkResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OkResponse != nil {
		l = m.OkResponse.Size()
		n += 1 + l + sovExternalAuth(uint64(l))
	}
	return n
}

func sovExternalAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExternalAuth(x uint64) (n int) {
	return sovExternalAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CheckRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attributes == nil {
				m.Attributes = &AttributeContext{}
			}
			if err := m.Attributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeniedHttpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeniedHttpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeniedHttpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &v3.HttpStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &v31.HeaderValueOption{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OkHttpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OkHttpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OkHttpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &v31.HeaderValueOption{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &rpc.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeniedResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeniedHttpResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.HttpResponse = &CheckResponse_DeniedResponse{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OkResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExternalAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OkHttpResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.HttpResponse = &CheckResponse_OkResponse{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternalAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExternalAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExternalAuth
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExternalAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthExternalAuth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExternalAuth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExternalAuth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExternalAuth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExternalAuth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExternalAuth = fmt.Errorf("proto: unexpected end of group")
)
