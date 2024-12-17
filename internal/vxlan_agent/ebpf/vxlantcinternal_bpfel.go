// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package ebpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type VxlanTCInternalExternalRouteInfo struct {
	ExternalIfaceIndex      uint32
	ExternalIfaceMac        struct{ Addr [6]uint8 }
	ExternalIfaceNextHopMac struct{ Addr [6]uint8 }
	ExternalIfaceIp         VxlanTCInternalInAddr
}

type VxlanTCInternalInAddr struct{ S_addr uint32 }

type VxlanTCInternalIpv4LpmKey struct {
	Prefixlen uint32
	Data      [4]uint8
}

type VxlanTCInternalNetworkVni struct {
	Vni                   uint32
	Network               VxlanTCInternalIpv4LpmKey
	InternalIfindexes     [10]uint32
	InternalIfindexesSize uint32
	BorderIps             [10]VxlanTCInternalInAddr
	BorderIpsSize         uint32
}

// LoadVxlanTCInternal returns the embedded CollectionSpec for VxlanTCInternal.
func LoadVxlanTCInternal() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_VxlanTCInternalBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load VxlanTCInternal: %w", err)
	}

	return spec, err
}

// LoadVxlanTCInternalObjects loads VxlanTCInternal and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*VxlanTCInternalObjects
//	*VxlanTCInternalPrograms
//	*VxlanTCInternalMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadVxlanTCInternalObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadVxlanTCInternal()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// VxlanTCInternalSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanTCInternalSpecs struct {
	VxlanTCInternalProgramSpecs
	VxlanTCInternalMapSpecs
}

// VxlanTCInternalSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanTCInternalProgramSpecs struct {
	VxlanTcInternal *ebpf.ProgramSpec `ebpf:"vxlan_tc_internal"`
}

// VxlanTCInternalMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanTCInternalMapSpecs struct {
	BorderIpToRouteInfoMap *ebpf.MapSpec `ebpf:"border_ip_to_route_info_map"`
	NetworksMap            *ebpf.MapSpec `ebpf:"networks_map"`
}

// VxlanTCInternalObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanTCInternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanTCInternalObjects struct {
	VxlanTCInternalPrograms
	VxlanTCInternalMaps
}

func (o *VxlanTCInternalObjects) Close() error {
	return _VxlanTCInternalClose(
		&o.VxlanTCInternalPrograms,
		&o.VxlanTCInternalMaps,
	)
}

// VxlanTCInternalMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanTCInternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanTCInternalMaps struct {
	BorderIpToRouteInfoMap *ebpf.Map `ebpf:"border_ip_to_route_info_map"`
	NetworksMap            *ebpf.Map `ebpf:"networks_map"`
}

func (m *VxlanTCInternalMaps) Close() error {
	return _VxlanTCInternalClose(
		m.BorderIpToRouteInfoMap,
		m.NetworksMap,
	)
}

// VxlanTCInternalPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanTCInternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanTCInternalPrograms struct {
	VxlanTcInternal *ebpf.Program `ebpf:"vxlan_tc_internal"`
}

func (p *VxlanTCInternalPrograms) Close() error {
	return _VxlanTCInternalClose(
		p.VxlanTcInternal,
	)
}

func _VxlanTCInternalClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed vxlantcinternal_bpfel.o
var _VxlanTCInternalBytes []byte
