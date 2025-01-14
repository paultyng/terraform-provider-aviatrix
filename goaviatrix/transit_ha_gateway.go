package goaviatrix

import (
	"golang.org/x/net/context"
)

type TransitHaGateway struct {
	Action                string `json:"action"`
	CID                   string `json:"CID"`
	AccountName           string `json:"account_name"`
	CloudType             int    `json:"cloud_type"`
	VpcID                 string `json:"vpc_id,omitempty"`
	VNetNameResourceGroup string `json:"vnet_and_resource_group_names"`
	PrimaryGwName         string `json:"primary_gw_name"`
	GwName                string `json:"ha_gw_name"`
	GwSize                string `json:"gw_size"`
	Subnet                string `json:"gw_subnet"`
	VpcRegion             string `json:"region"`
	Zone                  string `json:"zone"`
	AvailabilityDomain    string `json:"availability_domain"`
	FaultDomain           string `json:"fault_domain"`
	BgpLanVpcId           string `json:"bgp_lan_vpc"`
	BgpLanSubnet          string `json:"bgp_lan_specify_subnet"`
	Eip                   string `json:"eip,omitempty"`
	InsaneMode            string `json:"insane_mode"`
	TagList               string `json:"tag_string"`
	TagJson               string `json:"tag_json"`
	AutoGenHaGwName       string `json:"autogen_hagw_name"`
	BackupLinkList        []BackupLinkInterface
	BackupLinkConfig      string `json:"backup_link_config,omitempty"`
	InterfaceMapping      string `json:"interface_mapping,omitempty"`
	Interfaces            string `json:"interfaces,omitempty"`
	DeviceID              string `json:"device_id,omitempty"`
	ZtpFileDownloadPath   string `json:"-"`
}

type BackupLinkInterface struct {
	PeerGwName     string `json:"peer_gw_name"`
	PeerBackupPort string `json:"peer_backup_port"`
	SelfBackupPort string `json:"self_backup_port"`
	ConnectionType string `json:"connection_type"`
}

func (c *Client) CreateTransitHaGw(transitHaGateway *TransitHaGateway) (string, error) {
	transitHaGateway.CID = c.CID
	transitHaGateway.Action = "create_multicloud_ha_gateway"
	var data CreateEdgeEquinixResp
	resp, err := c.PostAPIContext2HaGw(context.Background(), &data, transitHaGateway.Action, transitHaGateway, BasicCheck)
	if err != nil {
		return "", err
	}
	// create the ZTP file for Equinix Edge transit gateway
	if transitHaGateway.CloudType == EDGEEQUINIX {
		fileName := getFileName(transitHaGateway.ZtpFileDownloadPath, transitHaGateway.GwName, transitHaGateway.VpcID)
		err = createZtpFile(fileName, data.Result)
		if err != nil {
			return "", err
		}
	}
	return resp, nil
}
