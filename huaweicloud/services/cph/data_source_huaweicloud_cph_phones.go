// Generated by PMS #381
package cph

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

func DataSourceCphPhones() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCphPhonesRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"phone_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the cloud phone name and support fuzzy query.`,
			},
			"server_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the cloud phone server ID.`,
			},
			"status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: `Specifies the cloud phone status.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the cloud phone type.`,
			},
			"phones": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The cloud phone list.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The cloud phone ID.`,
						},
						"phone_model_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The cloud phone flavor name.`,
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The cloud phone status.`,
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The create time.`,
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The update time.`,
						},
						"phone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The cloud phone name.`,
						},
						"server_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The cloud phone server ID.`,
						},
						"image_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The cloud phone image ID.`,
						},
						"vnc_enable": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Whether to enable the VNC service on the cloud phone.`,
						},
						"type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The cloud phone type.`,
						},
						"metadata": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `The order and product related information.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"order_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The order ID.`,
									},
									"product_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The product ID.`,
									},
								},
							},
						},
						"volume_mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Whether the physical disk of the mobile phone is independent.`,
						},
						"availability_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The availability zone where the cloud mobile server is located.`,
						},
						"image_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The image version.`,
						},
						"imei": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The IMEI of the phone.`,
						},
						"traffic_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The phone routing type.`,
						},
					},
				},
			},
		},
	}
}

type PhonesDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newPhonesDSWrapper(d *schema.ResourceData, meta interface{}) *PhonesDSWrapper {
	return &PhonesDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceCphPhonesRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newPhonesDSWrapper(d, meta)
	listCloudPhonesRst, err := wrapper.ListCloudPhones()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listCloudPhonesToSchema(listCloudPhonesRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API CPH GET /v1/{project_id}/cloud-phone/phones
func (w *PhonesDSWrapper) ListCloudPhones() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "cph")
	if err != nil {
		return nil, err
	}

	uri := "/v1/{project_id}/cloud-phone/phones"
	params := map[string]any{
		"phone_name": w.Get("phone_name"),
		"server_id":  w.Get("server_id"),
		"status":     w.Get("status"),
		"type":       w.Get("type"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		OffsetPager("phones", "offset", "limit", 200).
		Request().
		Result()
}

func (w *PhonesDSWrapper) listCloudPhonesToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("phones", schemas.SliceToList(body.Get("phones"),
			func(phones gjson.Result) any {
				return map[string]any{
					"phone_id":         phones.Get("phone_id").Value(),
					"phone_model_name": phones.Get("phone_model_name").Value(),
					"status":           phones.Get("status").Value(),
					"create_time":      phones.Get("create_time").Value(),
					"update_time":      phones.Get("update_time").Value(),
					"phone_name":       phones.Get("phone_name").Value(),
					"server_id":        phones.Get("server_id").Value(),
					"image_id":         phones.Get("image_id").Value(),
					"vnc_enable":       phones.Get("vnc_enable").Value(),
					"type":             phones.Get("type").Value(),
					"metadata": schemas.SliceToList(phones.Get("metadata"),
						func(metadata gjson.Result) any {
							return map[string]any{
								"order_id":   metadata.Get("order_id").Value(),
								"product_id": metadata.Get("product_id").Value(),
							}
						},
					),
					"volume_mode":       phones.Get("volume_mode").Value(),
					"availability_zone": phones.Get("availability_zone").Value(),
					"image_version":     phones.Get("image_version").Value(),
					"imei":              phones.Get("imei").Value(),
					"traffic_type":      phones.Get("traffic_type").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}