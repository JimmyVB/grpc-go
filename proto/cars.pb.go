// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/cars.proto

package go_car_grpc

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

type NewCarParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mark  string `protobuf:"bytes,1,opt,name=mark,proto3" json:"mark,omitempty"`
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Price int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *NewCarParams) Reset() {
	*x = NewCarParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCarParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCarParams) ProtoMessage() {}

func (x *NewCarParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCarParams.ProtoReflect.Descriptor instead.
func (*NewCarParams) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{0}
}

func (x *NewCarParams) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

func (x *NewCarParams) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *NewCarParams) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mark  string `protobuf:"bytes,2,opt,name=mark,proto3" json:"mark,omitempty"`
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Price int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{1}
}

func (x *Car) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Car) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

func (x *Car) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Car) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetCarsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCarsParams) Reset() {
	*x = GetCarsParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarsParams) ProtoMessage() {}

func (x *GetCarsParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarsParams.ProtoReflect.Descriptor instead.
func (*GetCarsParams) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{2}
}

type CarsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*Car `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
}

func (x *CarsList) Reset() {
	*x = CarsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarsList) ProtoMessage() {}

func (x *CarsList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarsList.ProtoReflect.Descriptor instead.
func (*CarsList) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{3}
}

func (x *CarsList) GetCars() []*Car {
	if x != nil {
		return x.Cars
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Data    []*Car `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{4}
}

func (x *Data) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Data) GetData() []*Car {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetOneCarParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOneCarParams) Reset() {
	*x = GetOneCarParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneCarParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneCarParams) ProtoMessage() {}

func (x *GetOneCarParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneCarParams.ProtoReflect.Descriptor instead.
func (*GetOneCarParams) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{5}
}

func (x *GetOneCarParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateCarParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mark  string `protobuf:"bytes,2,opt,name=mark,proto3" json:"mark,omitempty"`
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Price int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *UpdateCarParams) Reset() {
	*x = UpdateCarParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarParams) ProtoMessage() {}

func (x *UpdateCarParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarParams.ProtoReflect.Descriptor instead.
func (*UpdateCarParams) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCarParams) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCarParams) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

func (x *UpdateCarParams) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *UpdateCarParams) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_proto_cars_proto protoreflect.FileDescriptor

var file_proto_cars_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x4e, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x43,
	0x61, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x55, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x22, 0x29, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04,
	0x63, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x04, 0x63, 0x61, 0x72, 0x73, 0x22, 0x3f, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x21, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x61, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x61, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x32, 0xd0, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x77, 0x43, 0x61, 0x72, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65, 0x77, 0x43,
	0x61, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x61, 0x72, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73,
	0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x61, 0x72,
	0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x43, 0x61,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x61, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x61, 0x72, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x3b, 0x67, 0x6f, 0x5f, 0x63, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cars_proto_rawDescOnce sync.Once
	file_proto_cars_proto_rawDescData = file_proto_cars_proto_rawDesc
)

func file_proto_cars_proto_rawDescGZIP() []byte {
	file_proto_cars_proto_rawDescOnce.Do(func() {
		file_proto_cars_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cars_proto_rawDescData)
	})
	return file_proto_cars_proto_rawDescData
}

var file_proto_cars_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_cars_proto_goTypes = []interface{}{
	(*NewCarParams)(nil),    // 0: grpc.NewCarParams
	(*Car)(nil),             // 1: grpc.Car
	(*GetCarsParams)(nil),   // 2: grpc.GetCarsParams
	(*CarsList)(nil),        // 3: grpc.CarsList
	(*Data)(nil),            // 4: grpc.Data
	(*GetOneCarParams)(nil), // 5: grpc.GetOneCarParams
	(*UpdateCarParams)(nil), // 6: grpc.UpdateCarParams
}
var file_proto_cars_proto_depIdxs = []int32{
	1, // 0: grpc.CarsList.cars:type_name -> grpc.Car
	1, // 1: grpc.Data.data:type_name -> grpc.Car
	0, // 2: grpc.CarManagement.CreateNewCar:input_type -> grpc.NewCarParams
	2, // 3: grpc.CarManagement.GetCars:input_type -> grpc.GetCarsParams
	5, // 4: grpc.CarManagement.GetOneCar:input_type -> grpc.GetOneCarParams
	6, // 5: grpc.CarManagement.UpdateCar:input_type -> grpc.UpdateCarParams
	1, // 6: grpc.CarManagement.CreateNewCar:output_type -> grpc.Car
	4, // 7: grpc.CarManagement.GetCars:output_type -> grpc.Data
	1, // 8: grpc.CarManagement.GetOneCar:output_type -> grpc.Car
	1, // 9: grpc.CarManagement.UpdateCar:output_type -> grpc.Car
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_cars_proto_init() }
func file_proto_cars_proto_init() {
	if File_proto_cars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCarParams); i {
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
		file_proto_cars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_proto_cars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarsParams); i {
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
		file_proto_cars_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarsList); i {
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
		file_proto_cars_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_proto_cars_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneCarParams); i {
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
		file_proto_cars_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarParams); i {
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
			RawDescriptor: file_proto_cars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cars_proto_goTypes,
		DependencyIndexes: file_proto_cars_proto_depIdxs,
		MessageInfos:      file_proto_cars_proto_msgTypes,
	}.Build()
	File_proto_cars_proto = out.File
	file_proto_cars_proto_rawDesc = nil
	file_proto_cars_proto_goTypes = nil
	file_proto_cars_proto_depIdxs = nil
}
