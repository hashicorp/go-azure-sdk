// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_12_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/addons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/authorizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/cloudlinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/datastores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/globalreachconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/hcxenterprisesites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/placementpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/privateclouds"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/scripts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/workloadnetworks"
)

type Client struct {
	Addons                 *addons.AddonsClient
	Authorizations         *authorizations.AuthorizationsClient
	CloudLinks             *cloudlinks.CloudLinksClient
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
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	addonsClient := addons.NewAddonsClientWithBaseURI(endpoint)
	configureAuthFunc(&addonsClient.Client)

	authorizationsClient := authorizations.NewAuthorizationsClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationsClient.Client)

	cloudLinksClient := cloudlinks.NewCloudLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&cloudLinksClient.Client)

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

	return Client{
		Addons:                 &addonsClient,
		Authorizations:         &authorizationsClient,
		CloudLinks:             &cloudLinksClient,
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
	}
}
