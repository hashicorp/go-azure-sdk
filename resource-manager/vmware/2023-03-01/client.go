package v2023_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/addons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/authorizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/cloudlinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/cluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/datastores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/globalreachconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/hcxenterprisesites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/placementpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/privateclouds"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/scripts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/workloadnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-03-01/zone"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Addons                 *addons.AddonsClient
	Authorizations         *authorizations.AuthorizationsClient
	CloudLinks             *cloudlinks.CloudLinksClient
	Cluster                *cluster.ClusterClient
	Clusters               *clusters.ClustersClient
	DataStores             *datastores.DataStoresClient
	GlobalReachConnections *globalreachconnections.GlobalReachConnectionsClient
	HcxEnterpriseSites     *hcxenterprisesites.HcxEnterpriseSitesClient
	Locations              *locations.LocationsClient
	PlacementPolicies      *placementpolicies.PlacementPoliciesClient
	PrivateClouds          *privateclouds.PrivateCloudsClient
	Scripts                *scripts.ScriptsClient
	VirtualMachines        *virtualmachines.VirtualMachinesClient
	WorkloadNetworks       *workloadnetworks.WorkloadNetworksClient
	Zone                   *zone.ZoneClient
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

	clusterClient, err := cluster.NewClusterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Cluster client: %+v", err)
	}
	configureFunc(clusterClient.Client)

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

	scriptsClient, err := scripts.NewScriptsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Scripts client: %+v", err)
	}
	configureFunc(scriptsClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	workloadNetworksClient, err := workloadnetworks.NewWorkloadNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadNetworks client: %+v", err)
	}
	configureFunc(workloadNetworksClient.Client)

	zoneClient, err := zone.NewZoneClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Zone client: %+v", err)
	}
	configureFunc(zoneClient.Client)

	return &Client{
		Addons:                 addonsClient,
		Authorizations:         authorizationsClient,
		CloudLinks:             cloudLinksClient,
		Cluster:                clusterClient,
		Clusters:               clustersClient,
		DataStores:             dataStoresClient,
		GlobalReachConnections: globalReachConnectionsClient,
		HcxEnterpriseSites:     hcxEnterpriseSitesClient,
		Locations:              locationsClient,
		PlacementPolicies:      placementPoliciesClient,
		PrivateClouds:          privateCloudsClient,
		Scripts:                scriptsClient,
		VirtualMachines:        virtualMachinesClient,
		WorkloadNetworks:       workloadNetworksClient,
		Zone:                   zoneClient,
	}, nil
}
