package v2024_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/migrationconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/namespacesauthorizationrule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/queues"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/queuesauthorizationrule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/topics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2024-01-01/topicsauthorizationrule"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DisasterRecoveryConfigs              *disasterrecoveryconfigs.DisasterRecoveryConfigsClient
	MigrationConfigs                     *migrationconfigs.MigrationConfigsClient
	Namespaces                           *namespaces.NamespacesClient
	NamespacesAuthorizationRule          *namespacesauthorizationrule.NamespacesAuthorizationRuleClient
	NamespacesPrivateEndpointConnections *namespacesprivateendpointconnections.NamespacesPrivateEndpointConnectionsClient
	NamespacesPrivateLinkResources       *namespacesprivatelinkresources.NamespacesPrivateLinkResourcesClient
	Queues                               *queues.QueuesClient
	QueuesAuthorizationRule              *queuesauthorizationrule.QueuesAuthorizationRuleClient
	Rules                                *rules.RulesClient
	Subscriptions                        *subscriptions.SubscriptionsClient
	Topics                               *topics.TopicsClient
	TopicsAuthorizationRule              *topicsauthorizationrule.TopicsAuthorizationRuleClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	disasterRecoveryConfigsClient, err := disasterrecoveryconfigs.NewDisasterRecoveryConfigsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DisasterRecoveryConfigs client: %+v", err)
	}
	configureFunc(disasterRecoveryConfigsClient.Client)

	migrationConfigsClient, err := migrationconfigs.NewMigrationConfigsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrationConfigs client: %+v", err)
	}
	configureFunc(migrationConfigsClient.Client)

	namespacesAuthorizationRuleClient, err := namespacesauthorizationrule.NewNamespacesAuthorizationRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NamespacesAuthorizationRule client: %+v", err)
	}
	configureFunc(namespacesAuthorizationRuleClient.Client)

	namespacesClient, err := namespaces.NewNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Namespaces client: %+v", err)
	}
	configureFunc(namespacesClient.Client)

	namespacesPrivateEndpointConnectionsClient, err := namespacesprivateendpointconnections.NewNamespacesPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NamespacesPrivateEndpointConnections client: %+v", err)
	}
	configureFunc(namespacesPrivateEndpointConnectionsClient.Client)

	namespacesPrivateLinkResourcesClient, err := namespacesprivatelinkresources.NewNamespacesPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NamespacesPrivateLinkResources client: %+v", err)
	}
	configureFunc(namespacesPrivateLinkResourcesClient.Client)

	queuesAuthorizationRuleClient, err := queuesauthorizationrule.NewQueuesAuthorizationRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueuesAuthorizationRule client: %+v", err)
	}
	configureFunc(queuesAuthorizationRuleClient.Client)

	queuesClient, err := queues.NewQueuesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Queues client: %+v", err)
	}
	configureFunc(queuesClient.Client)

	rulesClient, err := rules.NewRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Rules client: %+v", err)
	}
	configureFunc(rulesClient.Client)

	subscriptionsClient, err := subscriptions.NewSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Subscriptions client: %+v", err)
	}
	configureFunc(subscriptionsClient.Client)

	topicsAuthorizationRuleClient, err := topicsauthorizationrule.NewTopicsAuthorizationRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TopicsAuthorizationRule client: %+v", err)
	}
	configureFunc(topicsAuthorizationRuleClient.Client)

	topicsClient, err := topics.NewTopicsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Topics client: %+v", err)
	}
	configureFunc(topicsClient.Client)

	return &Client{
		DisasterRecoveryConfigs:              disasterRecoveryConfigsClient,
		MigrationConfigs:                     migrationConfigsClient,
		Namespaces:                           namespacesClient,
		NamespacesAuthorizationRule:          namespacesAuthorizationRuleClient,
		NamespacesPrivateEndpointConnections: namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:       namespacesPrivateLinkResourcesClient,
		Queues:                               queuesClient,
		QueuesAuthorizationRule:              queuesAuthorizationRuleClient,
		Rules:                                rulesClient,
		Subscriptions:                        subscriptionsClient,
		Topics:                               topicsClient,
		TopicsAuthorizationRule:              topicsAuthorizationRuleClient,
	}, nil
}
