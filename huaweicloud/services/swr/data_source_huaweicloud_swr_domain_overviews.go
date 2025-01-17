// Generated by PMS #532
package swr

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

func DataSourceSwrDomainOverviews() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSwrDomainOverviewsRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"namspace_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of namespaces of the tenant.`,
			},
			"repo_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of repositories of the tenant.`,
			},
			"image_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of images of the tenant.`,
			},
			"store_space": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: `The storage size of the tenant.`,
			},
			"downflow_size": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: `The download traffic of the tenant.`,
			},
			"domain_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The domain ID.`,
			},
			"domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The domain name.`,
			},
		},
	}
}

type DomainOverviewsDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newDomainOverviewsDSWrapper(d *schema.ResourceData, meta interface{}) *DomainOverviewsDSWrapper {
	return &DomainOverviewsDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceSwrDomainOverviewsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newDomainOverviewsDSWrapper(d, meta)
	shoDomOveRst, err := wrapper.ShowDomainOverview()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.showDomainOverviewToSchema(shoDomOveRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API SWR GET /v2/manage/overview
func (w *DomainOverviewsDSWrapper) ShowDomainOverview() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "swr")
	if err != nil {
		return nil, err
	}

	uri := "/v2/manage/overview"
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Request().
		Result()
}

func (w *DomainOverviewsDSWrapper) showDomainOverviewToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("namspace_num", body.Get("namspace_num").Value()),
		d.Set("repo_num", body.Get("repo_num").Value()),
		d.Set("image_num", body.Get("image_num").Value()),
		d.Set("store_space", body.Get("store_space").Value()),
		d.Set("downflow_size", body.Get("downflow_size").Value()),
		d.Set("domain_id", body.Get("domain_id").Value()),
		d.Set("domain_name", body.Get("domain_name").Value()),
	)
	return mErr.ErrorOrNil()
}