// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_09_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2021-09-01/monitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2021-09-01/singlesignon"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2021-09-01/tagrules"
)

type Client struct {
	Monitors     *monitors.MonitorsClient
	SingleSignOn *singlesignon.SingleSignOnClient
	TagRules     *tagrules.TagRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	monitorsClient := monitors.NewMonitorsClientWithBaseURI(endpoint)
	configureAuthFunc(&monitorsClient.Client)

	singleSignOnClient := singlesignon.NewSingleSignOnClientWithBaseURI(endpoint)
	configureAuthFunc(&singleSignOnClient.Client)

	tagRulesClient := tagrules.NewTagRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&tagRulesClient.Client)

	return Client{
		Monitors:     &monitorsClient,
		SingleSignOn: &singleSignOnClient,
		TagRules:     &tagRulesClient,
	}
}
