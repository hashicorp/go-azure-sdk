package v2024_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/addons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/authorizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/cloudlinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/datastores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/globalreachconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/hcxenterprisesites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/hosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/iscsipaths"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/placementpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/privateclouds"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/provisionednetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/purestoragepolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/scriptcmdlets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/scriptexecutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/scriptpackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/vmwares"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/workloadnetworkgateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/workloadnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/workloadnetworksegments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/workloadnetworkvirtualmachines"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Addons                         *addons.AddonsClient
	Authorizations                 *authorizations.AuthorizationsClient
	CloudLinks                     *cloudlinks.CloudLinksClient
	Clusters                       *clusters.ClustersClient
	DataStores                     *datastores.DataStoresClient
	GlobalReachConnections         *globalreachconnections.GlobalReachConnectionsClient
	HcxEnterpriseSites             *hcxenterprisesites.HcxEnterpriseSitesClient
	Hosts                          *hosts.HostsClient
	IscsiPaths                     *iscsipaths.IscsiPathsClient
	Locations                      *locations.LocationsClient
	PlacementPolicies              *placementpolicies.PlacementPoliciesClient
	PrivateClouds                  *privateclouds.PrivateCloudsClient
	ProvisionedNetworks            *provisionednetworks.ProvisionedNetworksClient
	PureStoragePolicies            *purestoragepolicies.PureStoragePoliciesClient
	ScriptCmdlets                  *scriptcmdlets.ScriptCmdletsClient
	ScriptExecutions               *scriptexecutions.ScriptExecutionsClient
	ScriptPackages                 *scriptpackages.ScriptPackagesClient
	VMwares                        *vmwares.VMwaresClient
	VirtualMachines                *virtualmachines.VirtualMachinesClient
	WorkloadNetworkGateways        *workloadnetworkgateways.WorkloadNetworkGatewaysClient
	WorkloadNetworkSegments        *workloadnetworksegments.WorkloadNetworkSegmentsClient
	WorkloadNetworkVirtualMachines *workloadnetworkvirtualmachines.WorkloadNetworkVirtualMachinesClient
	WorkloadNetworks               *workloadnetworks.WorkloadNetworksClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	addonsClient, err := addons.NewAddonsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Addons client: %+v", err)
	}
	configureFunc(addonsClient.Client)

	authorizationsClient, err := authorizations.NewAuthorizationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Authorizations client: %+v", err)
	}
	configureFunc(authorizationsClient.Client)

	cloudLinksClient, err := cloudlinks.NewCloudLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudLinks client: %+v", err)
	}
	configureFunc(cloudLinksClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	dataStoresClient, err := datastores.NewDataStoresClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataStores client: %+v", err)
	}
	configureFunc(dataStoresClient.Client)

	globalReachConnectionsClient, err := globalreachconnections.NewGlobalReachConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GlobalReachConnections client: %+v", err)
	}
	configureFunc(globalReachConnectionsClient.Client)

	hcxEnterpriseSitesClient, err := hcxenterprisesites.NewHcxEnterpriseSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HcxEnterpriseSites client: %+v", err)
	}
	configureFunc(hcxEnterpriseSitesClient.Client)

	hostsClient, err := hosts.NewHostsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Hosts client: %+v", err)
	}
	configureFunc(hostsClient.Client)

	iscsiPathsClient, err := iscsipaths.NewIscsiPathsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IscsiPaths client: %+v", err)
	}
	configureFunc(iscsiPathsClient.Client)

	locationsClient, err := locations.NewLocationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Locations client: %+v", err)
	}
	configureFunc(locationsClient.Client)

	placementPoliciesClient, err := placementpolicies.NewPlacementPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlacementPolicies client: %+v", err)
	}
	configureFunc(placementPoliciesClient.Client)

	privateCloudsClient, err := privateclouds.NewPrivateCloudsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateClouds client: %+v", err)
	}
	configureFunc(privateCloudsClient.Client)

	provisionedNetworksClient, err := provisionednetworks.NewProvisionedNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProvisionedNetworks client: %+v", err)
	}
	configureFunc(provisionedNetworksClient.Client)

	pureStoragePoliciesClient, err := purestoragepolicies.NewPureStoragePoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PureStoragePolicies client: %+v", err)
	}
	configureFunc(pureStoragePoliciesClient.Client)

	scriptCmdletsClient, err := scriptcmdlets.NewScriptCmdletsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScriptCmdlets client: %+v", err)
	}
	configureFunc(scriptCmdletsClient.Client)

	scriptExecutionsClient, err := scriptexecutions.NewScriptExecutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScriptExecutions client: %+v", err)
	}
	configureFunc(scriptExecutionsClient.Client)

	scriptPackagesClient, err := scriptpackages.NewScriptPackagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScriptPackages client: %+v", err)
	}
	configureFunc(scriptPackagesClient.Client)

	vMwaresClient, err := vmwares.NewVMwaresClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VMwares client: %+v", err)
	}
	configureFunc(vMwaresClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	workloadNetworkGatewaysClient, err := workloadnetworkgateways.NewWorkloadNetworkGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadNetworkGateways client: %+v", err)
	}
	configureFunc(workloadNetworkGatewaysClient.Client)

	workloadNetworkSegmentsClient, err := workloadnetworksegments.NewWorkloadNetworkSegmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadNetworkSegments client: %+v", err)
	}
	configureFunc(workloadNetworkSegmentsClient.Client)

	workloadNetworkVirtualMachinesClient, err := workloadnetworkvirtualmachines.NewWorkloadNetworkVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadNetworkVirtualMachines client: %+v", err)
	}
	configureFunc(workloadNetworkVirtualMachinesClient.Client)

	workloadNetworksClient, err := workloadnetworks.NewWorkloadNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadNetworks client: %+v", err)
	}
	configureFunc(workloadNetworksClient.Client)

	return &Client{
		Addons:                         addonsClient,
		Authorizations:                 authorizationsClient,
		CloudLinks:                     cloudLinksClient,
		Clusters:                       clustersClient,
		DataStores:                     dataStoresClient,
		GlobalReachConnections:         globalReachConnectionsClient,
		HcxEnterpriseSites:             hcxEnterpriseSitesClient,
		Hosts:                          hostsClient,
		IscsiPaths:                     iscsiPathsClient,
		Locations:                      locationsClient,
		PlacementPolicies:              placementPoliciesClient,
		PrivateClouds:                  privateCloudsClient,
		ProvisionedNetworks:            provisionedNetworksClient,
		PureStoragePolicies:            pureStoragePoliciesClient,
		ScriptCmdlets:                  scriptCmdletsClient,
		ScriptExecutions:               scriptExecutionsClient,
		ScriptPackages:                 scriptPackagesClient,
		VMwares:                        vMwaresClient,
		VirtualMachines:                virtualMachinesClient,
		WorkloadNetworkGateways:        workloadNetworkGatewaysClient,
		WorkloadNetworkSegments:        workloadNetworkSegmentsClient,
		WorkloadNetworkVirtualMachines: workloadNetworkVirtualMachinesClient,
		WorkloadNetworks:               workloadNetworksClient,
	}, nil
}
