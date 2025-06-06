// Generated by PMS #521
package cfw

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

func DataSourceCfwIpsRuleDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCfwIpsRuleDetailsRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"fw_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the firewall ID.`,
			},
			"data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The IPS information.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ips_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The IPS type.`,
						},
						"ips_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The IPS version.`,
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The update time.`,
						},
					},
				},
			},
		},
	}
}

type IpsRuleDetailsDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newIpsRuleDetailsDSWrapper(d *schema.ResourceData, meta interface{}) *IpsRuleDetailsDSWrapper {
	return &IpsRuleDetailsDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceCfwIpsRuleDetailsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newIpsRuleDetailsDSWrapper(d, meta)
	showIpsUpdateTimeRst, err := wrapper.ShowIpsUpdateTime()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.showIpsUpdateTimeToSchema(showIpsUpdateTimeRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API CFW GET /v1/{project_id}/ips-rule/detail
func (w *IpsRuleDetailsDSWrapper) ShowIpsUpdateTime() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "cfw")
	if err != nil {
		return nil, err
	}

	uri := "/v1/{project_id}/ips-rule/detail"
	params := map[string]any{
		"fw_instance_id": w.Get("fw_instance_id"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		Request().
		Result()
}

func (w *IpsRuleDetailsDSWrapper) showIpsUpdateTimeToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("data", schemas.SliceToList(body.Get("data"),
			func(data gjson.Result) any {
				return map[string]any{
					"ips_type":    data.Get("ips_type").Value(),
					"ips_version": data.Get("ips_version").Value(),
					"update_time": w.setDataUpdateTime(data),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}

func (*IpsRuleDetailsDSWrapper) setDataUpdateTime(data gjson.Result) string {
	return utils.FormatTimeStampRFC3339(data.Get("update_time").Int()/1000, true)
}
