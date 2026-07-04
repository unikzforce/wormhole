## ADDED Requirements

### Requirement: TC programs SHALL use only TC-layer return codes
All eBPF programs attached via `SEC("tcx/ingress")` or any TC hook MUST return only valid TC action codes (`TC_ACT_OK = 0`, `TC_ACT_SHOT = 2`). XDP-layer enums (`AGENT_PASS`, `AGENT_ERROR_ABORT`, `AGENT_ERROR_DROP`, `AGENT_NO_ERROR`) are defined in `enum vxlan_agent_processing_error` for use exclusively within XDP helper functions and MUST NOT appear as return values in TC programs.

#### Scenario: Packet does not belong to an internal network in TC external program
- **WHEN** `vxlan_tc_external` receives a packet whose src/dst IPs do not belong to any configured internal network
- **THEN** the program returns `TC_ACT_OK` (0) to pass the packet up the stack, not `AGENT_PASS` (3)

#### Scenario: TC program drops a malformed packet
- **WHEN** a TC program encounters a packet it cannot process (e.g. undersized, missing headers)
- **THEN** the program returns `TC_ACT_SHOT` (2) to drop the packet, not `XDP_DROP` or any XDP enum value

#### Scenario: XDP helper functions may use AGENT_* enums internally
- **WHEN** an XDP helper function (called only from XDP entry points) signals a processing error
- **THEN** it returns a value from `enum vxlan_agent_processing_error`; the XDP entry point translates this to a valid XDP return code before returning to the kernel
