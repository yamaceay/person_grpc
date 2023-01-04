// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: person/person.proto

package person

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

// Choice is an enum type
type Person_Choice int32

const (
	// Default choice
	Person_DEFAULT Person_Choice = 0
	// Other choices
	Person_ONE   Person_Choice = 1
	Person_TWO   Person_Choice = 2
	Person_THREE Person_Choice = 3
)

// Enum value maps for Person_Choice.
var (
	Person_Choice_name = map[int32]string{
		0: "DEFAULT",
		1: "ONE",
		2: "TWO",
		3: "THREE",
	}
	Person_Choice_value = map[string]int32{
		"DEFAULT": 0,
		"ONE":     1,
		"TWO":     2,
		"THREE":   3,
	}
)

func (x Person_Choice) Enum() *Person_Choice {
	p := new(Person_Choice)
	*p = x
	return p
}

func (x Person_Choice) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_Choice) Descriptor() protoreflect.EnumDescriptor {
	return file_person_person_proto_enumTypes[0].Descriptor()
}

func (Person_Choice) Type() protoreflect.EnumType {
	return &file_person_person_proto_enumTypes[0]
}

func (x Person_Choice) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_Choice.Descriptor instead.
func (Person_Choice) EnumDescriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{1, 0}
}

// PersonKey is a struct with a single field
type PersonKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// idHash is a string field
	IdHash string `protobuf:"bytes,1,opt,name=idHash,proto3" json:"idHash,omitempty"`
}

func (x *PersonKey) Reset() {
	*x = PersonKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonKey) ProtoMessage() {}

func (x *PersonKey) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonKey.ProtoReflect.Descriptor instead.
func (*PersonKey) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{0}
}

func (x *PersonKey) GetIdHash() string {
	if x != nil {
		return x.IdHash
	}
	return ""
}

// Person is a struct with numerous fields
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Singular fields of different types
	IsMale bool    `protobuf:"varint,1,opt,name=isMale,proto3" json:"isMale,omitempty"`
	Age    int32   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Weight float32 `protobuf:"fixed32,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Height float64 `protobuf:"fixed64,4,opt,name=height,proto3" json:"height,omitempty"`
	Name   string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Optional fields
	Partner *Person        `protobuf:"bytes,6,opt,name=partner,proto3,oneof" json:"partner,omitempty"`
	Choice  *Person_Choice `protobuf:"varint,7,opt,name=choice,proto3,enum=person.Person_Choice,oneof" json:"choice,omitempty"`
	// occupation is an union of following fields
	//
	// Types that are assignable to Occupation:
	//
	//	*Person_JobTitle
	//	*Person_StudyField
	Occupation isPerson_Occupation `protobuf_oneof:"occupation"`
	// addresses contains key-value pairs of string type
	Addresses map[string]string `protobuf:"bytes,10,rep,name=addresses,proto3" json:"addresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// childs is an array of childs
	Childs []*Person_Child `protobuf:"bytes,11,rep,name=childs,proto3" json:"childs,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_person_person_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetIsMale() bool {
	if x != nil {
		return x.IsMale
	}
	return false
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Person) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetPartner() *Person {
	if x != nil {
		return x.Partner
	}
	return nil
}

func (x *Person) GetChoice() Person_Choice {
	if x != nil && x.Choice != nil {
		return *x.Choice
	}
	return Person_DEFAULT
}

func (m *Person) GetOccupation() isPerson_Occupation {
	if m != nil {
		return m.Occupation
	}
	return nil
}

func (x *Person) GetJobTitle() string {
	if x, ok := x.GetOccupation().(*Person_JobTitle); ok {
		return x.JobTitle
	}
	return ""
}

func (x *Person) GetStudyField() string {
	if x, ok := x.GetOccupation().(*Person_StudyField); ok {
		return x.StudyField
	}
	return ""
}

func (x *Person) GetAddresses() map[string]string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *Person) GetChilds() []*Person_Child {
	if x != nil {
		return x.Childs
	}
	return nil
}

type isPerson_Occupation interface {
	isPerson_Occupation()
}

type Person_JobTitle struct {
	JobTitle string `protobuf:"bytes,8,opt,name=jobTitle,proto3,oneof"`
}

type Person_StudyField struct {
	StudyField string `protobuf:"bytes,9,opt,name=studyField,proto3,oneof"`
}

func (*Person_JobTitle) isPerson_Occupation() {}

func (*Person_StudyField) isPerson_Occupation() {}

// Child is an inner struct extending Person
type Person_Child struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base class
	Self *Person `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	// Other fields
	Father *Person `protobuf:"bytes,2,opt,name=father,proto3" json:"father,omitempty"`
	Mother *Person `protobuf:"bytes,3,opt,name=mother,proto3" json:"mother,omitempty"`
}

func (x *Person_Child) Reset() {
	*x = Person_Child{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Child) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Child) ProtoMessage() {}

func (x *Person_Child) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Child.ProtoReflect.Descriptor instead.
func (*Person_Child) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Person_Child) GetSelf() *Person {
	if x != nil {
		return x.Self
	}
	return nil
}

func (x *Person_Child) GetFather() *Person {
	if x != nil {
		return x.Father
	}
	return nil
}

func (x *Person_Child) GetMother() *Person {
	if x != nil {
		return x.Mother
	}
	return nil
}

var File_person_person_proto protoreflect.FileDescriptor

var file_person_person_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x23, 0x0a,
	0x09, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x48, 0x61,
	0x73, 0x68, 0x22, 0xa4, 0x05, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x73, 0x4d, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x4d, 0x61, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x07,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x48, 0x02, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c,
	0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0a,
	0x73, 0x74, 0x75, 0x64, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3b,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x69, 0x6c,
	0x64, 0x52, 0x06, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x1a, 0x7b, 0x0a, 0x05, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06,
	0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x1a, 0x3c, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x06, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x54, 0x48, 0x52, 0x45, 0x45, 0x10, 0x03, 0x22, 0x04, 0x08, 0x04, 0x10, 0x04, 0x22, 0x04,
	0x08, 0x06, 0x10, 0x08, 0x42, 0x0c, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x80, 0x02, 0x0a, 0x0b, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x28, 0x01, 0x12, 0x3d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x28, 0x01, 0x30, 0x01, 0x42, 0x1c, 0x5a, 0x1a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x64,
	0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_person_person_proto_rawDescOnce sync.Once
	file_person_person_proto_rawDescData = file_person_person_proto_rawDesc
)

func file_person_person_proto_rawDescGZIP() []byte {
	file_person_person_proto_rawDescOnce.Do(func() {
		file_person_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_person_proto_rawDescData)
	})
	return file_person_person_proto_rawDescData
}

var file_person_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_person_person_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_person_person_proto_goTypes = []interface{}{
	(Person_Choice)(0),   // 0: person.Person.Choice
	(*PersonKey)(nil),    // 1: person.PersonKey
	(*Person)(nil),       // 2: person.Person
	(*Person_Child)(nil), // 3: person.Person.Child
	nil,                  // 4: person.Person.AddressesEntry
}
var file_person_person_proto_depIdxs = []int32{
	2,  // 0: person.Person.partner:type_name -> person.Person
	0,  // 1: person.Person.choice:type_name -> person.Person.Choice
	4,  // 2: person.Person.addresses:type_name -> person.Person.AddressesEntry
	3,  // 3: person.Person.childs:type_name -> person.Person.Child
	2,  // 4: person.Person.Child.self:type_name -> person.Person
	2,  // 5: person.Person.Child.father:type_name -> person.Person
	2,  // 6: person.Person.Child.mother:type_name -> person.Person
	1,  // 7: person.PersonQuery.GetPerson:input_type -> person.PersonKey
	1,  // 8: person.PersonQuery.GetPersonBlockedInput:input_type -> person.PersonKey
	1,  // 9: person.PersonQuery.GetPersonBlockedOutput:input_type -> person.PersonKey
	1,  // 10: person.PersonQuery.GetPersonBlockedInputOutput:input_type -> person.PersonKey
	2,  // 11: person.PersonQuery.GetPerson:output_type -> person.Person
	2,  // 12: person.PersonQuery.GetPersonBlockedInput:output_type -> person.Person
	2,  // 13: person.PersonQuery.GetPersonBlockedOutput:output_type -> person.Person
	2,  // 14: person.PersonQuery.GetPersonBlockedInputOutput:output_type -> person.Person
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_person_person_proto_init() }
func file_person_person_proto_init() {
	if File_person_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonKey); i {
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
		file_person_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_person_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Child); i {
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
	file_person_person_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Person_JobTitle)(nil),
		(*Person_StudyField)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_person_person_proto_goTypes,
		DependencyIndexes: file_person_person_proto_depIdxs,
		EnumInfos:         file_person_person_proto_enumTypes,
		MessageInfos:      file_person_person_proto_msgTypes,
	}.Build()
	File_person_person_proto = out.File
	file_person_person_proto_rawDesc = nil
	file_person_person_proto_goTypes = nil
	file_person_person_proto_depIdxs = nil
}
