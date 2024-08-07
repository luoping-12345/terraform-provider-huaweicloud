// Generated by PMS #271
package aom

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/httphelper"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/schemas"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func DataSourceAomAlarmActionRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAomAlarmActionRulesRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"action_rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Indicates the alarm action rules list.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the rule name.`,
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the action type.`,
						},
						"notification_template": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the message template.`,
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the rule description.`,
						},
						"created_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the user who created the rule.`,
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the create time.`,
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the update time.`,
						},
						"time_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the time zone.`,
						},
						"smn_topics": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Indicates the SMN topics.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topic_urn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Indicates the unique resource identifier of the topic.`,
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Indicates the name of the topic.`,
									},
									"display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Indicates the topic display name.`,
									},
									"push_policy": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `Indicates the SMN message push policy.`,
									},
									"status": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `Indicates the status of the topic subscriber.`,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

type AlarmActionRulesDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newAlarmActionRulesDSWrapper(d *schema.ResourceData, meta interface{}) *AlarmActionRulesDSWrapper {
	return &AlarmActionRulesDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceAomAlarmActionRulesRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newAlarmActionRulesDSWrapper(d, meta)
	listActionRuleRst, err := wrapper.ListActionRule()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listActionRuleToSchema(listActionRuleRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API AOM GET /v2/{project_id}/alert/action-rules
func (w *AlarmActionRulesDSWrapper) ListActionRule() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "aom")
	if err != nil {
		return nil, err
	}

	uri := "/v2/{project_id}/alert/action-rules"
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		OkCode(200).
		Request().
		Result()
}

func (w *AlarmActionRulesDSWrapper) listActionRuleToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("action_rules", schemas.SliceToList(body.Get("action_rules"),
			func(actionRules gjson.Result) any {
				return map[string]any{
					"name":                  actionRules.Get("rule_name").Value(),
					"type":                  actionRules.Get("type").Value(),
					"notification_template": actionRules.Get("notification_template").Value(),
					"description":           actionRules.Get("desc").Value(),
					"created_by":            actionRules.Get("user_name").Value(),
					"created_at":            w.setActRulCreTime(actionRules),
					"updated_at":            w.setActRulUpdTime(actionRules),
					"time_zone":             actionRules.Get("time_zone").Value(),
					"smn_topics": schemas.SliceToList(actionRules.Get("smn_topics"),
						func(smnTopics gjson.Result) any {
							return map[string]any{
								"topic_urn":    smnTopics.Get("topic_urn").Value(),
								"name":         smnTopics.Get("name").Value(),
								"display_name": smnTopics.Get("display_name").Value(),
								"push_policy":  smnTopics.Get("push_policy").Value(),
								"status":       smnTopics.Get("status").Value(),
							}
						},
					),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}

func (*AlarmActionRulesDSWrapper) setActRulCreTime(data gjson.Result) string {
	return utils.FormatTimeStampRFC3339((data.Get("create_time").Int())/1000, true)
}

func (*AlarmActionRulesDSWrapper) setActRulUpdTime(data gjson.Result) string {
	return utils.FormatTimeStampRFC3339((data.Get("update_time").Int())/1000, true)
}
