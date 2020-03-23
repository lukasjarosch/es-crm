// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crm/eventstore/v1/event_store.proto

package eventstorev1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Event is the core entity of the event-store.
type Event struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	AggregateId          string   `protobuf:"bytes,3,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateType        string   `protobuf:"bytes,4,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	Payload              string   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	TimestampUs          int64    `protobuf:"varint,6,opt,name=timestamp_us,json=timestampUs,proto3" json:"timestamp_us,omitempty"`
	Channel              string   `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd1bf46394471017, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *Event) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *Event) GetTimestampUs() int64 {
	if m != nil {
		return m.TimestampUs
	}
	return 0
}

func (m *Event) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

// EventFilter does filtering.
type EventFilter struct {
	EventId              string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	AggregateId          string   `protobuf:"bytes,2,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd1bf46394471017, []int{1}
}

func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFilter.Unmarshal(m, b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
}
func (m *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(m, src)
}
func (m *EventFilter) XXX_Size() int {
	return xxx_messageInfo_EventFilter.Size(m)
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EventFilter) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "crm.eventstore.v1.Event")
	proto.RegisterType((*EventFilter)(nil), "crm.eventstore.v1.EventFilter")
}

func init() {
	proto.RegisterFile("crm/eventstore/v1/event_store.proto", fileDescriptor_fd1bf46394471017)
}

var fileDescriptor_fd1bf46394471017 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0xba, 0xad, 0x9a, 0xce, 0xc9, 0x02, 0x42, 0xbc, 0x9b, 0x13, 0x61, 0x57, 0x1d,
	0xc5, 0x37, 0xd8, 0xa8, 0x30, 0xbc, 0x19, 0x53, 0x87, 0xc8, 0xa0, 0xc4, 0x26, 0xd6, 0x42, 0xd3,
	0x96, 0x34, 0x16, 0xf6, 0x3a, 0x5e, 0xfa, 0x28, 0xe2, 0x43, 0x49, 0x4e, 0xb5, 0xb9, 0xe8, 0x5d,
	0xfe, 0xef, 0xfc, 0xe7, 0x1c, 0xfe, 0x13, 0x7c, 0x9d, 0x28, 0xb9, 0x14, 0x8d, 0x28, 0x74, 0xad,
	0x4b, 0x25, 0x96, 0x4d, 0xd8, 0xaa, 0x18, 0x64, 0x50, 0xa9, 0x52, 0x97, 0x64, 0x9a, 0x28, 0x19,
	0x58, 0x53, 0xd0, 0x84, 0xf3, 0x1f, 0x07, 0x0f, 0x23, 0x43, 0xc8, 0x04, 0xa3, 0x8c, 0x53, 0x67,
	0xe6, 0x2c, 0x4e, 0x77, 0x28, 0xe3, 0x84, 0xe0, 0x81, 0x3e, 0x56, 0x82, 0x22, 0x20, 0xf0, 0x26,
	0x57, 0x78, 0xcc, 0xd2, 0x54, 0x89, 0x94, 0x69, 0x11, 0x67, 0x9c, 0xba, 0x50, 0xf3, 0x3b, 0xb6,
	0xe1, 0xe4, 0x06, 0x4f, 0xac, 0x05, 0x06, 0x0c, 0xc0, 0x74, 0xd6, 0xd1, 0x47, 0x33, 0x89, 0x62,
	0xaf, 0x62, 0xc7, 0xbc, 0x64, 0x9c, 0x0e, 0xa1, 0xfe, 0x2f, 0xcd, 0x0e, 0x9d, 0x49, 0x51, 0x6b,
	0x26, 0xab, 0xf8, 0xa3, 0xa6, 0xa3, 0x99, 0xb3, 0x70, 0x77, 0x7e, 0xc7, 0x9e, 0x6a, 0xd3, 0x9c,
	0xbc, 0xb3, 0xa2, 0x10, 0x39, 0xf5, 0xda, 0xe6, 0x3f, 0x39, 0xbf, 0xc7, 0x3e, 0xa4, 0xb9, 0xcb,
	0x72, 0x2d, 0x14, 0xb9, 0xc4, 0x27, 0xed, 0x15, 0xba, 0x64, 0x1e, 0xe8, 0x0d, 0xef, 0x45, 0x41,
	0xbd, 0x28, 0xab, 0x37, 0x7c, 0x91, 0x94, 0x32, 0xe8, 0x1d, 0x6d, 0x75, 0x0e, 0x3b, 0x1e, 0x8c,
	0xdc, 0x9a, 0xc3, 0x6e, 0x9d, 0x97, 0xb1, 0x75, 0x34, 0xe1, 0x27, 0x72, 0xd7, 0xd1, 0xf3, 0x17,
	0x9a, 0xae, 0x95, 0x0c, 0x22, 0xdb, 0xba, 0x0f, 0xbf, 0x81, 0x1d, 0x2c, 0x3b, 0xec, 0xc3, 0xd7,
	0x11, 0xfc, 0xce, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x18, 0xcc, 0x07, 0xc4, 0x01,
	0x00, 0x00,
}
