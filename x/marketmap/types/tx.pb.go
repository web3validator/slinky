// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/marketmap/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
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

// MsgCreateMarket defines a message carrying a payload for creating a new
// market in the x/marketmap module.
type MsgCreateMarket struct {
	// signer is the signer of this transaction (notice, this may not always be a
	// node from the SecondTier)
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// Ticker is the on-chain representation of the ticker. This is the target
	// ticker that the prices of the set of tickers will be converted to.
	Ticker Ticker `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker"`
	// ProvidersToOffChainTickers maps provider names to their off-chain
	// representations for the given ticker of the message.  This is used to
	// construct the updates MarketConfig resulting from this message execution.
	ProvidersToOffChainTickers map[string]string `protobuf:"bytes,3,rep,name=providers_to_off_chain_tickers,json=providersToOffChainTickers,proto3" json:"providers_to_off_chain_tickers" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paths is the list of convertable markets that will be used to convert the
	// prices of a set of tickers to a common ticker.
	Paths []Path `protobuf:"bytes,4,rep,name=paths,proto3" json:"paths"`
}

func (m *MsgCreateMarket) Reset()         { *m = MsgCreateMarket{} }
func (m *MsgCreateMarket) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMarket) ProtoMessage()    {}
func (*MsgCreateMarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9adadfc18297083, []int{0}
}
func (m *MsgCreateMarket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMarket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMarket.Merge(m, src)
}
func (m *MsgCreateMarket) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMarket) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMarket.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMarket proto.InternalMessageInfo

func (m *MsgCreateMarket) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgCreateMarket) GetTicker() Ticker {
	if m != nil {
		return m.Ticker
	}
	return Ticker{}
}

func (m *MsgCreateMarket) GetProvidersToOffChainTickers() map[string]string {
	if m != nil {
		return m.ProvidersToOffChainTickers
	}
	return nil
}

func (m *MsgCreateMarket) GetPaths() []Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

// MsgCreateMarketResponse is the response message for MsgCreateMarket.
type MsgCreateMarketResponse struct {
}

func (m *MsgCreateMarketResponse) Reset()         { *m = MsgCreateMarketResponse{} }
func (m *MsgCreateMarketResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMarketResponse) ProtoMessage()    {}
func (*MsgCreateMarketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9adadfc18297083, []int{1}
}
func (m *MsgCreateMarketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMarketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMarketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMarketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMarketResponse.Merge(m, src)
}
func (m *MsgCreateMarketResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMarketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMarketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMarketResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateMarket)(nil), "slinky.marketmap.v1.MsgCreateMarket")
	proto.RegisterMapType((map[string]string)(nil), "slinky.marketmap.v1.MsgCreateMarket.ProvidersToOffChainTickersEntry")
	proto.RegisterType((*MsgCreateMarketResponse)(nil), "slinky.marketmap.v1.MsgCreateMarketResponse")
}

func init() { proto.RegisterFile("slinky/marketmap/v1/tx.proto", fileDescriptor_e9adadfc18297083) }

var fileDescriptor_e9adadfc18297083 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xed, 0xba, 0x89, 0xd4, 0x2d, 0x12, 0x60, 0x22, 0xd5, 0x31, 0xc8, 0x89, 0x22, 0x0e,
	0x51, 0xd4, 0xda, 0x34, 0x08, 0x04, 0xbd, 0x91, 0x02, 0xb7, 0x88, 0xca, 0xf4, 0xc4, 0xc5, 0x72,
	0x92, 0xcd, 0x66, 0xe5, 0xda, 0xbb, 0xda, 0xd9, 0x5a, 0xcd, 0x0d, 0x71, 0xe4, 0x80, 0x78, 0x04,
	0x1e, 0x21, 0x07, 0x1e, 0xa2, 0x17, 0xa4, 0x8a, 0x13, 0x27, 0x84, 0x92, 0x43, 0x78, 0x0c, 0x14,
	0xef, 0x16, 0x41, 0x64, 0x44, 0x2e, 0xd1, 0xcc, 0xfc, 0x5f, 0xe6, 0xdf, 0x99, 0x31, 0xba, 0x07,
	0x67, 0x34, 0x4b, 0xa6, 0x41, 0x1a, 0x8b, 0x04, 0xcb, 0x34, 0xe6, 0x41, 0x7e, 0x18, 0xc8, 0x0b,
	0x9f, 0x0b, 0x26, 0x99, 0x7d, 0x47, 0xa9, 0xfe, 0x6f, 0xd5, 0xcf, 0x0f, 0xdd, 0xbd, 0x21, 0x83,
	0x94, 0x41, 0x90, 0x02, 0x59, 0xc1, 0x29, 0x10, 0x45, 0xbb, 0x35, 0xc2, 0x08, 0x2b, 0xc2, 0x60,
	0x15, 0xe9, 0x6a, 0x5d, 0xe1, 0x91, 0x12, 0x54, 0xa2, 0xa5, 0xdb, 0x71, 0x4a, 0x33, 0x16, 0x14,
	0xbf, 0xba, 0xd4, 0x2c, 0x7b, 0x8f, 0x4a, 0x14, 0xd1, 0xfa, 0x62, 0xa1, 0x9b, 0x7d, 0x20, 0xc7,
	0x02, 0xc7, 0x12, 0xf7, 0x0b, 0xc5, 0x7e, 0x80, 0xaa, 0x40, 0x49, 0x86, 0x85, 0x63, 0x36, 0xcd,
	0xf6, 0x4e, 0xcf, 0xf9, 0xfa, 0xf9, 0xa0, 0xa6, 0xad, 0x9e, 0x8d, 0x46, 0x02, 0x03, 0xbc, 0x96,
	0x82, 0x66, 0x24, 0xd4, 0x9c, 0xfd, 0x14, 0x55, 0x25, 0x1d, 0x26, 0x58, 0x38, 0x5b, 0x4d, 0xb3,
	0xbd, 0xdb, 0xbd, 0xeb, 0x97, 0x8c, 0xea, 0x9f, 0x16, 0x48, 0x6f, 0xfb, 0xf2, 0x7b, 0xc3, 0x08,
	0xf5, 0x1f, 0xec, 0x0f, 0x26, 0xf2, 0xb8, 0x60, 0x39, 0x1d, 0x61, 0x01, 0x91, 0x64, 0x11, 0x1b,
	0x8f, 0xa3, 0xe1, 0x24, 0xa6, 0x59, 0xa4, 0x08, 0x70, 0xac, 0xa6, 0xd5, 0xde, 0xed, 0x3e, 0x2f,
	0xed, 0xb9, 0xf6, 0x76, 0xff, 0xe4, 0xba, 0xd5, 0x29, 0x7b, 0x35, 0x1e, 0x1f, 0xaf, 0xfa, 0x28,
	0x5b, 0x78, 0x91, 0x49, 0x31, 0xd5, 0xe6, 0x2e, 0xff, 0x27, 0x66, 0x3f, 0x42, 0x15, 0x1e, 0xcb,
	0x09, 0x38, 0xdb, 0x85, 0x6d, 0xbd, 0xd4, 0xf6, 0x24, 0x96, 0x13, 0xdd, 0x4b, 0xd1, 0x6e, 0x1f,
	0x35, 0xfe, 0xe3, 0x6d, 0xdf, 0x42, 0x56, 0x82, 0xa7, 0x6a, 0xa9, 0xe1, 0x2a, 0xb4, 0x6b, 0xa8,
	0x92, 0xc7, 0x67, 0xe7, 0xb8, 0x58, 0xdb, 0x4e, 0xa8, 0x92, 0xa3, 0xad, 0x27, 0xe6, 0xd1, 0xe3,
	0x9f, 0x9f, 0x1a, 0xc6, 0xbb, 0xe5, 0xac, 0xa3, 0x57, 0xfc, 0x7e, 0x39, 0xeb, 0xb4, 0xf4, 0x35,
	0x2f, 0xfe, 0xb8, 0xe7, 0xda, 0xfc, 0xad, 0x3a, 0xda, 0x5b, 0x2b, 0x85, 0x18, 0x38, 0xcb, 0x00,
	0x77, 0x39, 0xb2, 0xfa, 0x40, 0xec, 0x01, 0xba, 0xf1, 0xd7, 0xb5, 0xef, 0x6f, 0xb2, 0x57, 0x77,
	0x7f, 0x13, 0xea, 0xda, 0xca, 0xad, 0xbc, 0x5d, 0xce, 0x3a, 0x66, 0xef, 0xe5, 0xe5, 0xdc, 0x33,
	0xaf, 0xe6, 0x9e, 0xf9, 0x63, 0xee, 0x99, 0x1f, 0x17, 0x9e, 0x71, 0xb5, 0xf0, 0x8c, 0x6f, 0x0b,
	0xcf, 0x78, 0xb3, 0x4f, 0xa8, 0x9c, 0x9c, 0x0f, 0xfc, 0x21, 0x4b, 0x03, 0x48, 0x28, 0x3f, 0x48,
	0x71, 0x1e, 0x94, 0x8c, 0x27, 0xa7, 0x1c, 0xc3, 0xa0, 0x5a, 0x7c, 0xab, 0x0f, 0x7f, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x06, 0x57, 0x94, 0x9f, 0x5f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateMarket creates a market from the given message.
	CreateMarket(ctx context.Context, in *MsgCreateMarket, opts ...grpc.CallOption) (*MsgCreateMarketResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateMarket(ctx context.Context, in *MsgCreateMarket, opts ...grpc.CallOption) (*MsgCreateMarketResponse, error) {
	out := new(MsgCreateMarketResponse)
	err := c.cc.Invoke(ctx, "/slinky.marketmap.v1.Msg/CreateMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateMarket creates a market from the given message.
	CreateMarket(context.Context, *MsgCreateMarket) (*MsgCreateMarketResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateMarket(ctx context.Context, req *MsgCreateMarket) (*MsgCreateMarketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMarket not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMarket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.marketmap.v1.Msg/CreateMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMarket(ctx, req.(*MsgCreateMarket))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.marketmap.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMarket",
			Handler:    _Msg_CreateMarket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/marketmap/v1/tx.proto",
}

func (m *MsgCreateMarket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMarket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMarket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for iNdEx := len(m.Paths) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Paths[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ProvidersToOffChainTickers) > 0 {
		for k := range m.ProvidersToOffChainTickers {
			v := m.ProvidersToOffChainTickers[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintTx(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTx(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTx(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Ticker.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateMarketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMarketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMarketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateMarket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Ticker.Size()
	n += 1 + l + sovTx(uint64(l))
	if len(m.ProvidersToOffChainTickers) > 0 {
		for k, v := range m.ProvidersToOffChainTickers {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTx(uint64(len(k))) + 1 + len(v) + sovTx(uint64(len(v)))
			n += mapEntrySize + 1 + sovTx(uint64(mapEntrySize))
		}
	}
	if len(m.Paths) > 0 {
		for _, e := range m.Paths {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateMarketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateMarket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateMarket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMarket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ticker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvidersToOffChainTickers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProvidersToOffChainTickers == nil {
				m.ProvidersToOffChainTickers = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTx
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTx
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTx
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthTx
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTx(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTx
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ProvidersToOffChainTickers[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Paths = append(m.Paths, Path{})
			if err := m.Paths[len(m.Paths)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateMarketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateMarketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMarketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)