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

// LoadDummyXdp returns the embedded CollectionSpec for DummyXdp.
func LoadDummyXdp() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_DummyXdpBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load DummyXdp: %w", err)
	}

	return spec, err
}

// LoadDummyXdpObjects loads DummyXdp and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*DummyXdpObjects
//	*DummyXdpPrograms
//	*DummyXdpMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadDummyXdpObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadDummyXdp()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// DummyXdpSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type DummyXdpSpecs struct {
	DummyXdpProgramSpecs
	DummyXdpMapSpecs
}

// DummyXdpSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type DummyXdpProgramSpecs struct {
	DummyXdp *ebpf.ProgramSpec `ebpf:"dummy_xdp"`
}

// DummyXdpMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type DummyXdpMapSpecs struct {
}

// DummyXdpObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadDummyXdpObjects or ebpf.CollectionSpec.LoadAndAssign.
type DummyXdpObjects struct {
	DummyXdpPrograms
	DummyXdpMaps
}

func (o *DummyXdpObjects) Close() error {
	return _DummyXdpClose(
		&o.DummyXdpPrograms,
		&o.DummyXdpMaps,
	)
}

// DummyXdpMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadDummyXdpObjects or ebpf.CollectionSpec.LoadAndAssign.
type DummyXdpMaps struct {
}

func (m *DummyXdpMaps) Close() error {
	return _DummyXdpClose()
}

// DummyXdpPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadDummyXdpObjects or ebpf.CollectionSpec.LoadAndAssign.
type DummyXdpPrograms struct {
	DummyXdp *ebpf.Program `ebpf:"dummy_xdp"`
}

func (p *DummyXdpPrograms) Close() error {
	return _DummyXdpClose(
		p.DummyXdp,
	)
}

func _DummyXdpClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed dummyxdp_bpfel.o
var _DummyXdpBytes []byte
