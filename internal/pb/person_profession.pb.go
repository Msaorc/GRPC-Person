// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: proto/person_profession.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Blank struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blank) Reset() {
	*x = Blank{}
	mi := &file_proto_person_profession_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blank) ProtoMessage() {}

func (x *Blank) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blank.ProtoReflect.Descriptor instead.
func (*Blank) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{0}
}

type Person struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Year          int32                  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_proto_person_profession_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type Profession struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Profession) Reset() {
	*x = Profession{}
	mi := &file_proto_person_profession_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Profession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profession) ProtoMessage() {}

func (x *Profession) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profession.ProtoReflect.Descriptor instead.
func (*Profession) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{2}
}

func (x *Profession) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Profession) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreatePersonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Year          int32                  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePersonRequest) Reset() {
	*x = CreatePersonRequest{}
	mi := &file_proto_person_profession_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonRequest) ProtoMessage() {}

func (x *CreatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonRequest.ProtoReflect.Descriptor instead.
func (*CreatePersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePersonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePersonRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type PersonResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Person        *Person                `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PersonResponse) Reset() {
	*x = PersonResponse{}
	mi := &file_proto_person_profession_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonResponse) ProtoMessage() {}

func (x *PersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonResponse.ProtoReflect.Descriptor instead.
func (*PersonResponse) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{4}
}

func (x *PersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type CreateProfessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Description   string                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProfessionRequest) Reset() {
	*x = CreateProfessionRequest{}
	mi := &file_proto_person_profession_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProfessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfessionRequest) ProtoMessage() {}

func (x *CreateProfessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfessionRequest.ProtoReflect.Descriptor instead.
func (*CreateProfessionRequest) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{5}
}

func (x *CreateProfessionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ProfessionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Profession    *Profession            `protobuf:"bytes,1,opt,name=profession,proto3" json:"profession,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfessionResponse) Reset() {
	*x = ProfessionResponse{}
	mi := &file_proto_person_profession_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfessionResponse) ProtoMessage() {}

func (x *ProfessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfessionResponse.ProtoReflect.Descriptor instead.
func (*ProfessionResponse) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{6}
}

func (x *ProfessionResponse) GetProfession() *Profession {
	if x != nil {
		return x.Profession
	}
	return nil
}

type PersonList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	People        []*Person              `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PersonList) Reset() {
	*x = PersonList{}
	mi := &file_proto_person_profession_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonList) ProtoMessage() {}

func (x *PersonList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonList.ProtoReflect.Descriptor instead.
func (*PersonList) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{7}
}

func (x *PersonList) GetPeople() []*Person {
	if x != nil {
		return x.People
	}
	return nil
}

type ProfessionList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Professions   []*Profession          `protobuf:"bytes,1,rep,name=professions,proto3" json:"professions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfessionList) Reset() {
	*x = ProfessionList{}
	mi := &file_proto_person_profession_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfessionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfessionList) ProtoMessage() {}

func (x *ProfessionList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfessionList.ProtoReflect.Descriptor instead.
func (*ProfessionList) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{8}
}

func (x *ProfessionList) GetProfessions() []*Profession {
	if x != nil {
		return x.Professions
	}
	return nil
}

type PersonGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PersonGetRequest) Reset() {
	*x = PersonGetRequest{}
	mi := &file_proto_person_profession_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonGetRequest) ProtoMessage() {}

func (x *PersonGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonGetRequest.ProtoReflect.Descriptor instead.
func (*PersonGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{9}
}

func (x *PersonGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProfessionGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfessionGetRequest) Reset() {
	*x = ProfessionGetRequest{}
	mi := &file_proto_person_profession_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfessionGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfessionGetRequest) ProtoMessage() {}

func (x *ProfessionGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_person_profession_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfessionGetRequest.ProtoReflect.Descriptor instead.
func (*ProfessionGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_person_profession_proto_rawDescGZIP(), []int{10}
}

func (x *ProfessionGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_person_profession_proto protoreflect.FileDescriptor

var file_proto_person_profession_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x22, 0x40, 0x0a, 0x06,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x3e,
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x34, 0x0a,
	0x0e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x44, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x22, 0x0a, 0x10,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x26, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xb3, 0x02, 0x0a, 0x0d, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x4c, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x29, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x1a, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x00, 0x32, 0xef,
	0x02, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x12, 0x58, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x31, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x1a, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_person_profession_proto_rawDescOnce sync.Once
	file_proto_person_profession_proto_rawDescData []byte
)

func file_proto_person_profession_proto_rawDescGZIP() []byte {
	file_proto_person_profession_proto_rawDescOnce.Do(func() {
		file_proto_person_profession_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_person_profession_proto_rawDesc), len(file_proto_person_profession_proto_rawDesc)))
	})
	return file_proto_person_profession_proto_rawDescData
}

var file_proto_person_profession_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_person_profession_proto_goTypes = []any{
	(*Blank)(nil),                   // 0: pb.blank
	(*Person)(nil),                  // 1: pb.Person
	(*Profession)(nil),              // 2: pb.Profession
	(*CreatePersonRequest)(nil),     // 3: pb.CreatePersonRequest
	(*PersonResponse)(nil),          // 4: pb.PersonResponse
	(*CreateProfessionRequest)(nil), // 5: pb.CreateProfessionRequest
	(*ProfessionResponse)(nil),      // 6: pb.ProfessionResponse
	(*PersonList)(nil),              // 7: pb.PersonList
	(*ProfessionList)(nil),          // 8: pb.ProfessionList
	(*PersonGetRequest)(nil),        // 9: pb.PersonGetRequest
	(*ProfessionGetRequest)(nil),    // 10: pb.ProfessionGetRequest
}
var file_proto_person_profession_proto_depIdxs = []int32{
	1,  // 0: pb.PersonResponse.person:type_name -> pb.Person
	2,  // 1: pb.ProfessionResponse.profession:type_name -> pb.Profession
	1,  // 2: pb.PersonList.people:type_name -> pb.Person
	2,  // 3: pb.ProfessionList.professions:type_name -> pb.Profession
	3,  // 4: pb.PersonService.CreatePerson:input_type -> pb.CreatePersonRequest
	3,  // 5: pb.PersonService.CreatePersonStream:input_type -> pb.CreatePersonRequest
	3,  // 6: pb.PersonService.CreatePersonStreamBidirectional:input_type -> pb.CreatePersonRequest
	0,  // 7: pb.PersonService.ListPerson:input_type -> pb.blank
	9,  // 8: pb.PersonService.GetPerson:input_type -> pb.PersonGetRequest
	5,  // 9: pb.ProfessionService.CreateProfession:input_type -> pb.CreateProfessionRequest
	5,  // 10: pb.ProfessionService.CreateProfessionStream:input_type -> pb.CreateProfessionRequest
	5,  // 11: pb.ProfessionService.CreateProfessionStreamBidirectional:input_type -> pb.CreateProfessionRequest
	0,  // 12: pb.ProfessionService.ListProfession:input_type -> pb.blank
	10, // 13: pb.ProfessionService.GetProfession:input_type -> pb.ProfessionGetRequest
	1,  // 14: pb.PersonService.CreatePerson:output_type -> pb.Person
	7,  // 15: pb.PersonService.CreatePersonStream:output_type -> pb.PersonList
	1,  // 16: pb.PersonService.CreatePersonStreamBidirectional:output_type -> pb.Person
	7,  // 17: pb.PersonService.ListPerson:output_type -> pb.PersonList
	1,  // 18: pb.PersonService.GetPerson:output_type -> pb.Person
	2,  // 19: pb.ProfessionService.CreateProfession:output_type -> pb.Profession
	8,  // 20: pb.ProfessionService.CreateProfessionStream:output_type -> pb.ProfessionList
	2,  // 21: pb.ProfessionService.CreateProfessionStreamBidirectional:output_type -> pb.Profession
	8,  // 22: pb.ProfessionService.ListProfession:output_type -> pb.ProfessionList
	2,  // 23: pb.ProfessionService.GetProfession:output_type -> pb.Profession
	14, // [14:24] is the sub-list for method output_type
	4,  // [4:14] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_person_profession_proto_init() }
func file_proto_person_profession_proto_init() {
	if File_proto_person_profession_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_person_profession_proto_rawDesc), len(file_proto_person_profession_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_person_profession_proto_goTypes,
		DependencyIndexes: file_proto_person_profession_proto_depIdxs,
		MessageInfos:      file_proto_person_profession_proto_msgTypes,
	}.Build()
	File_proto_person_profession_proto = out.File
	file_proto_person_profession_proto_goTypes = nil
	file_proto_person_profession_proto_depIdxs = nil
}
