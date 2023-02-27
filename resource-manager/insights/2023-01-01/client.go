// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2023_01_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2023-01-01/actiongroupsapis"
)

type Client struct {
	ActionGroupsAPIs *actiongroupsapis.ActionGroupsAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	actionGroupsAPIsClient := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&actionGroupsAPIsClient.Client)

	return Client{
		ActionGroupsAPIs: &actionGroupsAPIsClient,
	}
}
