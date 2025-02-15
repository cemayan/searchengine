// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: protos/event/event.proto

package event

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

type EntityType int32

const (
	EntityType_Record         EntityType = 0
	EntityType_RecordMetadata EntityType = 1
)

// Enum value maps for EntityType.
var (
	EntityType_name = map[int32]string{
		0: "Record",
		1: "RecordMetadata",
	}
	EntityType_value = map[string]int32{
		"Record":         0,
		"RecordMetadata": 1,
	}
)

func (x EntityType) Enum() *EntityType {
	p := new(EntityType)
	*p = x
	return p
}

func (x EntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_event_event_proto_enumTypes[0].Descriptor()
}

func (EntityType) Type() protoreflect.EnumType {
	return &file_protos_event_event_proto_enumTypes[0]
}

func (x EntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityType.Descriptor instead.
func (EntityType) EnumDescriptor() ([]byte, []int) {
	return file_protos_event_event_proto_rawDescGZIP(), []int{0}
}

type EventType int32

const (
	EventType_RECORD_CREATED         EventType = 0
	EventType_RECORDMETADATA_CREATED EventType = 1
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "RECORD_CREATED",
		1: "RECORDMETADATA_CREATED",
	}
	EventType_value = map[string]int32{
		"RECORD_CREATED":         0,
		"RECORDMETADATA_CREATED": 1,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_event_event_proto_enumTypes[1].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_protos_event_event_proto_enumTypes[1]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_protos_event_event_proto_rawDescGZIP(), []int{1}
}

type Db struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Db) Reset() {
	*x = Db{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Db) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Db) ProtoMessage() {}

func (x *Db) ProtoReflect() protoreflect.Message {
	mi := &file_protos_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Db.ProtoReflect.Descriptor instead.
func (*Db) Descriptor() ([]byte, []int) {
	return file_protos_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *Db) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Db) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SEError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DbName string `protobuf:"bytes,1,opt,name=dbName,proto3" json:"dbName,omitempty"`
	Kind   string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Key    string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Value  string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Date   int64  `protobuf:"varint,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *SEError) Reset() {
	*x = SEError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SEError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SEError) ProtoMessage() {}

func (x *SEError) ProtoReflect() protoreflect.Message {
	mi := &file_protos_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SEError.ProtoReflect.Descriptor instead.
func (*SEError) Descriptor() ([]byte, []int) {
	return file_protos_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *SEError) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *SEError) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *SEError) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SEError) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SEError) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SEError) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

// The request message containing the user's name.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type       EventType  `protobuf:"varint,2,opt,name=type,proto3,enum=protos.EventType" json:"type,omitempty"`
	EntityType EntityType `protobuf:"varint,3,opt,name=entityType,proto3,enum=protos.EntityType" json:"entityType,omitempty"`
	Date       int64      `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	Data       []byte     `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_event_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_protos_event_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_protos_event_event_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_RECORD_CREATED
}

func (x *Event) GetEntityType() EntityType {
	if x != nil {
		return x.EntityType
	}
	return EntityType_Record
}

func (x *Event) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Event) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protos_event_event_proto protoreflect.FileDescriptor

var file_protos_event_event_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x22, 0x2c, 0x0a, 0x02, 0x44, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x87, 0x01, 0x0a, 0x07, 0x53, 0x45, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x2c, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x10, 0x01, 0x2a, 0x3b, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x32, 0x3b, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65,
	0x6d, 0x61, 0x79, 0x61, 0x6e, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_event_event_proto_rawDescOnce sync.Once
	file_protos_event_event_proto_rawDescData = file_protos_event_event_proto_rawDesc
)

func file_protos_event_event_proto_rawDescGZIP() []byte {
	file_protos_event_event_proto_rawDescOnce.Do(func() {
		file_protos_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_event_event_proto_rawDescData)
	})
	return file_protos_event_event_proto_rawDescData
}

var file_protos_event_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protos_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_event_event_proto_goTypes = []interface{}{
	(EntityType)(0), // 0: protos.EntityType
	(EventType)(0),  // 1: protos.EventType
	(*Db)(nil),      // 2: protos.Db
	(*SEError)(nil), // 3: protos.SEError
	(*Event)(nil),   // 4: protos.Event
}
var file_protos_event_event_proto_depIdxs = []int32{
	1, // 0: protos.Event.type:type_name -> protos.EventType
	0, // 1: protos.Event.entityType:type_name -> protos.EntityType
	4, // 2: protos.EventHandler.SendEvent:input_type -> protos.Event
	4, // 3: protos.EventHandler.SendEvent:output_type -> protos.Event
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_event_event_proto_init() }
func file_protos_event_event_proto_init() {
	if File_protos_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Db); i {
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
		file_protos_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SEError); i {
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
		file_protos_event_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_protos_event_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_event_event_proto_goTypes,
		DependencyIndexes: file_protos_event_event_proto_depIdxs,
		EnumInfos:         file_protos_event_event_proto_enumTypes,
		MessageInfos:      file_protos_event_event_proto_msgTypes,
	}.Build()
	File_protos_event_event_proto = out.File
	file_protos_event_event_proto_rawDesc = nil
	file_protos_event_event_proto_goTypes = nil
	file_protos_event_event_proto_depIdxs = nil
}
