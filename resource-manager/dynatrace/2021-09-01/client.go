package v2021_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2021-09-01/monitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2021-09-01/singlesignon"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2021-09-01/tagrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Monitors     *monitors.MonitorsClient
	SingleSignOn *singlesignon.SingleSignOnClient
	TagRules     *tagrules.TagRulesClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	monitorsClient, err := monitors.NewMonitorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Monitors client: %+v", err)
	}
	configureFunc(monitorsClient.Client)

	singleSignOnClient, err := singlesignon.NewSingleSignOnClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SingleSignOn client: %+v", err)
	}
	configureFunc(singleSignOnClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	return &Client{
		Monitors:     monitorsClient,
		SingleSignOn: singleSignOnClient,
		TagRules:     tagRulesClient,
	}, nil
}
