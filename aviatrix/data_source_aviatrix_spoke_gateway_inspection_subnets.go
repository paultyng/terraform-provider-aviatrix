package aviatrix

import (
	"fmt"

	"github.com/AviatrixSystems/terraform-provider-aviatrix/v2/goaviatrix"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAviatrixSpokeGatewayInspectionSubnets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAviatrixSpokeGatewayInspectionSubnetsRead,

		Schema: map[string]*schema.Schema{
			"gw_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Spoke gateway name.",
			},
			"subnets_for_inspection": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of all subnets available for the subnet inspection feature.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceAviatrixSpokeGatewayInspectionSubnetsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)

	subnetsForInspection, err := client.GetSubnetsForInspection(d.Get("gw_name").(string))
	if err != nil {
		return fmt.Errorf("couldn't get subnets for inspection: %s", err)
	}
	d.Set("subnets_for_inspection", subnetsForInspection)

	d.SetId(d.Get("gw_name").(string) + "~inspection_subnets")
	return nil
}
