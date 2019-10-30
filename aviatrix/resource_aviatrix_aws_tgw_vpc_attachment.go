package aviatrix

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aviatrix/goaviatrix"
)

func resourceAviatrixAwsTgwVpcAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviatrixAwsTgwVpcAttachmentCreate,
		Read:   resourceAviatrixAwsTgwVpcAttachmentRead,
		Update: resourceAviatrixAwsTgwVpcAttachmentUpdate,
		Delete: resourceAviatrixAwsTgwVpcAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tgw_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the AWS TGW.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region of cloud provider.",
			},
			"security_domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the security domain.",
			},
			"vpc_account_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This parameter represents the name of a Cloud-Account in Aviatrix controller.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This parameter represents the ID of the VPC.",
			},
		},
	}
}

func resourceAviatrixAwsTgwVpcAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)

	awsTgwVpcAttachment := &goaviatrix.AwsTgwVpcAttachment{
		TgwName:            d.Get("tgw_name").(string),
		Region:             d.Get("region").(string),
		SecurityDomainName: d.Get("security_domain_name").(string),
		VpcAccountName:     d.Get("vpc_account_name").(string),
		VpcID:              d.Get("vpc_id").(string),
	}

	isFireNetVpc, err := client.IsFireNetVpc(awsTgwVpcAttachment.VpcID)
	if err != nil {
		if err == goaviatrix.ErrNotFound {
			return fmt.Errorf("could not find VPC with ID: " + awsTgwVpcAttachment.VpcID)
		}
		return fmt.Errorf(("could not find VPC due to: ") + err.Error())
	}

	isFirewallSecurityDomain, err := client.IsFirewallSecurityDomain(awsTgwVpcAttachment.TgwName, awsTgwVpcAttachment.SecurityDomainName)
	if err != nil {
		if err == goaviatrix.ErrNotFound {
			return fmt.Errorf("could not find Security Domain: " + awsTgwVpcAttachment.VpcID)
		}
		return fmt.Errorf(("could not find Security Domain due to: ") + err.Error())
	}

	if isFireNetVpc && !isFirewallSecurityDomain {
		return fmt.Errorf("could not attach a FireNet VPC to a Non Firewall Security Domain")
	} else if !isFireNetVpc && isFirewallSecurityDomain {
		return fmt.Errorf("could not attach an Non FireNet VPC to an Firewall Security Domain")
	}

	log.Printf("[INFO] Attaching vpc: %s to tgw %s", awsTgwVpcAttachment.VpcID, awsTgwVpcAttachment.TgwName)

	if isFireNetVpc && isFirewallSecurityDomain {
		err := client.CreateAwsTgwVpcAttachmentForFireNet(awsTgwVpcAttachment)
		if err != nil {
			return fmt.Errorf("failed to create Aviatrix Aws Tgw Vpc Attach for FireNet: %s", err)
		}
	} else if !isFireNetVpc && !isFirewallSecurityDomain {
		err := client.CreateAwsTgwVpcAttachment(awsTgwVpcAttachment)
		if err != nil {
			return fmt.Errorf("failed to create Aviatrix Aws Tgw Vpc Attach: %s", err)
		}
	}

	d.SetId(awsTgwVpcAttachment.TgwName + "~" + awsTgwVpcAttachment.SecurityDomainName + "~" + awsTgwVpcAttachment.VpcID)
	return resourceAviatrixAwsTgwVpcAttachmentRead(d, meta)
}

func resourceAviatrixAwsTgwVpcAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)

	tgwName := d.Get("tgw_name").(string)
	securityDomainName := d.Get("security_domain_name").(string)
	vpcID := d.Get("vpc_id").(string)

	if tgwName == "" || securityDomainName == "" || vpcID == "" {
		id := d.Id()
		log.Printf("[DEBUG] Looks like an import, no vpc names received. Import Id is %s", id)
		d.Set("tgw_name", strings.Split(id, "~")[0])
		d.Set("security_domain_name", strings.Split(id, "~")[1])
		d.Set("vpc_id", strings.Split(id, "~")[2])
	}
	awsTgwVpcAttachment := &goaviatrix.AwsTgwVpcAttachment{
		TgwName:            d.Get("tgw_name").(string),
		SecurityDomainName: d.Get("security_domain_name").(string),
		VpcID:              d.Get("vpc_id").(string),
	}

	aTVA, err := client.GetAwsTgwVpcAttachment(awsTgwVpcAttachment)
	if err != nil {
		if err == goaviatrix.ErrNotFound {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("failed to get Aviatrix Aws Tgw Vpc Attach: %s", err)
	}
	if aTVA != nil {
		d.Set("tgw_name", awsTgwVpcAttachment.TgwName)
		d.Set("region", awsTgwVpcAttachment.Region)
		d.Set("security_domain_name", awsTgwVpcAttachment.SecurityDomainName)
		d.Set("vpc_account_name", awsTgwVpcAttachment.VpcAccountName)
		d.Set("vpc_id", awsTgwVpcAttachment.VpcID)
		d.SetId(awsTgwVpcAttachment.TgwName + "~" + awsTgwVpcAttachment.SecurityDomainName + "~" + awsTgwVpcAttachment.VpcID)
		return nil
	}

	return fmt.Errorf("no Aviatrix Aws Tgw Vpc Attach found")
}

func resourceAviatrixAwsTgwVpcAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	if d.HasChange("tgw_name") {
		return fmt.Errorf("updating tgw_name is not allowed")
	}
	if d.HasChange("region") {
		return fmt.Errorf("updating region is not allowed")
	}
	if d.HasChange("vpc_account_name") {
		return fmt.Errorf("updating vpc_account_name is not allowed")
	}
	if d.HasChange("vpc_id") {
		return fmt.Errorf("updating vpc_id is not allowed")
	}

	d.Partial(false)
	return resourceAviatrixAwsTgwVpcAttachmentRead(d, meta)
}

func resourceAviatrixAwsTgwVpcAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)

	awsTgwVpcAttachment := &goaviatrix.AwsTgwVpcAttachment{
		TgwName:            d.Get("tgw_name").(string),
		Region:             d.Get("region").(string),
		SecurityDomainName: d.Get("security_domain_name").(string),
		VpcAccountName:     d.Get("vpc_account_name").(string),
		VpcID:              d.Get("vpc_id").(string),
	}

	isFireNetVpc, err := client.IsFireNetVpc(awsTgwVpcAttachment.VpcID)
	if err != nil {
		if err == goaviatrix.ErrNotFound {
			return fmt.Errorf("could not find VPC with ID: " + awsTgwVpcAttachment.VpcID)
		}
		return fmt.Errorf(("could not find VPC due to: ") + err.Error())
	}

	if isFireNetVpc {
		err := client.DeleteAwsTgwVpcAttachmentForFireNet(awsTgwVpcAttachment)
		if err != nil {
			return fmt.Errorf("failed to detach FireNet VPC from TGW: %s", err)
		}
	} else {
		err := client.DeleteAwsTgwVpcAttachment(awsTgwVpcAttachment)
		if err != nil {
			return fmt.Errorf("failed to detach VPC from TGW: %s", err)
		}
	}

	return nil
}
