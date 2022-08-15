package v2021_08_08

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2021-08-08/alertprocessingrules"
)

type Client struct {
	AlertProcessingRules *alertprocessingrules.AlertProcessingRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	alertProcessingRulesClient := alertprocessingrules.NewAlertProcessingRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&alertProcessingRulesClient.Client)

	return Client{
		AlertProcessingRules: &alertProcessingRulesClient,
	}
}
