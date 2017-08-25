// Code generated by capnpc-go. DO NOT EDIT.

package flowlogs

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type VpcFlowLogEntry struct{ capnp.Struct }

// VpcFlowLogEntry_TypeID is the unique identifier for the type VpcFlowLogEntry.
const VpcFlowLogEntry_TypeID = 0xf12c57574cc72cec

func NewVpcFlowLogEntry(s *capnp.Segment) (VpcFlowLogEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 48, PointerCount: 1})
	return VpcFlowLogEntry{st}, err
}

func NewRootVpcFlowLogEntry(s *capnp.Segment) (VpcFlowLogEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 48, PointerCount: 1})
	return VpcFlowLogEntry{st}, err
}

func ReadRootVpcFlowLogEntry(msg *capnp.Message) (VpcFlowLogEntry, error) {
	root, err := msg.RootPtr()
	return VpcFlowLogEntry{root.Struct()}, err
}

func (s VpcFlowLogEntry) String() string {
	str, _ := text.Marshal(0xf12c57574cc72cec, s.Struct)
	return str
}

func (s VpcFlowLogEntry) Version() int8 {
	return int8(s.Struct.Uint8(0))
}

func (s VpcFlowLogEntry) SetVersion(v int8) {
	s.Struct.SetUint8(0, uint8(v))
}

func (s VpcFlowLogEntry) AccountId() uint64 {
	return s.Struct.Uint64(8)
}

func (s VpcFlowLogEntry) SetAccountId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s VpcFlowLogEntry) InterfaceId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s VpcFlowLogEntry) HasInterfaceId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VpcFlowLogEntry) InterfaceIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s VpcFlowLogEntry) SetInterfaceId(v string) error {
	return s.Struct.SetText(0, v)
}

func (s VpcFlowLogEntry) SrcAddr() uint32 {
	return s.Struct.Uint32(4)
}

func (s VpcFlowLogEntry) SetSrcAddr(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s VpcFlowLogEntry) DstAddr() uint32 {
	return s.Struct.Uint32(16)
}

func (s VpcFlowLogEntry) SetDstAddr(v uint32) {
	s.Struct.SetUint32(16, v)
}

func (s VpcFlowLogEntry) SrcPort() uint16 {
	return s.Struct.Uint16(2)
}

func (s VpcFlowLogEntry) SetSrcPort(v uint16) {
	s.Struct.SetUint16(2, v)
}

func (s VpcFlowLogEntry) DstPort() uint16 {
	return s.Struct.Uint16(20)
}

func (s VpcFlowLogEntry) SetDstPort(v uint16) {
	s.Struct.SetUint16(20, v)
}

func (s VpcFlowLogEntry) Protocol() uint8 {
	return s.Struct.Uint8(1)
}

func (s VpcFlowLogEntry) SetProtocol(v uint8) {
	s.Struct.SetUint8(1, v)
}

func (s VpcFlowLogEntry) Packets() uint16 {
	return s.Struct.Uint16(22)
}

func (s VpcFlowLogEntry) SetPackets(v uint16) {
	s.Struct.SetUint16(22, v)
}

func (s VpcFlowLogEntry) Bytes() uint64 {
	return s.Struct.Uint64(24)
}

func (s VpcFlowLogEntry) SetBytes(v uint64) {
	s.Struct.SetUint64(24, v)
}

func (s VpcFlowLogEntry) Start() uint32 {
	return s.Struct.Uint32(32)
}

func (s VpcFlowLogEntry) SetStart(v uint32) {
	s.Struct.SetUint32(32, v)
}

func (s VpcFlowLogEntry) End() uint32 {
	return s.Struct.Uint32(36)
}

func (s VpcFlowLogEntry) SetEnd(v uint32) {
	s.Struct.SetUint32(36, v)
}

func (s VpcFlowLogEntry) Action() VpcFlowLogEntry_Action {
	return VpcFlowLogEntry_Action(s.Struct.Uint16(40))
}

func (s VpcFlowLogEntry) SetAction(v VpcFlowLogEntry_Action) {
	s.Struct.SetUint16(40, uint16(v))
}

func (s VpcFlowLogEntry) LogStatus() VpcFlowLogEntry_LogStatus {
	return VpcFlowLogEntry_LogStatus(s.Struct.Uint16(42))
}

func (s VpcFlowLogEntry) SetLogStatus(v VpcFlowLogEntry_LogStatus) {
	s.Struct.SetUint16(42, uint16(v))
}

// VpcFlowLogEntry_List is a list of VpcFlowLogEntry.
type VpcFlowLogEntry_List struct{ capnp.List }

// NewVpcFlowLogEntry creates a new list of VpcFlowLogEntry.
func NewVpcFlowLogEntry_List(s *capnp.Segment, sz int32) (VpcFlowLogEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 48, PointerCount: 1}, sz)
	return VpcFlowLogEntry_List{l}, err
}

func (s VpcFlowLogEntry_List) At(i int) VpcFlowLogEntry { return VpcFlowLogEntry{s.List.Struct(i)} }

func (s VpcFlowLogEntry_List) Set(i int, v VpcFlowLogEntry) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s VpcFlowLogEntry_List) String() string {
	str, _ := text.MarshalList(0xf12c57574cc72cec, s.List)
	return str
}

// VpcFlowLogEntry_Promise is a wrapper for a VpcFlowLogEntry promised by a client call.
type VpcFlowLogEntry_Promise struct{ *capnp.Pipeline }

func (p VpcFlowLogEntry_Promise) Struct() (VpcFlowLogEntry, error) {
	s, err := p.Pipeline.Struct()
	return VpcFlowLogEntry{s}, err
}

type VpcFlowLogEntry_Action uint16

// VpcFlowLogEntry_Action_TypeID is the unique identifier for the type VpcFlowLogEntry_Action.
const VpcFlowLogEntry_Action_TypeID = 0xe2172f125bfcee29

// Values of VpcFlowLogEntry_Action.
const (
	VpcFlowLogEntry_Action_accept VpcFlowLogEntry_Action = 0
	VpcFlowLogEntry_Action_reject VpcFlowLogEntry_Action = 1
)

// String returns the enum's constant name.
func (c VpcFlowLogEntry_Action) String() string {
	switch c {
	case VpcFlowLogEntry_Action_accept:
		return "accept"
	case VpcFlowLogEntry_Action_reject:
		return "reject"

	default:
		return ""
	}
}

// VpcFlowLogEntry_ActionFromString returns the enum value with a name,
// or the zero value if there's no such value.
func VpcFlowLogEntry_ActionFromString(c string) VpcFlowLogEntry_Action {
	switch c {
	case "accept":
		return VpcFlowLogEntry_Action_accept
	case "reject":
		return VpcFlowLogEntry_Action_reject

	default:
		return 0
	}
}

type VpcFlowLogEntry_Action_List struct{ capnp.List }

func NewVpcFlowLogEntry_Action_List(s *capnp.Segment, sz int32) (VpcFlowLogEntry_Action_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return VpcFlowLogEntry_Action_List{l.List}, err
}

func (l VpcFlowLogEntry_Action_List) At(i int) VpcFlowLogEntry_Action {
	ul := capnp.UInt16List{List: l.List}
	return VpcFlowLogEntry_Action(ul.At(i))
}

func (l VpcFlowLogEntry_Action_List) Set(i int, v VpcFlowLogEntry_Action) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type VpcFlowLogEntry_LogStatus uint16

// VpcFlowLogEntry_LogStatus_TypeID is the unique identifier for the type VpcFlowLogEntry_LogStatus.
const VpcFlowLogEntry_LogStatus_TypeID = 0xac3172069fc5f8ff

// Values of VpcFlowLogEntry_LogStatus.
const (
	VpcFlowLogEntry_LogStatus_ok       VpcFlowLogEntry_LogStatus = 0
	VpcFlowLogEntry_LogStatus_noData   VpcFlowLogEntry_LogStatus = 1
	VpcFlowLogEntry_LogStatus_skipData VpcFlowLogEntry_LogStatus = 2
)

// String returns the enum's constant name.
func (c VpcFlowLogEntry_LogStatus) String() string {
	switch c {
	case VpcFlowLogEntry_LogStatus_ok:
		return "ok"
	case VpcFlowLogEntry_LogStatus_noData:
		return "noData"
	case VpcFlowLogEntry_LogStatus_skipData:
		return "skipData"

	default:
		return ""
	}
}

// VpcFlowLogEntry_LogStatusFromString returns the enum value with a name,
// or the zero value if there's no such value.
func VpcFlowLogEntry_LogStatusFromString(c string) VpcFlowLogEntry_LogStatus {
	switch c {
	case "ok":
		return VpcFlowLogEntry_LogStatus_ok
	case "noData":
		return VpcFlowLogEntry_LogStatus_noData
	case "skipData":
		return VpcFlowLogEntry_LogStatus_skipData

	default:
		return 0
	}
}

type VpcFlowLogEntry_LogStatus_List struct{ capnp.List }

func NewVpcFlowLogEntry_LogStatus_List(s *capnp.Segment, sz int32) (VpcFlowLogEntry_LogStatus_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return VpcFlowLogEntry_LogStatus_List{l.List}, err
}

func (l VpcFlowLogEntry_LogStatus_List) At(i int) VpcFlowLogEntry_LogStatus {
	ul := capnp.UInt16List{List: l.List}
	return VpcFlowLogEntry_LogStatus(ul.At(i))
}

func (l VpcFlowLogEntry_LogStatus_List) Set(i int, v VpcFlowLogEntry_LogStatus) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_9f5a5320cda72e66 = "x\xda\x8c\xd2\xcdk\x14I\x18\x06\xf0\xe7\xa9\x9e\xcfd" +
	"\xc20[\x1d\x02aC\x96\xdd\xc3\xb2\x90\x8f\xcd\xe6\xb0" +
	"\xcb\\f#FH\x08\x98I\xa3!\xf1\xd4\xe9\xe9\x84" +
	"I\x86\xee\xa6\xbb\x920\xa0\x8c\x82B\x0e\x0a\x11\"(" +
	"D\x92\x83b\x84\x08\x11\x14\"xT<\xf9Ox\x16" +
	"\xc1\x9b\x07\xa1\xa5:\x8c\x13\x11\xc1[\xd5\xaf^\x9e\xaa" +
	"~\xfb\xfd{\x91\xff\x8b\xb1\xb4/\x80\xea\xaf\xe9L\x1c" +
	"\x7fz\xbd\x97\x09\xc7\x0eQ\xfaC\xc4\xef\x87\xde\xcc\xcc" +
	"\xcf\x0f}\x048>\xc0U\xca1f\x019\xcc\xf3`" +
	"\xfc\xd7\x87\xcf\x97~\x19\xed{\xf7]\xe5$\xe7(\x17" +
	"\x92\xca\x0b\xfc\x17\xec\x1cV%3\xf1\xf2\xc8\xa3\xb7\xbf" +
	"Y\x8b{H\xeb\x92\xf1\x05\xf6S\xd6\xf9'0~\x8d" +
	"}\x06\x86\xe3\x8d\xc0\x19]n\xf8\x9b\xa9\x86\xbf\x12\x8d" +
	"8v\xe0\x05\xe5\x8b\x81s\xae\xe1o\xce\xf8+\x93\x9e" +
	"\x0a\x9b#3\xfe\x8a\xa5\x8a\xb6Z\x8ff\xc9j\x81\x02" +
	"(\x0d\xf4\x03d\xa9\xb7\x0cP\x94z\xa6\x01\xc3_\xab" +
	"x\xfeY[\xd9q\xb4V\x0f\xf4\x02\xc0O]0\xe1" +
	"\xa8\xbaOO\xa7\xe7\x92\xf4R9I\xcf\x97\x81\x8a\xed" +
	"8n\xa0*\xa1\xbb\xea:\xeak\x9c\xf1\xa38#l" +
	"Vs<\xdd\xb1|\xf9T\xa3\xd3s\x95\xe46/N" +
	"\xbe\xcaV\xeb`T\xfd\xcfH\x01)\x02\xb2\xc93\x80" +
	"\xa5h\xd0\xbaJ\xc1\x12iR\xfb\x15\xce\x01\xd6e\xed" +
	"[\x14\xa40\xf5C\xe5\x0d.\x01\xd6u\xcd\xdb\xba\xdc" +
	"\xa0I\x03\x90\xb7\x92\x98-\xed;\xdaS)\x93)@" +
	"\xdeN\xfc\xa6\xf6\xbb\xda\xd34\x99\x06\xe4\x9d\xc4\xb7\xb5" +
	"\xefj\xcft\x99\xcc\x00\xf2^\xe2;\xda\xf7\xb5gi" +
	"&\xff\xfa>\xa7\x01kW\xfb\x81\xf6\\\xb7\xc9\x1c " +
	"\x1f&\xf5\xfb\xda\x0f\xb5\xe7\x0d\x93y@>\xe6?\x80" +
	"\xf5@\xfb\x91\xf6\xae\x9c\xc9.@>I\xfc@\xfb3" +
	"\xed\xddy\x93\xdd\x80|\xca\xdf\x01\xebP\xfb\xb1\xf6\x82" +
	"i\xb2\x00\xc8\xe7,\x03\xd6\x91\xf6\x97\xda{zM\xf6" +
	"\x00\xf2E\xd2\x9ec\xed\xaf(\xd8\xdap\xc3\xa8\xee{" +
	"\x14\x10\x14`l;\x8e\xbf\xee\xa9)\xb0\xc6<\x04\xf3" +
	"`\\\xf7\x94\x1b.\xdb\x0e\xb2\xeeT\x8d\x05\x08\x16\xc0" +
	"V\x14:\x13\xb5Z\xc8\x1c\x04s`\xab\x16\xa9o\xf6" +
	"Q\xe8\xcc\xfa\xa1b\x16\x82\xd9\x93\xf3\xd3\xfb8\x08}" +
	"\xe5;~\x03\x003\x10\xcc\x80\xad\xc0v\xd6\\\x15\xb5" +
	"k\x06\x97\x9a\xca\x8d\xda\xef\x18\x8c\x94\x1d\xaav~\xd6" +
	"\xf5j\xedu\xc5N&\x85\xc5\xce0\x81,\x82q\xa3" +
	"3;,v\xa6\xeb\xe4\xf4K\x00\x00\x00\xff\xff?\x02" +
	"\xe1\xb8"

func init() {
	schemas.Register(schema_9f5a5320cda72e66,
		0xac3172069fc5f8ff,
		0xe2172f125bfcee29,
		0xf12c57574cc72cec)
}
