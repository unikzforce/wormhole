// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package ebpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type VxlanXDPExternalIpv4LpmKey struct {
	Prefixlen uint32
	Data      [4]uint8
}

type VxlanXDPExternalMacAddress struct{ Addr [6]uint8 }

type VxlanXDPExternalMacTableEntry struct {
	ExpirationTimer     struct{ Opaque [2]uint64 }
	Ifindex             uint32
	_                   [4]byte
	LastSeenTimestampNs uint64
	BorderIp            struct{ S_addr uint32 }
	_                   [4]byte
}

type VxlanXDPExternalNetworkVniLight struct {
	Vni     uint32
	Network VxlanXDPExternalIpv4LpmKey
}

// LoadVxlanXDPExternal returns the embedded CollectionSpec for VxlanXDPExternal.
func LoadVxlanXDPExternal() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_VxlanXDPExternalBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load VxlanXDPExternal: %w", err)
	}

	return spec, err
}

// LoadVxlanXDPExternalObjects loads VxlanXDPExternal and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*VxlanXDPExternalObjects
//	*VxlanXDPExternalPrograms
//	*VxlanXDPExternalMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadVxlanXDPExternalObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadVxlanXDPExternal()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// VxlanXDPExternalSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanXDPExternalSpecs struct {
	VxlanXDPExternalProgramSpecs
	VxlanXDPExternalMapSpecs
}

// VxlanXDPExternalSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanXDPExternalProgramSpecs struct {
	VxlanXdpExternal *ebpf.ProgramSpec `ebpf:"vxlan_xdp_external"`
}

// VxlanXDPExternalMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanXDPExternalMapSpecs struct {
	MacTable         *ebpf.MapSpec `ebpf:"mac_table"`
	NetworksLightMap *ebpf.MapSpec `ebpf:"networks_light_map"`
}

// VxlanXDPExternalObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanXDPExternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanXDPExternalObjects struct {
	VxlanXDPExternalPrograms
	VxlanXDPExternalMaps
}

func (o *VxlanXDPExternalObjects) Close() error {
	return _VxlanXDPExternalClose(
		&o.VxlanXDPExternalPrograms,
		&o.VxlanXDPExternalMaps,
	)
}

// VxlanXDPExternalMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanXDPExternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanXDPExternalMaps struct {
	MacTable         *ebpf.Map `ebpf:"mac_table"`
	NetworksLightMap *ebpf.Map `ebpf:"networks_light_map"`
}

func (m *VxlanXDPExternalMaps) Close() error {
	return _VxlanXDPExternalClose(
		m.MacTable,
		m.NetworksLightMap,
	)
}

// VxlanXDPExternalPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanXDPExternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanXDPExternalPrograms struct {
	VxlanXdpExternal *ebpf.Program `ebpf:"vxlan_xdp_external"`
}

func (p *VxlanXDPExternalPrograms) Close() error {
	return _VxlanXDPExternalClose(
		p.VxlanXdpExternal,
	)
}

func _VxlanXDPExternalClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed vxlanxdpexternal_bpfeb.o
var _VxlanXDPExternalBytes []byte
