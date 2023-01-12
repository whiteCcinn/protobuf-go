// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Wrappers for primitive (non-message) types. These types are useful
// for embedding primitives in the `google.protobuf.Any` type and for places
// where we need to distinguish between the absence of a primitive
// typed field and its default value.
//
// These wrappers have no meaningful use within repeated fields as they lack
// the ability to detect presence on individual elements.
// These wrappers have no meaningful use within a map or a oneof since
// individual entries of a map or fields of a oneof can already detect presence.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/wrappers.proto

package wrapperspb

import (
	protoreflect "github.com/whiteCcinn/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/whiteCcinn/protobuf-go/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// Wrapper message for `double`.
//
// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The double value.
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

// Double stores v in a new DoubleValue and returns a pointer to it.
func Double(v float64) *DoubleValue {
	return &DoubleValue{Value: v}
}

func (x *DoubleValue) Reset() {
	*x = DoubleValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleValue) ProtoMessage() {}

func (x *DoubleValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleValue.ProtoReflect.Descriptor instead.
func (*DoubleValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{0}
}

func (x *DoubleValue) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
type FloatValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The float value.
	Value float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

// Float stores v in a new FloatValue and returns a pointer to it.
func Float(v float32) *FloatValue {
	return &FloatValue{Value: v}
}

func (x *FloatValue) Reset() {
	*x = FloatValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatValue) ProtoMessage() {}

func (x *FloatValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatValue.ProtoReflect.Descriptor instead.
func (*FloatValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{1}
}

func (x *FloatValue) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
type Int64Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The int64 value.
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

// Int64 stores v in a new Int64Value and returns a pointer to it.
func Int64(v int64) *Int64Value {
	return &Int64Value{Value: v}
}

func (x *Int64Value) Reset() {
	*x = Int64Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Value) ProtoMessage() {}

func (x *Int64Value) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Value.ProtoReflect.Descriptor instead.
func (*Int64Value) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{2}
}

func (x *Int64Value) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
type UInt64Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uint64 value.
	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

// UInt64 stores v in a new UInt64Value and returns a pointer to it.
func UInt64(v uint64) *UInt64Value {
	return &UInt64Value{Value: v}
}

func (x *UInt64Value) Reset() {
	*x = UInt64Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt64Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt64Value) ProtoMessage() {}

func (x *UInt64Value) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt64Value.ProtoReflect.Descriptor instead.
func (*UInt64Value) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{3}
}

func (x *UInt64Value) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `int32`.
//
// The JSON representation for `Int32Value` is JSON number.
type Int32Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The int32 value.
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

// Int32 stores v in a new Int32Value and returns a pointer to it.
func Int32(v int32) *Int32Value {
	return &Int32Value{Value: v}
}

func (x *Int32Value) Reset() {
	*x = Int32Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Value) ProtoMessage() {}

func (x *Int32Value) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Value.ProtoReflect.Descriptor instead.
func (*Int32Value) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{4}
}

func (x *Int32Value) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `uint32`.
//
// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uint32 value.
	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

// UInt32 stores v in a new UInt32Value and returns a pointer to it.
func UInt32(v uint32) *UInt32Value {
	return &UInt32Value{Value: v}
}

func (x *UInt32Value) Reset() {
	*x = UInt32Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt32Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt32Value) ProtoMessage() {}

func (x *UInt32Value) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt32Value.ProtoReflect.Descriptor instead.
func (*UInt32Value) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{5}
}

func (x *UInt32Value) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
type BoolValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bool value.
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

// Bool stores v in a new BoolValue and returns a pointer to it.
func Bool(v bool) *BoolValue {
	return &BoolValue{Value: v}
}

func (x *BoolValue) Reset() {
	*x = BoolValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolValue) ProtoMessage() {}

func (x *BoolValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolValue.ProtoReflect.Descriptor instead.
func (*BoolValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{6}
}

func (x *BoolValue) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

// Wrapper message for `string`.
//
// The JSON representation for `StringValue` is JSON string.
type StringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The string value.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

// String stores v in a new StringValue and returns a pointer to it.
func String(v string) *StringValue {
	return &StringValue{Value: v}
}

func (x *StringValue) Reset() {
	*x = StringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringValue) ProtoMessage() {}

func (x *StringValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringValue.ProtoReflect.Descriptor instead.
func (*StringValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{7}
}

func (x *StringValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Wrapper message for `bytes`.
//
// The JSON representation for `BytesValue` is JSON string.
type BytesValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes value.
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

// Bytes stores v in a new BytesValue and returns a pointer to it.
func Bytes(v []byte) *BytesValue {
	return &BytesValue{Value: v}
}

func (x *BytesValue) Reset() {
	*x = BytesValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_wrappers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesValue) ProtoMessage() {}

func (x *BytesValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_wrappers_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesValue.ProtoReflect.Descriptor instead.
func (*BytesValue) Descriptor() ([]byte, []int) {
	return file_google_protobuf_wrappers_proto_rawDescGZIP(), []int{8}
}

func (x *BytesValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_google_protobuf_wrappers_proto protoreflect.FileDescriptor

var file_google_protobuf_wrappers_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23,
	0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x23, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x83, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x42, 0x0d, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_wrappers_proto_rawDescOnce sync.Once
	file_google_protobuf_wrappers_proto_rawDescData = file_google_protobuf_wrappers_proto_rawDesc
)

func file_google_protobuf_wrappers_proto_rawDescGZIP() []byte {
	file_google_protobuf_wrappers_proto_rawDescOnce.Do(func() {
		file_google_protobuf_wrappers_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_wrappers_proto_rawDescData)
	})
	return file_google_protobuf_wrappers_proto_rawDescData
}

var file_google_protobuf_wrappers_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_google_protobuf_wrappers_proto_goTypes = []interface{}{
	(*DoubleValue)(nil), // 0: google.protobuf.DoubleValue
	(*FloatValue)(nil),  // 1: google.protobuf.FloatValue
	(*Int64Value)(nil),  // 2: google.protobuf.Int64Value
	(*UInt64Value)(nil), // 3: google.protobuf.UInt64Value
	(*Int32Value)(nil),  // 4: google.protobuf.Int32Value
	(*UInt32Value)(nil), // 5: google.protobuf.UInt32Value
	(*BoolValue)(nil),   // 6: google.protobuf.BoolValue
	(*StringValue)(nil), // 7: google.protobuf.StringValue
	(*BytesValue)(nil),  // 8: google.protobuf.BytesValue
}
var file_google_protobuf_wrappers_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_protobuf_wrappers_proto_init() }
func file_google_protobuf_wrappers_proto_init() {
	if File_google_protobuf_wrappers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_protobuf_wrappers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleValue); i {
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
		file_google_protobuf_wrappers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloatValue); i {
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
		file_google_protobuf_wrappers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Value); i {
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
		file_google_protobuf_wrappers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt64Value); i {
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
		file_google_protobuf_wrappers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32Value); i {
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
		file_google_protobuf_wrappers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt32Value); i {
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
		file_google_protobuf_wrappers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolValue); i {
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
		file_google_protobuf_wrappers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringValue); i {
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
		file_google_protobuf_wrappers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesValue); i {
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
			RawDescriptor: file_google_protobuf_wrappers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_wrappers_proto_goTypes,
		DependencyIndexes: file_google_protobuf_wrappers_proto_depIdxs,
		MessageInfos:      file_google_protobuf_wrappers_proto_msgTypes,
	}.Build()
	File_google_protobuf_wrappers_proto = out.File
	file_google_protobuf_wrappers_proto_rawDesc = nil
	file_google_protobuf_wrappers_proto_goTypes = nil
	file_google_protobuf_wrappers_proto_depIdxs = nil
}
