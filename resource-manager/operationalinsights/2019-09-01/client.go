package v2019_09_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2019-09-01/operationalinsights"
)

type Client struct {
	OperationalInsights *operationalinsights.OperationalInsightsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	operationalInsightsClient := operationalinsights.NewOperationalInsightsClientWithBaseURI(endpoint)
	configureAuthFunc(&operationalInsightsClient.Client)

	return Client{
		OperationalInsights: &operationalInsightsClient,
	}
}
