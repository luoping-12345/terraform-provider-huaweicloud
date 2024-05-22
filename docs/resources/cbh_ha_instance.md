---
subcategory: "Cloud Bastion Host (CBH)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cbh_ha_instance"
description: |-
  Manages a CBH HA instance resource within HuaweiCloud.
---

# huaweicloud_cbh_ha_instance

Manages a CBH HA instance resource within HuaweiCloud.

## Example Usage

```hcl
variable "name" {}
variable "flavor_id" {}
variable "vpc_id" {}
variable "subnet_id" {}
variable "security_group_id" {}
variable "master_availability_zone" {}
variable "slave_availability_zone" {}
variable "password" {}

resource "huaweicloud_cbh_ha_instance" "test" {
  name                     = var.name
  flavor_id                = var.flavor_id
  vpc_id                   = var.vpc_id
  subnet_id                = var.subnet_id
  security_group_id        = var.security_group_id
  master_availability_zone = var.master_availability_zone
  slave_availability_zone  = var.slave_availability_zone
  password                 = var.password
  charging_mode            = "prePaid"
  period_unit              = "month"
  period                   = 1
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the CBH HA instance.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `name` - (Required, String, ForceNew) Specifies the name of the CBH HA instance. The field can contain `1` to `64`
  characters. Only letters, digits, underscores (_), and hyphens (-) are allowed.

  Changing this parameter will create a new resource.

* `flavor_id` - (Required, String) Specifies the product ID of the CBH server. When updating the flavor, it can only be
  changed to a higher flavor.

  -> 1. The flavor change is a high-risk operation, with a certain risk of failure.
  <br/>2. Flavor change failing may impact the usability of the instance. Please be sure to back up your data.

* `vpc_id` - (Required, String, ForceNew) Specifies the ID of a VPC.

  Changing this parameter will create a new resource.

* `subnet_id` - (Required, String, ForceNew) Specifies the ID of a subnet.

  Changing this parameter will create a new resource.

* `security_group_id` - (Required, String) Specifies the IDs of the security group. Multiple security group IDs are
  separated by commas (,) without spaces.

* `master_availability_zone` - (Required, String, ForceNew) Specifies the availability zone name of the master instance.

  Changing this parameter will create a new resource.

* `slave_availability_zone` - (Required, String, ForceNew) Specifies the availability zone name of the slave instance.

  Changing this parameter will create a new resource.

* `password` - (Required, String) Specifies the password for logging in to the management console. The value of the
  field has the following restrictions:
  + The value of the field must contain `8` to `32` characters.
  + The value of the field must contain at least three of the following: letters, digits, and special characters
    (!@$%^-_=+[{}]:,./?~#*).
  + The value of the field cannot contain the username or the username spelled backwards.

* `charging_mode` - (Required, String, ForceNew) Specifies the charging mode of the CBH HA instance.
  The options are as follows:
  + **prePaid**: the yearly/monthly billing mode.

  Changing this parameter will create a new resource.

* `period_unit` - (Required, String, ForceNew) Specifies the charging period unit of the CBH HA instance.
  Valid values are *month* and *year*.

  Changing this parameter will create a new resource.

* `period` - (Required, Int, ForceNew) Specifies the charging period of the CBH HA instance.
  If `period_unit` is set to **month**, the value ranges from 1 to 9.
  If `period_unit` is set to **year**, the value ranges from 1 to 3.

  Changing this parameter will create a new resource.

* `master_private_ip` - (Optional, String, ForceNew) Specifies the private IP address of the master instance.

  Changing this parameter will create a new resource.

* `slave_private_ip` - (Optional, String, ForceNew) Specifies the private IP address of the slave instance.

  Changing this parameter will create a new resource.

* `floating_ip` - (Optional, String, ForceNew) Specifies the floating IP address of the CBH HA instance.

  Changing this parameter will create a new resource.

-> For the parameters `master_private_ip`, `slave_private_ip`, and `floating_ip`, if none of them are specified,
  a new IP address will be assigned to each. If one is specified, then the other two must also be specified.

* `auto_renew` - (Optional, String) Specifies whether auto-renew is enabled.
  Valid values are **true** and **false**. Defaults to **false**.

* `public_ip_id` - (Optional, String) Specifies the ID of the elastic IP.

* `ipv6_enable` - (Optional, Bool, ForceNew) Specifies whether the IPv6 network is enabled. Defaults to **false**.

  Changing this parameter will create a new resource.

* `attach_disk_size` - (Optional, Int) Specifies the size of the additional data disk for the CBH HA instance.
  The unit is TB. It refers to the additional disk size added on top of the existing disk. And the sum of the built-in
  disk of the instance flavor and the additional disk cannot exceed **300TB**.

  -> 1. Storage expansion is a high-risk operation, with a certain risk of failure.
  <br/>2. Expansion failure may affect the usability of the instance. Please ensure to back up your data.

* `tags` - (Optional, Map) Specifies the key/value pairs to associate with the CBH HA instance.

* `enterprise_project_id` - (Optional, String) Specifies the enterprise project ID to which the CBH HA instance belongs.
  For enterprise users, if omitted, default enterprise project will be used.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

* `public_ip` - The elastic IP address.

* `master_id` - The ID of the master instance.

* `slave_id` - The ID of the slave instance.

* `status` - The status of the CBH HA instance.

* `version` - The current version of the CBH HA instance image.

* `data_disk_size` - The data disk size of the CBH HA instance. The unit is TB. It represents the sum of the
  disks that come with the flavor and the disks that have already been expanded.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 60 minutes.
* `update` - Default is 60 minutes.
* `delete` - Default is 30 minutes.

## Import

The CBH HA instance can be imported using the master instance ID and the slave instance ID, separated by a slash, e.g.

```bash
$ terraform import huaweicloud_cbh_ha_instance.test <master_id>/<slave_id>
```

Note that the imported state may not be identical to your resource definition, due to some attributes missing from the
API response, security or some other reason. The missing attributes include: `charging_mode`, `period`, `period_unit`,
`auto_renew`, `password`, `ipv6_enable`, `attach_disk_size`.
It is generally recommended running `terraform plan` after importing an instance.
You can then decide if changes should be applied to the instance, or the resource definition should be updated
to align with the instance. Also, you can ignore changes as below.

```
resource "huaweicloud_cbh_ha_instance" "test" {
    ...

  lifecycle {
    ignore_changes = [
      charging_mode, period, period_unit, auto_renew, password, ipv6_enable, attach_disk_size,
    ]
  }
}
```