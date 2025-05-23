// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: user_service.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified bool                   `protobuf:"varint,5,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	IsGuest       bool                   `protobuf:"varint,6,opt,name=is_guest,json=isGuest,proto3" json:"is_guest,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetIsGuest() bool {
	if x != nil {
		return x.IsGuest
	}
	return false
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetUserByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIDRequest) Reset() {
	*x = GetUserByIDRequest{}
	mi := &file_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDRequest) ProtoMessage() {}

func (x *GetUserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIDRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserByIDResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIDResponse) Reset() {
	*x = GetUserByIDResponse{}
	mi := &file_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDResponse) ProtoMessage() {}

func (x *GetUserByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDResponse.ProtoReflect.Descriptor instead.
func (*GetUserByIDResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByIDResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserByEmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByEmailRequest) Reset() {
	*x = GetUserByEmailRequest{}
	mi := &file_user_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByEmailRequest) ProtoMessage() {}

func (x *GetUserByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByEmailRequest.ProtoReflect.Descriptor instead.
func (*GetUserByEmailRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetUserByEmailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByEmailResponse) Reset() {
	*x = GetUserByEmailResponse{}
	mi := &file_user_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByEmailResponse) ProtoMessage() {}

func (x *GetUserByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByEmailResponse.ProtoReflect.Descriptor instead.
func (*GetUserByEmailResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserByEmailResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_user_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_user_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{6}
}

var File_user_service_proto protoreflect.FileDescriptor

const file_user_service_proto_rawDesc = "" +
	"\n" +
	"\x12user_service.proto\x12\busers.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\x9c\x02\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12%\n" +
	"\x0eemail_verified\x18\x05 \x01(\bR\remailVerified\x12\x19\n" +
	"\bis_guest\x18\x06 \x01(\bR\aisGuest\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"$\n" +
	"\x12GetUserByIDRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"9\n" +
	"\x13GetUserByIDResponse\x12\"\n" +
	"\x04user\x18\x01 \x01(\v2\x0e.users.v1.UserR\x04user\"-\n" +
	"\x15GetUserByEmailRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\"<\n" +
	"\x16GetUserByEmailResponse\x12\"\n" +
	"\x04user\x18\x01 \x01(\v2\x0e.users.v1.UserR\x04user\"7\n" +
	"\x11CreateUserRequest\x12\"\n" +
	"\x04user\x18\x01 \x01(\v2\x0e.users.v1.UserR\x04user\"\x14\n" +
	"\x12CreateUserResponse2\xf7\x01\n" +
	"\vUserService\x12J\n" +
	"\vGetUserByID\x12\x1c.users.v1.GetUserByIDRequest\x1a\x1d.users.v1.GetUserByIDResponse\x12S\n" +
	"\x0eGetUserByEmail\x12\x1f.users.v1.GetUserByEmailRequest\x1a .users.v1.GetUserByEmailResponse\x12G\n" +
	"\n" +
	"CreateUser\x12\x1b.users.v1.CreateUserRequest\x1a\x1c.users.v1.CreateUserResponseB\x06Z\x04/genb\x06proto3"

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData []byte
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_service_proto_rawDesc), len(file_user_service_proto_rawDesc)))
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_service_proto_goTypes = []any{
	(*User)(nil),                   // 0: users.v1.User
	(*GetUserByIDRequest)(nil),     // 1: users.v1.GetUserByIDRequest
	(*GetUserByIDResponse)(nil),    // 2: users.v1.GetUserByIDResponse
	(*GetUserByEmailRequest)(nil),  // 3: users.v1.GetUserByEmailRequest
	(*GetUserByEmailResponse)(nil), // 4: users.v1.GetUserByEmailResponse
	(*CreateUserRequest)(nil),      // 5: users.v1.CreateUserRequest
	(*CreateUserResponse)(nil),     // 6: users.v1.CreateUserResponse
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
}
var file_user_service_proto_depIdxs = []int32{
	7, // 0: users.v1.User.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: users.v1.User.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: users.v1.GetUserByIDResponse.user:type_name -> users.v1.User
	0, // 3: users.v1.GetUserByEmailResponse.user:type_name -> users.v1.User
	0, // 4: users.v1.CreateUserRequest.user:type_name -> users.v1.User
	1, // 5: users.v1.UserService.GetUserByID:input_type -> users.v1.GetUserByIDRequest
	3, // 6: users.v1.UserService.GetUserByEmail:input_type -> users.v1.GetUserByEmailRequest
	5, // 7: users.v1.UserService.CreateUser:input_type -> users.v1.CreateUserRequest
	2, // 8: users.v1.UserService.GetUserByID:output_type -> users.v1.GetUserByIDResponse
	4, // 9: users.v1.UserService.GetUserByEmail:output_type -> users.v1.GetUserByEmailResponse
	6, // 10: users.v1.UserService.CreateUser:output_type -> users.v1.CreateUserResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_service_proto_rawDesc), len(file_user_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
