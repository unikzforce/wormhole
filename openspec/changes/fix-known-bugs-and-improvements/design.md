## Context

Wormhole is an eBPF XDP/TC-based VXLAN VTEP written in Go + C. Four eBPF programs (compiled via `bpf2go`) run in the Linux kernel fast path; a Go userspace runtime loads them, populates maps, and attaches them to network interfaces. The bugs are localized — five independent one-to-three-line fixes across three C files and one Go file — with no architectural changes required.

All eBPF source is under `internal/vxlan_agent/ebpf/c/` and `internal/switch_agent/ebpf/c/`. After any `.bpf.c` edit the corresponding `go generate` target must be re-run to regenerate the `_bpfel.go` / `_bpfeb.go` skeleton files and embedded object files.

## Goals / Non-Goals

**Goals:**
- Correct the `FIVE_MINUTES_IN_NS` constant to 300,000,000,000 ns (5 min)
- Replace `AGENT_PASS` with `TC_ACT_OK` in `vxlan_tc_external`
- Fix `prefixlen=24` → `prefixlen=32` in the IP-packet LPM lookup in `vxlan_xdp_external`
- Fix the ring buffer reader lifecycle in `switch_agent.go`
- Separate broadcast vs. unknown-unicast branches in `vxlan_xdp_internal`
- Regenerate eBPF skeletons after C changes

**Non-Goals:**
- Runtime behaviour changes observable by end users beyond bug fixes
- Any eBPF map schema changes (no migration needed)
- IPv6 support, multi-VNI routing, or control-plane distribution of MAC tables
- Full RFC 7348 compliance

## Decisions

### Decision 1 — Fix `FIVE_MINUTES_IN_NS` as a constant, not a config value

**Choice:** Change the `#define` in `vxlan_common.bpf.h` from `20000000000` to `300000000000ULL`.

**Rationale:** The TTL is a protocol-level constant (analogous to ARP cache TTL), not an operator knob. Making it configurable would require passing it through a BPF map or `.rodata` constant, adding complexity for no current benefit. A simple constant fix is sufficient.

**Alternative considered:** Make it a `.rodata` constant injectable from Go at load time. Rejected — over-engineering for a PoC.

---

### Decision 2 — Replace `AGENT_PASS` return with `TC_ACT_OK` in `vxlan_tc_external`

**Choice:** Single-line replacement at `vxlan_tc_external.bpf.c:141`.

**Rationale:** `AGENT_PASS = 3` is coincidentally equal to `TC_ACT_PIPE` (pass to next BPF program in chain) which works when there is no chained program, but is semantically wrong and fragile. `TC_ACT_OK = 0` is the correct "let the kernel continue processing" code.

**Alternative considered:** Remove the `AGENT_PASS` enum entirely and audit all usages. Deferred — the enum is legitimately used within XDP helper functions as an internal signal; removing it is a larger refactor.

---

### Decision 3 — Fix `prefixlen=24` → `prefixlen=32` in `vxlan_xdp_external` IP path only

**Choice:** Change only the IP-packet branch in `vxlan_xdp_external`. The ARP branch already uses 32.

**Rationale:** The LPM trie key `prefixlen` in a lookup controls which bits of `data` the kernel compares against stored keys. For a host-address query, all 32 bits should be presented so the trie can find the best (longest) matching prefix. Using 24 restricts matching to exactly 8-bit-aligned /24 entries.

**Alternative considered:** Use `prefixlen = 0` to match any stored prefix. Rejected — semantically wrong; 0 means "match everything regardless of address".

---

### Decision 4 — Move `ringbuf.NewReader` above the loop, let `defer` handle close

**Choice:** Restructure `handleNewDiscoveredEntriesRingBuffer` to create the reader once, defer its close, then loop on `rd.Read()`.

**Rationale:** `defer` is function-scoped. Calling it inside a loop body defers N closes to function exit, which never arrives during normal operation — effectively leaking one fd per event. The fix is idiomatic Go.

**Alternative considered:** Explicitly call `rd.Close()` at the bottom of each loop iteration. Rejected — loses the safety guarantee of `defer` under error paths and early returns.

---

### Decision 5 — Separate broadcast and unknown-unicast via explicit `if` branches

**Choice:** Add an `is_broadcast_address(dst_mac)` check before the `mac_table` lookup in `vxlan_xdp_internal`. Both paths currently return `XDP_PASS`, but they are now clearly labelled.

**Rationale:** The spec requires the distinction to be explicit in code so future rate-limiting or suppression of unknown-unicast flooding can be added without restructuring logic. No behaviour change yet — this is a code-clarity fix that enables future work.

**Alternative considered:** Add rate-limiting for unknown unicast immediately. Deferred — out of scope for this bug-fix change.

## Risks / Trade-offs

- **[Risk] eBPF regeneration step is manual** → After each `.bpf.c` change, `go generate` must be run inside the devcontainer. If skipped the embedded object files will be stale. Mitigation: tasks.md makes this explicit as a task step after every C file edit.

- **[Risk] TTL change from 20 s to 5 min may mask other bugs in testing** → Short TTL was useful for quickly observing eviction in development. Mitigation: developers can temporarily redefine `FIVE_MINUTES_IN_NS` locally; this is a debug concern, not a production concern.

- **[Risk] `TC_ACT_PIPE` (3) was accidentally working** → If another TCX program is chained after `vxlan_tc_external` in any deployment, changing to `TC_ACT_OK` will stop the chain. Mitigation: current codebase attaches only one TC program per interface, so no chain exists.

## Open Questions

- Should `FIVE_MINUTES_IN_NS` be promoted to a `.rodata` constant so it can be set per-agent from the YAML config? (Deferred to a future change.)
- Should unknown-unicast flooding be rate-limited or suppressed entirely after a configurable number of events? (Deferred to `flooding-policy` follow-up change.)
