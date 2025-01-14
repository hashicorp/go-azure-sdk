package v2023_10_20

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/agreements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/apikey"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/connectedresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/createresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/hosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/linkedresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/monitoredresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/monitoredsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/monitorsresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/refreshsetpasswordlink"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datadog/2023-10-20/singlesignon"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Agreements             *agreements.AgreementsClient
	ApiKey                 *apikey.ApiKeyClient
	ConnectedResources     *connectedresources.ConnectedResourcesClient
	CreateResource         *createresource.CreateResourceClient
	Hosts                  *hosts.HostsClient
	LinkedResources        *linkedresources.LinkedResourcesClient
	MonitoredResources     *monitoredresources.MonitoredResourcesClient
	MonitoredSubscriptions *monitoredsubscriptions.MonitoredSubscriptionsClient
	MonitorsResource       *monitorsresource.MonitorsResourceClient
	RefreshSetPasswordLink *refreshsetpasswordlink.RefreshSetPasswordLinkClient
	Rules                  *rules.RulesClient
	SingleSignOn           *singlesignon.SingleSignOnClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	agreementsClient, err := agreements.NewAgreementsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Agreements client: %+v", err)
	}
	configureFunc(agreementsClient.Client)

	apiKeyClient, err := apikey.NewApiKeyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApiKey client: %+v", err)
	}
	configureFunc(apiKeyClient.Client)

	connectedResourcesClient, err := connectedresources.NewConnectedResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectedResources client: %+v", err)
	}
	configureFunc(connectedResourcesClient.Client)

	createResourceClient, err := createresource.NewCreateResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CreateResource client: %+v", err)
	}
	configureFunc(createResourceClient.Client)

	hostsClient, err := hosts.NewHostsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Hosts client: %+v", err)
	}
	configureFunc(hostsClient.Client)

	linkedResourcesClient, err := linkedresources.NewLinkedResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LinkedResources client: %+v", err)
	}
	configureFunc(linkedResourcesClient.Client)

	monitoredResourcesClient, err := monitoredresources.NewMonitoredResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredResources client: %+v", err)
	}
	configureFunc(monitoredResourcesClient.Client)

	monitoredSubscriptionsClient, err := monitoredsubscriptions.NewMonitoredSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoredSubscriptions client: %+v", err)
	}
	configureFunc(monitoredSubscriptionsClient.Client)

	monitorsResourceClient, err := monitorsresource.NewMonitorsResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitorsResource client: %+v", err)
	}
	configureFunc(monitorsResourceClient.Client)

	refreshSetPasswordLinkClient, err := refreshsetpasswordlink.NewRefreshSetPasswordLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RefreshSetPasswordLink client: %+v", err)
	}
	configureFunc(refreshSetPasswordLinkClient.Client)

	rulesClient, err := rules.NewRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Rules client: %+v", err)
	}
	configureFunc(rulesClient.Client)

	singleSignOnClient, err := singlesignon.NewSingleSignOnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SingleSignOn client: %+v", err)
	}
	configureFunc(singleSignOnClient.Client)

	return &Client{
		Agreements:             agreementsClient,
		ApiKey:                 apiKeyClient,
		ConnectedResources:     connectedResourcesClient,
		CreateResource:         createResourceClient,
		Hosts:                  hostsClient,
		LinkedResources:        linkedResourcesClient,
		MonitoredResources:     monitoredResourcesClient,
		MonitoredSubscriptions: monitoredSubscriptionsClient,
		MonitorsResource:       monitorsResourceClient,
		RefreshSetPasswordLink: refreshSetPasswordLinkClient,
		Rules:                  rulesClient,
		SingleSignOn:           singleSignOnClient,
	}, nil
}
