package v2020_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2020-10-01/activitylogalertsapis"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ActivityLogAlertsAPIs *activitylogalertsapis.ActivityLogAlertsAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	activityLogAlertsAPIsClient, err := activitylogalertsapis.NewActivityLogAlertsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActivityLogAlertsAPIs client: %+v", err)
	}
	configureFunc(activityLogAlertsAPIsClient.Client)

	return &Client{
		ActivityLogAlertsAPIs: activityLogAlertsAPIsClient,
	}, nil
}
