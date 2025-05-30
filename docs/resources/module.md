---
# generated by https://github.com/fbreckle/terraform-plugin-docs
page_title: "netbox_module Resource - terraform-provider-netbox"
subcategory: "Data Center Inventory Management (DCIM)"
description: |-
  From the official documentation https://docs.netbox.dev/en/stable/models/dcim/module/:
  A module is a field-replaceable hardware component installed within a device which houses its own child components. The most common example is a chassis-based router or switch.
  Similar to devices, modules are instantiated from module types, and any components associated with the module type are automatically instantiated on the new model. Each module must be installed within a module bay on a device, and each module bay may have only one module installed in it.
---

# netbox_module (Resource)

From the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/module/):

> A module is a field-replaceable hardware component installed within a device which houses its own child components. The most common example is a chassis-based router or switch.

Similar to devices, modules are instantiated from module types, and any components associated with the module type are automatically instantiated on the new model. Each module must be installed within a module bay on a device, and each module bay may have only one module installed in it.

## Example Usage

```terraform
# Note that some terraform code is not included in the example for brevity

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name      = "SFP"
}

resource "netbox_manufacturer" "test" {
  name = "Dell"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "Networking"
}

resource "netbox_module" "test" {
  device_id      = netbox_device.test.id
  module_bay_id  = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status         = "active"

  description = "SFP card"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (Number)
- `module_bay_id` (Number)
- `module_type_id` (Number)
- `status` (String) One of [offline, active, planned, staged, failed, decommissioning].

### Optional

- `asset_tag` (String)
- `comments` (String)
- `custom_fields` (Map of String)
- `description` (String)
- `serial` (String)
- `tags` (Set of String)

### Read-Only

- `id` (String) The ID of this resource.
- `tags_all` (Set of String)


