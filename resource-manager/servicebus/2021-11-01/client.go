// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_11_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/migrationconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/namespacesauthorizationrule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/queues"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/queuesauthorizationrule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/topics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2021-11-01/topicsauthorizationrule"
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
