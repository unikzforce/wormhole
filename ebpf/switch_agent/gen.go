package switch_agent

//go:generated go run github.com/cilium/ebpf/cmd/bpf2go -type mac_address_iface_entry SwitchAgentXDP ./bpf/switch_agent_xdp.bpf.c
//go:generated go run github.com/cilium/ebpf/cmd/bpf2go SwitchAgentUnknownUnicastFlooding ./bpf/switch_agent_unknown_unicast_flooding.bpf.c
