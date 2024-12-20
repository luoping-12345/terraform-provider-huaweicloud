---
subcategory: "IAM Identity Center"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_identitycenter_provisioned_permission_sets"
description: |-
  Use this data source to get the Identity Center provisioned permission sets.
---

# huaweicloud_identitycenter_provisioned_permission_sets

Use this data source to get the Identity Center provisioned permission sets.

## Example Usage

```hcl
variable "instance_id" {}
variable "permission_set_id" {}
variable "target_type" {}

data "huaweicloud_identitycenter_provisioned_permission_sets" "test" {
  instance_id       = var.instance_id
  permission_set_id = var.permission_set_id
  target_type       = var.target_type
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `instance_id` - (Required, String) Specifies the ID of an IAM Identity Center instance.

* `permission_set_id` - (Required, String) Specifies the ID of a permission set.

* `target_id` - (Optional, String) Specifies the account ID.

* `target_type` - (Required, String) Specifies the type of the principal to be attached.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `permission_set_provisioning_status` - The authorization details of a permission set.

  The [permission_set_provisioning_status](#permission_set_provisioning_status_struct) structure is documented below.

<a name="permission_set_provisioning_status_struct"></a>
The `permission_set_provisioning_status` block supports:

* `status` - The authorization status of a permission set.

* `account_id` - The ID of a specified account.

* `created_at` - The time when a permission set was created.

* `failure_reason` - The failure reason.

* `permission_set_id` - The ID of a permission set.
