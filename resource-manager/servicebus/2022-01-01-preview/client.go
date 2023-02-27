package v2022_01_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/migrationconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/namespacesauthorizationrule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/queues"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/queuesauthorizationrule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/topics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2022-01-01-preview/topicsauthorizationrule"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	disasterRecoveryConfigsClient := disasterrecoveryconfigs.NewDisasterRecoveryConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&disasterRecoveryConfigsClient.Client)

	migrationConfigsClient := migrationconfigs.NewMigrationConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&migrationConfigsClient.Client)

	namespacesAuthorizationRuleClient := namespacesauthorizationrule.NewNamespacesAuthorizationRuleClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesAuthorizationRuleClient.Client)

	namespacesClient := namespaces.NewNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesClient.Client)

	namespacesPrivateEndpointConnectionsClient := namespacesprivateendpointconnections.NewNamespacesPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateEndpointConnectionsClient.Client)

	namespacesPrivateLinkResourcesClient := namespacesprivatelinkresources.NewNamespacesPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateLinkResourcesClient.Client)

	queuesAuthorizationRuleClient := queuesauthorizationrule.NewQueuesAuthorizationRuleClientWithBaseURI(endpoint)
	configureAuthFunc(&queuesAuthorizationRuleClient.Client)

	queuesClient := queues.NewQueuesClientWithBaseURI(endpoint)
	configureAuthFunc(&queuesClient.Client)

	rulesClient := rules.NewRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&rulesClient.Client)

	subscriptionsClient := subscriptions.NewSubscriptionsClientWithBaseURI(endpoint)
	configureAuthFunc(&subscriptionsClient.Client)

	topicsAuthorizationRuleClient := topicsauthorizationrule.NewTopicsAuthorizationRuleClientWithBaseURI(endpoint)
	configureAuthFunc(&topicsAuthorizationRuleClient.Client)

	topicsClient := topics.NewTopicsClientWithBaseURI(endpoint)
	configureAuthFunc(&topicsClient.Client)

	return Client{
		DisasterRecoveryConfigs:              &disasterRecoveryConfigsClient,
		MigrationConfigs:                     &migrationConfigsClient,
		Namespaces:                           &namespacesClient,
		NamespacesAuthorizationRule:          &namespacesAuthorizationRuleClient,
		NamespacesPrivateEndpointConnections: &namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:       &namespacesPrivateLinkResourcesClient,
		Queues:                               &queuesClient,
		QueuesAuthorizationRule:              &queuesAuthorizationRuleClient,
		Rules:                                &rulesClient,
		Subscriptions:                        &subscriptionsClient,
		Topics:                               &topicsClient,
		TopicsAuthorizationRule:              &topicsAuthorizationRuleClient,
	}
}
