## MODIFIED Requirements

### Requirement: LPM trie lookups SHALL use prefixlen=32 for host-address keys
All LPM trie (`BPF_MAP_TYPE_LPM_TRIE`) lookups that supply a specific host IP address as the key MUST set `prefixlen = 32`. The trie performs longest-prefix matching internally; a query with `prefixlen=32` correctly matches any stored network entry whose prefix covers that host address. Using any value less than 32 (e.g. 24) silently restricts matching to networks of exactly that prefix length and breaks lookup for all other prefix lengths.

#### Scenario: IP packet arrives on external interface for a /16 network
- **WHEN** `vxlan_xdp_external` receives an encapsulated IP packet whose inner destination IP belongs to a configured `/16` network
- **THEN** the LPM lookup with `prefixlen=32` matches the `/16` entry in `networks_light_map` and the packet is decapsulated and forwarded

#### Scenario: IP packet arrives on external interface for a /24 network
- **WHEN** `vxlan_xdp_external` receives an encapsulated IP packet whose inner destination IP belongs to a configured `/24` network
- **THEN** the LPM lookup with `prefixlen=32` matches the `/24` entry and the packet is decapsulated and forwarded

#### Scenario: ARP and IP paths use consistent prefixlen
- **WHEN** either an ARP or IP packet is processed by any vxlan XDP program
- **THEN** both paths set `prefixlen=32` in the LPM key before calling `bpf_map_lookup_elem`, ensuring consistent behavior regardless of packet type
