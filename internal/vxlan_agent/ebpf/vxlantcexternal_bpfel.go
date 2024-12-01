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

// LoadVxlanTCExternal returns the embedded CollectionSpec for VxlanTCExternal.
func LoadVxlanTCExternal() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_VxlanTCExternalBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load VxlanTCExternal: %w", err)
	}

	return spec, err
}

// LoadVxlanTCExternalObjects loads VxlanTCExternal and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*VxlanTCExternalObjects
//	*VxlanTCExternalPrograms
//	*VxlanTCExternalMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadVxlanTCExternalObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadVxlanTCExternal()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// VxlanTCExternalSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanTCExternalSpecs struct {
	VxlanTCExternalProgramSpecs
	VxlanTCExternalMapSpecs
}

// VxlanTCExternalSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanTCExternalProgramSpecs struct {
	VxlanTcExternal *ebpf.ProgramSpec `ebpf:"vxlan_tc_external"`
}

// VxlanTCExternalMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type VxlanTCExternalMapSpecs struct {
	InternalIfindexesArray       *ebpf.MapSpec `ebpf:"internal_ifindexes_array"`
	InternalIfindexesArrayLength *ebpf.MapSpec `ebpf:"internal_ifindexes_array_length"`
}

// VxlanTCExternalObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanTCExternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanTCExternalObjects struct {
	VxlanTCExternalPrograms
	VxlanTCExternalMaps
}

func (o *VxlanTCExternalObjects) Close() error {
	return _VxlanTCExternalClose(
		&o.VxlanTCExternalPrograms,
		&o.VxlanTCExternalMaps,
	)
}

// VxlanTCExternalMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanTCExternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanTCExternalMaps struct {
	InternalIfindexesArray       *ebpf.Map `ebpf:"internal_ifindexes_array"`
	InternalIfindexesArrayLength *ebpf.Map `ebpf:"internal_ifindexes_array_length"`
}

func (m *VxlanTCExternalMaps) Close() error {
	return _VxlanTCExternalClose(
		m.InternalIfindexesArray,
		m.InternalIfindexesArrayLength,
	)
}

// VxlanTCExternalPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadVxlanTCExternalObjects or ebpf.CollectionSpec.LoadAndAssign.
type VxlanTCExternalPrograms struct {
	VxlanTcExternal *ebpf.Program `ebpf:"vxlan_tc_external"`
}

func (p *VxlanTCExternalPrograms) Close() error {
	return _VxlanTCExternalClose(
		p.VxlanTcExternal,
	)
}

func _VxlanTCExternalClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed vxlantcexternal_bpfel.o
var _VxlanTCExternalBytes []byte
