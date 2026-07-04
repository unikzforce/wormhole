## Why

A full code review of the Wormhole eBPF VXLAN VTEP identified five bugs — ranging from a wrong constant value that causes premature MAC entry expiration, to an incorrect TC return code, a hardcoded LPM prefix length, a file-descriptor leak in the ring buffer reader, and overly aggressive flooding policy. Left unaddressed these bugs silently degrade correctness and performance in all deployments.

## What Changes

- Fix `FIVE_MINUTES_IN_NS` constant from 20 s to 300 s (5 min) in `vxlan_common.bpf.h`
- Sync `switch_agent` userspace bigcache TTL from `20*time.Second` to `5*time.Minute`
- Replace `AGENT_PASS` (XDP enum value 3) with `TC_ACT_OK` (0) in `vxlan_tc_external.bpf.c`
- Change hardcoded `prefixlen=24` to `prefixlen=32` for the IP-packet LPM lookup in `vxlan_xdp_external.bpf.c`
- Move `ringbuf.NewReader` outside the read loop in `switch_agent.go` to stop leaking file descriptors
- Separate broadcast flooding from unknown-unicast flooding in `vxlan_xdp_internal.bpf.c`

## Capabilities

### New Capabilities

- `flooding-policy`: Defines the rule for when a packet must be flooded vs. dropped vs. redirected at the XDP layer, distinguishing broadcast from unknown unicast.

### Modified Capabilities

- `mac-table-ttl`: MAC table entry lifetime changes from the current effective 20 s to the intended 5 minutes, affecting both the eBPF timer and the userspace eviction cache.
- `tc-return-codes`: TC programs must return `TC_ACT_OK` / `TC_ACT_SHOT` exclusively; XDP-layer enums (`AGENT_PASS`, `AGENT_ERROR_*`) are forbidden in TC context.
- `lpm-prefix-lookup`: The LPM trie key `prefixlen` for IP-packet destination lookups must be set to 32 (host address) so the trie performs longest-prefix matching correctly for any configured network prefix length.
- `ringbuf-reader-lifecycle`: The ring buffer reader must be created once per goroutine lifetime, not once per event, to prevent file-descriptor exhaustion.

## Impact

- `internal/vxlan_agent/ebpf/c/vxlan_common.bpf.h` — constant fix
- `internal/vxlan_agent/ebpf/c/vxlan_tc_external.bpf.c` — return code fix + flooding separation
- `internal/vxlan_agent/ebpf/c/vxlan_xdp_internal.bpf.c` — flooding policy logic
- `internal/vxlan_agent/ebpf/c/vxlan_xdp_external.bpf.c` — prefixlen fix
- `internal/switch_agent/switch_agent.go` — ringbuf reader lifecycle fix + TTL sync
- eBPF object files will need to be regenerated (`go generate ./internal/vxlan_agent/ebpf/ ./internal/switch_agent/ebpf/`) after any `.bpf.c` change
- No API or config-file format changes; no breaking changes to callers
