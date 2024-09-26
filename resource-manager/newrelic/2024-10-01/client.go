package v2024_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/connectedresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/linkedresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/monitoredsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/monitors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/organizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/plan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/resubscribe"
	"github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2024-10-01/tagrules"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Accounts               *accounts.AccountsClient
	ConnectedResources     *connectedresources.ConnectedResourcesClient
	LinkedResources        *linkedresources.LinkedResourcesClient
	MonitoredSubscriptions *monitoredsubscriptions.MonitoredSubscriptionsClient
	Monitors               *monitors.MonitorsClient
	Organizations          *organizations.OrganizationsClient
	Plan                   *plan.PlanClient
	Resubscribe            *resubscribe.ResubscribeClient
	TagRules               *tagrules.TagRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountsClient, err := accounts.NewAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

	connectedResourcesClient, err := connectedresources.NewConnectedResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectedResources client: %+v", err)
	}
	configureFunc(connectedResourcesClient.Client)

	linkedResourcesClient, err := linkedresources.NewLinkedResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LinkedResources client: %+v", err)
	}
	configureFunc(linkedResourcesClient.Client)

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

	resubscribeClient, err := resubscribe.NewResubscribeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Resubscribe client: %+v", err)
	}
	configureFunc(resubscribeClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	return &Client{
		Accounts:               accountsClient,
		ConnectedResources:     connectedResourcesClient,
		LinkedResources:        linkedResourcesClient,
		MonitoredSubscriptions: monitoredSubscriptionsClient,
		Monitors:               monitorsClient,
		Organizations:          organizationsClient,
		Plan:                   planClient,
		Resubscribe:            resubscribeClient,
		TagRules:               tagRulesClient,
	}, nil
}
