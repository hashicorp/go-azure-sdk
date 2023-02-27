// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/addons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/authorizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/cloudlinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/cluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/datastores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/globalreachconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/hcxenterprisesites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/placementpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/privateclouds"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/scripts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/workloadnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/zone"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	addonsClient := addons.NewAddonsClientWithBaseURI(endpoint)
	configureAuthFunc(&addonsClient.Client)

	authorizationsClient := authorizations.NewAuthorizationsClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationsClient.Client)

	cloudLinksClient := cloudlinks.NewCloudLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&cloudLinksClient.Client)

	clusterClient := cluster.NewClusterClientWithBaseURI(endpoint)
	configureAuthFunc(&clusterClient.Client)

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	dataStoresClient := datastores.NewDataStoresClientWithBaseURI(endpoint)
	configureAuthFunc(&dataStoresClient.Client)

	globalReachConnectionsClient := globalreachconnections.NewGlobalReachConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&globalReachConnectionsClient.Client)

	hcxEnterpriseSitesClient := hcxenterprisesites.NewHcxEnterpriseSitesClientWithBaseURI(endpoint)
	configureAuthFunc(&hcxEnterpriseSitesClient.Client)

	locationsClient := locations.NewLocationsClientWithBaseURI(endpoint)
	configureAuthFunc(&locationsClient.Client)

	placementPoliciesClient := placementpolicies.NewPlacementPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&placementPoliciesClient.Client)

	privateCloudsClient := privateclouds.NewPrivateCloudsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateCloudsClient.Client)

	scriptsClient := scripts.NewScriptsClientWithBaseURI(endpoint)
	configureAuthFunc(&scriptsClient.Client)

	virtualMachinesClient := virtualmachines.NewVirtualMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachinesClient.Client)

	workloadNetworksClient := workloadnetworks.NewWorkloadNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&workloadNetworksClient.Client)

	zoneClient := zone.NewZoneClientWithBaseURI(endpoint)
	configureAuthFunc(&zoneClient.Client)

	return Client{
		Addons:                 &addonsClient,
		Authorizations:         &authorizationsClient,
		CloudLinks:             &cloudLinksClient,
		Cluster:                &clusterClient,
		Clusters:               &clustersClient,
		DataStores:             &dataStoresClient,
		GlobalReachConnections: &globalReachConnectionsClient,
		HcxEnterpriseSites:     &hcxEnterpriseSitesClient,
		Locations:              &locationsClient,
		PlacementPolicies:      &placementPoliciesClient,
		PrivateClouds:          &privateCloudsClient,
		Scripts:                &scriptsClient,
		VirtualMachines:        &virtualMachinesClient,
		WorkloadNetworks:       &workloadNetworksClient,
		Zone:                   &zoneClient,
	}
}
