package aviatrix

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-aviatrix/goaviatrix"
)

func resourceAviatrixSplunkLogging() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviatrixSplunkLoggingCreate,
		Read:   resourceAviatrixSplunkLoggingRead,
		Update: resourceAviatrixSplunkLoggingUpdate,
		Delete: resourceAviatrixSplunkLoggingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Server IP",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     -1,
				Description: "Port number",
			},
			"cu_output_cfg_file_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Configuration file path",
			},
			"custom_input_cfg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Custom configuration",
			},
			"excluded_gateways": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of excluded gateways.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Enabled or not.",
			},
		},
	}
}

func marshalSplunkLoggingInput(d *schema.ResourceData, useCustomConfig bool) *goaviatrix.SplunkLogging {
	if useCustomConfig {
		return &goaviatrix.SplunkLogging{
			UseConfigFile:  true,
			ConfigFilePath: d.Get("cu_output_cfg_file_path").(string),
			CustomConfig:   d.Get("custom_input_cfg").(string),
		}
	} else {
		return &goaviatrix.SplunkLogging{
			UseConfigFile: false,
			Server:        d.Get("server").(string),
			Port:          d.Get("port").(int),
			CustomConfig:  d.Get("custom_input_cfg").(string),
		}
	}
}

func resourceAviatrixSplunkLoggingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)

	var splunkLogging *goaviatrix.SplunkLogging

	if d.Get("server").(string) == "" && d.Get("port").(int) == -1 && d.Get("cu_output_cfg_file_path").(string) == "" {
		return fmt.Errorf("please provide either server/port or configuration file path")
	} else if d.Get("cu_output_cfg_file_path").(string) != "" {
		splunkLogging = marshalSplunkLoggingInput(d, true)
	} else {
		if d.Get("port").(int) == -1 || d.Get("server").(string) == "" {
			return fmt.Errorf("please provide both server and port")
		}

		splunkLogging = marshalSplunkLoggingInput(d, false)
	}

	var excludedGateways []string
	for _, v := range d.Get("excluded_gateways").(*schema.Set).List() {
		excludedGateways = append(excludedGateways, v.(string))
	}
	if len(excludedGateways) != 0 {
		splunkLogging.ExcludedGatewaysInput = strings.Join(excludedGateways, ",")
	}

	if err := client.EnableSplunkLogging(splunkLogging); err != nil {
		return fmt.Errorf("could not enable splunk logging: %v", err)
	}

	d.SetId("splunk_logging")
	return nil
}
func resourceAviatrixSplunkLoggingRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)

	if d.Id() != "splunk_logging" {
		return fmt.Errorf("invalid ID, expected ID \"splunk_logging\", instead got %s", d.Id())
	}

	splunkLoggingStatus, err := client.GetSplunkLoggingStatus()
	if err != nil {
		return fmt.Errorf("could not get remote syslog status: %v", err)
	}

	d.Set("server", splunkLoggingStatus.Server)
	port, _ := strconv.Atoi(splunkLoggingStatus.Port)
	d.Set("port", port)
	d.Set("custom_input_cfg", splunkLoggingStatus.CustomConfig)
	d.Set("status", splunkLoggingStatus.Status)
	d.Set("excluded_gateways", splunkLoggingStatus.ExcludedGateways)

	d.SetId("splunk_logging")
	return nil
}

func resourceAviatrixSplunkLoggingUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAviatrixSplunkLoggingCreate(d, meta)
}

func resourceAviatrixSplunkLoggingDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goaviatrix.Client)

	if err := client.DisableSplunkLogging(); err != nil {
		return fmt.Errorf("could not disable remote syslog: %v", err)
	}

	d.SetId("")
	return nil
}
