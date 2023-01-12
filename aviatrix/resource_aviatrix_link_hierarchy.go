package aviatrix

import (
	"context"

	"github.com/AviatrixSystems/terraform-provider-aviatrix/v3/goaviatrix"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAviatrixLinkHierarchy() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceAviatrixLinkHierarchyCreate,
		ReadWithoutTimeout:   resourceAviatrixLinkHierarchyRead,
		UpdateWithoutTimeout: resourceAviatrixLinkHierarchyUpdate,
		DeleteWithoutTimeout: resourceAviatrixLinkHierarchyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"links": {
				Type:             schema.TypeList,
				Required:         true,
				Description:      "",
				DiffSuppressFunc: goaviatrix.DiffSuppressFuncLinkHierarchy,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"wan_link": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"wan_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
		},
	}
}

func marshalLinkHierarchyInput(d *schema.ResourceData) map[string]interface{} {
	var links []map[string]interface{}
	linksInput := d.Get("links").([]interface{})

	for _, v0 := range linksInput {
		link := make(map[string]interface{})
		for k1, v1 := range v0.(map[string]interface{}) {
			if k1 == "name" {
				link[k1] = v1
			} else {
				var wanLinkList []map[string]interface{}
				wanlink := make(map[string]interface{})
				for _, v2 := range v1.([]interface{}) {
					for k3, v3 := range v2.(map[string]interface{}) {
						wanlink[k3] = v3
					}
				}
				wanLinkList = append(wanLinkList, wanlink)
				link["wan_link"] = wanLinkList
			}
		}
		links = append(links, link)
	}

	linkHierarchy := make(map[string]interface{})
	linkHierarchy["name"] = d.Get("name").(string)
	linkHierarchy["links"] = links

	return linkHierarchy
}

func resourceAviatrixLinkHierarchyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*goaviatrix.Client)

	linkHierarchy := marshalLinkHierarchyInput(d)

	flag := false
	defer resourceAviatrixLinkHierarchyReadIfRequired(ctx, d, meta, &flag)

	uuid, err := client.CreateLinkHierarchy(ctx, linkHierarchy)
	if err != nil {
		return diag.Errorf("failed to create link hierarchy: %s", err)
	}
	d.SetId(uuid)
	return resourceAviatrixLinkHierarchyReadIfRequired(ctx, d, meta, &flag)
}

func resourceAviatrixLinkHierarchyReadIfRequired(ctx context.Context, d *schema.ResourceData, meta interface{}, flag *bool) diag.Diagnostics {
	if !(*flag) {
		*flag = true
		return resourceAviatrixLinkHierarchyRead(ctx, d, meta)
	}
	return nil
}

func resourceAviatrixLinkHierarchyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*goaviatrix.Client)

	uuid := d.Id()
	d.Set("uuid", uuid)

	linkHierarchy, err := client.GetLinkHierarchy(ctx, uuid)
	if err != nil {
		if err == goaviatrix.ErrNotFound {
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to read link hierarchy: %s", err)
	}

	d.Set("name", linkHierarchy.Name)

	var links []interface{}

	for _, l := range linkHierarchy.Links {
		link := make(map[string]interface{})
		var wanLinkList []map[string]interface{}

		for _, w := range l.WanLinkList {
			wanLink := make(map[string]interface{})
			wanLink["wan_tag"] = w.WanTag
			wanLinkList = append(wanLinkList, wanLink)
		}

		link["name"] = l.Name
		link["wan_link"] = wanLinkList
		links = append(links, link)
	}

	if err := d.Set("links", links); err != nil {
		return diag.Errorf("failed to set links: %s", err)
	}

	return nil
}

func resourceAviatrixLinkHierarchyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*goaviatrix.Client)

	uuid := d.Id()
	d.Partial(true)
	if d.HasChanges("name", "links") {
		linkHierarchy := marshalLinkHierarchyInput(d)

		err := client.UpdateLinkHierarchy(ctx, linkHierarchy, uuid)
		if err != nil {
			return diag.Errorf("failed to update link hierarchy: %s", err)
		}
	}

	d.Partial(false)
	return resourceAviatrixLinkHierarchyRead(ctx, d, meta)
}

func resourceAviatrixLinkHierarchyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*goaviatrix.Client)

	uuid := d.Id()
	err := client.DeleteLinkHierarchy(ctx, uuid)
	if err != nil {
		return diag.Errorf("failed to delete link hierarchy: %v", err)
	}

	return nil
}
