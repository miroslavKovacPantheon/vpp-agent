// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/bond.api.json

/*
Package bond is a generated from VPP binary API module 'bond'.

 The bond module consists of:
	 12 messages
	  6 services
*/
package bond

import api "git.fd.io/govpp.git/api"
import bytes "bytes"
import context "context"
import strconv "strconv"
import struc "github.com/lunixbochs/struc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = strconv.Itoa
var _ = struc.Pack

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

const (
	// ModuleName is the name of this module.
	ModuleName = "bond"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xd7e62f79
)

/* Messages */

// BondCreate represents VPP binary API message 'bond_create':
type BondCreate struct {
	ID           uint32
	UseCustomMac uint8
	MacAddress   []byte `struc:"[6]byte"`
	Mode         uint8
	Lb           uint8
}

func (*BondCreate) GetMessageName() string {
	return "bond_create"
}
func (*BondCreate) GetCrcString() string {
	return "3b21645d"
}
func (*BondCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BondCreateReply represents VPP binary API message 'bond_create_reply':
type BondCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*BondCreateReply) GetMessageName() string {
	return "bond_create_reply"
}
func (*BondCreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*BondCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BondDelete represents VPP binary API message 'bond_delete':
type BondDelete struct {
	SwIfIndex uint32
}

func (*BondDelete) GetMessageName() string {
	return "bond_delete"
}
func (*BondDelete) GetCrcString() string {
	return "529cb13f"
}
func (*BondDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BondDeleteReply represents VPP binary API message 'bond_delete_reply':
type BondDeleteReply struct {
	Retval int32
}

func (*BondDeleteReply) GetMessageName() string {
	return "bond_delete_reply"
}
func (*BondDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BondDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BondDetachSlave represents VPP binary API message 'bond_detach_slave':
type BondDetachSlave struct {
	SwIfIndex uint32
}

func (*BondDetachSlave) GetMessageName() string {
	return "bond_detach_slave"
}
func (*BondDetachSlave) GetCrcString() string {
	return "529cb13f"
}
func (*BondDetachSlave) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BondDetachSlaveReply represents VPP binary API message 'bond_detach_slave_reply':
type BondDetachSlaveReply struct {
	Retval int32
}

func (*BondDetachSlaveReply) GetMessageName() string {
	return "bond_detach_slave_reply"
}
func (*BondDetachSlaveReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BondDetachSlaveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BondEnslave represents VPP binary API message 'bond_enslave':
type BondEnslave struct {
	SwIfIndex     uint32
	BondSwIfIndex uint32
	IsPassive     uint8
	IsLongTimeout uint8
}

func (*BondEnslave) GetMessageName() string {
	return "bond_enslave"
}
func (*BondEnslave) GetCrcString() string {
	return "0ded34f6"
}
func (*BondEnslave) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BondEnslaveReply represents VPP binary API message 'bond_enslave_reply':
type BondEnslaveReply struct {
	Retval int32
}

func (*BondEnslaveReply) GetMessageName() string {
	return "bond_enslave_reply"
}
func (*BondEnslaveReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BondEnslaveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceBondDetails represents VPP binary API message 'sw_interface_bond_details':
type SwInterfaceBondDetails struct {
	SwIfIndex     uint32
	ID            uint32
	InterfaceName []byte `struc:"[64]byte"`
	Mode          uint8
	Lb            uint8
	ActiveSlaves  uint32
	Slaves        uint32
}

func (*SwInterfaceBondDetails) GetMessageName() string {
	return "sw_interface_bond_details"
}
func (*SwInterfaceBondDetails) GetCrcString() string {
	return "22e48daa"
}
func (*SwInterfaceBondDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceBondDump represents VPP binary API message 'sw_interface_bond_dump':
type SwInterfaceBondDump struct{}

func (*SwInterfaceBondDump) GetMessageName() string {
	return "sw_interface_bond_dump"
}
func (*SwInterfaceBondDump) GetCrcString() string {
	return "51077d14"
}
func (*SwInterfaceBondDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSlaveDetails represents VPP binary API message 'sw_interface_slave_details':
type SwInterfaceSlaveDetails struct {
	SwIfIndex     uint32
	InterfaceName []byte `struc:"[64]byte"`
	IsPassive     uint8
	IsLongTimeout uint8
}

func (*SwInterfaceSlaveDetails) GetMessageName() string {
	return "sw_interface_slave_details"
}
func (*SwInterfaceSlaveDetails) GetCrcString() string {
	return "d5c58e45"
}
func (*SwInterfaceSlaveDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceSlaveDump represents VPP binary API message 'sw_interface_slave_dump':
type SwInterfaceSlaveDump struct {
	SwIfIndex uint32
}

func (*SwInterfaceSlaveDump) GetMessageName() string {
	return "sw_interface_slave_dump"
}
func (*SwInterfaceSlaveDump) GetCrcString() string {
	return "529cb13f"
}
func (*SwInterfaceSlaveDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*BondCreate)(nil), "bond.BondCreate")
	api.RegisterMessage((*BondCreateReply)(nil), "bond.BondCreateReply")
	api.RegisterMessage((*BondDelete)(nil), "bond.BondDelete")
	api.RegisterMessage((*BondDeleteReply)(nil), "bond.BondDeleteReply")
	api.RegisterMessage((*BondDetachSlave)(nil), "bond.BondDetachSlave")
	api.RegisterMessage((*BondDetachSlaveReply)(nil), "bond.BondDetachSlaveReply")
	api.RegisterMessage((*BondEnslave)(nil), "bond.BondEnslave")
	api.RegisterMessage((*BondEnslaveReply)(nil), "bond.BondEnslaveReply")
	api.RegisterMessage((*SwInterfaceBondDetails)(nil), "bond.SwInterfaceBondDetails")
	api.RegisterMessage((*SwInterfaceBondDump)(nil), "bond.SwInterfaceBondDump")
	api.RegisterMessage((*SwInterfaceSlaveDetails)(nil), "bond.SwInterfaceSlaveDetails")
	api.RegisterMessage((*SwInterfaceSlaveDump)(nil), "bond.SwInterfaceSlaveDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*BondCreate)(nil),
		(*BondCreateReply)(nil),
		(*BondDelete)(nil),
		(*BondDeleteReply)(nil),
		(*BondDetachSlave)(nil),
		(*BondDetachSlaveReply)(nil),
		(*BondEnslave)(nil),
		(*BondEnslaveReply)(nil),
		(*SwInterfaceBondDetails)(nil),
		(*SwInterfaceBondDump)(nil),
		(*SwInterfaceSlaveDetails)(nil),
		(*SwInterfaceSlaveDump)(nil),
	}
}
