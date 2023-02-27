// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/deploymentscripts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/providers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/resourcegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/resources"
)

type Client struct {
	DeploymentScripts *deploymentscripts.DeploymentScriptsClient
	Deployments       *deployments.DeploymentsClient
	Providers         *providers.ProvidersClient
	ResourceGroups    *resourcegroups.ResourceGroupsClient
	Resources         *resources.ResourcesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	deploymentScriptsClient := deploymentscripts.NewDeploymentScriptsClientWithBaseURI(endpoint)
	configureAuthFunc(&deploymentScriptsClient.Client)

	deploymentsClient := deployments.NewDeploymentsClientWithBaseURI(endpoint)
	configureAuthFunc(&deploymentsClient.Client)

	providersClient := providers.NewProvidersClientWithBaseURI(endpoint)
	configureAuthFunc(&providersClient.Client)

	resourceGroupsClient := resourcegroups.NewResourceGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceGroupsClient.Client)

	resourcesClient := resources.NewResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&resourcesClient.Client)

	return Client{
		DeploymentScripts: &deploymentScriptsClient,
		Deployments:       &deploymentsClient,
		Providers:         &providersClient,
		ResourceGroups:    &resourceGroupsClient,
		Resources:         &resourcesClient,
	}
}
