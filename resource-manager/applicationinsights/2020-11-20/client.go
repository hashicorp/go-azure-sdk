package v2020_11_20

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2020-11-20/applicationinsights"
)

type Client struct {
	ApplicationInsights *applicationinsights.ApplicationInsightsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	applicationInsightsClient := applicationinsights.NewApplicationInsightsClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationInsightsClient.Client)

	return Client{
		ApplicationInsights: &applicationInsightsClient,
	}
}
