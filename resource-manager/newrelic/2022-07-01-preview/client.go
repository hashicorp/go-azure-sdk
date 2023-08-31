package v2022_07_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/monitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/organizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/plan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2022-07-01-preview/tagrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Accounts      *accounts.AccountsClient
	Monitors      *monitors.MonitorsClient
	Organizations *organizations.OrganizationsClient
	Plan          *plan.PlanClient
	TagRules      *tagrules.TagRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountsClient, err := accounts.NewAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

	monitorsClient, err := monitors.NewMonitorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Monitors client: %+v", err)
	}
	configureFunc(monitorsClient.Client)

	organizationsClient, err := organizations.NewOrganizationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Organizations client: %+v", err)
	}
	configureFunc(organizationsClient.Client)

	planClient, err := plan.NewPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Plan client: %+v", err)
	}
	configureFunc(planClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	return &Client{
		Accounts:      accountsClient,
		Monitors:      monitorsClient,
		Organizations: organizationsClient,
		Plan:          planClient,
		TagRules:      tagRulesClient,
	}, nil
}
