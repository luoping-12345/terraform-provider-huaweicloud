// Generated by PMS #332
package as

import (
	"context"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/httphelper"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/schemas"
)

func DataSourceAsGroupQuotas() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAsGroupQuotasRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the AS group ID.`,
			},
			"quotas": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Specifies quota details.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resources": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Specifies resources.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `Specifies the quota upper limit.`,
									},
									"min": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `Specifies the quota lower limit.`,
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Specifies the quota type.`,
									},
									"used": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `Specifies the used quota.`,
									},
									"quota": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `Specifies the total quota.`,
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

type GroupQuotasDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newGroupQuotasDSWrapper(d *schema.ResourceData, meta interface{}) *GroupQuotasDSWrapper {
	return &GroupQuotasDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceAsGroupQuotasRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newGroupQuotasDSWrapper(d, meta)
	shoPolAndInsQuoRst, err := wrapper.ShowPolicyAndInstanceQuota()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.showPolicyAndInstanceQuotaToSchema(shoPolAndInsQuoRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API AS GET /autoscaling-api/v1/{project_id}/quotas/{scaling_group_id}
func (w *GroupQuotasDSWrapper) ShowPolicyAndInstanceQuota() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "autoscaling")
	if err != nil {
		return nil, err
	}

	uri := "/autoscaling-api/v1/{project_id}/quotas/{scaling_group_id}"
	uri = strings.ReplaceAll(uri, "{scaling_group_id}", w.Get("scaling_group_id").(string))
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Request().
		Result()
}

func (w *GroupQuotasDSWrapper) showPolicyAndInstanceQuotaToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("quotas", schemas.ObjectToList(body.Get("quotas"),
			func(quotas gjson.Result) any {
				return map[string]any{
					"resources": schemas.SliceToList(quotas.Get("resources"),
						func(resources gjson.Result) any {
							return map[string]any{
								"max":   resources.Get("max").Value(),
								"min":   resources.Get("min").Value(),
								"type":  resources.Get("type").Value(),
								"used":  resources.Get("used").Value(),
								"quota": resources.Get("quota").Value(),
							}
						},
					),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}