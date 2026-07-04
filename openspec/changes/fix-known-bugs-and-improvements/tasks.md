## 1. Fix MAC Table TTL Constant

- [ ] 1.1 In `internal/vxlan_agent/ebpf/c/vxlan_common.bpf.h` line 27, change `#define FIVE_MINUTES_IN_NS 20000000000` to `#define FIVE_MINUTES_IN_NS 300000000000ULL`
- [ ] 1.2 In `internal/switch_agent/switch_agent.go` line 117, change `bigcache.DefaultConfig(20 * time.Second)` to `bigcache.DefaultConfig(5 * time.Minute)` and update the inline `Config` struct `LifeWindow` accordingly
- [ ] 1.3 Regenerate eBPF skeletons: `go generate ./internal/vxlan_agent/ebpf/`

## 2. Fix TC Return Code in vxlan_tc_external

- [ ] 2.1 In `internal/vxlan_agent/ebpf/c/vxlan_tc_external.bpf.c` line 141, replace `return AGENT_PASS;` with `return TC_ACT_OK;`
- [ ] 2.2 Audit the rest of `vxlan_tc_external.bpf.c` and `vxlan_tc_internal.bpf.c` for any other `AGENT_*` enum values used as TC return codes and replace them
- [ ] 2.3 Regenerate eBPF skeletons: `go generate ./internal/vxlan_agent/ebpf/`

## 3. Fix LPM Prefix Length in vxlan_xdp_external

- [ ] 3.1 In `internal/vxlan_agent/ebpf/c/vxlan_xdp_external.bpf.c` inside `handle_packet_received_by_external_iface__ip_packet`, change `struct ipv4_lpm_key key = {.prefixlen = 24};` to `struct ipv4_lpm_key key = {.prefixlen = 32};`
- [ ] 3.2 Verify the ARP path in the same file already uses `prefixlen = 32` (no change needed, just confirm)
- [ ] 3.3 Regenerate eBPF skeletons: `go generate ./internal/vxlan_agent/ebpf/`

## 4. Fix Ring Buffer Reader Lifecycle in switch_agent

- [ ] 4.1 In `internal/switch_agent/switch_agent.go`, refactor `handleNewDiscoveredEntriesRingBuffer` to call `ringbuf.NewReader` once before the `for` loop and `defer rd.Close()` at function scope
- [ ] 4.2 Add error handling for `rd.Read()` errors (log and break the loop rather than silently continuing)
- [ ] 4.3 Verify there are no other goroutines in the codebase that repeat the same pattern

## 5. Separate Broadcast and Unknown-Unicast Flooding Branches

- [ ] 5.1 In `internal/vxlan_agent/ebpf/c/vxlan_xdp_internal.bpf.c`, inside `handle_packet_received_by_internal_iface`, replace the single `else { return XDP_PASS; }` block with two explicit branches: one for `is_broadcast_address(&dst_mac)` and one for unknown unicast, both returning `XDP_PASS` but with distinct `my_bpf_printk` log messages labelling each case
- [ ] 5.2 Regenerate eBPF skeletons: `go generate ./internal/vxlan_agent/ebpf/`

## 6. Validate and Clean Up

- [ ] 6.1 Delete the old `openspecs/` folder and its contents (replaced by the proper `openspec/` structure)
- [ ] 6.2 Run `openspec validate --change fix-known-bugs-and-improvements` and resolve any issues
- [ ] 6.3 Build the project inside the devcontainer to confirm no compilation errors: `go build ./...`
- [ ] 6.4 Run the e2e tests inside the devcontainer: `./scripts/vxlan_agent_run_tests.sh`
