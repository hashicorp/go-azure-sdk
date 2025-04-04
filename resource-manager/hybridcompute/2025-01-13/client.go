package v2025_01_13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/agentversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/extensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/gateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/hybrididentitymetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/licenseprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/licenses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machineextensionssetup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machineextensionsupgrade"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machinenetworkprofile"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machineruncommands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/networkconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/networksecurityperimeterconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/privatelinkscopes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/settings"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AgentVersions                         *agentversions.AgentVersionsClient
	Extensions                            *extensions.ExtensionsClient
	Gateways                              *gateways.GatewaysClient
	HybridIdentityMetadata                *hybrididentitymetadata.HybridIdentityMetadataClient
	LicenseProfiles                       *licenseprofiles.LicenseProfilesClient
	Licenses                              *licenses.LicensesClient
	MachineExtensions                     *machineextensions.MachineExtensionsClient
	MachineExtensionsSetup                *machineextensionssetup.MachineExtensionsSetupClient
	MachineExtensionsUpgrade              *machineextensionsupgrade.MachineExtensionsUpgradeClient
	MachineNetworkProfile                 *machinenetworkprofile.MachineNetworkProfileClient
	MachineRunCommands                    *machineruncommands.MachineRunCommandsClient
	Machines                              *machines.MachinesClient
	NetworkConfigurations                 *networkconfigurations.NetworkConfigurationsClient
	NetworkSecurityPerimeterConfiguration *networksecurityperimeterconfiguration.NetworkSecurityPerimeterConfigurationClient
	PrivateEndpointConnections            *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                  *privatelinkresources.PrivateLinkResourcesClient
	PrivateLinkScopes                     *privatelinkscopes.PrivateLinkScopesClient
	Settings                              *settings.SettingsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	agentVersionsClient, err := agentversions.NewAgentVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AgentVersions client: %+v", err)
	}
	configureFunc(agentVersionsClient.Client)

	extensionsClient, err := extensions.NewExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Extensions client: %+v", err)
	}
	configureFunc(extensionsClient.Client)

	gatewaysClient, err := gateways.NewGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Gateways client: %+v", err)
	}
	configureFunc(gatewaysClient.Client)

	hybridIdentityMetadataClient, err := hybrididentitymetadata.NewHybridIdentityMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridIdentityMetadata client: %+v", err)
	}
	configureFunc(hybridIdentityMetadataClient.Client)

	licenseProfilesClient, err := licenseprofiles.NewLicenseProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LicenseProfiles client: %+v", err)
	}
	configureFunc(licenseProfilesClient.Client)

	licensesClient, err := licenses.NewLicensesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Licenses client: %+v", err)
	}
	configureFunc(licensesClient.Client)

	machineExtensionsClient, err := machineextensions.NewMachineExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineExtensions client: %+v", err)
	}
	configureFunc(machineExtensionsClient.Client)

	machineExtensionsSetupClient, err := machineextensionssetup.NewMachineExtensionsSetupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineExtensionsSetup client: %+v", err)
	}
	configureFunc(machineExtensionsSetupClient.Client)

	machineExtensionsUpgradeClient, err := machineextensionsupgrade.NewMachineExtensionsUpgradeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineExtensionsUpgrade client: %+v", err)
	}
	configureFunc(machineExtensionsUpgradeClient.Client)

	machineNetworkProfileClient, err := machinenetworkprofile.NewMachineNetworkProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineNetworkProfile client: %+v", err)
	}
	configureFunc(machineNetworkProfileClient.Client)

	machineRunCommandsClient, err := machineruncommands.NewMachineRunCommandsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineRunCommands client: %+v", err)
	}
	configureFunc(machineRunCommandsClient.Client)

	machinesClient, err := machines.NewMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Machines client: %+v", err)
	}
	configureFunc(machinesClient.Client)

	networkConfigurationsClient, err := networkconfigurations.NewNetworkConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkConfigurations client: %+v", err)
	}
	configureFunc(networkConfigurationsClient.Client)

	networkSecurityPerimeterConfigurationClient, err := networksecurityperimeterconfiguration.NewNetworkSecurityPerimeterConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterConfiguration client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterConfigurationClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	privateLinkScopesClient, err := privatelinkscopes.NewPrivateLinkScopesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkScopes client: %+v", err)
	}
	configureFunc(privateLinkScopesClient.Client)

	settingsClient, err := settings.NewSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Settings client: %+v", err)
	}
	configureFunc(settingsClient.Client)

	return &Client{
		AgentVersions:                         agentVersionsClient,
		Extensions:                            extensionsClient,
		Gateways:                              gatewaysClient,
		HybridIdentityMetadata:                hybridIdentityMetadataClient,
		LicenseProfiles:                       licenseProfilesClient,
		Licenses:                              licensesClient,
		MachineExtensions:                     machineExtensionsClient,
		MachineExtensionsSetup:                machineExtensionsSetupClient,
		MachineExtensionsUpgrade:              machineExtensionsUpgradeClient,
		MachineNetworkProfile:                 machineNetworkProfileClient,
		MachineRunCommands:                    machineRunCommandsClient,
		Machines:                              machinesClient,
		NetworkConfigurations:                 networkConfigurationsClient,
		NetworkSecurityPerimeterConfiguration: networkSecurityPerimeterConfigurationClient,
		PrivateEndpointConnections:            privateEndpointConnectionsClient,
		PrivateLinkResources:                  privateLinkResourcesClient,
		PrivateLinkScopes:                     privateLinkScopesClient,
		Settings:                              settingsClient,
	}, nil
}
