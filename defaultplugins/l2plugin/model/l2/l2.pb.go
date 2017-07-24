// Code generated by protoc-gen-gogo.
// source: l2.proto
// DO NOT EDIT!

/*
Package l2 is a generated protocol buffer package.

It is generated from these files:
	l2.proto

It has these top-level messages:
	BridgeDomains
	FibTableEntries
	XConnectPairs
	BridgeDomainErrors
*/
package l2

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type FibTableEntries_FibTableEntry_Action int32

const (
	FibTableEntries_FibTableEntry_FORWARD FibTableEntries_FibTableEntry_Action = 0
	FibTableEntries_FibTableEntry_DROP    FibTableEntries_FibTableEntry_Action = 1
)

var FibTableEntries_FibTableEntry_Action_name = map[int32]string{
	0: "FORWARD",
	1: "DROP",
}
var FibTableEntries_FibTableEntry_Action_value = map[string]int32{
	"FORWARD": 0,
	"DROP":    1,
}

func (x FibTableEntries_FibTableEntry_Action) String() string {
	return proto.EnumName(FibTableEntries_FibTableEntry_Action_name, int32(x))
}

type BridgeDomains struct {
	BridgeDomains []*BridgeDomains_BridgeDomain `protobuf:"bytes,1,rep,name=bridge_domains" json:"bridge_domains,omitempty"`
}

func (m *BridgeDomains) Reset()         { *m = BridgeDomains{} }
func (m *BridgeDomains) String() string { return proto.CompactTextString(m) }
func (*BridgeDomains) ProtoMessage()    {}

func (m *BridgeDomains) GetBridgeDomains() []*BridgeDomains_BridgeDomain {
	if m != nil {
		return m.BridgeDomains
	}
	return nil
}

type BridgeDomains_BridgeDomain struct {
	Name                string                                            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Flood               bool                                              `protobuf:"varint,2,opt,name=flood,proto3" json:"flood,omitempty"`
	UnknownUnicastFlood bool                                              `protobuf:"varint,3,opt,name=unknown_unicast_flood,proto3" json:"unknown_unicast_flood,omitempty"`
	Forward             bool                                              `protobuf:"varint,4,opt,name=forward,proto3" json:"forward,omitempty"`
	Learn               bool                                              `protobuf:"varint,5,opt,name=learn,proto3" json:"learn,omitempty"`
	ArpTermination      bool                                              `protobuf:"varint,6,opt,name=arp_termination,proto3" json:"arp_termination,omitempty"`
	MacAge              uint32                                            `protobuf:"varint,7,opt,name=mac_age,proto3" json:"mac_age,omitempty"`
	Interfaces          []*BridgeDomains_BridgeDomain_Interfaces          `protobuf:"bytes,100,rep,name=interfaces" json:"interfaces,omitempty"`
	ArpTerminationTable []*BridgeDomains_BridgeDomain_ArpTerminationTable `protobuf:"bytes,102,rep,name=arp_termination_table" json:"arp_termination_table,omitempty"`
}

func (m *BridgeDomains_BridgeDomain) Reset()         { *m = BridgeDomains_BridgeDomain{} }
func (m *BridgeDomains_BridgeDomain) String() string { return proto.CompactTextString(m) }
func (*BridgeDomains_BridgeDomain) ProtoMessage()    {}

func (m *BridgeDomains_BridgeDomain) GetInterfaces() []*BridgeDomains_BridgeDomain_Interfaces {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *BridgeDomains_BridgeDomain) GetArpTerminationTable() []*BridgeDomains_BridgeDomain_ArpTerminationTable {
	if m != nil {
		return m.ArpTerminationTable
	}
	return nil
}

type BridgeDomains_BridgeDomain_Interfaces struct {
	Name                    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BridgedVirtualInterface bool   `protobuf:"varint,2,opt,name=bridged_virtual_interface,proto3" json:"bridged_virtual_interface,omitempty"`
	SplitHorizonGroup       uint32 `protobuf:"varint,3,opt,name=split_horizon_group,proto3" json:"split_horizon_group,omitempty"`
}

func (m *BridgeDomains_BridgeDomain_Interfaces) Reset()         { *m = BridgeDomains_BridgeDomain_Interfaces{} }
func (m *BridgeDomains_BridgeDomain_Interfaces) String() string { return proto.CompactTextString(m) }
func (*BridgeDomains_BridgeDomain_Interfaces) ProtoMessage()    {}

type BridgeDomains_BridgeDomain_ArpTerminationTable struct {
	IpAddress   string `protobuf:"bytes,1,opt,name=ip_address,proto3" json:"ip_address,omitempty"`
	PhysAddress string `protobuf:"bytes,2,opt,name=phys_address,proto3" json:"phys_address,omitempty"`
}

func (m *BridgeDomains_BridgeDomain_ArpTerminationTable) Reset() {
	*m = BridgeDomains_BridgeDomain_ArpTerminationTable{}
}
func (m *BridgeDomains_BridgeDomain_ArpTerminationTable) String() string {
	return proto.CompactTextString(m)
}
func (*BridgeDomains_BridgeDomain_ArpTerminationTable) ProtoMessage() {}

type FibTableEntries struct {
	FibTableEntry []*FibTableEntries_FibTableEntry `protobuf:"bytes,100,rep,name=fib_table_entry" json:"fib_table_entry,omitempty"`
}

func (m *FibTableEntries) Reset()         { *m = FibTableEntries{} }
func (m *FibTableEntries) String() string { return proto.CompactTextString(m) }
func (*FibTableEntries) ProtoMessage()    {}

func (m *FibTableEntries) GetFibTableEntry() []*FibTableEntries_FibTableEntry {
	if m != nil {
		return m.FibTableEntry
	}
	return nil
}

type FibTableEntries_FibTableEntry struct {
	PhysAddress             string                               `protobuf:"bytes,1,opt,name=phys_address,proto3" json:"phys_address,omitempty"`
	BridgeDomain            string                               `protobuf:"bytes,2,opt,name=bridge_domain,proto3" json:"bridge_domain,omitempty"`
	Action                  FibTableEntries_FibTableEntry_Action `protobuf:"varint,3,opt,name=action,proto3,enum=l2.FibTableEntries_FibTableEntry_Action" json:"action,omitempty"`
	OutgoingInterface       string                               `protobuf:"bytes,4,opt,name=outgoing_interface,proto3" json:"outgoing_interface,omitempty"`
	StaticConfig            bool                                 `protobuf:"varint,5,opt,name=static_config,proto3" json:"static_config,omitempty"`
	BridgedVirtualInterface bool                                 `protobuf:"varint,6,opt,name=bridged_virtual_interface,proto3" json:"bridged_virtual_interface,omitempty"`
}

func (m *FibTableEntries_FibTableEntry) Reset()         { *m = FibTableEntries_FibTableEntry{} }
func (m *FibTableEntries_FibTableEntry) String() string { return proto.CompactTextString(m) }
func (*FibTableEntries_FibTableEntry) ProtoMessage()    {}

type XConnectPairs struct {
	XConnectPairs []*XConnectPairs_XConnectPair `protobuf:"bytes,100,rep,name=x_connect_pairs" json:"x_connect_pairs,omitempty"`
}

func (m *XConnectPairs) Reset()         { *m = XConnectPairs{} }
func (m *XConnectPairs) String() string { return proto.CompactTextString(m) }
func (*XConnectPairs) ProtoMessage()    {}

func (m *XConnectPairs) GetXConnectPairs() []*XConnectPairs_XConnectPair {
	if m != nil {
		return m.XConnectPairs
	}
	return nil
}

type XConnectPairs_XConnectPair struct {
	ReceiveInterface  string `protobuf:"bytes,2,opt,name=receive_interface,proto3" json:"receive_interface,omitempty"`
	TransmitInterface string `protobuf:"bytes,3,opt,name=transmit_interface,proto3" json:"transmit_interface,omitempty"`
}

func (m *XConnectPairs_XConnectPair) Reset()         { *m = XConnectPairs_XConnectPair{} }
func (m *XConnectPairs_XConnectPair) String() string { return proto.CompactTextString(m) }
func (*XConnectPairs_XConnectPair) ProtoMessage()    {}

type BridgeDomainErrors struct {
	BridgeDomain []*BridgeDomainErrors_BridgeDomainError `protobuf:"bytes,1,rep,name=bridge_domain" json:"bridge_domain,omitempty"`
}

func (m *BridgeDomainErrors) Reset()         { *m = BridgeDomainErrors{} }
func (m *BridgeDomainErrors) String() string { return proto.CompactTextString(m) }
func (*BridgeDomainErrors) ProtoMessage()    {}

func (m *BridgeDomainErrors) GetBridgeDomain() []*BridgeDomainErrors_BridgeDomainError {
	if m != nil {
		return m.BridgeDomain
	}
	return nil
}

type BridgeDomainErrors_BridgeDomainError struct {
	BdName       string `protobuf:"bytes,1,opt,name=bd_name,proto3" json:"bd_name,omitempty"`
	ChangeType   string `protobuf:"bytes,2,opt,name=change_type,proto3" json:"change_type,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,proto3" json:"error_message,omitempty"`
	LastChange   int64  `protobuf:"varint,4,opt,name=last_change,proto3" json:"last_change,omitempty"`
}

func (m *BridgeDomainErrors_BridgeDomainError) Reset()         { *m = BridgeDomainErrors_BridgeDomainError{} }
func (m *BridgeDomainErrors_BridgeDomainError) String() string { return proto.CompactTextString(m) }
func (*BridgeDomainErrors_BridgeDomainError) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("l2.FibTableEntries_FibTableEntry_Action", FibTableEntries_FibTableEntry_Action_name, FibTableEntries_FibTableEntry_Action_value)
}
