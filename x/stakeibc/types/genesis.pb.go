// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the stakeibc module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// list of zones that are registered by the protocol
	ICAAccount         *ICAAccount     `protobuf:"bytes,4,opt,name=iCAAccount,proto3" json:"iCAAccount,omitempty"`
	HostZoneList       []HostZone      `protobuf:"bytes,5,rep,name=hostZoneList,proto3" json:"hostZoneList"`
	HostZoneCount      uint64          `protobuf:"varint,6,opt,name=hostZoneCount,proto3" json:"hostZoneCount,omitempty"`
	DepositRecordList  []DepositRecord `protobuf:"bytes,7,rep,name=depositRecordList,proto3" json:"depositRecordList"`
	DepositRecordCount uint64          `protobuf:"varint,8,opt,name=depositRecordCount,proto3" json:"depositRecordCount,omitempty"`
	// stores a map from hostZone base denom to hostZone
	DenomToHostZone  map[string]string `protobuf:"bytes,9,rep,name=denomToHostZone,proto3" json:"denomToHostZone,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EpochTrackerList []EpochTracker    `protobuf:"bytes,10,rep,name=epochTrackerList,proto3" json:"epochTrackerList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b132bbaf7441a735, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetICAAccount() *ICAAccount {
	if m != nil {
		return m.ICAAccount
	}
	return nil
}

func (m *GenesisState) GetHostZoneList() []HostZone {
	if m != nil {
		return m.HostZoneList
	}
	return nil
}

func (m *GenesisState) GetHostZoneCount() uint64 {
	if m != nil {
		return m.HostZoneCount
	}
	return 0
}

func (m *GenesisState) GetDepositRecordList() []DepositRecord {
	if m != nil {
		return m.DepositRecordList
	}
	return nil
}

func (m *GenesisState) GetDepositRecordCount() uint64 {
	if m != nil {
		return m.DepositRecordCount
	}
	return 0
}

func (m *GenesisState) GetDenomToHostZone() map[string]string {
	if m != nil {
		return m.DenomToHostZone
	}
	return nil
}

func (m *GenesisState) GetEpochTrackerList() []EpochTracker {
	if m != nil {
		return m.EpochTrackerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "Stridelabs.stride.stakeibc.GenesisState")
	proto.RegisterMapType((map[string]string)(nil), "Stridelabs.stride.stakeibc.GenesisState.DenomToHostZoneEntry")
}

func init() { proto.RegisterFile("stakeibc/genesis.proto", fileDescriptor_b132bbaf7441a735) }

var fileDescriptor_b132bbaf7441a735 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x6d, 0xd6, 0xb4, 0x5b, 0xbd, 0x21, 0x8a, 0x55, 0x20, 0x8a, 0x20, 0x54, 0xd3, 0x84, 0xc2,
	0x01, 0x47, 0x1a, 0x17, 0x84, 0x84, 0xc4, 0xba, 0x0d, 0x36, 0x34, 0x21, 0x94, 0xed, 0x54, 0x09,
	0x45, 0x4e, 0x62, 0xa5, 0x56, 0xd7, 0x38, 0xb2, 0x5d, 0x44, 0xf9, 0x15, 0xfc, 0x29, 0xa4, 0x1d,
	0x77, 0xe4, 0x84, 0x50, 0xfb, 0x47, 0x50, 0x6d, 0x37, 0x6a, 0x58, 0xc9, 0xcd, 0xf6, 0xfb, 0xde,
	0xfb, 0xde, 0xfb, 0xf2, 0x05, 0x3c, 0x12, 0x12, 0x8f, 0x09, 0x8d, 0x93, 0x20, 0x23, 0x39, 0x11,
	0x54, 0xa0, 0x82, 0x33, 0xc9, 0xa0, 0x7b, 0x29, 0x39, 0x4d, 0xc9, 0x35, 0x8e, 0x05, 0x12, 0xea,
	0x88, 0x56, 0x95, 0x6e, 0x2f, 0x63, 0x19, 0x53, 0x65, 0xc1, 0xf2, 0xa4, 0x19, 0xee, 0xc3, 0x52,
	0xa9, 0xc0, 0x1c, 0x4f, 0x8c, 0x90, 0xeb, 0x96, 0xcf, 0x34, 0xc1, 0x11, 0x4e, 0x12, 0x36, 0xcd,
	0xa5, 0xc1, 0x9c, 0x12, 0x1b, 0x31, 0x21, 0xa3, 0xef, 0x2c, 0x27, 0x06, 0x79, 0x5a, 0x22, 0x29,
	0x29, 0x98, 0xa0, 0x32, 0xe2, 0x24, 0x61, 0x3c, 0x35, 0xf0, 0x93, 0x12, 0x26, 0x05, 0x4b, 0x46,
	0x91, 0xe4, 0x38, 0x19, 0x13, 0xae, 0xd1, 0xfd, 0x9f, 0x2d, 0xb0, 0xf7, 0x41, 0xa7, 0xb9, 0x94,
	0x58, 0x12, 0xf8, 0x0e, 0xb4, 0xb5, 0x27, 0xc7, 0xea, 0x5b, 0xfe, 0xee, 0xe1, 0x3e, 0xfa, 0x7f,
	0x3a, 0xf4, 0x59, 0x55, 0x0e, 0xec, 0x9b, 0xdf, 0xcf, 0x1a, 0xa1, 0xe1, 0xc1, 0xc7, 0x60, 0xbb,
	0x60, 0x5c, 0x46, 0x34, 0x75, 0xb6, 0xfa, 0x96, 0xdf, 0x09, 0xdb, 0xcb, 0xeb, 0x79, 0x0a, 0xdf,
	0x03, 0x40, 0x8f, 0x8f, 0x8e, 0x74, 0x2c, 0xc7, 0x56, 0xf2, 0xcf, 0xeb, 0xe4, 0xcf, 0xcb, 0xea,
	0x70, 0x8d, 0x09, 0x3f, 0x81, 0xbd, 0xe5, 0x0c, 0x86, 0x2c, 0x27, 0x17, 0x54, 0x48, 0xa7, 0xd5,
	0x6f, 0xfa, 0xbb, 0x87, 0x07, 0x75, 0x4a, 0x67, 0xa6, 0xde, 0x58, 0xad, 0xf0, 0xe1, 0x01, 0xb8,
	0xb7, 0xba, 0x1f, 0x2b, 0x6b, 0xed, 0xbe, 0xe5, 0xdb, 0x61, 0xf5, 0x11, 0x7e, 0x01, 0x0f, 0xcc,
	0x7c, 0x43, 0x35, 0x5e, 0xd5, 0x7a, 0x5b, 0xb5, 0x7e, 0x51, 0xd7, 0xfa, 0x64, 0x9d, 0x64, 0xfa,
	0xdf, 0x55, 0x82, 0x08, 0xc0, 0xca, 0xa3, 0x76, 0xb2, 0xa3, 0x9c, 0x6c, 0x40, 0x60, 0x06, 0xee,
	0xa7, 0x24, 0x67, 0x93, 0x2b, 0xb6, 0xca, 0xe6, 0x74, 0x94, 0x99, 0xb7, 0x75, 0x66, 0xd6, 0x3f,
	0x35, 0x3a, 0xa9, 0xf2, 0x4f, 0x73, 0xc9, 0x67, 0xe1, 0xbf, 0xaa, 0x70, 0x08, 0xba, 0x6a, 0x71,
	0xae, 0xf4, 0xde, 0xa8, 0xd8, 0x40, 0x75, 0xf2, 0xeb, 0x3a, 0x9d, 0xae, 0x71, 0x4c, 0xea, 0x3b,
	0x3a, 0xee, 0x00, 0xf4, 0x36, 0x99, 0x80, 0x5d, 0xd0, 0x1c, 0x93, 0x99, 0xda, 0xc0, 0x4e, 0xb8,
	0x3c, 0xc2, 0x1e, 0x68, 0x7d, 0xc5, 0xd7, 0x53, 0x62, 0x56, 0x4a, 0x5f, 0xde, 0x6c, 0xbd, 0xb6,
	0x3e, 0xda, 0x3b, 0xcd, 0xae, 0x3d, 0x38, 0xbb, 0x99, 0x7b, 0xd6, 0xed, 0xdc, 0xb3, 0xfe, 0xcc,
	0x3d, 0xeb, 0xc7, 0xc2, 0x6b, 0xdc, 0x2e, 0xbc, 0xc6, 0xaf, 0x85, 0xd7, 0x18, 0xa2, 0x8c, 0xca,
	0xd1, 0x34, 0x46, 0x09, 0x9b, 0x04, 0xda, 0xef, 0xcb, 0x0b, 0x1c, 0x8b, 0x40, 0x1b, 0x0e, 0xbe,
	0x05, 0xe5, 0xff, 0x21, 0x67, 0x05, 0x11, 0x71, 0x5b, 0xfd, 0x18, 0xaf, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x91, 0x68, 0xd9, 0x24, 0xee, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EpochTrackerList) > 0 {
		for iNdEx := len(m.EpochTrackerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EpochTrackerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.DenomToHostZone) > 0 {
		for k := range m.DenomToHostZone {
			v := m.DenomToHostZone[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenesis(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.DepositRecordCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.DepositRecordCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.DepositRecordList) > 0 {
		for iNdEx := len(m.DepositRecordList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DepositRecordList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.HostZoneCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.HostZoneCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.HostZoneList) > 0 {
		for iNdEx := len(m.HostZoneList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HostZoneList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.ICAAccount != nil {
		{
			size, err := m.ICAAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ICAAccount != nil {
		l = m.ICAAccount.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.HostZoneList) > 0 {
		for _, e := range m.HostZoneList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.HostZoneCount != 0 {
		n += 1 + sovGenesis(uint64(m.HostZoneCount))
	}
	if len(m.DepositRecordList) > 0 {
		for _, e := range m.DepositRecordList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.DepositRecordCount != 0 {
		n += 1 + sovGenesis(uint64(m.DepositRecordCount))
	}
	if len(m.DenomToHostZone) > 0 {
		for k, v := range m.DenomToHostZone {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + len(v) + sovGenesis(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	if len(m.EpochTrackerList) > 0 {
		for _, e := range m.EpochTrackerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ICAAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ICAAccount == nil {
				m.ICAAccount = &ICAAccount{}
			}
			if err := m.ICAAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostZoneList = append(m.HostZoneList, HostZone{})
			if err := m.HostZoneList[len(m.HostZoneList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostZoneCount", wireType)
			}
			m.HostZoneCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HostZoneCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositRecordList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositRecordList = append(m.DepositRecordList, DepositRecord{})
			if err := m.DepositRecordList[len(m.DepositRecordList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositRecordCount", wireType)
			}
			m.DepositRecordCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositRecordCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomToHostZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DenomToHostZone == nil {
				m.DenomToHostZone = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
							return ErrIntOverflowGenesis
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
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
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
							return ErrIntOverflowGenesis
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
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.DenomToHostZone[mapkey] = mapvalue
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochTrackerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochTrackerList = append(m.EpochTrackerList, EpochTracker{})
			if err := m.EpochTrackerList[len(m.EpochTrackerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
