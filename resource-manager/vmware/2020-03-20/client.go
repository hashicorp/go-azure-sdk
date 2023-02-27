package v2020_03_20

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2020-03-20/authorizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2020-03-20/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2020-03-20/hcxenterprisesites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2020-03-20/locations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2020-03-20/privateclouds"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Authorizations     *authorizations.AuthorizationsClient
	Clusters           *clusters.ClustersClient
	HcxEnterpriseSites *hcxenterprisesites.HcxEnterpriseSitesClient
	Locations          *locations.LocationsClient
	PrivateClouds      *privateclouds.PrivateCloudsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	authorizationsClient := authorizations.NewAuthorizationsClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationsClient.Client)

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	hcxEnterpriseSitesClient := hcxenterprisesites.NewHcxEnterpriseSitesClientWithBaseURI(endpoint)
	configureAuthFunc(&hcxEnterpriseSitesClient.Client)

	locationsClient := locations.NewLocationsClientWithBaseURI(endpoint)
	configureAuthFunc(&locationsClient.Client)

	privateCloudsClient := privateclouds.NewPrivateCloudsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateCloudsClient.Client)

	return Client{
		Authorizations:     &authorizationsClient,
		Clusters:           &clustersClient,
		HcxEnterpriseSites: &hcxEnterpriseSitesClient,
		Locations:          &locationsClient,
		PrivateClouds:      &privateCloudsClient,
	}
}
