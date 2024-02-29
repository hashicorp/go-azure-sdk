package v2021_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/actiongroupsapis"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ActionGroupsAPIs *actiongroupsapis.ActionGroupsAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	actionGroupsAPIsClient, err := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActionGroupsAPIs client: %+v", err)
	}
	configureFunc(actionGroupsAPIsClient.Client)

	return &Client{
		ActionGroupsAPIs: actionGroupsAPIsClient,
	}, nil
}
