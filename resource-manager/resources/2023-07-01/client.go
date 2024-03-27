package v2023_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-07-01/deploymentoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-07-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-07-01/providers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-07-01/resourcegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-07-01/resources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-07-01/tags"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeploymentOperations *deploymentoperations.DeploymentOperationsClient
	Deployments          *deployments.DeploymentsClient
	Providers            *providers.ProvidersClient
	ResourceGroups       *resourcegroups.ResourceGroupsClient
	Resources            *resources.ResourcesClient
	Tags                 *tags.TagsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deploymentOperationsClient, err := deploymentoperations.NewDeploymentOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentOperations client: %+v", err)
	}
	configureFunc(deploymentOperationsClient.Client)

	deploymentsClient, err := deployments.NewDeploymentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Deployments client: %+v", err)
	}
	configureFunc(deploymentsClient.Client)

	providersClient, err := providers.NewProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Providers client: %+v", err)
	}
	configureFunc(providersClient.Client)

	resourceGroupsClient, err := resourcegroups.NewResourceGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceGroups client: %+v", err)
	}
	configureFunc(resourceGroupsClient.Client)

	resourcesClient, err := resources.NewResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Resources client: %+v", err)
	}
	configureFunc(resourcesClient.Client)

	tagsClient, err := tags.NewTagsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Tags client: %+v", err)
	}
	configureFunc(tagsClient.Client)

	return &Client{
		DeploymentOperations: deploymentOperationsClient,
		Deployments:          deploymentsClient,
		Providers:            providersClient,
		ResourceGroups:       resourceGroupsClient,
		Resources:            resourcesClient,
		Tags:                 tagsClient,
	}, nil
}
