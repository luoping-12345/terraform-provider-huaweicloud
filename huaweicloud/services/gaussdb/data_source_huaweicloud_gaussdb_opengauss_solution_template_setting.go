// Generated by PMS #519
package gaussdb

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

func DataSourceGaussdbOpengaussSolutionTemplateSetting() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGaussdbOpengaussSolutionTemplateSettingRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"solution": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"solution", "instance_id"},
				Description:  `Specifies the solution template name.`,
			},
			"instance_id": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"solution", "instance_id"},
				Description:  `Specifies the GaussDB OpenGauss instance ID.`,
			},
			"shard_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Indicates the number of shards.`,
			},
			"replica_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Indicates the number of replicas.`,
			},
			"initial_node_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Indicates the number of initial nodes.`,
			},
		},
	}
}

type OpengaussSolutionTemplateSettingDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newOpengaussSolutionTemplateSettingDSWrapper(d *schema.ResourceData, meta interface{}) *OpengaussSolutionTemplateSettingDSWrapper {
	return &OpengaussSolutionTemplateSettingDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceGaussdbOpengaussSolutionTemplateSettingRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newOpengaussSolutionTemplateSettingDSWrapper(d, meta)
	shoDepForRst, err := wrapper.ShowDeploymentForm()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.showDeploymentFormToSchema(shoDepForRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API GaussDB GET /v3/{project_id}/deployment-form
func (w *OpengaussSolutionTemplateSettingDSWrapper) ShowDeploymentForm() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "opengauss")
	if err != nil {
		return nil, err
	}

	uri := "/v3/{project_id}/deployment-form"
	params := map[string]any{
		"solution":    w.Get("solution"),
		"instance_id": w.Get("instance_id"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		Request().
		Result()
}

func (w *OpengaussSolutionTemplateSettingDSWrapper) showDeploymentFormToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("shard_num", body.Get("shard_num").Value()),
		d.Set("replica_num", body.Get("replica_num").Value()),
		d.Set("initial_node_num", body.Get("initial_node_num").Value()),
	)
	return mErr.ErrorOrNil()
}