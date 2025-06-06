---
# generated by https://github.com/fbreckle/terraform-plugin-docs
page_title: "netbox_vlan_group Resource - terraform-provider-netbox"
subcategory: "IP Address Management (IPAM)"
description: |-
  A VLAN Group represents a collection of VLANs. Generally, these are limited by one of a number of scopes such as "Site" or "Virtualization Cluster".
---

# netbox_vlan_group (Resource)

> A VLAN Group represents a collection of VLANs. Generally, these are limited by one of a number of scopes such as "Site" or "Virtualization Cluster".

## Example Usage

```terraform
#Basic VLAN Group example
resource "netbox_vlan_group" "example1" {
  name = "example1"
  slug = "example1"
}

#Full VLAN Group example
resource "netbox_vlan_group" "example2" {
  name        = "Second Example"
  slug        = "example2"
  scope_type  = "dcim.site"
  scope_id    = netbox_site.example.id
  description = "Second Example VLAN Group"
  tags        = [netbox_tag.example.id]
  vid_ranges  = [[1, 2], [3, 4]]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `slug` (String)
- `vid_ranges` (List of List of Number)

### Optional

- `description` (String) Defaults to `""`.
- `scope_id` (Number) Required when `scope_type` is set.
- `scope_type` (String) Valid values are `dcim.location`, `dcim.site`, `dcim.sitegroup`, `dcim.region`, `dcim.rack`, `virtualization.cluster` and `virtualization.clustergroup`.
- `tags` (Set of String)

### Read-Only

- `id` (String) The ID of this resource.
- `tags_all` (Set of String)


