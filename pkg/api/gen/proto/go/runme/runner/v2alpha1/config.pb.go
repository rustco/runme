// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: runme/runner/v2alpha1/config.proto

package runnerv2alpha1

import (
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

type CommandMode int32

const (
	CommandMode_COMMAND_MODE_UNSPECIFIED CommandMode = 0
	CommandMode_COMMAND_MODE_INLINE      CommandMode = 1
	CommandMode_COMMAND_MODE_FILE        CommandMode = 2
	CommandMode_COMMAND_MODE_TERMINAL    CommandMode = 3
)

// Enum value maps for CommandMode.
var (
	CommandMode_name = map[int32]string{
		0: "COMMAND_MODE_UNSPECIFIED",
		1: "COMMAND_MODE_INLINE",
		2: "COMMAND_MODE_FILE",
		3: "COMMAND_MODE_TERMINAL",
	}
	CommandMode_value = map[string]int32{
		"COMMAND_MODE_UNSPECIFIED": 0,
		"COMMAND_MODE_INLINE":      1,
		"COMMAND_MODE_FILE":        2,
		"COMMAND_MODE_TERMINAL":    3,
	}
)

func (x CommandMode) Enum() *CommandMode {
	p := new(CommandMode)
	*p = x
	return p
}

func (x CommandMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandMode) Descriptor() protoreflect.EnumDescriptor {
	return file_runme_runner_v2alpha1_config_proto_enumTypes[0].Descriptor()
}

func (CommandMode) Type() protoreflect.EnumType {
	return &file_runme_runner_v2alpha1_config_proto_enumTypes[0]
}

func (x CommandMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandMode.Descriptor instead.
func (CommandMode) EnumDescriptor() ([]byte, []int) {
	return file_runme_runner_v2alpha1_config_proto_rawDescGZIP(), []int{0}
}

// ProgramConfig is a configuration for a program to execute.
// From this configuration, any program can be built.
type ProgramConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// program_name is a name of the program to execute.
	// If it's not a path (relative or absolute), the runner
	// will try to resolve the name.
	// For example: "sh", "/bin/bash".
	ProgramName string `protobuf:"bytes,1,opt,name=program_name,json=programName,proto3" json:"program_name,omitempty"`
	// arguments is a list of arguments passed to the program.
	Arguments []string `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// directory to execute the program in.
	Directory string `protobuf:"bytes,3,opt,name=directory,proto3" json:"directory,omitempty"`
	// language_id associated with script to allow interpreter
	// detection if no specific interpreter program is provided.
	LanguageId string `protobuf:"bytes,4,opt,name=language_id,json=languageId,proto3" json:"language_id,omitempty"`
	// background indicates a background process
	// required to handle running background tasks via CLI in C/I
	Background bool `protobuf:"varint,5,opt,name=background,proto3" json:"background,omitempty"`
	// file_extension associated with script. Some interpreters are strict
	// about file extensions, such as tsc which requires "ts" extension.
	FileExtension string `protobuf:"bytes,6,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
	// env is a list of additional environment variables
	// that will be injected to the executed program.
	Env []string `protobuf:"bytes,7,rep,name=env,proto3" json:"env,omitempty"`
	// Types that are assignable to Source:
	//
	//	*ProgramConfig_Commands
	//	*ProgramConfig_Script
	Source isProgramConfig_Source `protobuf_oneof:"source"`
	// interactive, if true, uses a pseudo-tty to execute the program.
	// Otherwise, the program is executed using in-memory buffers for I/O.
	Interactive bool `protobuf:"varint,10,opt,name=interactive,proto3" json:"interactive,omitempty"`
	// TODO(adamb): understand motivation for this. In theory, source
	// should tell whether to execute it inline or as a file.
	Mode CommandMode `protobuf:"varint,11,opt,name=mode,proto3,enum=runme.runner.v2alpha1.CommandMode" json:"mode,omitempty"`
	// optional well known id for cell/block
	KnownId string `protobuf:"bytes,12,opt,name=known_id,json=knownId,proto3" json:"known_id,omitempty"`
	// optional well known name for cell/block
	KnownName string `protobuf:"bytes,13,opt,name=known_name,json=knownName,proto3" json:"known_name,omitempty"`
}

func (x *ProgramConfig) Reset() {
	*x = ProgramConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_runner_v2alpha1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgramConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramConfig) ProtoMessage() {}

func (x *ProgramConfig) ProtoReflect() protoreflect.Message {
	mi := &file_runme_runner_v2alpha1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramConfig.ProtoReflect.Descriptor instead.
func (*ProgramConfig) Descriptor() ([]byte, []int) {
	return file_runme_runner_v2alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProgramConfig) GetProgramName() string {
	if x != nil {
		return x.ProgramName
	}
	return ""
}

func (x *ProgramConfig) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *ProgramConfig) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

func (x *ProgramConfig) GetLanguageId() string {
	if x != nil {
		return x.LanguageId
	}
	return ""
}

func (x *ProgramConfig) GetBackground() bool {
	if x != nil {
		return x.Background
	}
	return false
}

func (x *ProgramConfig) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

func (x *ProgramConfig) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (m *ProgramConfig) GetSource() isProgramConfig_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *ProgramConfig) GetCommands() *ProgramConfig_CommandList {
	if x, ok := x.GetSource().(*ProgramConfig_Commands); ok {
		return x.Commands
	}
	return nil
}

func (x *ProgramConfig) GetScript() string {
	if x, ok := x.GetSource().(*ProgramConfig_Script); ok {
		return x.Script
	}
	return ""
}

func (x *ProgramConfig) GetInteractive() bool {
	if x != nil {
		return x.Interactive
	}
	return false
}

func (x *ProgramConfig) GetMode() CommandMode {
	if x != nil {
		return x.Mode
	}
	return CommandMode_COMMAND_MODE_UNSPECIFIED
}

func (x *ProgramConfig) GetKnownId() string {
	if x != nil {
		return x.KnownId
	}
	return ""
}

func (x *ProgramConfig) GetKnownName() string {
	if x != nil {
		return x.KnownName
	}
	return ""
}

type isProgramConfig_Source interface {
	isProgramConfig_Source()
}

type ProgramConfig_Commands struct {
	// commands are commands to be executed by the program.
	// The commands are joined and executed as a script.
	Commands *ProgramConfig_CommandList `protobuf:"bytes,8,opt,name=commands,proto3,oneof"`
}

type ProgramConfig_Script struct {
	// script is code to be executed by the program.
	// Individual lines are joined with the new line character.
	Script string `protobuf:"bytes,9,opt,name=script,proto3,oneof"`
}

func (*ProgramConfig_Commands) isProgramConfig_Source() {}

func (*ProgramConfig_Script) isProgramConfig_Source() {}

type ProgramConfig_CommandList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// commands are commands to be executed by the program.
	// The commands are joined and executed as a script.
	// For example: ["echo 'Hello, World'", "ls -l /etc"].
	Items []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ProgramConfig_CommandList) Reset() {
	*x = ProgramConfig_CommandList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runme_runner_v2alpha1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgramConfig_CommandList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgramConfig_CommandList) ProtoMessage() {}

func (x *ProgramConfig_CommandList) ProtoReflect() protoreflect.Message {
	mi := &file_runme_runner_v2alpha1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgramConfig_CommandList.ProtoReflect.Descriptor instead.
func (*ProgramConfig_CommandList) Descriptor() ([]byte, []int) {
	return file_runme_runner_v2alpha1_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProgramConfig_CommandList) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_runme_runner_v2alpha1_config_proto protoreflect.FileDescriptor

var file_runme_runner_v2alpha1_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x95, 0x04, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x4e, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65,
	0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x23, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2a, 0x76, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x49, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4d,
	0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x42, 0x58, 0x5a, 0x56, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66,
	0x75, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x72, 0x75, 0x6e, 0x6d, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runme_runner_v2alpha1_config_proto_rawDescOnce sync.Once
	file_runme_runner_v2alpha1_config_proto_rawDescData = file_runme_runner_v2alpha1_config_proto_rawDesc
)

func file_runme_runner_v2alpha1_config_proto_rawDescGZIP() []byte {
	file_runme_runner_v2alpha1_config_proto_rawDescOnce.Do(func() {
		file_runme_runner_v2alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_runme_runner_v2alpha1_config_proto_rawDescData)
	})
	return file_runme_runner_v2alpha1_config_proto_rawDescData
}

var file_runme_runner_v2alpha1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_runme_runner_v2alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_runme_runner_v2alpha1_config_proto_goTypes = []any{
	(CommandMode)(0),                  // 0: runme.runner.v2alpha1.CommandMode
	(*ProgramConfig)(nil),             // 1: runme.runner.v2alpha1.ProgramConfig
	(*ProgramConfig_CommandList)(nil), // 2: runme.runner.v2alpha1.ProgramConfig.CommandList
}
var file_runme_runner_v2alpha1_config_proto_depIdxs = []int32{
	2, // 0: runme.runner.v2alpha1.ProgramConfig.commands:type_name -> runme.runner.v2alpha1.ProgramConfig.CommandList
	0, // 1: runme.runner.v2alpha1.ProgramConfig.mode:type_name -> runme.runner.v2alpha1.CommandMode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_runme_runner_v2alpha1_config_proto_init() }
func file_runme_runner_v2alpha1_config_proto_init() {
	if File_runme_runner_v2alpha1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runme_runner_v2alpha1_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProgramConfig); i {
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
		file_runme_runner_v2alpha1_config_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProgramConfig_CommandList); i {
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
	file_runme_runner_v2alpha1_config_proto_msgTypes[0].OneofWrappers = []any{
		(*ProgramConfig_Commands)(nil),
		(*ProgramConfig_Script)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_runme_runner_v2alpha1_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_runme_runner_v2alpha1_config_proto_goTypes,
		DependencyIndexes: file_runme_runner_v2alpha1_config_proto_depIdxs,
		EnumInfos:         file_runme_runner_v2alpha1_config_proto_enumTypes,
		MessageInfos:      file_runme_runner_v2alpha1_config_proto_msgTypes,
	}.Build()
	File_runme_runner_v2alpha1_config_proto = out.File
	file_runme_runner_v2alpha1_config_proto_rawDesc = nil
	file_runme_runner_v2alpha1_config_proto_goTypes = nil
	file_runme_runner_v2alpha1_config_proto_depIdxs = nil
}
