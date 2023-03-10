package v2022_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-09-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-09-01/providers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-09-01/resourcegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-09-01/resources"
)

type Client struct {
	Deployments    *deployments.DeploymentsClient
	Providers      *providers.ProvidersClient
	ResourceGroups *resourcegroups.ResourceGroupsClient
	Resources      *resources.ResourcesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	deploymentsClient := deployments.NewDeploymentsClientWithBaseURI(endpoint)
	configureAuthFunc(&deploymentsClient.Client)

	providersClient := providers.NewProvidersClientWithBaseURI(endpoint)
	configureAuthFunc(&providersClient.Client)

	resourceGroupsClient := resourcegroups.NewResourceGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceGroupsClient.Client)

	resourcesClient := resources.NewResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&resourcesClient.Client)

	return Client{
		Deployments:    &deploymentsClient,
		Providers:      &providersClient,
		ResourceGroups: &resourceGroupsClient,
		Resources:      &resourcesClient,
	}
}
