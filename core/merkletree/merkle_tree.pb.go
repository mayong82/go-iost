// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/merkletree/merkle_tree.proto

package merkletree

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MerkleTree struct {
	HashList             [][]byte         `protobuf:"bytes,1,rep,name=hash_list,json=hashList" json:"hash_list,omitempty"`
	Hash2Idx             map[string]int32 `protobuf:"bytes,2,rep,name=hash2_idx,json=hash2Idx" json:"hash2_idx,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	LeafNum              int32            `protobuf:"varint,3,opt,name=leaf_num,json=leafNum,proto3" json:"leaf_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MerkleTree) Reset()         { *m = MerkleTree{} }
func (m *MerkleTree) String() string { return proto.CompactTextString(m) }
func (*MerkleTree) ProtoMessage()    {}
func (*MerkleTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_merkle_tree_15dc9b8f1778676f, []int{0}
}
func (m *MerkleTree) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MerkleTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MerkleTree.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *MerkleTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerkleTree.Merge(dst, src)
}
func (m *MerkleTree) XXX_Size() int {
	return m.Size()
}
func (m *MerkleTree) XXX_DiscardUnknown() {
	xxx_messageInfo_MerkleTree.DiscardUnknown(m)
}

var xxx_messageInfo_MerkleTree proto.InternalMessageInfo

func (m *MerkleTree) GetHashList() [][]byte {
	if m != nil {
		return m.HashList
	}
	return nil
}

func (m *MerkleTree) GetHash2Idx() map[string]int32 {
	if m != nil {
		return m.Hash2Idx
	}
	return nil
}

func (m *MerkleTree) GetLeafNum() int32 {
	if m != nil {
		return m.LeafNum
	}
	return 0
}

type TXRMerkleTree struct {
	Mt                   *MerkleTree       `protobuf:"bytes,1,opt,name=mt" json:"mt,omitempty"`
	Tx2Txr               map[string][]byte `protobuf:"bytes,2,rep,name=tx2_txr,json=tx2Txr" json:"tx2_txr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TXRMerkleTree) Reset()         { *m = TXRMerkleTree{} }
func (m *TXRMerkleTree) String() string { return proto.CompactTextString(m) }
func (*TXRMerkleTree) ProtoMessage()    {}
func (*TXRMerkleTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_merkle_tree_15dc9b8f1778676f, []int{1}
}
func (m *TXRMerkleTree) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TXRMerkleTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TXRMerkleTree.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TXRMerkleTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TXRMerkleTree.Merge(dst, src)
}
func (m *TXRMerkleTree) XXX_Size() int {
	return m.Size()
}
func (m *TXRMerkleTree) XXX_DiscardUnknown() {
	xxx_messageInfo_TXRMerkleTree.DiscardUnknown(m)
}

var xxx_messageInfo_TXRMerkleTree proto.InternalMessageInfo

func (m *TXRMerkleTree) GetMt() *MerkleTree {
	if m != nil {
		return m.Mt
	}
	return nil
}

func (m *TXRMerkleTree) GetTx2Txr() map[string][]byte {
	if m != nil {
		return m.Tx2Txr
	}
	return nil
}

func init() {
	proto.RegisterType((*MerkleTree)(nil), "merkletree.MerkleTree")
	proto.RegisterMapType((map[string]int32)(nil), "merkletree.MerkleTree.Hash2IdxEntry")
	proto.RegisterType((*TXRMerkleTree)(nil), "merkletree.TXRMerkleTree")
	proto.RegisterMapType((map[string][]byte)(nil), "merkletree.TXRMerkleTree.Tx2TxrEntry")
}
func (m *MerkleTree) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MerkleTree) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.HashList) > 0 {
		for _, b := range m.HashList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMerkleTree(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.Hash2Idx) > 0 {
		for k, _ := range m.Hash2Idx {
			dAtA[i] = 0x12
			i++
			v := m.Hash2Idx[k]
			mapSize := 1 + len(k) + sovMerkleTree(uint64(len(k))) + 1 + sovMerkleTree(uint64(v))
			i = encodeVarintMerkleTree(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintMerkleTree(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintMerkleTree(dAtA, i, uint64(v))
		}
	}
	if m.LeafNum != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMerkleTree(dAtA, i, uint64(m.LeafNum))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TXRMerkleTree) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TXRMerkleTree) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mt != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMerkleTree(dAtA, i, uint64(m.Mt.Size()))
		n1, err := m.Mt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Tx2Txr) > 0 {
		for k, _ := range m.Tx2Txr {
			dAtA[i] = 0x12
			i++
			v := m.Tx2Txr[k]
			byteSize := 0
			if len(v) > 0 {
				byteSize = 1 + len(v) + sovMerkleTree(uint64(len(v)))
			}
			mapSize := 1 + len(k) + sovMerkleTree(uint64(len(k))) + byteSize
			i = encodeVarintMerkleTree(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintMerkleTree(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if len(v) > 0 {
				dAtA[i] = 0x12
				i++
				i = encodeVarintMerkleTree(dAtA, i, uint64(len(v)))
				i += copy(dAtA[i:], v)
			}
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMerkleTree(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MerkleTree) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.HashList) > 0 {
		for _, b := range m.HashList {
			l = len(b)
			n += 1 + l + sovMerkleTree(uint64(l))
		}
	}
	if len(m.Hash2Idx) > 0 {
		for k, v := range m.Hash2Idx {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovMerkleTree(uint64(len(k))) + 1 + sovMerkleTree(uint64(v))
			n += mapEntrySize + 1 + sovMerkleTree(uint64(mapEntrySize))
		}
	}
	if m.LeafNum != 0 {
		n += 1 + sovMerkleTree(uint64(m.LeafNum))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TXRMerkleTree) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mt != nil {
		l = m.Mt.Size()
		n += 1 + l + sovMerkleTree(uint64(l))
	}
	if len(m.Tx2Txr) > 0 {
		for k, v := range m.Tx2Txr {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovMerkleTree(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovMerkleTree(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovMerkleTree(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMerkleTree(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMerkleTree(x uint64) (n int) {
	return sovMerkleTree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MerkleTree) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMerkleTree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MerkleTree: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MerkleTree: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashList", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkleTree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMerkleTree
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashList = append(m.HashList, make([]byte, postIndex-iNdEx))
			copy(m.HashList[len(m.HashList)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash2Idx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkleTree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMerkleTree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Hash2Idx == nil {
				m.Hash2Idx = make(map[string]int32)
			}
			var mapkey string
			var mapvalue int32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMerkleTree
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMerkleTree
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMerkleTree
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMerkleTree
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMerkleTree(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMerkleTree
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Hash2Idx[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeafNum", wireType)
			}
			m.LeafNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkleTree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeafNum |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMerkleTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMerkleTree
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
func (m *TXRMerkleTree) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMerkleTree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TXRMerkleTree: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TXRMerkleTree: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkleTree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMerkleTree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Mt == nil {
				m.Mt = &MerkleTree{}
			}
			if err := m.Mt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx2Txr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkleTree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMerkleTree
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx2Txr == nil {
				m.Tx2Txr = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMerkleTree
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMerkleTree
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMerkleTree
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMerkleTree
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthMerkleTree
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMerkleTree(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthMerkleTree
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tx2Txr[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMerkleTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMerkleTree
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
func skipMerkleTree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMerkleTree
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
					return 0, ErrIntOverflowMerkleTree
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMerkleTree
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMerkleTree
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMerkleTree
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMerkleTree(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMerkleTree = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMerkleTree   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("core/merkletree/merkle_tree.proto", fileDescriptor_merkle_tree_15dc9b8f1778676f)
}

var fileDescriptor_merkle_tree_15dc9b8f1778676f = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xce, 0x2f, 0x4a,
	0xd5, 0xcf, 0x4d, 0x2d, 0xca, 0xce, 0x49, 0x2d, 0x29, 0x4a, 0x85, 0x31, 0xe3, 0x41, 0x6c, 0xbd,
	0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x2e, 0x84, 0xac, 0xd2, 0x11, 0x46, 0x2e, 0x2e, 0x5f, 0x30,
	0x37, 0xa4, 0x28, 0x35, 0x55, 0x48, 0x9a, 0x8b, 0x33, 0x23, 0xb1, 0x38, 0x23, 0x3e, 0x27, 0xb3,
	0xb8, 0x44, 0x82, 0x51, 0x81, 0x59, 0x83, 0x27, 0x88, 0x03, 0x24, 0xe0, 0x93, 0x59, 0x5c, 0x22,
	0xe4, 0x08, 0x91, 0x34, 0x8a, 0xcf, 0x4c, 0xa9, 0x90, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x52,
	0xd1, 0x43, 0x98, 0xa5, 0x87, 0x30, 0x47, 0xcf, 0x03, 0xa4, 0xce, 0x33, 0xa5, 0xc2, 0x35, 0xaf,
	0xa4, 0xa8, 0x12, 0x62, 0x04, 0x88, 0x2b, 0x24, 0xc9, 0xc5, 0x91, 0x93, 0x9a, 0x98, 0x16, 0x9f,
	0x57, 0x9a, 0x2b, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x0e, 0xe2, 0xfb, 0x95, 0xe6, 0x4a,
	0x59, 0x73, 0xf1, 0xa2, 0xe8, 0x12, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98,
	0xc0, 0x5a, 0x21, 0x1c, 0x2b, 0x26, 0x0b, 0x46, 0xa5, 0x4d, 0x8c, 0x5c, 0xbc, 0x21, 0x11, 0x41,
	0x48, 0x3e, 0x51, 0xe3, 0x62, 0xca, 0x2d, 0x01, 0x6b, 0xe6, 0x36, 0x12, 0xc3, 0xee, 0xca, 0x20,
	0xa6, 0xdc, 0x12, 0x21, 0x3b, 0x2e, 0xf6, 0x92, 0x0a, 0xa3, 0xf8, 0x92, 0x8a, 0x22, 0xa8, 0x97,
	0x54, 0x91, 0x15, 0xa3, 0x98, 0xa9, 0x17, 0x52, 0x61, 0x14, 0x52, 0x51, 0x04, 0xf1, 0x13, 0x5b,
	0x09, 0x98, 0x23, 0x65, 0xc9, 0xc5, 0x8d, 0x24, 0x4c, 0xc8, 0xd1, 0x3c, 0x48, 0x8e, 0x76, 0x12,
	0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96,
	0x63, 0x48, 0x62, 0x03, 0x47, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x0f, 0x55, 0xa4,
	0xc5, 0x01, 0x00, 0x00,
}
