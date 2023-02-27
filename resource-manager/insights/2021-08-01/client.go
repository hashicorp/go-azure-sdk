// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_08_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-08-01/scheduledqueryrules"
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
