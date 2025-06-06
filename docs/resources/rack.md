---
# generated by https://github.com/fbreckle/terraform-plugin-docs
page_title: "netbox_rack Resource - terraform-provider-netbox"
subcategory: "Data Center Inventory Management (DCIM)"
description: |-
  From the official documentation https://docs.netbox.dev/en/stable/models/dcim/rack/:
  The rack model represents a physical two- or four-post equipment rack in which devices can be installed. Each rack must be assigned to a site, and may optionally be assigned to a location within that site. Racks can also be organized by user-defined functional roles. The name and facility ID of each rack within a location must be unique.
  Rack height is measured in rack units (U); racks are commonly between 42U and 48U tall, but NetBox allows you to define racks of arbitrary height. A toggle is provided to indicate whether rack units are in ascending (from the ground up) or descending order.
  Each rack is assigned a name and (optionally) a separate facility ID. This is helpful when leasing space in a data center your organization does not own: The facility will often assign a seemingly arbitrary ID to a rack (for example, "M204.313") whereas internally you refer to is simply as "R113." A unique serial number and asset tag may also be associated with each rack.
---

# netbox_rack (Resource)

From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rack/):

> The rack model represents a physical two- or four-post equipment rack in which devices can be installed. Each rack must be assigned to a site, and may optionally be assigned to a location within that site. Racks can also be organized by user-defined functional roles. The name and facility ID of each rack within a location must be unique.

Rack height is measured in rack units (U); racks are commonly between 42U and 48U tall, but NetBox allows you to define racks of arbitrary height. A toggle is provided to indicate whether rack units are in ascending (from the ground up) or descending order.

Each rack is assigned a name and (optionally) a separate facility ID. This is helpful when leasing space in a data center your organization does not own: The facility will often assign a seemingly arbitrary ID to a rack (for example, "M204.313") whereas internally you refer to is simply as "R113." A unique serial number and asset tag may also be associated with each rack.

## Example Usage

```terraform
resource "netbox_site" "test" {
  name   = "test"
  status = "active"
}

resource "netbox_rack" "test" {
  name     = "test"
  site_id  = netbox_site.test.id
  status   = "reserved"
  width    = 19
  u_height = 48
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `site_id` (Number)
- `status` (String) Valid values are `reserved`, `available`, `planned`, `active` and `deprecated`.

### Optional

- `asset_tag` (String)
- `comments` (String)
- `custom_fields` (Map of String)
- `desc_units` (Boolean) If rack units are descending. Defaults to `false`.
- `description` (String)
- `facility_id` (String)
- `form_factor` (String) Valid values are `2-post-frame`, `4-post-frame`, `4-post-cabinet`, `wall-frame`, `wall-frame-vertical`, `wall-cabinet` and `wall-cabinet-vertical`.
- `location_id` (Number)
- `max_weight` (Number)
- `mounting_depth` (Number)
- `outer_depth` (Number)
- `outer_unit` (String) Valid values are `mm` and `in`. Required when `outer_width` and `outer_depth` is set.
- `outer_width` (Number)
- `role_id` (Number)
- `serial` (String)
- `tags` (Set of String)
- `tenant_id` (Number)
- `u_height` (Number)
- `weight` (Number)
- `weight_unit` (String) Valid values are `kg`, `g`, `lb` and `oz`. Required when `weight` and `max_weight` is set.
- `width` (Number) Valid values are `10`, `19`, `21` and `23`.

### Read-Only

- `id` (String) The ID of this resource.
- `tags_all` (Set of String)


