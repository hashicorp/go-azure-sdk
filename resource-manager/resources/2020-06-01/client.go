package v2020_06_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/deployments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/providers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/resourcegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/resources"
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
