// Generated by PMS #514
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
)

func DataSourceCfwTags() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCfwTagsRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The tag list.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The tag key.`,
						},
						"values": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: `The tag values.`,
						},
					},
				},
			},
		},
	}
}

type TagsDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newTagsDSWrapper(d *schema.ResourceData, meta interface{}) *TagsDSWrapper {
	return &TagsDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceCfwTagsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newTagsDSWrapper(d, meta)
	listProjectTagsRst, err := wrapper.ListProjectTags()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listProjectTagsToSchema(listProjectTagsRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API CFW GET /v2/{project_id}/cfw-cfw/tags
func (w *TagsDSWrapper) ListProjectTags() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "cfw")
	if err != nil {
		return nil, err
	}

	uri := "/v2/{project_id}/cfw-cfw/tags"
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		OffsetPager("tags", "offset", "limit", 1000).
		Request().
		Result()
}

func (w *TagsDSWrapper) listProjectTagsToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("tags", schemas.SliceToList(body.Get("tags"),
			func(tags gjson.Result) any {
				return map[string]any{
					"key":    tags.Get("key").Value(),
					"values": schemas.SliceToStrList(tags.Get("values")),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}