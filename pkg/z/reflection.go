package z

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ReflectInterface interface {
	ServiceReflectionFromMessage(msg proto.Message) (ServiceReflectionInterface, error)
	ServiceReflectionFromName(name protoreflect.FullName) (ServiceReflectionInterface, bool)
	MethodReflectionFromMessage(msg proto.Message) (MethodReflectionInterface, error)
	AddService(name string)
	AddProtoExtension(*protoimpl.ExtensionInfo)
	Populate() error
}

type ReflectAccessor interface {
	Get(key *protoimpl.ExtensionInfo) (interface{}, bool)
	Has(key *protoimpl.ExtensionInfo) bool
	Bool(key *protoimpl.ExtensionInfo) bool
}

type ServiceReflectionInterface interface {
	ReflectAccessor
	FullName() string
	Descriptor() protoreflect.ServiceDescriptor
	Methods() []MethodReflectionInterface
	Shadow(svc *protoimpl.ExtensionInfo, method *protoimpl.ExtensionInfo, mref MethodReflectionInterface) (v interface{}, has bool)
}

type MethodReflectionInterface interface {
	ReflectAccessor
	Descriptor() protoreflect.MethodDescriptor
	SlashFullName() string
	Request() MessageReflectionInterface
	Response() MessageReflectionInterface
	FromMessage(proto.Message) MessageReflectionInterface
	IsRequest(proto.Message) bool
	IsResponse(proto.Message) bool
	IsValidMessage(proto.Message) bool
}

type MessageReflectionInterface interface {
	ReflectAccessor
	Descriptor() protoreflect.MessageDescriptor
	Value() proto.Message
	WithValue(v proto.Message)
	Fields() []FieldReflectInterface
}

type FieldReflectInterface interface {
	ReflectAccessor
	Descriptor() protoreflect.FieldDescriptor
}
