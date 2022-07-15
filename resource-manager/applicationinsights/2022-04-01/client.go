package v2022_04_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2022-04-01/applicationinsights"
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
