# AGENTS.md — Wormhole Project

This file is the authoritative reference for AI coding agents working in this repository. Read it fully before making any changes.

---

## Project Overview

**Wormhole** is a proof-of-concept eBPF XDP/TC-based VXLAN VTEP (Virtual Tunnel Endpoint) written in Go + C. It demonstrates that VXLAN encapsulation, decapsulation, MAC learning, and unknown-unicast flooding can be implemented entirely in the Linux kernel fast path using eBPF, bypassing the Linux networking stack.

This is a **learning/PoC project** — the author acknowledges possible VXLAN protocol non-compliance. Suggestions and changes should prioritize clarity and correctness over full RFC compliance.

---

## Repository Layout

```
wormhole/
├── cmd/
│   ├── vxlan_agent/          # VXLAN VTEP binary entry point (reads YAML config)
│   ├── switch_agent/         # L2 switch binary entry point
│   ├── dummy_xdp/            # Trivial XDP passthrough (used on src/dst test nodes)
│   └── test_agent/           # gRPC test harness server (runs inside test containers)
│       ├── generated/        # protobuf-generated Go code (do not edit manually)
│       ├── impl/             # gRPC service implementation
│       └── test_agent.proto  # service definition
├── internal/
│   ├── vxlan_agent/
│   │   ├── vxlan_agent.go    # Go runtime: load eBPF, init maps, attach programs
│   │   ├── mapper.go         # MAC address conversion helpers
│   │   └── ebpf/
│   │       ├── gen/          # bpf2go-generated Go skeletons (do not edit manually)
│   │       ├── gen.go        # go:generate directives
│   │       └── c/            # eBPF C source files (edit these)
│   │           ├── vxlan_common.bpf.h          # shared structs, maps, constants
│   │           ├── vxlan_xdp_internal.bpf.c    # XDP on internal NICs
│   │           ├── vxlan_xdp_external.bpf.c    # XDP on external NICs
│   │           ├── vxlan_tc_internal.bpf.c     # TC on internal NICs (flooding)
│   │           └── vxlan_tc_external.bpf.c     # TC on external NICs (flooding)
│   └── switch_agent/
│       ├── switch_agent.go   # Go runtime: load eBPF, userspace MAC table, ring buf
│       ├── mapper.go         # MAC address conversion helpers
│       └── ebpf/
│           ├── gen.go        # go:generate directives
│           ├── (generated)   # bpf2go-generated skeletons
│           └── c/
│               ├── switch_agent_xdp.bpf.c                        # XDP: MAC learning + redirect
│               └── switch_agent_unknown_unicast_flooding.bpf.c   # TC: clone + flood
├── include/
│   └── vmlinux.h             # kernel type definitions (generated via bpftool)
├── test/e2e/
│   ├── vxlan_agent/          # Ginkgo e2e tests for vxlan_agent
│   │   ├── vxlan_agent_test.go
│   │   └── clab-topologies/vxlan.clab.yml
│   └── switch_agent/         # Ginkgo e2e tests for switch_agent
│       └── clab-topologies/switch.clab.yml
├── scripts/
│   ├── build_images.sh           # build wormhole/test_agent Docker image
│   ├── vxlan_agent_run_tests.sh  # run e2e tests
│   ├── generate_grpc.sh          # regenerate protobuf/gRPC Go code
│   ├── generate_vmlinux_header.sh
│   ├── ebpf_explorer_start.sh
│   └── edgeshark_start.sh
├── openspec/                 # OpenSpec change management (see below)
│   ├── config.yaml
│   ├── changes/              # in-flight change proposals
│   └── specs/                # stable per-capability specs
├── readme/                   # diagrams referenced by README.md
├── .devcontainer/            # VSCode devcontainer definition
└── .vscode/
```

---

## Development Environment

**All development MUST happen inside the devcontainer.** The host machine only needs:
- Linux kernel 6.5 or newer
- Docker
- VSCode (or any IDE with devcontainer support)

Open the project in VSCode → accept the "Reopen in Container" prompt. The devcontainer installs:
- Go, Clang, LLVM, libbpf, bpftool, xdp-tools
- ContainerLab
- All other eBPF build dependencies

Do **not** attempt to compile eBPF C files or run tests on macOS or a non-Linux host.

---

## Build & Generate Commands

### Regenerate eBPF Go skeletons (required after any `.bpf.c` edit)
```sh
go generate ./internal/vxlan_agent/ebpf/
go generate ./internal/switch_agent/ebpf/
go generate ./internal/dummy_xdp/ebpf/
```
This runs `bpf2go` to compile the C files and embed the resulting ELF as Go files in `ebpf/gen/`.

### Regenerate protobuf/gRPC (required after editing `test_agent.proto`)
```sh
./scripts/generate_grpc.sh
```

### Regenerate vmlinux.h (rarely needed)
```sh
./scripts/generate_vmlinux_header.sh
```

### Build all binaries
```sh
go build ./...
```

### Build the test Docker image (must be done before running e2e tests)
```sh
./scripts/build_images.sh
```
This calls both generate scripts and then runs `docker build` producing `wormhole/test_agent:latest`.

---

## Running Tests

### e2e tests (require devcontainer + Docker-in-Docker)
```sh
./scripts/vxlan_agent_run_tests.sh
```
Tests use **ContainerLab** to spin up a 4-node topology (`src`, `border1`, `border2`, `dst`), communicate with each container via **gRPC** (`test_agent` server), run pings, and assert connectivity.

### Interactive testing (manual)
```sh
cd test/e2e/vxlan_agent/clab-topologies/
clab deploy
# then exec into containers and run binaries manually
# see README.md for the full walkthrough
clab destroy --all   # clean up when done
```

---

## Architecture — VXLAN Agent

### Roles of each eBPF program

| Program | Hook | Attached to | Responsibility |
|---|---|---|---|
| `vxlan_xdp_internal` | XDP (`xdp.frags`) | Internal NICs | MAC learning; lookup dst MAC → redirect (fast path) or XDP_PASS (flood) |
| `vxlan_xdp_external` | XDP (`xdp.frags`) | External NICs | Decapsulate VXLAN; MAC learning (records remote border IP); redirect or XDP_PASS |
| `vxlan_tc_internal` | TCX ingress | Internal NICs | Unknown unicast flooding: clone to all internal ifaces + encapsulate + send to all remote border IPs |
| `vxlan_tc_external` | TCX ingress | External NICs | Flooding of decapsulated packet to all internal ifaces |

### Packet lifecycle (internal → external)
1. Internal host sends packet → hits `vxlan_xdp_internal`
2. Source MAC learned; destination MAC looked up in `mac_table`
3. **MAC found + internal iface** → `bpf_redirect` (fast path, stays in XDP)
4. **MAC found + external iface** → encapsulate (add outer Eth/IP/UDP/VXLAN via `bpf_xdp_adjust_head`) + `bpf_redirect`
5. **MAC not found** → `XDP_PASS` → `vxlan_tc_internal` clones and floods to all internal ifaces and all remote border IPs (slow path)

### Packet lifecycle (external → internal)
1. Encapsulated packet arrives at external NIC → hits `vxlan_xdp_external`
2. Inner src MAC + outer src IP learned into `mac_table` (with `border_ip` set)
3. Inner dst MAC looked up
4. **MAC found** → strip outer headers (`bpf_xdp_adjust_head`) + `bpf_redirect`
5. **MAC not found** → `XDP_PASS` → `vxlan_tc_external` decapsulates + floods to all internal ifaces

### Shared eBPF maps (pinned to `/sys/fs/bpf`, named-pinned so all 4 programs share them)

| Map | Type | Key | Value | Purpose |
|---|---|---|---|---|
| `mac_table` | HASH | `mac_address` (6 bytes) | `mac_table_entry` (ifindex + timestamp + border_ip + bpf_timer) | MAC → outgoing iface/border mapping |
| `networks_map` | LPM_TRIE | `ipv4_lpm_key` (prefixlen + 4 bytes) | `network_vni` (vni + internal ifindexes[] + border_ips[]) | CIDR → network info; used by internal XDP + both TC programs |
| `networks_light_map` | LPM_TRIE | `ipv4_lpm_key` | `network_vni_light` (vni + network) | Lighter version for external XDP |
| `border_ip_to_route_info_map` | HASH | `in_addr` (u32) | `external_route_info` (iface_index + iface_mac + next_hop_mac + iface_ip) | Border IP → L2 routing info for encapsulation |
| `ifindex_is_internal_map` | HASH | u32 (ifindex) | bool | True = internal iface, False = external iface |

### MAC table entry TTL
Each `mac_table_entry` embeds a `bpf_timer`. On every packet, the source MAC timer is cancelled and restarted. If no packet is seen for **5 minutes**, the timer callback (`mac_table_expiration_callback`) deletes the entry. The constant is `FIVE_MINUTES_IN_NS` in `vxlan_common.bpf.h`.

### Jumbo frame support
All VXLAN XDP programs use `SEC("xdp.frags")` instead of `SEC("xdp")` to support MTUs above 1500 bytes. Only the first page of the packet is directly accessible, which is sufficient for header processing.

---

## Architecture — Switch Agent

A simpler L2 switch (no VXLAN) using two eBPF programs:
- `switch_agent_xdp` (XDP): MAC learning via ring buffer → `new_discovered_entries_rb`; redirect known MACs; `XDP_PASS` for unknown
- `switch_agent_unknown_unicast_flooding` (TC): `bpf_clone_redirect` to all switch ports except ingress

The Go runtime (`switch_agent.go`) maintains a **userspace MAC table** (bigcache, 5-minute TTL) that mirrors the kernel map and drives eviction: on expiry it checks the kernel timestamp and either re-inserts (if recently active) or deletes the kernel entry.

---

## eBPF Coding Rules

1. **Always validate packet bounds** before accessing any header field: `if ((void*)(hdr + 1) > data_end) return XDP_DROP;`
2. **XDP programs** return: `XDP_PASS`, `XDP_DROP`, `XDP_REDIRECT`, `XDP_ABORTED` only.
3. **TC programs** return: `TC_ACT_OK` (0) or `TC_ACT_SHOT` (2) only. Never use `AGENT_PASS` or any `enum vxlan_agent_processing_error` value as a TC return.
4. **`enum vxlan_agent_processing_error`** (`AGENT_NO_ERROR`, `AGENT_ERROR_ABORT`, `AGENT_ERROR_DROP`, `AGENT_PASS`) is only for XDP internal helper function return types, not for TC programs.
5. **LPM trie lookups** for host addresses must set `prefixlen = 32` in the key so the kernel finds the longest matching prefix.
6. **Packet cloning** (`bpf_clone_redirect`) is only available in TC programs, not XDP.
7. **Encapsulation** uses `bpf_xdp_adjust_head(ctx, -NEW_HDR_LEN)` (negative = grow at head) in XDP and `bpf_skb_change_head(skb, NEW_HDR_LEN, 0)` in TC.
8. **Decapsulation** uses `bpf_xdp_adjust_head(ctx, NEW_HDR_LEN)` (positive = shrink at head) in XDP and `bpf_skb_adjust_room(skb, -NEW_HDR_LEN, BPF_ADJ_ROOM_MAC, 0)` in TC.
9. After any `.bpf.c` change, **always regenerate** the Go skeletons via `go generate`.

---

## Go Coding Rules

1. **Ring buffer readers** (`ringbuf.NewReader`) must be created once before a read loop and closed via `defer` at function scope — never inside a loop body.
2. **Error handling**: never silently discard errors from `ringbuf.Read()`, `bpf_map.Lookup()`, or `link.Attach*()`.
3. **`logrus.Fatal`** terminates the process — use it only for unrecoverable startup errors, not inside goroutines that should stay alive.
4. **MAC address helpers** (`ConvertMacToString`, `ConvertStringToMac`, `ConvertMacBytesToMac`) are defined in `mapper.go` for each agent package — use them, do not re-implement.

---

## OpenSpec Workflow

This project uses [`@fission-ai/openspec`](https://www.npmjs.com/package/@fission-ai/openspec) for spec-driven development.

```
openspec/
├── config.yaml         # schema: spec-driven
├── specs/              # stable, long-lived per-capability specs (one subfolder each)
└── changes/            # in-flight changes (proposal + design + specs + tasks)
    └── <change-name>/
        ├── proposal.md
        ├── design.md
        ├── tasks.md
        └── specs/      # delta specs for this change
```

### Key commands
```sh
openspec init --tools claude         # already done; do not re-run
openspec new change <name>           # start a new change
openspec status --change <name>      # check artifact completion
openspec instructions <artifact> --change <name> --json   # get writing instructions
openspec validate <change-name>      # validate a change
openspec validate --specs            # validate all main specs
openspec list --specs                # list all stable specs
```

### Active changes

| Change | Status | Description |
|---|---|---|
| `fix-known-bugs-and-improvements` | Open | Fix 5 bugs found in code review (TTL constant, TC return code, LPM prefixlen, ringbuf lifecycle, flooding policy separation) |

---

## Known Bugs (Open)

See `openspec/changes/fix-known-bugs-and-improvements/` for full details. Summary:

| # | File | Bug |
|---|---|---|
| 1 | `vxlan_common.bpf.h:27` | `FIVE_MINUTES_IN_NS = 20000000000` (20 s) should be `300000000000` (5 min) |
| 2 | `vxlan_tc_external.bpf.c:141` | Returns `AGENT_PASS` (XDP enum = 3) inside a TC program; must be `TC_ACT_OK` (0) |
| 3 | `vxlan_xdp_external.bpf.c:~301` | IP-packet LPM lookup uses `prefixlen=24`; must be `prefixlen=32` |
| 4 | `switch_agent.go:241` | `ringbuf.NewReader` created inside loop body — leaks one fd per MAC event |
| 5 | `vxlan_xdp_internal.bpf.c` | Broadcast and unknown-unicast packets collapsed into single `XDP_PASS` branch |

---

## Test Topology (vxlan_agent e2e)

```
[src 192.168.1.10/24]──eth1──[border1]──eth2──[border2]──eth1──[dst 192.168.1.11/24]
                               eth2 IP:            eth2 IP:
                               3.3.3.1/24          3.3.3.2/24
```

- `src` and `dst` are in the same logical subnet but physically separated by the border nodes.
- `border1` and `border2` run `vxlan_agent` as the VTEP software.
- `src` and `dst` run `dummy_xdp` (trivial XDP passthrough) on their internal interface.
- Each container runs `test_agent` (gRPC server on port 9000) that the test harness uses to start programs and issue ping requests remotely.
- The e2e test asserts: **without** vxlan_agent → src/dst cannot ping; **with** vxlan_agent → src/dst can ping.

---

## Useful Debugging Tools

| Tool | How to start | Purpose |
|---|---|---|
| `bpftool` | available in devcontainer | inspect loaded eBPF programs and maps |
| Edgeshark | `./scripts/edgeshark_start.sh` | browser-based packet capture on container interfaces |
| eBPF Explorer | `./scripts/ebpf_explorer_start.sh` | visualize eBPF maps and programs |
| `bpf_printk` / `my_bpf_printk` | `cat /sys/kernel/debug/tracing/trace_pipe` | kernel-side eBPF debug logging |
| Zellij layouts | `.zellij/` | pre-configured multi-pane terminal layouts for interactive testing |
