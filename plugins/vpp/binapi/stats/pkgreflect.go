// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package stats

import "reflect"

var Types = map[string]reflect.Type{
	"IP4FibCounter":                      reflect.TypeOf((*IP4FibCounter)(nil)).Elem(),
	"IP4MfibCounter":                     reflect.TypeOf((*IP4MfibCounter)(nil)).Elem(),
	"IP4NbrCounter":                      reflect.TypeOf((*IP4NbrCounter)(nil)).Elem(),
	"IP6FibCounter":                      reflect.TypeOf((*IP6FibCounter)(nil)).Elem(),
	"IP6MfibCounter":                     reflect.TypeOf((*IP6MfibCounter)(nil)).Elem(),
	"IP6NbrCounter":                      reflect.TypeOf((*IP6NbrCounter)(nil)).Elem(),
	"StatsGetPollerDelay":                reflect.TypeOf((*StatsGetPollerDelay)(nil)).Elem(),
	"StatsGetPollerDelayReply":           reflect.TypeOf((*StatsGetPollerDelayReply)(nil)).Elem(),
	"UDPEncapCounter":                    reflect.TypeOf((*UDPEncapCounter)(nil)).Elem(),
	"VlibCounter":                        reflect.TypeOf((*VlibCounter)(nil)).Elem(),
	"VnetCombinedCounter":                reflect.TypeOf((*VnetCombinedCounter)(nil)).Elem(),
	"VnetGetSummaryStats":                reflect.TypeOf((*VnetGetSummaryStats)(nil)).Elem(),
	"VnetGetSummaryStatsReply":           reflect.TypeOf((*VnetGetSummaryStatsReply)(nil)).Elem(),
	"VnetIP4FibCounters":                 reflect.TypeOf((*VnetIP4FibCounters)(nil)).Elem(),
	"VnetIP4MfibCounters":                reflect.TypeOf((*VnetIP4MfibCounters)(nil)).Elem(),
	"VnetIP4NbrCounters":                 reflect.TypeOf((*VnetIP4NbrCounters)(nil)).Elem(),
	"VnetIP6FibCounters":                 reflect.TypeOf((*VnetIP6FibCounters)(nil)).Elem(),
	"VnetIP6MfibCounters":                reflect.TypeOf((*VnetIP6MfibCounters)(nil)).Elem(),
	"VnetIP6NbrCounters":                 reflect.TypeOf((*VnetIP6NbrCounters)(nil)).Elem(),
	"VnetInterfaceCombinedCounters":      reflect.TypeOf((*VnetInterfaceCombinedCounters)(nil)).Elem(),
	"VnetInterfaceSimpleCounters":        reflect.TypeOf((*VnetInterfaceSimpleCounters)(nil)).Elem(),
	"VnetPerInterfaceCombinedCounters":   reflect.TypeOf((*VnetPerInterfaceCombinedCounters)(nil)).Elem(),
	"VnetPerInterfaceSimpleCounters":     reflect.TypeOf((*VnetPerInterfaceSimpleCounters)(nil)).Elem(),
	"VnetSimpleCounter":                  reflect.TypeOf((*VnetSimpleCounter)(nil)).Elem(),
	"VnetUDPEncapCounters":               reflect.TypeOf((*VnetUDPEncapCounters)(nil)).Elem(),
	"WantIP4FibStats":                    reflect.TypeOf((*WantIP4FibStats)(nil)).Elem(),
	"WantIP4FibStatsReply":               reflect.TypeOf((*WantIP4FibStatsReply)(nil)).Elem(),
	"WantIP4MfibStats":                   reflect.TypeOf((*WantIP4MfibStats)(nil)).Elem(),
	"WantIP4MfibStatsReply":              reflect.TypeOf((*WantIP4MfibStatsReply)(nil)).Elem(),
	"WantIP4NbrStats":                    reflect.TypeOf((*WantIP4NbrStats)(nil)).Elem(),
	"WantIP4NbrStatsReply":               reflect.TypeOf((*WantIP4NbrStatsReply)(nil)).Elem(),
	"WantIP6FibStats":                    reflect.TypeOf((*WantIP6FibStats)(nil)).Elem(),
	"WantIP6FibStatsReply":               reflect.TypeOf((*WantIP6FibStatsReply)(nil)).Elem(),
	"WantIP6MfibStats":                   reflect.TypeOf((*WantIP6MfibStats)(nil)).Elem(),
	"WantIP6MfibStatsReply":              reflect.TypeOf((*WantIP6MfibStatsReply)(nil)).Elem(),
	"WantIP6NbrStats":                    reflect.TypeOf((*WantIP6NbrStats)(nil)).Elem(),
	"WantIP6NbrStatsReply":               reflect.TypeOf((*WantIP6NbrStatsReply)(nil)).Elem(),
	"WantInterfaceCombinedStats":         reflect.TypeOf((*WantInterfaceCombinedStats)(nil)).Elem(),
	"WantInterfaceCombinedStatsReply":    reflect.TypeOf((*WantInterfaceCombinedStatsReply)(nil)).Elem(),
	"WantInterfaceSimpleStats":           reflect.TypeOf((*WantInterfaceSimpleStats)(nil)).Elem(),
	"WantInterfaceSimpleStatsReply":      reflect.TypeOf((*WantInterfaceSimpleStatsReply)(nil)).Elem(),
	"WantPerInterfaceCombinedStats":      reflect.TypeOf((*WantPerInterfaceCombinedStats)(nil)).Elem(),
	"WantPerInterfaceCombinedStatsReply": reflect.TypeOf((*WantPerInterfaceCombinedStatsReply)(nil)).Elem(),
	"WantPerInterfaceSimpleStats":        reflect.TypeOf((*WantPerInterfaceSimpleStats)(nil)).Elem(),
	"WantPerInterfaceSimpleStatsReply":   reflect.TypeOf((*WantPerInterfaceSimpleStatsReply)(nil)).Elem(),
	"WantStats":                          reflect.TypeOf((*WantStats)(nil)).Elem(),
	"WantStatsReply":                     reflect.TypeOf((*WantStatsReply)(nil)).Elem(),
	"WantUDPEncapStats":                  reflect.TypeOf((*WantUDPEncapStats)(nil)).Elem(),
	"WantUDPEncapStatsReply":             reflect.TypeOf((*WantUDPEncapStatsReply)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewStatsGetPollerDelay":                reflect.ValueOf(NewStatsGetPollerDelay),
	"NewStatsGetPollerDelayReply":           reflect.ValueOf(NewStatsGetPollerDelayReply),
	"NewVnetGetSummaryStats":                reflect.ValueOf(NewVnetGetSummaryStats),
	"NewVnetGetSummaryStatsReply":           reflect.ValueOf(NewVnetGetSummaryStatsReply),
	"NewVnetIP4FibCounters":                 reflect.ValueOf(NewVnetIP4FibCounters),
	"NewVnetIP4MfibCounters":                reflect.ValueOf(NewVnetIP4MfibCounters),
	"NewVnetIP4NbrCounters":                 reflect.ValueOf(NewVnetIP4NbrCounters),
	"NewVnetIP6FibCounters":                 reflect.ValueOf(NewVnetIP6FibCounters),
	"NewVnetIP6MfibCounters":                reflect.ValueOf(NewVnetIP6MfibCounters),
	"NewVnetIP6NbrCounters":                 reflect.ValueOf(NewVnetIP6NbrCounters),
	"NewVnetInterfaceCombinedCounters":      reflect.ValueOf(NewVnetInterfaceCombinedCounters),
	"NewVnetInterfaceSimpleCounters":        reflect.ValueOf(NewVnetInterfaceSimpleCounters),
	"NewVnetPerInterfaceCombinedCounters":   reflect.ValueOf(NewVnetPerInterfaceCombinedCounters),
	"NewVnetPerInterfaceSimpleCounters":     reflect.ValueOf(NewVnetPerInterfaceSimpleCounters),
	"NewVnetUDPEncapCounters":               reflect.ValueOf(NewVnetUDPEncapCounters),
	"NewWantIP4FibStats":                    reflect.ValueOf(NewWantIP4FibStats),
	"NewWantIP4FibStatsReply":               reflect.ValueOf(NewWantIP4FibStatsReply),
	"NewWantIP4MfibStats":                   reflect.ValueOf(NewWantIP4MfibStats),
	"NewWantIP4MfibStatsReply":              reflect.ValueOf(NewWantIP4MfibStatsReply),
	"NewWantIP4NbrStats":                    reflect.ValueOf(NewWantIP4NbrStats),
	"NewWantIP4NbrStatsReply":               reflect.ValueOf(NewWantIP4NbrStatsReply),
	"NewWantIP6FibStats":                    reflect.ValueOf(NewWantIP6FibStats),
	"NewWantIP6FibStatsReply":               reflect.ValueOf(NewWantIP6FibStatsReply),
	"NewWantIP6MfibStats":                   reflect.ValueOf(NewWantIP6MfibStats),
	"NewWantIP6MfibStatsReply":              reflect.ValueOf(NewWantIP6MfibStatsReply),
	"NewWantIP6NbrStats":                    reflect.ValueOf(NewWantIP6NbrStats),
	"NewWantIP6NbrStatsReply":               reflect.ValueOf(NewWantIP6NbrStatsReply),
	"NewWantInterfaceCombinedStats":         reflect.ValueOf(NewWantInterfaceCombinedStats),
	"NewWantInterfaceCombinedStatsReply":    reflect.ValueOf(NewWantInterfaceCombinedStatsReply),
	"NewWantInterfaceSimpleStats":           reflect.ValueOf(NewWantInterfaceSimpleStats),
	"NewWantInterfaceSimpleStatsReply":      reflect.ValueOf(NewWantInterfaceSimpleStatsReply),
	"NewWantPerInterfaceCombinedStats":      reflect.ValueOf(NewWantPerInterfaceCombinedStats),
	"NewWantPerInterfaceCombinedStatsReply": reflect.ValueOf(NewWantPerInterfaceCombinedStatsReply),
	"NewWantPerInterfaceSimpleStats":        reflect.ValueOf(NewWantPerInterfaceSimpleStats),
	"NewWantPerInterfaceSimpleStatsReply":   reflect.ValueOf(NewWantPerInterfaceSimpleStatsReply),
	"NewWantStats":                          reflect.ValueOf(NewWantStats),
	"NewWantStatsReply":                     reflect.ValueOf(NewWantStatsReply),
	"NewWantUDPEncapStats":                  reflect.ValueOf(NewWantUDPEncapStats),
	"NewWantUDPEncapStatsReply":             reflect.ValueOf(NewWantUDPEncapStatsReply),
}

var Variables = map[string]reflect.Value{}

var Consts = map[string]reflect.Value{}
