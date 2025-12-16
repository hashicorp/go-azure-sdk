package v2024_04_24

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/dynatraces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/dynatracesinglesignonresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/monitoredsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/monitorresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/tagrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DynatraceSingleSignOnResources *dynatracesinglesignonresources.DynatraceSingleSignOnResourcesClient
	Dynatraces                     *dynatraces.DynatracesClient
	MonitorResources               *monitorresources.MonitorResourcesClient
	MonitoredSubscriptions         *monitoredsubscriptions.MonitoredSubscriptionsClient
	TagRules                       *tagrules.TagRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dynatraceSingleSignOnResourcesClient, err := dynatracesinglesignonresources.NewDynatraceSingleSignOnResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DynatraceSingleSignOnResources client: %+v", err)
	}
	configureFunc(dynatraceSingleSignOnResourcesClient.Client)

	dynatracesClient, err := dynatraces.NewDynatracesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Dynatraces client: %+v", err)
	}
	configureFunc(dynatracesClient.Client)

	monitorResourcesClient, err := monitorresources.NewMonitorResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitorResources client: %+v", err)
	}
	configureFunc(monitorResourcesClient.Client)

	monitoredSubscriptionsClient, err := monitoredsubscriptions.NewMonitoredSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredSubscriptions client: %+v", err)
	}
	configureFunc(monitoredSubscriptionsClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	return &Client{
		DynatraceSingleSignOnResources: dynatraceSingleSignOnResourcesClient,
		Dynatraces:                     dynatracesClient,
		MonitorResources:               monitorResourcesClient,
		MonitoredSubscriptions:         monitoredSubscriptionsClient,
		TagRules:                       tagRulesClient,
	}, nil
}
