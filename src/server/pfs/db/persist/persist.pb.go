// Code generated by protoc-gen-go.
// source: server/pfs/db/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pfs/db/persist/persist.proto

It has these top-level messages:
	Clock
	ClockID
	BranchClock
	Repo
	Branch
	BlockRef
	Diff
	Commit
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileType int32

const (
	FileType_NONE FileType = 0
	FileType_FILE FileType = 1
	FileType_DIR  FileType = 2
)

var FileType_name = map[int32]string{
	0: "NONE",
	1: "FILE",
	2: "DIR",
}
var FileType_value = map[string]int32{
	"NONE": 0,
	"FILE": 1,
	"DIR":  2,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}
func (FileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Clock struct {
	// a document either has these two fields
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,2,opt,name=clock" json:"clock,omitempty"`
}

func (m *Clock) Reset()                    { *m = Clock{} }
func (m *Clock) String() string            { return proto.CompactTextString(m) }
func (*Clock) ProtoMessage()               {}
func (*Clock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClockID struct {
	ID     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Branch string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,4,opt,name=clock" json:"clock,omitempty"`
}

func (m *ClockID) Reset()                    { *m = ClockID{} }
func (m *ClockID) String() string            { return proto.CompactTextString(m) }
func (*ClockID) ProtoMessage()               {}
func (*ClockID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BranchClock struct {
	Clocks []*Clock `protobuf:"bytes,1,rep,name=clocks" json:"clocks,omitempty"`
}

func (m *BranchClock) Reset()                    { *m = BranchClock{} }
func (m *BranchClock) String() string            { return proto.CompactTextString(m) }
func (*BranchClock) ProtoMessage()               {}
func (*BranchClock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BranchClock) GetClocks() []*Clock {
	if m != nil {
		return m.Clocks
	}
	return nil
}

type Repo struct {
	Name    string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Repo) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type Branch struct {
	ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Branch) Reset()                    { *m = Branch{} }
func (m *Branch) String() string            { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()               {}
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type BlockRef struct {
	Hash  string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Lower uint64 `protobuf:"varint,2,opt,name=lower" json:"lower,omitempty"`
	Upper uint64 `protobuf:"varint,3,opt,name=upper" json:"upper,omitempty"`
}

func (m *BlockRef) Reset()                    { *m = BlockRef{} }
func (m *BlockRef) String() string            { return proto.CompactTextString(m) }
func (*BlockRef) ProtoMessage()               {}
func (*BlockRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Diff struct {
	ID       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo     string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	CommitID string `protobuf:"bytes,3,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	Path     string `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	// block_refs and delete cannot both be set
	BlockRefs    []*BlockRef    `protobuf:"bytes,5,rep,name=block_refs,json=blockRefs" json:"block_refs,omitempty"`
	Delete       bool           `protobuf:"varint,6,opt,name=delete" json:"delete,omitempty"`
	Size         uint64         `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	BranchClocks []*BranchClock `protobuf:"bytes,8,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	FileType     FileType       `protobuf:"varint,9,opt,name=file_type,json=fileType,enum=FileType" json:"file_type,omitempty"`
}

func (m *Diff) Reset()                    { *m = Diff{} }
func (m *Diff) String() string            { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()               {}
func (*Diff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Diff) GetBlockRefs() []*BlockRef {
	if m != nil {
		return m.BlockRefs
	}
	return nil
}

func (m *Diff) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

type Commit struct {
	ID           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo         string                     `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	BranchClocks []*BranchClock             `protobuf:"bytes,3,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	Started      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started" json:"started,omitempty"`
	Finished     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=finished" json:"finished,omitempty"`
	Cancelled    bool                       `protobuf:"varint,6,opt,name=cancelled" json:"cancelled,omitempty"`
	Provenance   []string                   `protobuf:"bytes,7,rep,name=provenance" json:"provenance,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Commit) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Commit) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Commit) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func init() {
	proto.RegisterType((*Clock)(nil), "Clock")
	proto.RegisterType((*ClockID)(nil), "ClockID")
	proto.RegisterType((*BranchClock)(nil), "BranchClock")
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*Branch)(nil), "Branch")
	proto.RegisterType((*BlockRef)(nil), "BlockRef")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Commit)(nil), "Commit")
	proto.RegisterEnum("FileType", FileType_name, FileType_value)
}

func init() { proto.RegisterFile("server/pfs/db/persist/persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x5f, 0x8b, 0xd3, 0x5e,
	0x10, 0xfd, 0xb5, 0x49, 0xd3, 0x64, 0xba, 0xbf, 0xa5, 0x5c, 0x44, 0xc2, 0x2a, 0x2a, 0x11, 0xb4,
	0x08, 0x26, 0xb8, 0xfe, 0x79, 0x96, 0xdd, 0xee, 0x42, 0x45, 0x56, 0xb9, 0xec, 0x9b, 0x0f, 0x25,
	0x49, 0x27, 0xdb, 0x60, 0xd2, 0x84, 0x7b, 0xb3, 0x2b, 0xfa, 0x15, 0xf4, 0x43, 0x3b, 0x77, 0x92,
	0xd8, 0x82, 0x42, 0xfb, 0x94, 0x99, 0x93, 0x39, 0x33, 0xe7, 0x9e, 0x19, 0x78, 0xaa, 0x51, 0xdd,
	0xa1, 0x8a, 0xea, 0x4c, 0x47, 0xab, 0x24, 0xaa, 0x51, 0xe9, 0x5c, 0x37, 0xfd, 0x37, 0xac, 0x55,
	0xd5, 0x54, 0x27, 0x8f, 0x6f, 0xaa, 0xea, 0xa6, 0xc0, 0x88, 0xb3, 0xe4, 0x36, 0x8b, 0x9a, 0xbc,
	0x44, 0xdd, 0xc4, 0x65, 0xdd, 0x16, 0x04, 0x6f, 0x61, 0x74, 0x5e, 0x54, 0xe9, 0x57, 0x71, 0x1f,
	0x9c, 0x44, 0xc5, 0x9b, 0x74, 0xed, 0x0f, 0x9e, 0x0c, 0x66, 0x9e, 0xec, 0x32, 0x71, 0x0f, 0x46,
	0xa9, 0x29, 0xf0, 0x87, 0x04, 0xdb, 0xb2, 0x4d, 0x82, 0x2f, 0x30, 0x66, 0xda, 0x62, 0x2e, 0x8e,
	0x61, 0x98, 0xaf, 0x3a, 0x12, 0x45, 0x42, 0x80, 0xad, 0xb0, 0xae, 0xb8, 0xde, 0x93, 0x1c, 0xef,
	0x34, 0xb7, 0xfe, 0xdd, 0xdc, 0xde, 0x6d, 0xfe, 0x12, 0x26, 0x67, 0xfc, 0xbf, 0x55, 0xf6, 0x08,
	0x1c, 0xc6, 0x35, 0x0d, 0xb1, 0x66, 0x93, 0x53, 0x27, 0x64, 0x5c, 0x76, 0x68, 0xf0, 0x19, 0x6c,
	0x69, 0x86, 0xd0, 0xe0, 0x4d, 0x5c, 0x62, 0x27, 0x85, 0x63, 0xf1, 0x06, 0xc6, 0xa9, 0xc2, 0xb8,
	0xc1, 0x15, 0xeb, 0x99, 0x9c, 0x9e, 0x84, 0xad, 0x23, 0x61, 0xef, 0x48, 0x78, 0xdd, 0x3b, 0x22,
	0xfb, 0xd2, 0xe0, 0x3d, 0x38, 0xad, 0x80, 0x83, 0x1e, 0xd7, 0xcf, 0xb5, 0xb6, 0x73, 0x83, 0x0f,
	0xe0, 0x9e, 0xb1, 0x48, 0xcc, 0xcc, 0xff, 0x75, 0xac, 0x7b, 0x5f, 0x39, 0x36, 0x0f, 0x2f, 0xaa,
	0x6f, 0xa8, 0x7a, 0x57, 0x39, 0x31, 0xe8, 0x6d, 0x4d, 0x0b, 0xe4, 0x56, 0x84, 0x72, 0x12, 0xfc,
	0x1a, 0x82, 0x3d, 0xcf, 0xb3, 0xec, 0x20, 0x31, 0x0f, 0xc0, 0x4b, 0xab, 0xb2, 0xcc, 0x9b, 0x25,
	0x95, 0xb6, 0x8a, 0xdc, 0x16, 0x58, 0x30, 0xa1, 0x8e, 0x9b, 0x35, 0xbb, 0x4d, 0x04, 0x13, 0x8b,
	0x19, 0x40, 0x62, 0x94, 0x2e, 0x15, 0x66, 0xda, 0x1f, 0xb1, 0xc3, 0x5e, 0xd8, 0x8b, 0x97, 0x5e,
	0xd2, 0x45, 0xda, 0x2c, 0x71, 0x85, 0x05, 0x36, 0xe8, 0x3b, 0xc4, 0x77, 0x65, 0x97, 0x99, 0xae,
	0x3a, 0xff, 0x81, 0xfe, 0x98, 0x45, 0x73, 0x2c, 0x5e, 0xc1, 0xff, 0xed, 0x8a, 0x97, 0xdd, 0xea,
	0x5c, 0x6e, 0x7c, 0x14, 0xee, 0x2c, 0x56, 0x1e, 0x25, 0xdb, 0x44, 0x8b, 0x67, 0xe0, 0x65, 0x79,
	0x81, 0xcb, 0xe6, 0x7b, 0x8d, 0xbe, 0x47, 0xbd, 0x8e, 0x49, 0xc7, 0x25, 0x21, 0xd7, 0x04, 0x48,
	0x37, 0xeb, 0xa2, 0xe0, 0xe7, 0x10, 0x9c, 0x73, 0x7e, 0xd1, 0x41, 0x86, 0xfc, 0xa5, 0xc4, 0xda,
	0xab, 0x84, 0x8e, 0x86, 0x0e, 0x42, 0x99, 0xa3, 0xb1, 0xf7, 0x1f, 0x4d, 0x57, 0x2a, 0xde, 0x01,
	0x69, 0xdc, 0xe4, 0x7a, 0x4d, 0xb4, 0xd1, 0x5e, 0xda, 0x9f, 0x5a, 0xf1, 0x90, 0x36, 0x46, 0xc3,
	0xb1, 0x28, 0x88, 0xd8, 0x3a, 0xbb, 0x05, 0xe8, 0xf8, 0x81, 0xd8, 0x77, 0xb8, 0x31, 0x08, 0x59,
	0x6c, 0xd1, 0xc3, 0x76, 0x90, 0x17, 0xcf, 0xc1, 0xed, 0x3d, 0x12, 0x2e, 0xd8, 0x57, 0x9f, 0xae,
	0x2e, 0xa6, 0xff, 0x99, 0xe8, 0x72, 0xf1, 0xf1, 0x62, 0x3a, 0x10, 0x63, 0xb0, 0xe6, 0x0b, 0x39,
	0x1d, 0x26, 0x0e, 0x8b, 0x78, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xd9, 0xd9, 0x63, 0x37,
	0x04, 0x00, 0x00,
}