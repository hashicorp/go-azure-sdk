package v2021_08_08

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2021-08-08/alertsmanagement"
)

type Client struct {
	AlertsManagement *alertsmanagement.AlertsManagementClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	alertsManagementClient := alertsmanagement.NewAlertsManagementClientWithBaseURI(endpoint)
	configureAuthFunc(&alertsManagementClient.Client)

	return Client{
		AlertsManagement: &alertsManagementClient,
	}
}
