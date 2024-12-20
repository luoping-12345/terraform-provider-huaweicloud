package identitycenter

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceIdentitycenterProvisionedPermissionSets_basic(t *testing.T) {
	dataSource := "data.huaweicloud_identitycenter_provisioned_permission_sets.test"
	rName := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckMultiAccount(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceDataSourceIdentitycenterProvisionedPermissionSets_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(rName, "permission_set_provisioning_status"),
				),
			},
		},
	})
}

func testDataSourceDataSourceIdentitycenterProvisionedPermissionSets_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_identitycenter_provisioned_permission_sets" "test" {
  instance_id       = data.huaweicloud_identitycenter_instance.system.id
  permission_set_id = huaweicloud_identitycenter_permission_set.test.id
  target_type       = "ACCOUNT"
}
`, testPermissionSet_basic(name))
}
