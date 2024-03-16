// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package switch_agent

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type SwitchAgentXDPIfaceIndex struct {
	InterfaceIndex uint32
	_              [4]byte
	Timestamp      uint64
}

type SwitchAgentXDPMacAddress struct{ Mac [6]uint8 }

type SwitchAgentXDPMacAddressIfaceEntry struct {
	Mac   SwitchAgentXDPMacAddress
	_     [2]byte
	Iface SwitchAgentXDPIfaceIndex
}

// LoadSwitchAgentXDP returns the embedded CollectionSpec for SwitchAgentXDP.
func LoadSwitchAgentXDP() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_SwitchAgentXDPBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load SwitchAgentXDP: %w", err)
	}

	return spec, err
}

// LoadSwitchAgentXDPObjects loads SwitchAgentXDP and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*SwitchAgentXDPObjects
//	*SwitchAgentXDPPrograms
//	*SwitchAgentXDPMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadSwitchAgentXDPObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadSwitchAgentXDP()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// SwitchAgentXDPSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type SwitchAgentXDPSpecs struct {
	SwitchAgentXDPProgramSpecs
	SwitchAgentXDPMapSpecs
}

// SwitchAgentXDPSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type SwitchAgentXDPProgramSpecs struct {
	SwitchAgentXdp *ebpf.ProgramSpec `ebpf:"switch_agent_xdp"`
}

// SwitchAgentXDPMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type SwitchAgentXDPMapSpecs struct {
	MacTable               *ebpf.MapSpec `ebpf:"mac_table"`
	NewDiscoveredEntriesRb *ebpf.MapSpec `ebpf:"new_discovered_entries_rb"`
}

// SwitchAgentXDPObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadSwitchAgentXDPObjects or ebpf.CollectionSpec.LoadAndAssign.
type SwitchAgentXDPObjects struct {
	SwitchAgentXDPPrograms
	SwitchAgentXDPMaps
}

func (o *SwitchAgentXDPObjects) Close() error {
	return _SwitchAgentXDPClose(
		&o.SwitchAgentXDPPrograms,
		&o.SwitchAgentXDPMaps,
	)
}

// SwitchAgentXDPMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadSwitchAgentXDPObjects or ebpf.CollectionSpec.LoadAndAssign.
type SwitchAgentXDPMaps struct {
	MacTable               *ebpf.Map `ebpf:"mac_table"`
	NewDiscoveredEntriesRb *ebpf.Map `ebpf:"new_discovered_entries_rb"`
}

func (m *SwitchAgentXDPMaps) Close() error {
	return _SwitchAgentXDPClose(
		m.MacTable,
		m.NewDiscoveredEntriesRb,
	)
}

// SwitchAgentXDPPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadSwitchAgentXDPObjects or ebpf.CollectionSpec.LoadAndAssign.
type SwitchAgentXDPPrograms struct {
	SwitchAgentXdp *ebpf.Program `ebpf:"switch_agent_xdp"`
}

func (p *SwitchAgentXDPPrograms) Close() error {
	return _SwitchAgentXDPClose(
		p.SwitchAgentXdp,
	)
}

func _SwitchAgentXDPClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed switchagentxdp_bpfel.o
var _SwitchAgentXDPBytes []byte
