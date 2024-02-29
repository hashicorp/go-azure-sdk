package v2023_03_15_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2023-03-15-preview/scheduledqueryrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ScheduledQueryRules *scheduledqueryrules.ScheduledQueryRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	scheduledQueryRulesClient, err := scheduledqueryrules.NewScheduledQueryRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScheduledQueryRules client: %+v", err)
	}
	configureFunc(scheduledQueryRulesClient.Client)

	return &Client{
		ScheduledQueryRules: scheduledQueryRulesClient,
	}, nil
}
