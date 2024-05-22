// Generated by PMS #165
package cfw

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/filters"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/httphelper"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/schemas"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func DataSourceCfwAddressGroupMembers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCfwAddressGroupMembersRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the ID of the IP address group.`,
			},
			"key_word": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the keyword.`,
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the IP address`,
			},
			"item_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the address group member ID.`,
			},
			"fw_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the firewall instance ID.`,
			},
			"query_address_set_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the query address group type.`,
			},
			"records": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The IP address group member list.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"item_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The ID of an address group member.`,
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The address group member description.`,
						},
						"address_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The address type.`,
						},
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The IP address.`,
						},
					},
				},
			},
		},
	}
}

type AddressGroupMembersDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newAddressGroupMembersDSWrapper(d *schema.ResourceData, meta interface{}) *AddressGroupMembersDSWrapper {
	return &AddressGroupMembersDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceCfwAddressGroupMembersRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newAddressGroupMembersDSWrapper(d, meta)
	lisAddIteRst, err := wrapper.ListAddressItems()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listAddressItemsToSchema(lisAddIteRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API CFW GET /v1/{project_id}/address-items
func (w *AddressGroupMembersDSWrapper) ListAddressItems() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "cfw")
	if err != nil {
		return nil, err
	}

	uri := "/v1/{project_id}/address-items"
	params := map[string]any{
		"set_id":                 w.Get("group_id"),
		"key_word":               w.Get("key_word"),
		"address":                w.Get("address"),
		"fw_instance_id":         w.Get("fw_instance_id"),
		"query_address_set_type": w.Get("query_address_set_type"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		OffsetPager("data.records", "offset", "limit", 1024).
		Filter(
			filters.New().From("data.records").
				Where("item_id", "=", w.Get("item_id")),
		).
		OkCode(200).
		Request().
		Result()
}

func (w *AddressGroupMembersDSWrapper) listAddressItemsToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("records", schemas.SliceToList(body.Get("data.records"),
			func(record gjson.Result) any {
				return map[string]any{
					"item_id":      record.Get("item_id").Value(),
					"description":  record.Get("description").Value(),
					"address_type": record.Get("address_type").Value(),
					"address":      record.Get("address").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}