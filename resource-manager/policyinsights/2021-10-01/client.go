package v2021_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2021-10-01/policyinsights"
)

type Client struct {
	PolicyInsights *policyinsights.PolicyInsightsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	policyInsightsClient := policyinsights.NewPolicyInsightsClientWithBaseURI(endpoint)
	configureAuthFunc(&policyInsightsClient.Client)

	return Client{
		PolicyInsights: &policyInsightsClient,
	}
}
