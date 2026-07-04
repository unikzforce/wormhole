# Wormhole — Product Scenarios & Teaching Guide

A self-contained, interactive HTML guide exploring how the Wormhole eBPF XDP data plane
(this repository) could grow into a commercial product. Written as a teaching artifact:
every scenario explains the networking concepts it uses.

## How to open

No server, no build step — it is one file:

```sh
open docs/tutorial/index.html        # macOS
xdg-open docs/tutorial/index.html    # Linux
```

Navigate with the tab bar or ←/→ arrow keys. Deep-link a tab with the URL hash,
e.g. `index.html#crypto`.

## What's inside

| Tab | Scenario | Core teaching |
|-----|----------|---------------|
| 00 Fundamentals | The repo's own containerlab lab, wire by wire | VTEP, VNI, underlay/overlay, flood-and-learn, XDP vs kernel path, encapsulation anatomy |
| 01 The Landscape | Competitive positioning | L2/L3/L4 virtualization × data-plane location; why the (L2, XDP) corner is empty |
| 02 Hybrid Multi-Cloud | One L2 segment across AWS + GCP + colo | tap-device host mode vs gateway mode, cloud L2 restrictions, MSS/PMTUD, NAT traversal |
| 03 Cloud Provider VPC | Regional cloud building a VPC product | Spine-leaf Clos, /31 underlay, VNI multi-tenancy with overlapping CIDRs, OVS economics |
| 04 Industrial Edge | 12 factories, identical golden-image subnets | Non-routable OT protocols (PROFINET DCP), IEC 62443 zones, ARP suppression, dual-WAN |
| 05 Private 5G UPF | XDP as a 5G user plane | GTP-U ≅ VXLAN, N3/N4/N6 reference points, PFCP→BPF-map translation, ULCL breakout |
| 06 Security & Observability | East-west NDR from the same agent | Flow maps + ringbuf export, first-seen-edge detection, SPAN blind spots |
| 07 Crypto Deep-Dive | bpf_crypto in the data plane | Wire format, key_id in VXLAN reserved bits, GCM nonce discipline, zero-loss key rotation |

## Diagram conventions

All topology diagrams are drawn at **wire level**: every node shows its real interfaces
(`eth0`, `tap0`, `ens1`, switch ports `g1/1`…) with per-interface IPs and MACs; lines are
cables between specific ports; dashed boxes delimit subnets/broadcast domains. Color code:

- **cyan** — overlay-side ports and tunnels
- **amber** — underlay/WAN-side ports
- **green** — encrypted paths, data-plane components
- **violet** — control plane (Controller, PFCP, mTLS)
- **red** — threats / hard-real-time exclusion zones

Diagrams are inline SVG sharing one `<defs>` sprite (markers, filters, symbols) at the top
of the file — extract any `<figure>` block and it will render standalone if you copy the
defs block along with it.

## Regenerating / extending

Hand-written file; edit directly. Conventions to preserve:

- One `<section class="panel" id="tab-…">` per tab + matching `data-tab` button in the nav
- SVG styling lives in the page CSS (classes `nb`, `tt12`, `tmono10`, `lnk-*`, …), not inline
- Animation uses CSS `stroke-dasharray` flows + SMIL `animateMotion`; both are stripped
  automatically when the viewer has `prefers-reduced-motion` set
- Ballpark performance numbers are teaching aids — benchmark before quoting them anywhere

Related planning: `openspec/changes/fix-known-bugs-and-improvements/` tracks the data-plane
bug fixes referenced by the guide; `AGENTS.md` documents the repository itself.
