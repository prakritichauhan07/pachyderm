// Code generated by protoc-gen-gogo.
// source: server/pfs/fuse/fuse.proto
// DO NOT EDIT!

/*
Package fuse is a generated protocol buffer package.

It is generated from these files:
	server/pfs/fuse/fuse.proto

It has these top-level messages:
	CommitMount
	Filesystem
	Node
	Attr
	Dirent
	Root
	DirectoryAttr
	DirectoryLookup
	DirectoryReadDirAll
	DirectoryCreate
	DirectoryMkdir
	FileAttr
	FileSetAttr
	FileRead
	FileOpen
	FileWrite
	FileRemove
*/
package fuse

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import google_protobuf1 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CommitMount struct {
	Commit *pfs.Commit `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	Alias  string      `protobuf:"bytes,4,opt,name=alias,proto3" json:"alias,omitempty"`
	Lazy   bool        `protobuf:"varint,6,opt,name=lazy,proto3" json:"lazy,omitempty"`
}

func (m *CommitMount) Reset()                    { *m = CommitMount{} }
func (m *CommitMount) String() string            { return proto.CompactTextString(m) }
func (*CommitMount) ProtoMessage()               {}
func (*CommitMount) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{0} }

func (m *CommitMount) GetCommit() *pfs.Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *CommitMount) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *CommitMount) GetLazy() bool {
	if m != nil {
		return m.Lazy
	}
	return false
}

type Filesystem struct {
	CommitMounts []*CommitMount `protobuf:"bytes,2,rep,name=commit_mounts,json=commitMounts" json:"commit_mounts,omitempty"`
}

func (m *Filesystem) Reset()                    { *m = Filesystem{} }
func (m *Filesystem) String() string            { return proto.CompactTextString(m) }
func (*Filesystem) ProtoMessage()               {}
func (*Filesystem) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{1} }

func (m *Filesystem) GetCommitMounts() []*CommitMount {
	if m != nil {
		return m.CommitMounts
	}
	return nil
}

type Node struct {
	File      *pfs.File                   `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	RepoAlias string                      `protobuf:"bytes,2,opt,name=repo_alias,json=repoAlias,proto3" json:"repo_alias,omitempty"`
	Write     bool                        `protobuf:"varint,3,opt,name=write,proto3" json:"write,omitempty"`
	Modified  *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=modified" json:"modified,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{2} }

func (m *Node) GetFile() *pfs.File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Node) GetRepoAlias() string {
	if m != nil {
		return m.RepoAlias
	}
	return ""
}

func (m *Node) GetWrite() bool {
	if m != nil {
		return m.Write
	}
	return false
}

func (m *Node) GetModified() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Attr struct {
	Mode uint32 `protobuf:"varint,1,opt,name=Mode,json=mode,proto3" json:"Mode,omitempty"`
}

func (m *Attr) Reset()                    { *m = Attr{} }
func (m *Attr) String() string            { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()               {}
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{3} }

func (m *Attr) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type Dirent struct {
	Inode uint64 `protobuf:"varint,1,opt,name=inode,proto3" json:"inode,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Dirent) Reset()                    { *m = Dirent{} }
func (m *Dirent) String() string            { return proto.CompactTextString(m) }
func (*Dirent) ProtoMessage()               {}
func (*Dirent) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{4} }

func (m *Dirent) GetInode() uint64 {
	if m != nil {
		return m.Inode
	}
	return 0
}

func (m *Dirent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Root struct {
	Filesystem *Filesystem `protobuf:"bytes,1,opt,name=filesystem" json:"filesystem,omitempty"`
	Result     *Node       `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error      string      `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *Root) Reset()                    { *m = Root{} }
func (m *Root) String() string            { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()               {}
func (*Root) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{5} }

func (m *Root) GetFilesystem() *Filesystem {
	if m != nil {
		return m.Filesystem
	}
	return nil
}

func (m *Root) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Root) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DirectoryAttr struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *DirectoryAttr) Reset()                    { *m = DirectoryAttr{} }
func (m *DirectoryAttr) String() string            { return proto.CompactTextString(m) }
func (*DirectoryAttr) ProtoMessage()               {}
func (*DirectoryAttr) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{6} }

func (m *DirectoryAttr) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DirectoryAttr) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DirectoryLookup struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Result    *Node  `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Err       string `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *DirectoryLookup) Reset()                    { *m = DirectoryLookup{} }
func (m *DirectoryLookup) String() string            { return proto.CompactTextString(m) }
func (*DirectoryLookup) ProtoMessage()               {}
func (*DirectoryLookup) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{7} }

func (m *DirectoryLookup) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryLookup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DirectoryLookup) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DirectoryLookup) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type DirectoryReadDirAll struct {
	Directory *Node     `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    []*Dirent `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
	Error     string    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *DirectoryReadDirAll) Reset()                    { *m = DirectoryReadDirAll{} }
func (m *DirectoryReadDirAll) String() string            { return proto.CompactTextString(m) }
func (*DirectoryReadDirAll) ProtoMessage()               {}
func (*DirectoryReadDirAll) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{8} }

func (m *DirectoryReadDirAll) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryReadDirAll) GetResult() []*Dirent {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DirectoryReadDirAll) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DirectoryCreate struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *DirectoryCreate) Reset()                    { *m = DirectoryCreate{} }
func (m *DirectoryCreate) String() string            { return proto.CompactTextString(m) }
func (*DirectoryCreate) ProtoMessage()               {}
func (*DirectoryCreate) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{9} }

func (m *DirectoryCreate) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryCreate) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DirectoryCreate) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DirectoryMkdir struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *DirectoryMkdir) Reset()                    { *m = DirectoryMkdir{} }
func (m *DirectoryMkdir) String() string            { return proto.CompactTextString(m) }
func (*DirectoryMkdir) ProtoMessage()               {}
func (*DirectoryMkdir) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{10} }

func (m *DirectoryMkdir) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryMkdir) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DirectoryMkdir) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FileAttr struct {
	File   *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Result *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *FileAttr) Reset()                    { *m = FileAttr{} }
func (m *FileAttr) String() string            { return proto.CompactTextString(m) }
func (*FileAttr) ProtoMessage()               {}
func (*FileAttr) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{11} }

func (m *FileAttr) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *FileAttr) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FileSetAttr struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *FileSetAttr) Reset()                    { *m = FileSetAttr{} }
func (m *FileSetAttr) String() string            { return proto.CompactTextString(m) }
func (*FileSetAttr) ProtoMessage()               {}
func (*FileSetAttr) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{12} }

func (m *FileSetAttr) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileSetAttr) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FileRead struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *FileRead) Reset()                    { *m = FileRead{} }
func (m *FileRead) String() string            { return proto.CompactTextString(m) }
func (*FileRead) ProtoMessage()               {}
func (*FileRead) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{13} }

func (m *FileRead) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileRead) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *FileRead) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FileOpen struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *FileOpen) Reset()                    { *m = FileOpen{} }
func (m *FileOpen) String() string            { return proto.CompactTextString(m) }
func (*FileOpen) ProtoMessage()               {}
func (*FileOpen) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{14} }

func (m *FileOpen) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileOpen) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FileWrite struct {
	File   *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Data   string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Error  string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *FileWrite) Reset()                    { *m = FileWrite{} }
func (m *FileWrite) String() string            { return proto.CompactTextString(m) }
func (*FileWrite) ProtoMessage()               {}
func (*FileWrite) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{15} }

func (m *FileWrite) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileWrite) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *FileWrite) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FileWrite) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FileRemove struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dir   bool   `protobuf:"varint,3,opt,name=dir,proto3" json:"dir,omitempty"`
	Error string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *FileRemove) Reset()                    { *m = FileRemove{} }
func (m *FileRemove) String() string            { return proto.CompactTextString(m) }
func (*FileRemove) ProtoMessage()               {}
func (*FileRemove) Descriptor() ([]byte, []int) { return fileDescriptorFuse, []int{16} }

func (m *FileRemove) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileRemove) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileRemove) GetDir() bool {
	if m != nil {
		return m.Dir
	}
	return false
}

func (m *FileRemove) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*CommitMount)(nil), "fuse.CommitMount")
	proto.RegisterType((*Filesystem)(nil), "fuse.Filesystem")
	proto.RegisterType((*Node)(nil), "fuse.Node")
	proto.RegisterType((*Attr)(nil), "fuse.Attr")
	proto.RegisterType((*Dirent)(nil), "fuse.Dirent")
	proto.RegisterType((*Root)(nil), "fuse.Root")
	proto.RegisterType((*DirectoryAttr)(nil), "fuse.DirectoryAttr")
	proto.RegisterType((*DirectoryLookup)(nil), "fuse.DirectoryLookup")
	proto.RegisterType((*DirectoryReadDirAll)(nil), "fuse.DirectoryReadDirAll")
	proto.RegisterType((*DirectoryCreate)(nil), "fuse.DirectoryCreate")
	proto.RegisterType((*DirectoryMkdir)(nil), "fuse.DirectoryMkdir")
	proto.RegisterType((*FileAttr)(nil), "fuse.FileAttr")
	proto.RegisterType((*FileSetAttr)(nil), "fuse.FileSetAttr")
	proto.RegisterType((*FileRead)(nil), "fuse.FileRead")
	proto.RegisterType((*FileOpen)(nil), "fuse.FileOpen")
	proto.RegisterType((*FileWrite)(nil), "fuse.FileWrite")
	proto.RegisterType((*FileRemove)(nil), "fuse.FileRemove")
}

func init() { proto.RegisterFile("server/pfs/fuse/fuse.proto", fileDescriptorFuse) }

var fileDescriptorFuse = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0x96, 0x2c, 0x5a, 0x2e, 0x1b, 0x0c, 0x33, 0xa1, 0xa8, 0xd2, 0xa0, 0x0a, 0x3c, 0xf4,
	0x29, 0x45, 0x45, 0xda, 0x33, 0xd3, 0x2a, 0x9e, 0x28, 0x48, 0x66, 0x12, 0x2f, 0x48, 0x53, 0xd6,
	0x5c, 0x86, 0xb5, 0x24, 0x8e, 0x6c, 0x67, 0xa8, 0xf0, 0xcc, 0x1f, 0xe0, 0x17, 0x23, 0xdb, 0x69,
	0x1a, 0xc4, 0xaa, 0x76, 0x45, 0xe2, 0xa5, 0xba, 0xf3, 0x5d, 0xbf, 0xef, 0xbb, 0x2f, 0x67, 0xc3,
	0x40, 0xa2, 0xb8, 0x43, 0x31, 0xae, 0x73, 0x39, 0xce, 0x1b, 0x89, 0xe6, 0x27, 0xa9, 0x05, 0x57,
	0x9c, 0x78, 0x3a, 0x1e, 0x9c, 0xcc, 0x0b, 0x86, 0x95, 0x32, 0x1d, 0x75, 0x2e, 0x6d, 0x6d, 0xf0,
	0xe2, 0x86, 0xf3, 0x9b, 0x02, 0xc7, 0x26, 0xbb, 0x6e, 0xf2, 0xb1, 0x62, 0x25, 0x4a, 0x95, 0x96,
	0xb5, 0x6d, 0x88, 0xbf, 0x40, 0x78, 0xc1, 0xcb, 0x92, 0xa9, 0x19, 0x6f, 0x2a, 0x45, 0x5e, 0x82,
	0x3f, 0x37, 0x69, 0xe4, 0x0c, 0x9d, 0x51, 0x38, 0x09, 0x13, 0x8d, 0x65, 0x3b, 0x68, 0x5b, 0x22,
	0x27, 0xb0, 0x9f, 0x16, 0x2c, 0x95, 0x91, 0x37, 0x74, 0x46, 0x01, 0xb5, 0x09, 0x21, 0xe0, 0x15,
	0xe9, 0xf7, 0x45, 0xe4, 0x0f, 0x9d, 0xd1, 0x01, 0x35, 0x71, 0x3c, 0x05, 0x78, 0xc7, 0x0a, 0x94,
	0x0b, 0xa9, 0xb0, 0x24, 0x67, 0x70, 0x64, 0x11, 0xae, 0x4a, 0x4d, 0x26, 0xa3, 0xbd, 0xa1, 0x3b,
	0x0a, 0x27, 0x4f, 0x12, 0x33, 0x4c, 0x4f, 0x06, 0x3d, 0x9c, 0xaf, 0x12, 0x19, 0xff, 0x72, 0xc0,
	0xfb, 0xc0, 0x33, 0x24, 0xa7, 0xe0, 0xe5, 0xac, 0xc0, 0x56, 0x5b, 0x60, 0xb4, 0x69, 0x7c, 0x6a,
	0x8e, 0xc9, 0x29, 0x80, 0xc0, 0x9a, 0x5f, 0x59, 0x71, 0x7b, 0x46, 0x5c, 0xa0, 0x4f, 0xce, 0x8d,
	0xc0, 0x13, 0xd8, 0xff, 0x26, 0x98, 0xc2, 0xc8, 0x35, 0x0a, 0x6d, 0x42, 0xce, 0xe0, 0xa0, 0xe4,
	0x19, 0xcb, 0x19, 0x66, 0xd1, 0xbe, 0xc1, 0x1d, 0x24, 0xd6, 0xb4, 0x64, 0x69, 0x5a, 0x72, 0xb9,
	0x34, 0x8d, 0x76, 0xbd, 0xf1, 0x00, 0xbc, 0x73, 0xa5, 0x84, 0x1e, 0x7b, 0xc6, 0x33, 0xab, 0xe9,
	0x88, 0x7a, 0x25, 0xcf, 0x30, 0x9e, 0x80, 0x3f, 0x65, 0x02, 0x2b, 0x63, 0x15, 0xab, 0x96, 0x65,
	0x8f, 0xda, 0x44, 0xff, 0xa7, 0x4a, 0x4b, 0x6c, 0x25, 0x9a, 0x38, 0x16, 0xe0, 0x51, 0xce, 0x15,
	0x79, 0x0d, 0x90, 0x77, 0x96, 0xb5, 0x93, 0x1e, 0x5b, 0x87, 0x56, 0x56, 0xd2, 0x5e, 0x0f, 0x89,
	0xc1, 0x17, 0x28, 0x9b, 0x42, 0x19, 0xbc, 0x70, 0x02, 0xb6, 0x5b, 0x3b, 0x46, 0xdb, 0x8a, 0xd6,
	0x81, 0x42, 0x70, 0x61, 0x66, 0x0f, 0xa8, 0x4d, 0x62, 0x09, 0x47, 0x5a, 0xe7, 0x5c, 0x71, 0xb1,
	0x30, 0xc3, 0x8c, 0x20, 0xc8, 0x96, 0x07, 0x2d, 0x77, 0x1f, 0x6d, 0x55, 0x5c, 0x47, 0xaa, 0x51,
	0x36, 0x90, 0xfe, 0x74, 0xe0, 0x71, 0xc7, 0xfa, 0x9e, 0xf3, 0xdb, 0xa6, 0x7e, 0x00, 0xef, 0x3d,
	0xd6, 0xf5, 0xb4, 0xb8, 0x6b, 0x0d, 0x38, 0x06, 0x17, 0x85, 0x68, 0x37, 0x56, 0x87, 0xf1, 0x0f,
	0x78, 0xda, 0xc9, 0xa0, 0x98, 0x66, 0x53, 0x26, 0xce, 0x8b, 0xe2, 0x01, 0x52, 0x5e, 0xf5, 0x2c,
	0xd0, 0x7b, 0x7c, 0x68, 0xdb, 0xec, 0x97, 0xdf, 0x60, 0x42, 0xd3, 0xf3, 0xe0, 0x42, 0x60, 0xaa,
	0xf0, 0xdf, 0xbd, 0xdf, 0xe2, 0x83, 0x2b, 0x78, 0xd4, 0xd1, 0xce, 0x6e, 0x33, 0x26, 0xfe, 0x0b,
	0x6b, 0x06, 0x07, 0x7a, 0x75, 0xcd, 0x86, 0x3d, 0xff, 0xe3, 0x0a, 0xf7, 0x31, 0xec, 0x1d, 0xde,
	0x7d, 0xaf, 0x2e, 0x20, 0xd4, 0x2c, 0x9f, 0x50, 0x6d, 0x45, 0xd4, 0x81, 0xec, 0xf5, 0x41, 0x2e,
	0xad, 0x54, 0xbd, 0x0f, 0x1b, 0x11, 0x08, 0x78, 0x59, 0xaa, 0xd2, 0xe5, 0x2a, 0xea, 0x78, 0x8d,
	0xb4, 0xb7, 0x16, 0xf5, 0x63, 0x8d, 0xd5, 0x8e, 0xba, 0x4a, 0x08, 0x34, 0xc2, 0x67, 0xf3, 0x64,
	0xed, 0x22, 0xec, 0x19, 0xf8, 0x3c, 0xcf, 0x25, 0xda, 0x3b, 0xe2, 0xd2, 0x36, 0x5b, 0xd1, 0x79,
	0x7d, 0xba, 0xaf, 0xf6, 0xdd, 0xa6, 0x58, 0xf2, 0xbb, 0xad, 0xf8, 0xfe, 0xba, 0x93, 0xc7, 0xe0,
	0x66, 0x4c, 0xb4, 0x4f, 0xad, 0x0e, 0xef, 0x67, 0xba, 0xf6, 0xcd, 0x23, 0xfb, 0xe6, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8d, 0x0f, 0x1c, 0x33, 0xe1, 0x06, 0x00, 0x00,
}
