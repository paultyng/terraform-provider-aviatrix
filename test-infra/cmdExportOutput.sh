ARM_APPLICATION_ID=$(terraform output ARM_APPLICATION_ID)
ARM_APPLICATION_KEY=$(terraform output ARM_APPLICATION_KEY)
ARM_DIRECTORY_ID=$(terraform output ARM_DIRECTORY_ID)
ARM_SUBSCRIPTION_ID=$(terraform output ARM_SUBSCRIPTION_ID)
AZURE_GW_SIZE=$(terraform output AZURE_GW_SIZE)
AZURE_REGION=$(terraform output AZURE_REGION)
AZURE_REGION2=$(terraform output AZURE_REGION2)
AZURE_SUBNET=$(terraform output AZURE_SUBNET)
AZURE_SUBNET2=$(terraform output AZURE_SUBNET2)
AZURE_VNET_ID=$(terraform output AZURE_VNET_ID)
AZURE_VNET_ID2=$(terraform output AZURE_VNET_ID2)
AVIATRIX_CONTROLLER_IP=$(terraform output AVIATRIX_CONTROLLER_IP)
AVIATRIX_PASSWORD=$(terraform output AVIATRIX_PASSWORD)
AVIATRIX_USERNAME=$(terraform output AVIATRIX_USERNAME)
AWS_ACCESS_KEY=$(terraform output AWS_ACCESS_KEY)
AWS_ACCOUNT_NUMBER=$(terraform output AWS_ACCOUNT_NUMBER)
AWS_BGP_VGW_ID=$(terraform output AWS_BGP_VGW_ID)
AWS_REGION=$(terraform output AWS_REGION)
AWS_REGION2=$(terraform output AWS_REGION2)
AWS_SECRET_KEY=$(terraform output AWS_SECRET_KEY)
AWS_SUBNET=$(terraform output AWS_SUBNET)
AWS_SUBNET2=$(terraform output AWS_SUBNET2)
AWS_VPC_ID=$(terraform output AWS_VPC_ID)
AWS_VPC_ID2=$(terraform output AWS_VPC_ID2)
AWS_DX_GATEWAY_ID=$(terraform output AWS_DX_GATEWAY_ID)
DOMAIN_NAME=$(terraform output DOMAIN_NAME)
AWSGOV_ACCESS_KEY=$(terraform output AWSGOV_ACCESS_KEY)
AWSGOV_SECRET_KEY=$(terraform output AWSGOV_SECRET_KEY)
AWSGOV_ACCOUNT_NUMBER=$(terraform output AWSGOV_ACCOUNT_NUMBER)
AWSGOV_REGION=$(terraform output AWSGOV_REGION)
AWSGOV_SUBNET=$(terraform output AWSGOV_SUBNET)
AWSGOV_VPC_ID=$(terraform output AWSGOV_VPC_ID)
DEVICE_PUBLIC_IP=$(terraform output DEVICE_PUBLIC_IP)
DEVICE_KEY_FILE_PATH=$(terraform output DEVICE_KEY_FILE_PATH)
GCP_CREDENTIALS_FILEPATH=$(terraform output GCP_CREDENTIALS_FILEPATH)
GCP_ID=$(terraform output GCP_ID)
GCP_SUBNET=$(terraform output GCP_SUBNET)
GCP_VPC_ID=$(terraform output GCP_VPC_ID)
GCP_ZONE=$(terraform output GCP_ZONE)
IDP_METADATA=$(terraform output IDP_METADATA)
IDP_METADATA_TYPE=$(terraform output IDP_METADATA_TYPE)
OCI_API_KEY_FILEPATH=$(terraform output OCI_API_KEY_FILEPATH)
OCI_COMPARTMENT_ID=$(terraform output OCI_COMPARTMENT_ID)
OCI_REGION=$(terraform output OCI_REGION)
OCI_SUBNET=$(terraform output OCI_SUBNET)
OCI_TENANCY_ID=$(terraform output OCI_TENANCY_ID)
OCI_USER_ID=$(terraform output OCI_USER_ID)
OCI_VPC_ID=$(terraform output OCI_VPC_ID)
controller_private_ip=$(terraform output controller_private_ip)
TRANSIT_GATEWAY_NAME=$(terraform output TRANSIT_GATEWAY_NAME)
ARM_RESOURCE_GROUP=$(terraform output ARM_RESOURCE_GROUP)
ARM_HUB_NAME=$(terraform output ARM_HUB_NAME)
AWS_TGW_NAME=$(terraform output AWS_TGW_NAME)
DATADOG_API_KEY=$(terraform output DATADOG_API_KEY)

SKIP_DATA_ACCOUNT="no"
SKIP_DATA_CALLER_IDENTITY="no"
SKIP_DATA_FIRENET="no"
SKIP_DATA_FIRENET_VENDOR_INTEGRATION="no"
SKIP_DATA_FIREWALL="no"
SKIP_DATA_GATEWAY="no"
SKIP_DATA_SPOKE_GATEWAY="no"
SKIP_DATA_SPOKE_GATEWAY_AWS="no"
SKIP_DATA_SPOKE_GATEWAY_AZURE="no"
SKIP_DATA_SPOKE_GATEWAY_GCP="no"
SKIP_DATA_TRANSIT_GATEWAY="no"
SKIP_DATA_TRANSIT_GATEWAY_AWS="no"
SKIP_DATA_TRANSIT_GATEWAY_AZURE="no"
SKIP_DATA_TRANSIT_GATEWAY_GCP="no"
SKIP_DATA_VPC="no"
SKIP_DATA_VPC_TRACKER="no"
SKIP_ACCOUNT="no"
SKIP_ACCOUNT_AWS="no"
SKIP_ACCOUNT_AZURE="no"
SKIP_ACCOUNT_GCP="no"
SKIP_ACCOUNT_OCI="yes"
SKIP_ACCOUNT_AWSGOV="no"
SKIP_ACCOUNT_USER="no"
SKIP_ARM_PEER="yes"
SKIP_AWS_GUARD_DUTY="no"
SKIP_AWS_PEER="no"
SKIP_AWS_TGW="no"
SKIP_AWS_TGW_DIRECTCONNECT="no"
SKIP_AWS_TGW_PEERING="no"
SKIP_AWS_TGW_PEERING_DOMAIN_CONN="no"
SKIP_AWS_TGW_TRANSIT_GATEWAY_ATTACHMENT="no"
SKIP_AWS_TGW_VPC_ATTACHMENT="no"
SKIP_AWS_TGW_VPN_CONN="no"
SKIP_AZURE_PEER="no"
SKIP_AZURE_SPOKE_NATIVE_PEERING="no"
SKIP_CONTROLLER_CONFIG="no"
SKIP_DATADOG_AGENT="no"
SKIP_DEVICE_AWS_TGW_ATTACHMENT="no"
SKIP_DEVICE_INTERFACE_CONFIG="no"
SKIP_DEVICE_REGISTRATION="no"
SKIP_DEVICE_TAG="no"
SKIP_DEVICE_TRANSIT_GATEWAY_ATTACHMENT="no"
SKIP_DEVICE_VIRTUAL_WAN_ATTACHMENT="no"
SKIP_FIRENET="no"
SKIP_FIREWALL="no"
SKIP_FIREWALL_INSTANCE="no"
SKIP_FIREWALL_INSTANCE_ASSOCIATION="no"
SKIP_FIREWALL_MANAGEMENT_ACCESS="no"
SKIP_FIREWALL_MANAGEMENT_ACCESS_AWS="no"
SKIP_FIREWALL_MANAGEMENT_ACCESS_AZURE="no"
SKIP_FIREWALL_POLICY="no"
SKIP_FIREWALL_TAG="no"
SKIP_FQDN="no"
SKIP_FQDN_PASS_THROUGH="no"
SKIP_FQDN_TAG_RULE="no"
SKIP_GATEWAY="no"
SKIP_GATEWAY_AWS="no"
SKIP_GATEWAY_GCP="no"
SKIP_GATEWAY_AZURE="no"
SKIP_GATEWAY_OCI="yes"
SKIP_GATEWAY_AWSGOV="yes"
SKIP_GATEWAY_DNAT="no"
SKIP_GATEWAY_DNAT_AWS="no"
SKIP_GATEWAY_DNAT_AZURE="no"
SKIP_GATEWAY_SNAT="no"
SKIP_GATEWAY_SNAT_AWS="no"
SKIP_GATEWAY_SNAT_AZURE="yes"
SKIP_GEO_VPN="no"
SKIP_PERIODIC_PING="no"
SKIP_RBAC_GROUP="no"
SKIP_RBAC_GROUP_ACCESS_ACCOUNT_ATTACHMENT="no"
SKIP_RBAC_GROUP_PERMISSION_ATTACHMENT="no"
SKIP_RBAC_GROUP_USER_ATTACHMENT="no"
SKIP_REMOTE_SYSLOG="no"
SKIP_SAML_ENDPOINT="no"
SKIP_S2C="no"
SKIP_SEGMENTATION_SECURITY_DOMAIN="no"
SKIP_SEGMENTATION_SECURITY_DOMAIN_ASSOCIATION="no"
SKIP_SEGMENTATION_SECURITY_DOMAIN_CONNECTION_POLICY="no"
SKIP_SPLUNK_LOGGING="no"
SKIP_SPOKE="yes"
SKIP_SPOKE_AZURE="yes"
SKIP_SPOKE_AWS="yes"
SKIP_SPOKE_GCP="yes"
SKIP_SPOKE_GATEWAY="no"
SKIP_SPOKE_GATEWAY_AZURE="no"
SKIP_SPOKE_GATEWAY_AWS="no"
SKIP_SPOKE_GATEWAY_GCP="no"
SKIP_SPOKE_GATEWAY_OCI="yes"
SKIP_SUMOLOGIC_FORWARDER="no"
SKIP_TRANS_PEER="no"
SKIP_TRANSIT="yes"
SKIP_TRANSIT_AWS="yes"
SKIP_TRANSIT_AZURE="yes"
SKIP_TRANSIT_EXTERNAL_DEVICE_CONN="no"
SKIP_TRANSIT_FIRENET_POLICY="no"
SKIP_TRANSIT_FIRENET_POLICY_AWS="no"
SKIP_TRANSIT_FIRENET_POLICY_AZURE="no"
SKIP_TRANSIT_GATEWAY="no"
SKIP_TRANSIT_GATEWAY_AWS="no"
SKIP_TRANSIT_GATEWAY_AZURE="no"
SKIP_TRANSIT_GATEWAY_GCP="no"
SKIP_TRANSIT_GATEWAY_OCI="yes"
SKIP_TRANSIT_GATEWAY_PEERING="no"
SKIP_TUNNEL="no"
SKIP_VPC="no"
SKIP_VGW_CONN="no"
SKIP_VPN_PROFILE="no"
SKIP_VPN_USER="no"
SKIP_VPN_USER_ACCELERATOR="no"

echo "ARM_APPLICATION_ID=$ARM_APPLICATION_ID"
echo "ARM_APPLICATION_KEY=$ARM_APPLICATION_KEY"
echo "ARM_DIRECTORY_ID=$ARM_DIRECTORY_ID"
echo "ARM_SUBSCRIPTION_ID=$ARM_SUBSCRIPTION_ID"
echo "AZURE_GW_SIZE=$AZURE_GW_SIZE"
echo "AZURE_REGION=$AZURE_REGION"
echo "AZURE_REGION2=$AZURE_REGION2"
echo "AZURE_SUBNET=$AZURE_SUBNET"
echo "AZURE_SUBNET2=$AZURE_SUBNET2"
echo "AZURE_VNET_ID=$AZURE_VNET_ID"
echo "AZURE_VNET_ID2=$AZURE_VNET_ID2"
echo "AVIATRIX_CONTROLLER_IP=$AVIATRIX_CONTROLLER_IP"
echo "AVIATRIX_PASSWORD=$AVIATRIX_PASSWORD"
echo "AVIATRIX_USERNAME=$AVIATRIX_USERNAME"
echo "AWS_ACCESS_KEY=$AWS_ACCESS_KEY"
echo "AWS_ACCOUNT_NUMBER=$AWS_ACCOUNT_NUMBER"
echo "AWS_BGP_VGW_ID=$AWS_BGP_VGW_ID"
echo "AWS_REGION=$AWS_REGION"
echo "AWS_REGION2=$AWS_REGION2"
echo "AWS_SECRET_KEY=$AWS_SECRET_KEY"
echo "AWS_SUBNET=$AWS_SUBNET"
echo "AWS_SUBNET2=$AWS_SUBNET2"
echo "AWS_VPC_ID=$AWS_VPC_ID"
echo "AWS_VPC_ID2=$AWS_VPC_ID2"
echo "AWS_DX_GATEWAY_ID=$AWS_DX_GATEWAY_ID"
echo "DOMAIN_NAME=$DOMAIN_NAME"
echo "AWSGOV_ACCESS_KEY=$AWSGOV_ACCESS_KEY"
echo "AWSGOV_ACCOUNT_NUMBER=$AWSGOV_ACCOUNT_NUMBER"
echo "AWSGOV_SECRET_KEY=$AWSGOV_SECRET_KEY"
echo "DEVICE_PUBLIC_IP=$DEVICE_PUBLIC_IP"
echo "DEVICE_KEY_FILE_PATH=$DEVICE_KEY_FILE_PATH"
echo "GCP_CREDENTIALS_FILEPATH=$GCP_CREDENTIALS_FILEPATH"
echo "GCP_ID=$GCP_ID"
echo "GCP_SUBNET=$GCP_SUBNET"
echo "GCP_VPC_ID=$GCP_VPC_ID"
echo "GCP_ZONE=$GCP_ZONE"
echo "IDP_METADATA=$IDP_METADATA"
echo "IDP_METADATA_TYPE=$IDP_METADATA_TYPE"
echo "OCI_API_KEY_FILEPATH=$OCI_API_KEY_FILEPATH"
echo "OCI_COMPARTMENT_ID=$OCI_COMPARTMENT_ID"
echo "OCI_REGION=$OCI_REGION"
echo "OCI_SUBNET=$OCI_SUBNET"
echo "OCI_TENANCY_ID=$OCI_TENANCY_ID"
echo "OCI_USER_ID=$OCI_USER_ID"
echo "OCI_VPC_ID=$OCI_VPC_ID"
echo "controller_private_ip=$controller_private_ip"
echo "TRANSIT_GATEWAY_NAME=$TRANSIT_GATEWAY_NAME"
echo "ARM_RESOURCE_GROUP=$ARM_RESOURCE_GROUP"
echo "ARM_HUB_NAME=$ARM_HUB_NAME"
echo "AWS_TGW_NAME=$AWS_TGW_NAME"
echo "DATADOG_API_KEY=$DATADOG_API_KEY"

echo "SKIP_DATA_ACCOUNT=$SKIP_DATA_ACCOUNT"
echo "SKIP_DATA_CALLER_IDENTITY=$SKIP_DATA_CALLER_IDENTITY"
echo "SKIP_DATA_FIRENET=$SKIP_DATA_FIRENET"
echo "SKIP_DATA_FIRENET_VENDOR_INTEGRATION=$SKIP_DATA_FIRENET_VENDOR_INTEGRATION"
echo "SKIP_DATA_FIREWALL=$SKIP_DATA_FIREWALL"
echo "SKIP_DATA_GATEWAY=$SKIP_DATA_GATEWAY"
echo "SKIP_DATA_SPOKE_GATEWAY=$SKIP_DATA_SPOKE_GATEWAY"
echo "SKIP_DATA_SPOKE_GATEWAY_AWS=$SKIP_DATA_SPOKE_GATEWAY_AWS"
echo "SKIP_DATA_SPOKE_GATEWAY_AZURE=$SKIP_DATA_SPOKE_GATEWAY_AZURE"
echo "SKIP_DATA_SPOKE_GATEWAY_GCP=$SKIP_DATA_SPOKE_GATEWAY_GCP"
echo "SKIP_DATA_TRANSIT_GATEWAY=$SKIP_DATA_TRANSIT_GATEWAY"
echo "SKIP_DATA_TRANSIT_GATEWAY_AWS=$SKIP_DATA_TRANSIT_GATEWAY_AWS"
echo "SKIP_DATA_TRANSIT_GATEWAY_AZURE=$SKIP_DATA_TRANSIT_GATEWAY_AZURE"
echo "SKIP_DATA_TRANSIT_GATEWAY_GCP=$SKIP_DATA_TRANSIT_GATEWAY_GCP"
echo "SKIP_DATA_VPC=$SKIP_DATA_VPC"
echo "SKIP_DATA_VPC_TRACKER=$SKIP_DATA_VPC_TRACKER"
echo "SKIP_ACCOUNT=$SKIP_ACCOUNT"
echo "SKIP_ACCOUNT_AZURE=$SKIP_ACCOUNT_AZURE"
echo "SKIP_ACCOUNT_AWS=$SKIP_ACCOUNT_AWS"
echo "SKIP_ACCOUNT_GCP=$SKIP_ACCOUNT_GCP"
echo "SKIP_ACCOUNT_OCI=$SKIP_ACCOUNT_OCI"
echo "SKIP_ACCOUNT_AWSGOV=$SKIP_ACCOUNT_AWSGOV"
echo "SKIP_ACCOUNT_USER=$SKIP_ACCOUNT_USER"
echo "SKIP_ARM_PEER=$SKIP_ARM_PEER"
echo "SKIP_AWS_GUARD_DUTY=$SKIP_AWS_GUARD_DUTY"
echo "SKIP_AWS_PEER=$SKIP_AWS_PEER"
echo "SKIP_AWS_TGW=$SKIP_AWS_TGW"
echo "SKIP_AWS_TGW_DIRECTCONNECT=$SKIP_AWS_TGW_DIRECTCONNECT"
echo "SKIP_AWS_TGW_PEERING=$SKIP_AWS_TGW_PEERING"
echo "SKIP_AWS_TGW_PEERING_DOMAIN_CONN=$SKIP_AWS_TGW_PEERING_DOMAIN_CONN"
echo "SKIP_AWS_TGW_TRANSIT_GATEWAY_ATTACHMENT=$SKIP_AWS_TGW_TRANSIT_GATEWAY_ATTACHMENT"
echo "SKIP_AWS_TGW_VPC_ATTACHMENT=$SKIP_AWS_TGW_VPC_ATTACHMENT"
echo "SKIP_AWS_TGW_VPN_CONN=$SKIP_AWS_TGW_VPN_CONN"
echo "SKIP_AZURE_PEER=$SKIP_AZURE_PEER"
echo "SKIP_AZURE_SPOKE_NATIVE_PEERING=$SKIP_AZURE_SPOKE_NATIVE_PEERING"
echo "SKIP_CONTROLLER_CONFIG=$SKIP_CONTROLLER_CONFIG"
echo "SKIP_DATADOG_AGENT=$SKIP_DATADOG_AGENT"
echo "SKIP_DEVICE_AWS_TGW_ATTACHMENT=$SKIP_DEVICE_AWS_TGW_ATTACHMENT"
echo "SKIP_DEVICE_INTERFACE_CONFIG=$SKIP_DEVICE_INTERFACE_CONFIG"
echo "SKIP_DEVICE_REGISTRATION=$SKIP_DEVICE_REGISTRATION"
echo "SKIP_DEVICE_TAG=$SKIP_DEVICE_TAG"
echo "SKIP_DEVICE_TRANSIT_GATEWAY_ATTACHMENT=$SKIP_DEVICE_TRANSIT_GATEWAY_ATTACHMENT"
echo "SKIP_DEVICE_VIRTUAL_WAN_ATTACHMENT=$SKIP_DEVICE_VIRTUAL_WAN_ATTACHMENT"
echo "SKIP_FIRENET=$SKIP_FIRENET"
echo "SKIP_FIREWALL=$SKIP_FIREWALL"
echo "SKIP_FIREWALL_INSTANCE=$SKIP_FIREWALL_INSTANCE"
echo "SKIP_FIREWALL_INSTANCE_ASSOCIATION=$SKIP_FIREWALL_INSTANCE_ASSOCIATION"
echo "SKIP_FIREWALL_MANAGEMENT_ACCESS=$SKIP_FIREWALL_MANAGEMENT_ACCESS"
echo "SKIP_FIREWALL_MANAGEMENT_ACCESS_AWS=$SKIP_FIREWALL_MANAGEMENT_ACCESS_AWS"
echo "SKIP_FIREWALL_MANAGEMENT_ACCESS_AZURE=$SKIP_FIREWALL_MANAGEMENT_ACCESS_AZURE"
echo "SKIP_FIREWALL_POLICY=$SKIP_FIREWALL_POLICY"
echo "SKIP_FIREWALL_TAG=$SKIP_FIREWALL_TAG"
echo "SKIP_FQDN=$SKIP_FQDN"
echo "SKIP_FQDN_PASS_THROUGH=$SKIP_FQDN_PASS_THROUGH"
echo "SKIP_FQDN_TAG_RULE=$SKIP_FQDN_TAG_RULE"
echo "SKIP_GATEWAY=$SKIP_GATEWAY"
echo "SKIP_GATEWAY_AZURE=$SKIP_GATEWAY_AZURE"
echo "SKIP_GATEWAY_AWS=$SKIP_GATEWAY_AWS"
echo "SKIP_GATEWAY_GCP=$SKIP_GATEWAY_GCP"
echo "SKIP_GATEWAY_OCI=$SKIP_GATEWAY_OCI"
echo "SKIP_GATEWAY_AWSGOV=$SKIP_GATEWAY_AWSGOV"
echo "SKIP_GATEWAY_DNAT=$SKIP_GATEWAY_DNAT"
echo "SKIP_GATEWAY_DNAT_AZURE=$SKIP_GATEWAY_DNAT_AZURE"
echo "SKIP_GATEWAY_DNAT_AWS=$SKIP_GATEWAY_DNAT_AWS"
echo "SKIP_GATEWAY_SNAT=$SKIP_GATEWAY_SNAT"
echo "SKIP_GATEWAY_SNAT_AZURE=$SKIP_GATEWAY_SNAT_AZURE"
echo "SKIP_GATEWAY_SNAT_AWS=$SKIP_GATEWAY_SNAT_AWS"
echo "SKIP_GEO_VPN=$SKIP_GEO_VPN"
echo "SKIP_PERIODIC_PING=$SKIP_PERIODIC_PING"
echo "SKIP_RBAC_GROUP=$SKIP_RBAC_GROUP"
echo "SKIP_RBAC_GROUP_ACCESS_ACCOUNT_ATTACHMENT=$SKIP_RBAC_GROUP_ACCESS_ACCOUNT_ATTACHMENT"
echo "SKIP_RBAC_GROUP_PERMISSION_ATTACHMENT=$SKIP_RBAC_GROUP_PERMISSION_ATTACHMENT"
echo "SKIP_RBAC_GROUP_USER_ATTACHMENT=$SKIP_RBAC_GROUP_USER_ATTACHMENT"
echo "SKIP_REMOTE_SYSLOG=$SKIP_REMOTE_SYSLOG"
echo "SKIP_SAML_ENDPOINT=$SKIP_SAML_ENDPOINT"
echo "SKIP_S2C=$SKIP_S2C"
echo "SKIP_SEGMENTATION_SECURITY_DOMAIN=$SKIP_SEGMENTATION_SECURITY_DOMAIN"
echo "SKIP_SEGMENTATION_SECURITY_DOMAIN_ASSOCIATION=$SKIP_SEGMENTATION_SECURITY_DOMAIN_ASSOCIATION"
echo "SKIP_SEGMENTATION_SECURITY_DOMAIN_CONNECTION_POLICY=$SKIP_SEGMENTATION_SECURITY_DOMAIN_CONNECTION_POLICY"
echo "SKIP_SPLUNK_LOGGING=$SKIP_SPLUNK_LOGGING"
echo "SKIP_SPOKE=$SKIP_SPOKE"
echo "SKIP_SPOKE_AZURE=$SKIP_SPOKE_AZURE"
echo "SKIP_SPOKE_AWS=$SKIP_SPOKE_AWS"
echo "SKIP_SPOKE_GCP=$SKIP_SPOKE_GCP"
echo "SKIP_SPOKE_GATEWAY=$SKIP_SPOKE_GATEWAY"
echo "SKIP_SPOKE_GATEWAY_AZURE=$SKIP_SPOKE_GATEWAY_AZURE"
echo "SKIP_SPOKE_GATEWAY_AWS=$SKIP_SPOKE_GATEWAY_AWS"
echo "SKIP_SPOKE_GATEWAY_GCP=$SKIP_SPOKE_GATEWAY_GCP"
echo "SKIP_SPOKE_GATEWAY_OCI=$SKIP_SPOKE_GATEWAY_OCI"
echo "SKIP_SUMOLOGIC_FORWARDER=$SKIP_SUMOLOGIC_FORWARDER"
echo "SKIP_TRANSIT=$SKIP_TRANSIT"
echo "SKIP_TRANSIT_AWS=$SKIP_TRANSIT_AWS"
echo "SKIP_TRANSIT_AZURE=$SKIP_TRANSIT_AZURE"
echo "SKIP_TRANSIT_EXTERNAL_DEVICE_CONN=$SKIP_TRANSIT_EXTERNAL_DEVICE_CONN"
echo "SKIP_TRANSIT_FIRENET_POLICY=$SKIP_TRANSIT_FIRENET_POLICY"
echo "SKIP_TRANSIT_FIRENET_POLICY_AWS=$SKIP_TRANSIT_FIRENET_POLICY_AWS"
echo "SKIP_TRANSIT_FIRENET_POLICY_AZURE=$SKIP_TRANSIT_FIRENET_POLICY_AZURE"
echo "SKIP_TRANSIT_GATEWAY=$SKIP_TRANSIT_GATEWAY"
echo "SKIP_TRANSIT_GATEWAY_AWS=$SKIP_TRANSIT_GATEWAY_AWS"
echo "SKIP_TRANSIT_GATEWAY_AZURE=$SKIP_TRANSIT_GATEWAY_AZURE"
echo "SKIP_TRANSIT_GATEWAY_GCP=$SKIP_TRANSIT_GATEWAY_GCP"
echo "SKIP_TRANSIT_GATEWAY_OCI=$SKIP_TRANSIT_GATEWAY_OCI"
echo "SKIP_TRANSIT_GATEWAY_PEERING=$SKIP_TRANSIT_GATEWAY_PEERING"
echo "SKIP_TRANS_PEER=$SKIP_TRANS_PEER"
echo "SKIP_TUNNEL=$SKIP_TUNNEL"
echo "SKIP_VPC=$SKIP_VPC"
echo "SKIP_VGW_CONN=$SKIP_VGW_CONN"
echo "SKIP_VPN_PROFILE=$SKIP_VPN_PROFILE"
echo "SKIP_VPN_USER=$SKIP_VPN_USER"
echo "SKIP_VPN_USER_ACCELERATOR=$SKIP_VPN_USER_ACCELERATOR"

export ARM_APPLICATION_ID
export ARM_APPLICATION_KEY
export ARM_DIRECTORY_ID
export ARM_SUBSCRIPTION_ID
export AZURE_GW_SIZE
export AZURE_REGION
export AZURE_REGION2
export AZURE_SUBNET
export AZURE_SUBNET2
export AZURE_VNET_ID
export AZURE_VNET_ID2
export AVIATRIX_CONTROLLER_IP
export AVIATRIX_PASSWORD
export AVIATRIX_USERNAME
export AWS_ACCESS_KEY
export AWS_ACCOUNT_NUMBER
export AWS_BGP_VGW_ID
export AWS_REGION
export AWS_REGION2
export AWS_SECRET_KEY
export AWS_SUBNET
export AWS_SUBNET2
export AWS_VPC_ID
export AWS_VPC_ID2
export AWS_DX_GATEWAY_ID
export DOMAIN_NAME
export AWSGOV_ACCESS_KEY
export AWSGOV_ACCOUNT_NUMBER
export AWSGOV_SECRET_KEY
export DEVICE_PUBLIC_IP
export DEVICE_KEY_FILE_PATH
export GCP_CREDENTIALS_FILEPATH
export GCP_ID
export GCP_SUBNET
export GCP_VPC_ID
export GCP_ZONE
export IDP_METADATA
export IDP_METADATA_TYPE
export OCI_API_KEY_FILEPATH
export OCI_COMPARTMENT_ID
export OCI_REGION
export OCI_SUBNET
export OCI_TENANCY_ID
export OCI_USER_ID
export OCI_VPC_ID
export controller_private_ip
export TRANSIT_GATEWAY_NAME
export ARM_RESOURCE_GROUP
export ARM_HUB_NAME
export AWS_TGW_NAME
export DATADOG_API_KEY

export SKIP_DATA_ACCOUNT
export SKIP_DATA_CALLER_IDENTITY
export SKIP_DATA_FIRENET
export SKIP_DATA_FIRENET_VENDOR_INTEGRATION
export SKIP_DATA_FIREWALL
export SKIP_DATA_GATEWAY
export SKIP_DATA_SPOKE_GATEWAY
export SKIP_DATA_SPOKE_GATEWAY_AWS
export SKIP_DATA_SPOKE_GATEWAY_AZURE
export SKIP_DATA_SPOKE_GATEWAY_GCP
export SKIP_DATA_TRANSIT_GATEWAY
export SKIP_DATA_TRANSIT_GATEWAY_AWS
export SKIP_DATA_TRANSIT_GATEWAY_AZURE
export SKIP_DATA_TRANSIT_GATEWAY_GCP
export SKIP_DATA_VPC
export SKIP_DATA_VPC_TRACKER
export SKIP_ACCOUNT
export SKIP_ACCOUNT_AWS
export SKIP_ACCOUNT_AZURE
export SKIP_ACCOUNT_GCP
export SKIP_ACCOUNT_OCI
export SKIP_ACCOUNT_AWSGOV
export SKIP_ACCOUNT_USER
export SKIP_ARM_PEER
export SKIP_AWS_GUARD_DUTY
export SKIP_AWS_PEER
export SKIP_AWS_TGW
export SKIP_AWS_TGW_DIRECTCONNECT
export SKIP_AWS_TGW_PEERING
export SKIP_AWS_TGW_PEERING_DOMAIN_CONN
export SKIP_AWS_TGW_TRANSIT_GATEWAY_ATTACHMENT
export SKIP_AWS_TGW_VPC_ATTACHMENT
export SKIP_AWS_TGW_VPN_CONN
export SKIP_AZURE_PEER
export SKIP_AZURE_SPOKE_NATIVE_PEERING
export SKIP_CONTROLLER_CONFIG
export SKIP_DATADOG_AGENT
export SKIP_DEVICE_AWS_TGW_ATTACHMENT
export SKIP_DEVICE_INTERFACE_CONFIG
export SKIP_DEVICE_REGISTRATION
export SKIP_DEVICE_TAG
export SKIP_DEVICE_TRANSIT_GATEWAY_ATTACHMENT
export SKIP_DEVICE_VIRTUAL_WAN_ATTACHMENT
export SKIP_FIRENET
export SKIP_FIREWALL
export SKIP_FIREWALL_INSTANCE
export SKIP_FIREWALL_INSTANCE_ASSOCIATION
export SKIP_FIREWALL_MANAGEMENT_ACCESS
export SKIP_FIREWALL_MANAGEMENT_ACCESS_AWS
export SKIP_FIREWALL_MANAGEMENT_ACCESS_AZURE
export SKIP_FIREWALL_POLICY
export SKIP_FIREWALL_TAG
export SKIP_FQDN
export SKIP_FQDN_PASS_THROUGH
export SKIP_FQDN_TAG_RULE
export SKIP_GATEWAY
export SKIP_GATEWAY_AWS
export SKIP_GATEWAY_GCP
export SKIP_GATEWAY_AZURE
export SKIP_GATEWAY_OCI
export SKIP_GATEWAY_DNAT
export SKIP_GATEWAY_DNAT_AWS
export SKIP_GATEWAY_DNAT_AZURE
export SKIP_GATEWAY_SNAT
export SKIP_GATEWAY_SNAT_AWS
export SKIP_GATEWAY_SNAT_AZURE
export SKIP_GATEWAY_AWSGOV
export SKIP_GEO_VPN
export SKIP_PERIODIC_PING
export SKIP_RBAC_GROUP
export SKIP_RBAC_GROUP_ACCESS_ACCOUNT_ATTACHMENT
export SKIP_RBAC_GROUP_PERMISSION_ATTACHMENT
export SKIP_RBAC_GROUP_USER_ATTACHMENT
export SKIP_REMOTE_SYSLOG
export SKIP_SAML_ENDPOINT
export SKIP_S2C
export SKIP_SEGMENTATION_SECURITY_DOMAIN
export SKIP_SEGMENTATION_SECURITY_DOMAIN_ASSOCIATION
export SKIP_SEGMENTATION_SECURITY_DOMAIN_CONNECTION_POLICY
export SKIP_SPLUNK_LOGGING
export SKIP_SPOKE
export SKIP_SPOKE_AZURE
export SKIP_SPOKE_AWS
export SKIP_SPOKE_GCP
export SKIP_SPOKE_GATEWAY
export SKIP_SPOKE_GATEWAY_AZURE
export SKIP_SPOKE_GATEWAY_AWS
export SKIP_SPOKE_GATEWAY_GCP
export SKIP_SPOKE_GATEWAY_OCI
export SKIP_SUMOLOGIC_FORWARDER
export SKIP_TRANS_PEER
export SKIP_TRANSIT
export SKIP_TRANSIT_AWS
export SKIP_TRANSIT_AZURE
export SKIP_TRANSIT_EXTERNAL_DEVICE_CONN
export SKIP_TRANSIT_FIRENET_POLICY
export SKIP_TRANSIT_FIRENET_POLICY_AWS
export SKIP_TRANSIT_FIRENET_POLICY_AZURE
export SKIP_TRANSIT_GATEWAY
export SKIP_TRANSIT_GATEWAY_AWS
export SKIP_TRANSIT_GATEWAY_AZURE
export SKIP_TRANSIT_GATEWAY_GCP
export SKIP_TRANSIT_GATEWAY_OCI
export SKIP_TRANSIT_GATEWAY_PEERING
export SKIP_TUNNEL
export SKIP_VPC
export SKIP_VGW_CONN
export SKIP_VPN_PROFILE
export SKIP_VPN_USER
export SKIP_VPN_USER_ACCELERATOR
