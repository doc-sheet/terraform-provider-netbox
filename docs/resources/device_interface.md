---
# generated by https://github.com/fbreckle/terraform-plugin-docs
page_title: "netbox_device_interface Resource - terraform-provider-netbox"
subcategory: "Data Center Inventory Management (DCIM)"
description: |-
  From the official documentation https://docs.netbox.dev/en/stable/features/device/#interface:
  Interfaces in NetBox represent network interfaces used to exchange data with connected devices. On modern networks, these are most commonly Ethernet, but other types are supported as well. IP addresses and VLANs can be assigned to interfaces.
---

# netbox_device_interface (Resource)

From the [official documentation](https://docs.netbox.dev/en/stable/features/device/#interface):

> Interfaces in NetBox represent network interfaces used to exchange data with connected devices. On modern networks, these are most commonly Ethernet, but other types are supported as well. IP addresses and VLANs can be assigned to interfaces.

## Example Usage

```terraform
// Assumes a device with ID 123 exists
resource "netbox_device_interface" "test" {
  name      = "testinterface"
  device_id = 123
  type      = "1000base-t"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (Number)
- `name` (String)
- `type` (String)

### Optional

- `description` (String)
- `enabled` (Boolean) Defaults to `true`.
- `label` (String)
- `lag_device_interface_id` (Number) If this device is a member of a LAG group, you can reference the LAG interface here.
- `mgmtonly` (Boolean)
- `mode` (String) Valid values are `access`, `tagged`, `tagged-all` and `q-in-q`.
- `mtu` (Number)
- `parent_device_interface_id` (Number) The netbox_device_interface id of the parent interface. Useful if this interface is a logical interface.
- `speed` (Number)
- `tagged_vlans` (Set of Number)
- `tags` (Set of String)
- `untagged_vlan` (Number)

### Read-Only

- `id` (String) The ID of this resource.
- `mac_address` (String) The MAC address as string from the first MAC address assigned to this interface, if any.
- `mac_addresses` (Set of Object) (see [below for nested schema](#nestedatt--mac_addresses))
- `tags_all` (Set of String)

<a id="nestedatt--mac_addresses"></a>
### Nested Schema for `mac_addresses`

Read-Only:

- `description` (String)
- `id` (Number)
- `mac_address` (String)


