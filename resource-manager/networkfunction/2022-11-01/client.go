package v2022_11_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkfunction/2022-11-01/azuretrafficcollectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkfunction/2022-11-01/collectorpolicies"
)

type Client struct {
	AzureTrafficCollectors *azuretrafficcollectors.AzureTrafficCollectorsClient
	CollectorPolicies      *collectorpolicies.CollectorPoliciesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	azureTrafficCollectorsClient := azuretrafficcollectors.NewAzureTrafficCollectorsClientWithBaseURI(endpoint)
	configureAuthFunc(&azureTrafficCollectorsClient.Client)

	collectorPoliciesClient := collectorpolicies.NewCollectorPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&collectorPoliciesClient.Client)

	return Client{
		AzureTrafficCollectors: &azureTrafficCollectorsClient,
		CollectorPolicies:      &collectorPoliciesClient,
	}
}
