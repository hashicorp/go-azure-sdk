package v2024_04_24

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/createresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/monitoredsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/monitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/singlesignon"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/tagrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CreateResource         *createresource.CreateResourceClient
	MonitoredSubscriptions *monitoredsubscriptions.MonitoredSubscriptionsClient
	Monitors               *monitors.MonitorsClient
	SingleSignOn           *singlesignon.SingleSignOnClient
	TagRules               *tagrules.TagRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	createResourceClient, err := createresource.NewCreateResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CreateResource client: %+v", err)
	}
	configureFunc(createResourceClient.Client)

	monitoredSubscriptionsClient, err := monitoredsubscriptions.NewMonitoredSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredSubscriptions client: %+v", err)
	}
	configureFunc(monitoredSubscriptionsClient.Client)

	monitorsClient, err := monitors.NewMonitorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Monitors client: %+v", err)
	}
	configureFunc(monitorsClient.Client)

	singleSignOnClient, err := singlesignon.NewSingleSignOnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SingleSignOn client: %+v", err)
	}
	configureFunc(singleSignOnClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	return &Client{
		CreateResource:         createResourceClient,
		MonitoredSubscriptions: monitoredSubscriptionsClient,
		Monitors:               monitorsClient,
		SingleSignOn:           singleSignOnClient,
		TagRules:               tagRulesClient,
	}, nil
}
