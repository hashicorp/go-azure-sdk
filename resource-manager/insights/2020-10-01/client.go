package v2020_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2020-10-01/activitylogalertsapis"
)

type Client struct {
	ActivityLogAlertsAPIs *activitylogalertsapis.ActivityLogAlertsAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	activityLogAlertsAPIsClient := activitylogalertsapis.NewActivityLogAlertsAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&activityLogAlertsAPIsClient.Client)

	return Client{
		ActivityLogAlertsAPIs: &activityLogAlertsAPIsClient,
	}
}
