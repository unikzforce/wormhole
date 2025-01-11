#ifndef VXLAN_COMMON_BPF_H
#define VXLAN_COMMON_BPF_H

#include "../../../../include/vmlinux.h"

#include <bpf/bpf_endian.h>
#include <bpf/bpf_helpers.h>

#define TC_ACT_SHOT 2
#define TC_ACT_OK 0

#define ETH_P_IP 0x0800  /* Internet Protocol packet */
#define ETH_P_ARP 0x0806 /* Address Resolution packet */

#define ARPOP_REQUEST 1 /* ARP request */
#define ARPOP_REPLY 2   /* ARP reply */

#define CLOCK_BOOTTIME 7 /* Monotonic system-wide clock that includes time spent in suspension.  */

#define ETH_ALEN 6  /* Ethernet address length */
#define ETH_HLEN 14 /* Total octets in header.	 */
#define IP_HDR_LEN (int)sizeof(struct iphdr)
#define UDP_HDR_LEN (int)sizeof(struct udphdr)
#define VXLAN_HDR_LEN (int)sizeof(struct vxlanhdr)
#define NEW_HDR_LEN (ETH_HLEN + IP_HDR_LEN + UDP_HDR_LEN + VXLAN_HDR_LEN)

#define FIVE_MINUTES_IN_NS 20000000000

static volatile const struct
{
    char hostname[64];
    char type[64];
    char sub_type[64];
} vxlan_agent_metadata SEC(".rodata");

#define my_bpf_printk(fmt, args...) bpf_printk("%s :: %s :: %s :::: " fmt, vxlan_agent_metadata.hostname, vxlan_agent_metadata.type, vxlan_agent_metadata.sub_type, ##args)

enum vxlan_agent_processing_error
{
    AGENT_ERROR_ABORT = 0,
    AGENT_ERROR_DROP = 1,
    AGENT_NO_ERROR = 2,
    AGENT_PASS = 3,
};

struct mac_table_entry
{
    struct bpf_timer expiration_timer; // the timer object to expire this mac entry from the map in 5 minutes
    __u32 ifindex;                     // interface which mac address is learned from
    __u64 last_seen_timestamp_ns;      // last time this mac address was seen
    struct in_addr border_ip;          // remote agent border ip address that this mac address is learned from.
                                       // - in case of an internal mac address, this field is not used and set to 0.0.0.0
                                       // - in case of an external mac address, this field is used and set to the remote agent border ip address
};

struct external_route_info
{
    __u32 external_iface_index;
    struct mac_address external_iface_mac;
    struct mac_address external_iface_next_hop_mac;
    struct in_addr external_iface_ip;
};

struct arp_payload
{
    unsigned char ar_sha[ETH_ALEN]; // Sender hardware address
    unsigned char ar_sip[4];        // Sender IP address
    unsigned char ar_tha[ETH_ALEN]; // Target hardware address
    unsigned char ar_tip[4];        // Target IP address
};

struct ipv4_lpm_key
{
    __u32 prefixlen;
    __u8 data[4];
};

#define MAX_INTERNAL_IFINDEXES 10
#define MAX_BORDER_IPS 10

struct network_vni
{
    __u32 vni;
    struct ipv4_lpm_key network;
    __u32 internal_ifindexes[MAX_INTERNAL_IFINDEXES];
    __u32 internal_ifindexes_size;
    struct in_addr border_ips[MAX_BORDER_IPS];
    __u32 border_ips_size;
};

struct network_vni_light
{
    __u32 vni;
    struct ipv4_lpm_key network;
};

// Function to get a random port within the ephemeral range
static __always_inline __u16
get_ephemeral_port()
{
    return 49152 + bpf_get_prandom_u32() % (65535 - 49152 + 1);
}

// --------------------------------------------------------

// Function to check if the MAC address is the broadcast address
static __always_inline bool is_broadcast_address(const struct mac_address *mac)
{
    // Check if the MAC address is the broadcast address (FF:FF:FF:FF:FF:FF)
    if (mac->addr[0] != 0xFF)
        return false;
    if (mac->addr[1] != 0xFF)
        return false;
    if (mac->addr[2] != 0xFF)
        return false;
    if (mac->addr[3] != 0xFF)
        return false;
    if (mac->addr[4] != 0xFF)
        return false;
    if (mac->addr[5] != 0xFF)
        return false;
    return true;
}

// the mac table expiration callback function
static int mac_table_expiration_callback(void *map, struct mac_address *key, struct mac_table_entry *value)
{
    __u64 passed_time = bpf_ktime_get_tai_ns() - value->last_seen_timestamp_ns;

    if (passed_time >= FIVE_MINUTES_IN_NS)
    {
        // if the mac entry is expired
        bpf_map_delete_elem(map, key);
    }
    else
    {
        // if the mac entry is not expired we need to restart the timer according to the remaining time
        bpf_timer_start(&value->expiration_timer, FIVE_MINUTES_IN_NS - passed_time, 0);
    }

    return 0;
}

// bool is_ip_in_network2(const struct ipv4_lpm_key *key, const struct in_addr *src_ip_addr)
// {

//     if (key == NULL || src_ip_addr == NULL)
//     {
//         return false; // Ensure pointers are valid
//     }

//     // Convert the network address (key->data) to a __u32 using bpf_ntohl
//     __u32 network_address = (key->data[0] << 24) |
//                             (key->data[1] << 16) |
//                             (key->data[2] << 8) |
//                             key->data[3];

//     // Convert the source IP address from network byte order to host byte order using bpf_ntohl
//     my_bpf_printk("Raw Source IP Address (s_addr): 0x%08x\n", src_ip_addr->s_addr);

//     my_bpf_printk("Raw Network Address: %d.%d.%d.%d\n",
//                   key->data[0],
//                   key->data[1],
//                   key->data[2],
//                   key->data[3]);

//     __u32 src_ip_fixed = bpf_ntohl(src_ip_addr->s_addr);

//     // Create the network mask based on the prefix length
//     __u32 mask = (key->prefixlen == 32) ? 0 : (~0U >> (32 - key->prefixlen));

//     // Apply the mask to the source IP address and the network address
//     __u32 masked_src_ip = src_ip_fixed & mask;
//     __u32 masked_network_address = network_address & mask;

//     // Debug prints
//     my_bpf_printk("Network Address: %d.%d.%d.%d\n",
//                   (network_address >> 24) & 0xFF,
//                   (network_address >> 16) & 0xFF,
//                   (network_address >> 8) & 0xFF,
//                   network_address & 0xFF);
//     my_bpf_printk("PASHM Source IP Address: %d.%d.%d.%d\n",
//                   (src_ip_fixed >> 24) & 0xFF,
//                   (src_ip_fixed >> 16) & 0xFF,
//                   (src_ip_fixed >> 8) & 0xFF,
//                   src_ip_fixed & 0xFF);
//     my_bpf_printk("Mask: 0x%08x\n", mask);
//     my_bpf_printk("Masked Source IP: 0x%08x\n", masked_src_ip);
//     my_bpf_printk("Masked Network Address: 0x%08x\n", masked_network_address);

//     // Check if they match
//     return masked_src_ip == masked_network_address;
// }

static __always_inline uint32_t create_mask(uint32_t prefixlen)
{
    switch (prefixlen)
    {
    case 0:
        return 0b00000000000000000000000000000000;
    case 1:
        return 0b00000000000000000000000000000001;
    case 2:
        return 0b00000000000000000000000000000011;
    case 3:
        return 0b00000000000000000000000000000111;
    case 4:
        return 0b00000000000000000000000000001111;
    case 5:
        return 0b00000000000000000000000000011111;
    case 6:
        return 0b00000000000000000000000000111111;
    case 7:
        return 0b00000000000000000000000001111111;
    case 8:
        return 0b00000000000000000000000011111111;
    case 9:
        return 0b00000000000000000000000111111111;
    case 10:
        return 0b00000000000000000000001111111111;
    case 11:
        return 0b00000000000000000000011111111111;
    case 12:
        return 0b00000000000000000000111111111111;
    case 13:
        return 0b00000000000000000001111111111111;
    case 14:
        return 0b00000000000000000011111111111111;
    case 15:
        return 0b00000000000000000111111111111111;
    case 16:
        return 0b00000000000000001111111111111111;
    case 17:
        return 0b00000000000000011111111111111111;
    case 18:
        return 0b00000000000000111111111111111111;
    case 19:
        return 0b00000000000001111111111111111111;
    case 20:
        return 0b00000000000011111111111111111111;
    case 21:
        return 0b00000000000111111111111111111111;
    case 22:
        return 0b00000000001111111111111111111111;
    case 23:
        return 0b00000000011111111111111111111111;
    case 24:
        return 0b00000000111111111111111111111111;
    case 25:
        return 0b00000001111111111111111111111111;
    case 26:
        return 0b00000011111111111111111111111111;
    case 27:
        return 0b00000111111111111111111111111111;
    case 28:
        return 0b00001111111111111111111111111111;
    case 29:
        return 0b00011111111111111111111111111111;
    case 30:
        return 0b00111111111111111111111111111111;
    case 31:
        return 0b01111111111111111111111111111111;
    case 32:
        return 0b11111111111111111111111111111111;
    default:
        return 0b00000000000000000000000000000000; // Invalid prefix length
    }
}

bool __always_inline is_ip_in_network(const struct ipv4_lpm_key *key, const struct in_addr *src_ip_addr)
{
    if (key == NULL || src_ip_addr == NULL)
    {
        return false; // Ensure pointers are valid
    }

    // create a copy of src_ip_addr
    struct in_addr src_ip_addr_copy;
    __builtin_memcpy(&src_ip_addr_copy, src_ip_addr, sizeof(src_ip_addr_copy));

    // create a copy of key
    struct ipv4_lpm_key key_copy;
    __builtin_memcpy(&key_copy, key, sizeof(key_copy));

    uint32_t network = (((uint32_t)key_copy.data[3]) << 24) |
                       (((uint32_t)key_copy.data[2]) << 16) |
                       (((uint32_t)key_copy.data[1]) << 8) |
                       ((uint32_t)key_copy.data[0]);

    // Ensure prefixlen is valid
    if (key_copy.prefixlen > 32)
    {
        my_bpf_printk("Invalid prefix length\n");
        return false;
    }

    uint32_t mask = create_mask(key_copy.prefixlen);

    uint32_t mask_copy;
    __builtin_memcpy(&mask_copy, &mask, sizeof(mask_copy));

    uint32_t src_ip_addr_copy_copy = bpf_ntohl(src_ip_addr_copy.s_addr);

    uint32_t src_ip_addr_copy_copy_copy;
    __builtin_memcpy(&src_ip_addr_copy_copy_copy, &src_ip_addr_copy_copy, sizeof(src_ip_addr_copy_copy));

    // Print out the source IP address in network byte order
    my_bpf_printk("Source IP original: 0x%08x\n", src_ip_addr_copy.s_addr);
    my_bpf_printk("Source IP original: 0x%08x\n", src_ip_addr_copy_copy);
    my_bpf_printk("Source IP original: 0x%08x\n", src_ip_addr_copy_copy_copy);

    // Print the manually converted host order source IP
    my_bpf_printk("Network Address: 0x%08x\n", network);
    my_bpf_printk("Prefix Length: %d, Mask: 0x%08x\n", key_copy.prefixlen, mask);

    // Apply the mask to both source IP and network address
    uint32_t src_masked = src_ip_addr_copy_copy_copy & mask_copy;
    uint32_t src_masked_2 = src_ip_addr->s_addr & mask_copy;
    uint32_t network_masked = network & mask;

    // Print out the masked values for debugging
    my_bpf_printk("Masked Source IP: 0x%08x\n", src_masked);
    my_bpf_printk("Masked Source IP 2: 0x%08x\n", src_masked_2);
    my_bpf_printk(" Masked Network Address: 0x%08x\n", network_masked);

    // Compare the masked source IP with the network address
    if (src_masked == network_masked || src_masked_2 == network_masked)
    {
        my_bpf_printk("TRUEEEEEEEEEEEEEEE \n");
    }

    return (src_masked == network_masked || src_masked_2 == network_masked);
}

void create_in_addr_from_arp_ip(const unsigned char ar_ip[4], struct in_addr *src_ip_addr)
{
    if (src_ip_addr == NULL)
    {
        return; // Ensure the pointer is valid
    }

    // Copy the 4-byte IP address into the in_addr structure
    __builtin_memcpy(&src_ip_addr->s_addr, ar_ip, sizeof(src_ip_addr->s_addr));
}

#define GENERATE_DUMMY_MAP(type_name)    \
    struct                               \
    {                                    \
        __uint(type, BPF_MAP_TYPE_HASH); \
        __uint(max_entries, 1);          \
        __type(key, int);                \
        __type(value, struct type_name); \
    } dummy_##type_name SEC(".maps");

#endif