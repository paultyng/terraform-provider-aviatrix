package aviatrix

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/AviatrixSystems/terraform-provider-aviatrix/v3/goaviatrix"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func preAvxTransitGatewayPeeringCheck(t *testing.T, msgCommon string) {
	preGatewayCheck(t, msgCommon)
	preGateway2Check(t, msgCommon)
}

func TestAccAviatrixTransitGatewayPeering_basic(t *testing.T) {
	rName := acctest.RandString(5)
	vpcID1 := os.Getenv("AWS_VPC_ID")
	region1 := os.Getenv("AWS_REGION")
	subnet1 := os.Getenv("AWS_SUBNET")
	haSubnet1 := os.Getenv("AWS_SUBNET")

	vpcID2 := os.Getenv("AWS_VPC_ID2")
	region2 := os.Getenv("AWS_REGION2")
	subnet2 := os.Getenv("AWS_SUBNET2")
	haSubnet2 := os.Getenv("AWS_SUBNET2")

	resourceName := "aviatrix_transit_gateway_peering.foo"

	accountName := "megaport-" + rName
	transit1GwName := "spoke-" + rName
	transit1SiteID := "site-" + rName
	path, _ := os.Getwd()
	transit2GwName := "transit-" + rName
	transit2SiteID := "site-" + rName
	resourceNameEdge := "aviatrix_transit_gateway_peering.test_transit_gateway_peering"

	skipAcc := os.Getenv("SKIP_TRANSIT_GATEWAY_PEERING")
	if skipAcc == "yes" {
		t.Skip("Skipping Aviatrix transit gateway peering test as SKIP_TRANSIT_GATEWAY_PEERING is set")
	}
	msgCommon := ". Set SKIP_TRANSIT_GATEWAY_PEERING to yes to skip Aviatrix transit gateway peering tests"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			preAvxTransitGatewayPeeringCheck(t, msgCommon)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTransitGatewayPeeringDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransitGatewayPeeringConfigBasic(rName, vpcID1, vpcID2, region1, region2,
					subnet1, subnet2, haSubnet1, haSubnet2),
				Check: resource.ComposeTestCheckFunc(
					tesAccCheckTransitGatewayPeeringExists("aviatrix_transit_gateway_peering.foo"),
					resource.TestCheckResourceAttr(resourceName, "transit_gateway_name1", fmt.Sprintf("tfg-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "transit_gateway_name2", fmt.Sprintf("tfg2-%s", rName)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccTransitGatewayPeeringConfigEdge(accountName, transit1GwName, transit1SiteID, path, transit2GwName, transit2SiteID),
				Check: resource.ComposeTestCheckFunc(
					tesAccCheckTransitGatewayPeeringExists(resourceNameEdge),
					resource.TestCheckResourceAttr(resourceNameEdge, "transit_gateway_name1", transit1GwName),
					resource.TestCheckResourceAttr(resourceNameEdge, "transit_gateway_name2", transit2GwName),
					resource.TestCheckResourceAttr(resourceNameEdge, "enable_peering_over_private_network", "true"),
					resource.TestCheckResourceAttr(resourceNameEdge, "jumbo_frame", "false"),
					resource.TestCheckResourceAttr(resourceNameEdge, "insane_mode", "true"),
					resource.TestCheckResourceAttr(resourceNameEdge, "gateway1_logical_ifnames.0", "wan1"),
					resource.TestCheckResourceAttr(resourceNameEdge, "gateway2_logical_ifnames.0", "wan1"),
				),
			},
		},
	})
}

func testAccTransitGatewayPeeringConfigBasic(rName string, vpcID1 string, vpcID2 string, region1 string, region2 string,
	subnet1 string, subnet2 string, haSubnet1 string, haSubnet2 string,
) string {
	return fmt.Sprintf(`
resource "aviatrix_account" "test" {
	account_name       = "tfa-%s"
	cloud_type         = 1
	aws_account_number = "%s"
	aws_iam            = false
	aws_access_key     = "%s"
	aws_secret_key     = "%s"
}
resource "aviatrix_transit_gateway" "transitGw1" {
	cloud_type   = 1
	account_name = aviatrix_account.test.account_name
	gw_name      = "tfg-%s"
	vpc_id       = "%s"
	vpc_reg      = "%s"
	gw_size      = "t2.micro"
	subnet       = "%s"
	ha_subnet    = "%s"
	ha_gw_size   = "t2.micro"
}
resource "aviatrix_transit_gateway" "transitGw2" {
	cloud_type   = 1
	account_name = aviatrix_account.test.account_name
	gw_name      = "tfg2-%s"
	vpc_id       = "%s"
	vpc_reg      = "%s"
	gw_size      = "t2.micro"
	subnet       = "%s"
	ha_subnet    = "%s"
	ha_gw_size   = "t2.micro"
}
resource "aviatrix_transit_gateway_peering" "foo" {
	transit_gateway_name1 = aviatrix_transit_gateway.transitGw1.gw_name
	transit_gateway_name2 = aviatrix_transit_gateway.transitGw2.gw_name
}
	`, rName, os.Getenv("AWS_ACCOUNT_NUMBER"), os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"),
		rName, vpcID1, region1, subnet1, haSubnet1, rName, vpcID2, region2, subnet2, haSubnet2)
}

func tesAccCheckTransitGatewayPeeringExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("aviatrix transit gateway peering Not Created: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no aviatrix transit gateway peering ID is set")
		}

		client := testAccProvider.Meta().(*goaviatrix.Client)

		foundTransitGatewayPeering := &goaviatrix.TransitGatewayPeering{
			TransitGatewayName1: rs.Primary.Attributes["transit_gateway_name1"],
			TransitGatewayName2: rs.Primary.Attributes["transit_gateway_name2"],
		}

		err := client.GetTransitGatewayPeering(foundTransitGatewayPeering)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccTransitGatewayPeeringConfigEdge(accountName, transit1GwName, transit1SiteID, path, transit2GwName, transit2SiteID string) string {
	return fmt.Sprintf(`
	resource "aviatrix_account" "test_acc_edge_megaport" {
		account_name       = "edge-%s"
		cloud_type         = 1048576
	}

	resource "aviatrix_transit_gateway" "test_edge_transit_1" {
		cloud_type   = 1048576
		account_name = aviatrix_account.test_acc_edge_megaport.account_name
		gw_name      = "%s"
		vpc_id       = "%s"
		gw_size      = "SMALL"
		ztp_file_download_path = "%s"
		interfaces {
			gateway_ip     = "192.168.20.1"
			ip_address     = "192.168.20.11/24"
			public_ip      = "67.207.104.19"
			logical_ifname = "wan0"
			secondary_private_cidr_list = ["192.168.20.16/29"]
		}
		interfaces {
			gateway_ip     = "192.168.21.1"
			ip_address     = "192.168.21.11/24"
			public_ip      = "67.71.12.148"
			logical_ifname = "wan1"
			secondary_private_cidr_list = ["192.168.21.16/29"]
		}
		interfaces {
			dhcp           = true
			logical_ifname = "mgmt0"
		}
		interfaces {
			gateway_ip     = "192.168.22.1"
			ip_address     = "192.168.22.11/24"
			logical_ifname = "wan2"
		}
		interfaces {
			gateway_ip     = "192.168.23.1"
			ip_address     = "192.168.23.11/24"
			logical_ifname = "wan3"
		}
	}

	resource "aviatrix_transit_gateway" "test_edge_transit_2" {
		cloud_type   = 1048576
		account_name = aviatrix_account.test_acc_edge_megaport.account_name
		gw_name      = "%s"
		vpc_id       = "%s"
		gw_size      = "SMALL"
		ztp_file_download_path = "%s"
		interfaces {
			gateway_ip     = "192.168.24.1"
			ip_address     = "192.168.24.11/24"
			public_ip      = "67.207.104.24"
			logical_ifname = "wan0"
			secondary_private_cidr_list = ["192.168.24.16/29"]
		}
		interfaces {
			gateway_ip     = "192.168.25.1"
			ip_address     = "192.168.25.11/24"
			public_ip      = "67.71.12.25"
			logical_ifname = "wan1"
			secondary_private_cidr_list = ["192.168.25.16/29"]
		}
		interfaces {
			dhcp           = true
			logical_ifname = "mgmt0"
		}
		interfaces {
			gateway_ip     = "192.168.26.1"
			ip_address     = "192.168.26.11/24"
			logical_ifname = "wan2"
		}
		interfaces {
			gateway_ip     = "192.168.27.1"
			ip_address     = "192.168.27.11/24"
			logical_ifname = "wan3"
		}
	}

	resource "aviatrix_transit_gateway_peering" "test_transit_gateway_peering" {
		transit_gateway_name1 = aviatrix_transit_gateway.test_edge_transit_1.gw_name
		transit_gateway_name2 = aviatrix_transit_gateway.test_edge_transit_2.gw_name
		enable_peering_over_private_network = true
		jumbo_frame = false
		insane_mode = true
		gateway1_logical_ifnames = ["wan1"]
		gateway2_logical_ifnames = ["wan1"]
	}
		`, accountName, transit1GwName, transit1SiteID, path, transit2GwName, transit2SiteID, path)
}

func testAccCheckTransitGatewayPeeringDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*goaviatrix.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aviatrix_transit_gateway_peering" {
			continue
		}

		foundTransitGatewayPeering := &goaviatrix.TransitGatewayPeering{
			TransitGatewayName1: rs.Primary.Attributes["transit_gateway_name1"],
			TransitGatewayName2: rs.Primary.Attributes["transit_gateway_name2"],
		}

		err := client.GetTransitGatewayPeering(foundTransitGatewayPeering)
		if err != goaviatrix.ErrNotFound {
			return fmt.Errorf("aviatrix transit gateway peering still exists")
		}
	}

	return nil
}

func TestReverseIfnameTranslation(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected map[string]string
	}{
		{
			name: "Basic translation reversal",
			input: map[string]string{
				"eth0": "wan.0",
				"eth1": "wan.1",
			},
			expected: map[string]string{
				"wan0": "eth0",
				"wan1": "eth1",
			},
		},
		{
			name:     "Handles empty map",
			input:    map[string]string{},
			expected: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseIfnameTranslation(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test %s failed. Expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestSetWanInterfaces(t *testing.T) {
	tests := []struct {
		name                   string
		logicalIfNames         []interface{}
		reversedInterfaceNames map[string]string
		expected               string
		expectErr              bool
	}{
		{
			name:           "Valid input - single interface",
			logicalIfNames: []interface{}{"wan0"},
			reversedInterfaceNames: map[string]string{
				"wan0": "eth0",
			},
			expected:  "eth0",
			expectErr: false,
		},
		{
			name:           "Valid input - multiple interfaces",
			logicalIfNames: []interface{}{"wan0", "wan1"},
			reversedInterfaceNames: map[string]string{
				"wan0": "eth0",
				"wan1": "eth1",
			},
			expected:  "eth0,eth1",
			expectErr: false,
		},
		{
			name:           "Interface name not found in map",
			logicalIfNames: []interface{}{"wan2"},
			reversedInterfaceNames: map[string]string{
				"wan0": "eth0",
			},
			expected:  "",
			expectErr: true,
		},
		{
			name:           "Empty logicalIfNames",
			logicalIfNames: []interface{}{},
			reversedInterfaceNames: map[string]string{
				"wan0": "eth0",
			},
			expected:  "",
			expectErr: false, // Empty input should return an empty string
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SetWanInterfaces(tt.logicalIfNames, tt.reversedInterfaceNames)
			if (err != nil) != tt.expectErr {
				t.Errorf("Test %s failed: expected error=%v, got error=%v", tt.name, tt.expectErr, err != nil)
			}
			if !tt.expectErr && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Test %s failed: expected %q, got %q", tt.name, tt.expected, result)
			}
		})
	}
}

func TestSetWanInterfaceNames(t *testing.T) {
	tests := []struct {
		name           string
		logicalIfNames []string
		cloudType      int
		gatewayDetails *goaviatrix.Gateway
		gatewayPrefix  string
		expectedError  bool
	}{
		{
			name:           "Valid logical interfaces for Equinix cloud type",
			logicalIfNames: []string{"wan0", "wan1"},
			cloudType:      goaviatrix.EDGEEQUINIX,
			gatewayDetails: &goaviatrix.Gateway{IfNamesTranslation: map[string]string{"eth0": "wan.0", "eth1": "wan.1"}},
			gatewayPrefix:  "gateway1",
			expectedError:  false,
		},
		{
			name:           "Valid logical interfaces for Megaport cloud type",
			logicalIfNames: []string{"wan0", "wan1"},
			cloudType:      goaviatrix.EDGEMEGAPORT,
			gatewayDetails: &goaviatrix.Gateway{},
			gatewayPrefix:  "gateway1",
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transitGatewayPeering := &goaviatrix.TransitGatewayPeering{}

			// Call the function with actual cloud type
			err := setWanInterfaceNames(tt.logicalIfNames, tt.cloudType, tt.gatewayDetails, tt.gatewayPrefix, transitGatewayPeering)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.cloudType == goaviatrix.EDGEEQUINIX || tt.cloudType == goaviatrix.EDGENEO {
					assert.NotEmpty(t, transitGatewayPeering.SrcWanInterfaces)
				} else if tt.cloudType == goaviatrix.EDGEMEGAPORT {
					// Check the logical interfaces set for Megaport
					if tt.gatewayPrefix == "gateway1" {
						assert.Equal(t, tt.logicalIfNames, transitGatewayPeering.Gateway1LogicalIfNames)
					} else {
						assert.Equal(t, tt.logicalIfNames, transitGatewayPeering.Gateway2LogicalIfNames)
					}
				}
			}
		})
	}
}

func TestGetLogicalIfNames(t *testing.T) {
	gateway := &goaviatrix.Gateway{
		IfNamesTranslation: map[string]string{
			"eth0": "wan.0",
			"eth1": "wan.1",
			"eth2": "mgmt.0",
			"eth3": "wan.2",
			"eth4": "wan.3",
		},
	}

	tests := []struct {
		name          string
		interfaceList []string
		expected      []string
		expectErr     bool
	}{
		{
			name:          "Valid interface list",
			interfaceList: []string{"eth0", "eth1", "eth2", "eth3", "eth4"},
			expected:      []string{"wan0", "wan1", "mgmt0", "wan2", "wan3"},
			expectErr:     false,
		},
		{
			name:          "Single valid interface",
			interfaceList: []string{"eth0"},
			expected:      []string{"wan0"},
			expectErr:     false,
		},
		{
			name:          "Interface not found",
			interfaceList: []string{"eth5"},
			expected:      nil,
			expectErr:     true,
		},
		{
			name:          "Mixed valid and invalid interfaces",
			interfaceList: []string{"eth0", "eth5"},
			expected:      nil,
			expectErr:     true,
		},
		{
			name:          "Empty interface list",
			interfaceList: []string{},
			expected:      []string{},
			expectErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getLogicalIfNames(gateway, tt.interfaceList)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if result == nil {
				result = []string{}
			}
			if tt.expected == nil {
				tt.expected = []string{}
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
