---
subcategory: "Cloud Certificate Manager (CCM)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_ccm_private_certificate_revoke"
description: |
  Manages CCM private certificate revoke resource within HuaweiCloud.
---

# huaweicloud_ccm_private_certificate_revoke

Manages CCM private certificate revoke resource within HuaweiCloud.

-> The current resource is a one-time resource, and destroying this resource will not change the current status.

## Example Usage

```hcl
variable "certificate_id" {}

resource "huaweicloud_ccm_private_certificate_revoke" "test" {
  certificate_id = var.certificate_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region to which the CCM private certificate belongs.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `certificate_id` - (Required, String, ForceNew) Specifies the ID of the private certificate to be revoked.
  Changing this parameter will create a new resource.

* `reason` - (Optional, String, ForceNew) Specifies the reason for revoking the private certificate.
  Changing this parameter will create a new resource.
  The valid values are as follows:
  + **UNSPECIFIED**: The default value. No reason is specified for revocation.
  + **KEY_COMPROMISE**: The certificate key material has been leaked.
  + **CERTIFICATE_AUTHORITY_COMPROMISE**: The CA key meterial may be leaked in the certificate chain.
  + **AFFILIATION_CHANGED**: The subject or other information in the certificate has been changed.
  + **SUPERSEDED**: The certificate has been replaced.
  + **CESSATION_OF_OPERATION**: The entity in the certificate or  certificate chain has ceased operation.
  + **CERTIFICATE_HOLD**: The certificate should not be considered valid currently and may take effect in the future.
  + **PRIVILEGE_WITHDRAWN**: The certificate no longer has permissions on the properties it claims.
  + **ATTRIBUTE_AUTHORITY_COMPROMISE**: The authority that guarantee attributes of the certificate may have been
    compromised.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.