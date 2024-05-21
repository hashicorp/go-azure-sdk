package v2015_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/deploymentoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/policyassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/policydefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/providers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/resourcegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/resources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/tags"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2015-11-01/tenants"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeploymentOperations *deploymentoperations.DeploymentOperationsClient
	Deployments          *deployments.DeploymentsClient
	PolicyAssignments    *policyassignments.PolicyAssignmentsClient
	PolicyDefinitions    *policydefinitions.PolicyDefinitionsClient
	Providers            *providers.ProvidersClient
	ResourceGroups       *resourcegroups.ResourceGroupsClient
	Resources            *resources.ResourcesClient
	Subscriptions        *subscriptions.SubscriptionsClient
	Tags                 *tags.TagsClient
	Tenants              *tenants.TenantsClient
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

	policyAssignmentsClient, err := policyassignments.NewPolicyAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAssignments client: %+v", err)
	}
	configureFunc(policyAssignmentsClient.Client)

	policyDefinitionsClient, err := policydefinitions.NewPolicyDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDefinitions client: %+v", err)
	}
	configureFunc(policyDefinitionsClient.Client)

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

	subscriptionsClient, err := subscriptions.NewSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Subscriptions client: %+v", err)
	}
	configureFunc(subscriptionsClient.Client)

	tagsClient, err := tags.NewTagsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Tags client: %+v", err)
	}
	configureFunc(tagsClient.Client)

	tenantsClient, err := tenants.NewTenantsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Tenants client: %+v", err)
	}
	configureFunc(tenantsClient.Client)

	return &Client{
		DeploymentOperations: deploymentOperationsClient,
		Deployments:          deploymentsClient,
		PolicyAssignments:    policyAssignmentsClient,
		PolicyDefinitions:    policyDefinitionsClient,
		Providers:            providersClient,
		ResourceGroups:       resourceGroupsClient,
		Resources:            resourcesClient,
		Subscriptions:        subscriptionsClient,
		Tags:                 tagsClient,
		Tenants:              tenantsClient,
	}, nil
}
