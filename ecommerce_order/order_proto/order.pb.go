// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: order_proto/order.proto

package ecommerce_order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_order_proto_order_proto protoreflect.FileDescriptor

var file_order_proto_order_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xae, 0x03, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1c, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x41, 0x56, 0x41, 0x4e,
	0x45, 0x45, 0x53, 0x48, 0x4b, 0x4d, 0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_order_proto_order_proto_goTypes = []interface{}{
	(*RemoveOrderRequest)(nil),  // 0: order_proto.RemoveOrderRequest
	(*CustomerOrder)(nil),       // 1: order_proto.CustomerOrder
	(*GetOrderRequest)(nil),     // 2: order_proto.GetOrderRequest
	(*UpdateOrderRequest)(nil),  // 3: order_proto.UpdateOrderRequest
	(*RemoveOrderResponse)(nil), // 4: order_proto.RemoveOrderResponse
	(*CustomerResponse)(nil),    // 5: order_proto.CustomerResponse
	(*UpdateOrderResponse)(nil), // 6: order_proto.UpdateOrderResponse
}
var file_order_proto_order_proto_depIdxs = []int32{
	0, // 0: order_proto.OrderService.RemoveOrderCustomer:input_type -> order_proto.RemoveOrderRequest
	1, // 1: order_proto.OrderService.CreateOrder:input_type -> order_proto.CustomerOrder
	2, // 2: order_proto.OrderService.GetOrderDetails:input_type -> order_proto.GetOrderRequest
	3, // 3: order_proto.OrderService.UpdateOrderDetails:input_type -> order_proto.UpdateOrderRequest
	3, // 4: order_proto.OrderService.AddOrderDetails:input_type -> order_proto.UpdateOrderRequest
	4, // 5: order_proto.OrderService.RemoveOrderCustomer:output_type -> order_proto.RemoveOrderResponse
	5, // 6: order_proto.OrderService.CreateOrder:output_type -> order_proto.CustomerResponse
	1, // 7: order_proto.OrderService.GetOrderDetails:output_type -> order_proto.CustomerOrder
	6, // 8: order_proto.OrderService.UpdateOrderDetails:output_type -> order_proto.UpdateOrderResponse
	6, // 9: order_proto.OrderService.AddOrderDetails:output_type -> order_proto.UpdateOrderResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_proto_order_proto_init() }
func file_order_proto_order_proto_init() {
	if File_order_proto_order_proto != nil {
		return
	}
	file_order_proto_remove_order_proto_init()
	file_order_proto_create_order_proto_init()
	file_order_proto_get_order_proto_init()
	file_order_proto_update_order_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_order_proto_goTypes,
		DependencyIndexes: file_order_proto_order_proto_depIdxs,
	}.Build()
	File_order_proto_order_proto = out.File
	file_order_proto_order_proto_rawDesc = nil
	file_order_proto_order_proto_goTypes = nil
	file_order_proto_order_proto_depIdxs = nil
}
