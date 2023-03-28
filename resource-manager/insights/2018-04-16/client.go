package v2018_04_16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2018-04-16/scheduledqueryrules"
)

type Client struct {
	ScheduledQueryRules *scheduledqueryrules.ScheduledQueryRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	scheduledQueryRulesClient := scheduledqueryrules.NewScheduledQueryRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduledQueryRulesClient.Client)

	return Client{
		ScheduledQueryRules: &scheduledQueryRulesClient,
	}
}
