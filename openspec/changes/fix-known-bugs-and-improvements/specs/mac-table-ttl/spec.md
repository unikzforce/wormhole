## MODIFIED Requirements

### Requirement: MAC table entries SHALL expire after 5 minutes of inactivity
The eBPF `mac_table` and the userspace eviction cache SHALL use a TTL of exactly 5 minutes (300 seconds). The constant `FIVE_MINUTES_IN_NS` in `vxlan_common.bpf.h` SHALL be set to `300000000000` (300 × 10⁹ nanoseconds). The `switch_agent` userspace bigcache lifetime window SHALL be set to `5 * time.Minute`. Both values MUST remain in sync.

#### Scenario: MAC entry expires after 5 minutes of silence
- **WHEN** a MAC address is learned and no further packet from that MAC is seen for 5 minutes
- **THEN** the `bpf_timer` callback fires, the entry is removed from `mac_table`, and subsequent packets to that MAC trigger unknown-unicast flooding

#### Scenario: MAC entry timer resets on activity
- **WHEN** a packet from a known source MAC arrives before the 5-minute timer fires
- **THEN** the timer is cancelled and restarted from zero, keeping the entry alive for another 5 minutes

#### Scenario: Userspace eviction cache stays in sync with kernel
- **WHEN** the userspace bigcache TTL for a MAC entry expires
- **THEN** the eviction callback checks the kernel `mac_table` timestamp; if the entry was seen within the last 5 minutes the key is re-inserted, otherwise the kernel entry is also deleted
