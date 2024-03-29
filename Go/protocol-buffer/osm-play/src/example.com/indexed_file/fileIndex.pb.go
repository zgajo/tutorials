// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: fileIndex.proto

package indexed_file

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StoredIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// System.currentTimeMillis()
	DateCreated int64        `protobuf:"varint,18,opt,name=dateCreated,proto3" json:"dateCreated,omitempty"`
	FileIndex   []*FileIndex `protobuf:"bytes,7,rep,name=fileIndex,proto3" json:"fileIndex,omitempty"`
}

func (x *StoredIndex) Reset() {
	*x = StoredIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileIndex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredIndex) ProtoMessage() {}

func (x *StoredIndex) ProtoReflect() protoreflect.Message {
	mi := &file_fileIndex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredIndex.ProtoReflect.Descriptor instead.
func (*StoredIndex) Descriptor() ([]byte, []int) {
	return file_fileIndex_proto_rawDescGZIP(), []int{0}
}

func (x *StoredIndex) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *StoredIndex) GetDateCreated() int64 {
	if x != nil {
		return x.DateCreated
	}
	return 0
}

func (x *StoredIndex) GetFileIndex() []*FileIndex {
	if x != nil {
		return x.FileIndex
	}
	return nil
}

type FileIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size         int64       `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	DateModified int64       `protobuf:"varint,2,opt,name=dateModified,proto3" json:"dateModified,omitempty"`
	FileName     string      `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Version      int32       `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	CityIndex    []*CityPart `protobuf:"bytes,8,rep,name=cityIndex,proto3" json:"cityIndex,omitempty"`
}

func (x *FileIndex) Reset() {
	*x = FileIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileIndex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileIndex) ProtoMessage() {}

func (x *FileIndex) ProtoReflect() protoreflect.Message {
	mi := &file_fileIndex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileIndex.ProtoReflect.Descriptor instead.
func (*FileIndex) Descriptor() ([]byte, []int) {
	return file_fileIndex_proto_rawDescGZIP(), []int{1}
}

func (x *FileIndex) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileIndex) GetDateModified() int64 {
	if x != nil {
		return x.DateModified
	}
	return 0
}

func (x *FileIndex) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileIndex) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *FileIndex) GetCityIndex() []*CityPart {
	if x != nil {
		return x.CityIndex
	}
	return nil
}

type CityPart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size            int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Offset          int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Name            string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	NameEn          string   `protobuf:"bytes,4,opt,name=nameEn,proto3" json:"nameEn,omitempty"`
	IndexNameOffset int32    `protobuf:"varint,5,opt,name=indexNameOffset,proto3" json:"indexNameOffset,omitempty"`
	AdditionalTags  []string `protobuf:"bytes,6,rep,name=additionalTags,proto3" json:"additionalTags,omitempty"`
}

func (x *CityPart) Reset() {
	*x = CityPart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileIndex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityPart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityPart) ProtoMessage() {}

func (x *CityPart) ProtoReflect() protoreflect.Message {
	mi := &file_fileIndex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityPart.ProtoReflect.Descriptor instead.
func (*CityPart) Descriptor() ([]byte, []int) {
	return file_fileIndex_proto_rawDescGZIP(), []int{2}
}

func (x *CityPart) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CityPart) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CityPart) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CityPart) GetNameEn() string {
	if x != nil {
		return x.NameEn
	}
	return ""
}

func (x *CityPart) GetIndexNameOffset() int32 {
	if x != nil {
		return x.IndexNameOffset
	}
	return 0
}

func (x *CityPart) GetAdditionalTags() []string {
	if x != nil {
		return x.AdditionalTags
	}
	return nil
}

var File_fileIndex_proto protoreflect.FileDescriptor

var file_fileIndex_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x80, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x22, 0xaf, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x43, 0x69, 0x74, 0x79, 0x50, 0x61, 0x72, 0x74, 0x52, 0x09, 0x63, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0xb4, 0x01, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x50, 0x61, 0x72,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x54, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x42, 0x2b, 0x0a, 0x0d, 0x63,
	0x72, 0x6f, 0x73, 0x62, 0x79, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x48, 0x03, 0x5a, 0x18,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fileIndex_proto_rawDescOnce sync.Once
	file_fileIndex_proto_rawDescData = file_fileIndex_proto_rawDesc
)

func file_fileIndex_proto_rawDescGZIP() []byte {
	file_fileIndex_proto_rawDescOnce.Do(func() {
		file_fileIndex_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileIndex_proto_rawDescData)
	})
	return file_fileIndex_proto_rawDescData
}

var file_fileIndex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fileIndex_proto_goTypes = []interface{}{
	(*StoredIndex)(nil), // 0: indexed.file.StoredIndex
	(*FileIndex)(nil),   // 1: indexed.file.FileIndex
	(*CityPart)(nil),    // 2: indexed.file.CityPart
}
var file_fileIndex_proto_depIdxs = []int32{
	1, // 0: indexed.file.StoredIndex.fileIndex:type_name -> indexed.file.FileIndex
	2, // 1: indexed.file.FileIndex.cityIndex:type_name -> indexed.file.CityPart
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fileIndex_proto_init() }
func file_fileIndex_proto_init() {
	if File_fileIndex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fileIndex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredIndex); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileIndex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileIndex); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fileIndex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityPart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fileIndex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fileIndex_proto_goTypes,
		DependencyIndexes: file_fileIndex_proto_depIdxs,
		MessageInfos:      file_fileIndex_proto_msgTypes,
	}.Build()
	File_fileIndex_proto = out.File
	file_fileIndex_proto_rawDesc = nil
	file_fileIndex_proto_goTypes = nil
	file_fileIndex_proto_depIdxs = nil
}
