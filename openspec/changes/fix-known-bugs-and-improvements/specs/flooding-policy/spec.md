## ADDED Requirements

### Requirement: XDP programs SHALL distinguish broadcast from unknown-unicast packets
The XDP internal program (`vxlan_xdp_internal`) MUST apply different handling to broadcast frames and unknown-unicast frames:

- **Broadcast** (`FF:FF:FF:FF:FF:FF` destination MAC): MUST always be passed to the TC flooding path via `XDP_PASS`, as all nodes on the virtual segment need to receive it.
- **Unknown unicast** (destination MAC not in `mac_table`, destination is not broadcast): MUST also be passed to `XDP_PASS` for TC flooding during the learning phase, but this path SHALL be documented separately and may in future revisions be suppressed or rate-limited.

The current behavior of treating both cases identically is acceptable as a first implementation, but the code MUST make the distinction explicit via separate conditional branches or clearly labelled comments so future differentiation is straightforward.

#### Scenario: Broadcast ARP request arrives on internal interface
- **WHEN** a packet with destination MAC `FF:FF:FF:FF:FF:FF` arrives on an internal NIC
- **THEN** `vxlan_xdp_internal` returns `XDP_PASS` so `vxlan_tc_internal` clones and floods it to all internal interfaces and all remote border IPs

#### Scenario: Unknown unicast arrives on internal interface
- **WHEN** a packet with a unicast destination MAC not present in `mac_table` arrives on an internal NIC
- **THEN** `vxlan_xdp_internal` returns `XDP_PASS` so `vxlan_tc_internal` floods it (learning phase behaviour)

#### Scenario: Known unicast takes the fast path
- **WHEN** a packet with a destination MAC present in `mac_table` arrives on an internal NIC
- **THEN** `vxlan_xdp_internal` redirects it directly via `bpf_redirect` without involving the TC layer

#### Scenario: Code separates broadcast and unknown-unicast branches
- **WHEN** a developer reads the XDP internal program source
- **THEN** broadcast and unknown-unicast cases are handled in distinct `if`/`else if` branches or clearly labelled sections, not collapsed into a single fallthrough
