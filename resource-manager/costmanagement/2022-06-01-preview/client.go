package v2022_06_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2022-06-01-preview/scheduledactions"
)

type Client struct {
	ScheduledActions *scheduledactions.ScheduledActionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	scheduledActionsClient := scheduledactions.NewScheduledActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduledActionsClient.Client)

	return Client{
		ScheduledActions: &scheduledActionsClient,
	}
}
