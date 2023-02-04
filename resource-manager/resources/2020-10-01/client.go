package v2020_10_01

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/deploymentscripts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/providers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/resourcegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-10-01/resources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeploymentScripts *deploymentscripts.DeploymentScriptsClient
	Deployments       *deployments.DeploymentsClient
	Providers         *providers.ProvidersClient
	ResourceGroups    *resourcegroups.ResourceGroupsClient
	Resources         *resources.ResourcesClient
}

func NewClientWithBaseURI(api environments.Api, configureAuthFunc func(c *resourcemanager.Client)) (*Client, error) {
	deploymentScriptsClient, err := deploymentscripts.NewDeploymentScriptsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for DeploymentScripts: %+v", err)
	}
	configureAuthFunc(deploymentScriptsClient.Client)

	deploymentsClient, err := deployments.NewDeploymentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Deployments: %+v", err)
	}
	configureAuthFunc(deploymentsClient.Client)

	providersClient, err := providers.NewProvidersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Providers: %+v", err)
	}
	configureAuthFunc(providersClient.Client)

	resourceGroupsClient, err := resourcegroups.NewResourceGroupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ResourceGroups: %+v", err)
	}
	configureAuthFunc(resourceGroupsClient.Client)

	resourcesClient, err := resources.NewResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Resources: %+v", err)
	}
	configureAuthFunc(resourcesClient.Client)

	return &Client{
		DeploymentScripts: deploymentScriptsClient,
		Deployments:       deploymentsClient,
		Providers:         providersClient,
		ResourceGroups:    resourceGroupsClient,
		Resources:         resourcesClient,
	}, nil
}
