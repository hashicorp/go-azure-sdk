package v2026_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2026-06-01/monitoredsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2026-06-01/newrelicmonitorresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2026-06-01/newrelics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2026-06-01/tagrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	MonitoredSubscriptions   *monitoredsubscriptions.MonitoredSubscriptionsClient
	NewRelicMonitorResources *newrelicmonitorresources.NewRelicMonitorResourcesClient
	NewRelics                *newrelics.NewRelicsClient
	TagRules                 *tagrules.TagRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	monitoredSubscriptionsClient, err := monitoredsubscriptions.NewMonitoredSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredSubscriptions client: %+v", err)
	}
	configureFunc(monitoredSubscriptionsClient.Client)

	newRelicMonitorResourcesClient, err := newrelicmonitorresources.NewNewRelicMonitorResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NewRelicMonitorResources client: %+v", err)
	}
	configureFunc(newRelicMonitorResourcesClient.Client)

	newRelicsClient, err := newrelics.NewNewRelicsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NewRelics client: %+v", err)
	}
	configureFunc(newRelicsClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	return &Client{
		MonitoredSubscriptions:   monitoredSubscriptionsClient,
		NewRelicMonitorResources: newRelicMonitorResourcesClient,
		NewRelics:                newRelicsClient,
		TagRules:                 tagRulesClient,
	}, nil
}
