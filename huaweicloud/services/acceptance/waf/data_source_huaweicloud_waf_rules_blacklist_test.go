package waf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceRulesBlacklist_basic(t *testing.T) {
	var (
		dataSourceName = "data.huaweicloud_waf_rules_blacklist.test"
		rName          = acceptance.RandomAccResourceName()
		dc             = acceptance.InitDataSourceCheck(dataSourceName)

		byRuleId   = "data.huaweicloud_waf_rules_blacklist.filter_by_rule_id"
		dcByRuleId = acceptance.InitDataSourceCheck(byRuleId)

		byName   = "data.huaweicloud_waf_rules_blacklist.filter_by_rule_id"
		dcByName = acceptance.InitDataSourceCheck(byName)

		byNameNotFound   = "data.huaweicloud_waf_rules_blacklist.not_found"
		dcByNameNotFound = acceptance.InitDataSourceCheck(byNameNotFound)

		byStatus   = "data.huaweicloud_waf_rules_blacklist.filter_by_status"
		dcByStatus = acceptance.InitDataSourceCheck(byStatus)

		byAction   = "data.huaweicloud_waf_rules_blacklist.filter_by_action"
		dcByAction = acceptance.InitDataSourceCheck(byAction)
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPrecheckWafInstance(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceRulesBlacklist_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSourceName, "rules.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "rules.0.id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "rules.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "rules.0.status"),
					resource.TestCheckResourceAttrSet(dataSourceName, "rules.0.description"),
					resource.TestCheckResourceAttrSet(dataSourceName, "rules.0.action"),
					resource.TestCheckResourceAttrSet(dataSourceName, "rules.0.created_at"),

					dcByRuleId.CheckResourceExists(),
					resource.TestCheckOutput("rule_id_filter_is_useful", "true"),

					dcByName.CheckResourceExists(),
					resource.TestCheckOutput("name_filter_is_useful", "true"),

					dcByNameNotFound.CheckResourceExists(),
					resource.TestCheckOutput("is_not_found", "true"),

					dcByStatus.CheckResourceExists(),
					resource.TestCheckOutput("status_filter_is_useful", "true"),

					dcByAction.CheckResourceExists(),
					resource.TestCheckOutput("action_filter_is_useful", "true"),
				),
			},
		},
	})
}
func testDataSourceRulesBlacklist_base(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_waf_rule_blacklist" "test" {
  policy_id   = huaweicloud_waf_policy.policy_1.id
  ip_address  = "192.160.0.0/24"
  name        = "%[2]s"
  description = "test description"
}
`, testAccWafPolicyV1_basic(name), name)
}

func testDataSourceRulesBlacklist_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_waf_rules_blacklist" "test" {
  depends_on = [
    huaweicloud_waf_rule_blacklist.test
  ]

  policy_id = huaweicloud_waf_policy.policy_1.id
}

locals {
  rule_id = data.huaweicloud_waf_rules_blacklist.test.rules[0].id
}

data "huaweicloud_waf_rules_blacklist" "filter_by_rule_id" {
  policy_id = huaweicloud_waf_policy.policy_1.id
  rule_id   = local.rule_id
}

locals {
  rule_id_filter_result = [
    for v in data.huaweicloud_waf_rules_blacklist.filter_by_rule_id.rules[*].id : v == local.rule_id
  ]
}

output "rule_id_filter_is_useful" {
  value = alltrue(local.rule_id_filter_result) && length(local.rule_id_filter_result) > 0
}

locals {
  name = data.huaweicloud_waf_rules_blacklist.test.rules[0].name
}

data "huaweicloud_waf_rules_blacklist" "filter_by_name" {
  policy_id = huaweicloud_waf_policy.policy_1.id
  name      = local.name
}

locals {
  name_filter_result = [
    for v in data.huaweicloud_waf_rules_blacklist.filter_by_name.rules[*].name : v == local.name
  ]
}

output "name_filter_is_useful" {
  value = alltrue(local.name_filter_result) && length(local.name_filter_result) > 0
}

data "huaweicloud_waf_rules_blacklist" "not_found" {
  policy_id = huaweicloud_waf_policy.policy_1.id
  name      = "not_found"
}
  
output "is_not_found" {
  value = length(data.huaweicloud_waf_rules_blacklist.not_found.rules) == 0
}

locals {
  status = data.huaweicloud_waf_rules_blacklist.test.rules[0].status
}

data "huaweicloud_waf_rules_blacklist" "filter_by_status" {
  policy_id = huaweicloud_waf_policy.policy_1.id
  status    = local.status
}

locals {
  status_filter_result = [ 
    for v in data.huaweicloud_waf_rules_blacklist.filter_by_status.rules[*].status : v == local.status
  ]
}

output "status_filter_is_useful" {
  value = alltrue(local.status_filter_result) && length(local.status_filter_result) > 0
}

locals {
  action = data.huaweicloud_waf_rules_blacklist.test.rules[0].action
}

data "huaweicloud_waf_rules_blacklist" "filter_by_action" {
  policy_id = huaweicloud_waf_policy.policy_1.id
  action    = local.action
}

locals {
  action_filter_result = [ 
    for v in data.huaweicloud_waf_rules_blacklist.filter_by_action.rules[*].action : v == local.action
  ]
}

output "action_filter_is_useful" {
  value = alltrue(local.action_filter_result) && length(local.action_filter_result) > 0
}
`, testDataSourceRulesBlacklist_base(name))
}
