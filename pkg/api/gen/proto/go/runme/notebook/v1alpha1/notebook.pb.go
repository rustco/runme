// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: runme/notebook/v1alpha1/notebook.proto

package notebookv1alpha1

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	v1 "github.com/stateful/runme/v3/pkg/api/gen/proto/go/runme/parser/v1"
	v11 "github.com/stateful/runme/v3/pkg/api/gen/proto/go/runme/runner/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResolveNotebookRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Notebook      *v1.Notebook            `protobuf:"bytes,1,opt,name=notebook,proto3" json:"notebook,omitempty"`
	CommandMode   v11.CommandMode         `protobuf:"varint,2,opt,name=command_mode,json=commandMode,proto3,enum=runme.runner.v1.CommandMode" json:"command_mode,omitempty"`
	CellIndex     *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=cell_index,json=cellIndex,proto3" json:"cell_index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResolveNotebookRequest) Reset() {
	*x = ResolveNotebookRequest{}
	mi := &file_runme_notebook_v1alpha1_notebook_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveNotebookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveNotebookRequest) ProtoMessage() {}

func (x *ResolveNotebookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runme_notebook_v1alpha1_notebook_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveNotebookRequest.ProtoReflect.Descriptor instead.
func (*ResolveNotebookRequest) Descriptor() ([]byte, []int) {
	return file_runme_notebook_v1alpha1_notebook_proto_rawDescGZIP(), []int{0}
}

func (x *ResolveNotebookRequest) GetNotebook() *v1.Notebook {
	if x != nil {
		return x.Notebook
	}
	return nil
}

func (x *ResolveNotebookRequest) GetCommandMode() v11.CommandMode {
	if x != nil {
		return x.CommandMode
	}
	return v11.CommandMode(0)
}

func (x *ResolveNotebookRequest) GetCellIndex() *wrapperspb.UInt32Value {
	if x != nil {
		return x.CellIndex
	}
	return nil
}

type ResolveNotebookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Script        string                 `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResolveNotebookResponse) Reset() {
	*x = ResolveNotebookResponse{}
	mi := &file_runme_notebook_v1alpha1_notebook_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResolveNotebookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveNotebookResponse) ProtoMessage() {}

func (x *ResolveNotebookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runme_notebook_v1alpha1_notebook_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveNotebookResponse.ProtoReflect.Descriptor instead.
func (*ResolveNotebookResponse) Descriptor() ([]byte, []int) {
	return file_runme_notebook_v1alpha1_notebook_proto_rawDescGZIP(), []int{1}
}

func (x *ResolveNotebookResponse) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

var File_runme_notebook_v1alpha1_notebook_proto protoreflect.FileDescriptor

var file_runme_notebook_v1alpha1_notebook_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01,
	0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x75, 0x6e,
	0x6d, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12,
	0x3f, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x72, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x31, 0x0a,
	0x17, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x32, 0x89, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x2f, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5c, 0x5a, 0x5a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x66, 0x75, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_runme_notebook_v1alpha1_notebook_proto_rawDescOnce sync.Once
	file_runme_notebook_v1alpha1_notebook_proto_rawDescData []byte
)

func file_runme_notebook_v1alpha1_notebook_proto_rawDescGZIP() []byte {
	file_runme_notebook_v1alpha1_notebook_proto_rawDescOnce.Do(func() {
		file_runme_notebook_v1alpha1_notebook_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_runme_notebook_v1alpha1_notebook_proto_rawDesc), len(file_runme_notebook_v1alpha1_notebook_proto_rawDesc)))
	})
	return file_runme_notebook_v1alpha1_notebook_proto_rawDescData
}

var file_runme_notebook_v1alpha1_notebook_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_runme_notebook_v1alpha1_notebook_proto_goTypes = []any{
	(*ResolveNotebookRequest)(nil),  // 0: runme.notebook.v1alpha1.ResolveNotebookRequest
	(*ResolveNotebookResponse)(nil), // 1: runme.notebook.v1alpha1.ResolveNotebookResponse
	(*v1.Notebook)(nil),             // 2: runme.parser.v1.Notebook
	(v11.CommandMode)(0),            // 3: runme.runner.v1.CommandMode
	(*wrapperspb.UInt32Value)(nil),  // 4: google.protobuf.UInt32Value
}
var file_runme_notebook_v1alpha1_notebook_proto_depIdxs = []int32{
	2, // 0: runme.notebook.v1alpha1.ResolveNotebookRequest.notebook:type_name -> runme.parser.v1.Notebook
	3, // 1: runme.notebook.v1alpha1.ResolveNotebookRequest.command_mode:type_name -> runme.runner.v1.CommandMode
	4, // 2: runme.notebook.v1alpha1.ResolveNotebookRequest.cell_index:type_name -> google.protobuf.UInt32Value
	0, // 3: runme.notebook.v1alpha1.NotebookService.ResolveNotebook:input_type -> runme.notebook.v1alpha1.ResolveNotebookRequest
	1, // 4: runme.notebook.v1alpha1.NotebookService.ResolveNotebook:output_type -> runme.notebook.v1alpha1.ResolveNotebookResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_runme_notebook_v1alpha1_notebook_proto_init() }
func file_runme_notebook_v1alpha1_notebook_proto_init() {
	if File_runme_notebook_v1alpha1_notebook_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_runme_notebook_v1alpha1_notebook_proto_rawDesc), len(file_runme_notebook_v1alpha1_notebook_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runme_notebook_v1alpha1_notebook_proto_goTypes,
		DependencyIndexes: file_runme_notebook_v1alpha1_notebook_proto_depIdxs,
		MessageInfos:      file_runme_notebook_v1alpha1_notebook_proto_msgTypes,
	}.Build()
	File_runme_notebook_v1alpha1_notebook_proto = out.File
	file_runme_notebook_v1alpha1_notebook_proto_goTypes = nil
	file_runme_notebook_v1alpha1_notebook_proto_depIdxs = nil
}
