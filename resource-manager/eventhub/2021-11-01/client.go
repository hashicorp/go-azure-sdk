// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_11_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/authorizationrulesdisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/authorizationruleseventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/authorizationrulesnamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/checknameavailabilitydisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/checknameavailabilitynamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/consumergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubsclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubsclustersavailableclusterregions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubsclustersconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubsclustersnamespace"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/networkrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/schemaregistry"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	AuthorizationRulesDisasterRecoveryConfigs    *authorizationrulesdisasterrecoveryconfigs.AuthorizationRulesDisasterRecoveryConfigsClient
	AuthorizationRulesEventHubs                  *authorizationruleseventhubs.AuthorizationRulesEventHubsClient
	AuthorizationRulesNamespaces                 *authorizationrulesnamespaces.AuthorizationRulesNamespacesClient
	CheckNameAvailabilityDisasterRecoveryConfigs *checknameavailabilitydisasterrecoveryconfigs.CheckNameAvailabilityDisasterRecoveryConfigsClient
	CheckNameAvailabilityNamespaces              *checknameavailabilitynamespaces.CheckNameAvailabilityNamespacesClient
	ConsumerGroups                               *consumergroups.ConsumerGroupsClient
	DisasterRecoveryConfigs                      *disasterrecoveryconfigs.DisasterRecoveryConfigsClient
	EventHubs                                    *eventhubs.EventHubsClient
	EventHubsClusters                            *eventhubsclusters.EventHubsClustersClient
	EventHubsClustersAvailableClusterRegions     *eventhubsclustersavailableclusterregions.EventHubsClustersAvailableClusterRegionsClient
	EventHubsClustersConfiguration               *eventhubsclustersconfiguration.EventHubsClustersConfigurationClient
	EventHubsClustersNamespace                   *eventhubsclustersnamespace.EventHubsClustersNamespaceClient
	Namespaces                                   *namespaces.NamespacesClient
	NamespacesPrivateEndpointConnections         *namespacesprivateendpointconnections.NamespacesPrivateEndpointConnectionsClient
	NamespacesPrivateLinkResources               *namespacesprivatelinkresources.NamespacesPrivateLinkResourcesClient
	NetworkRuleSets                              *networkrulesets.NetworkRuleSetsClient
	SchemaRegistry                               *schemaregistry.SchemaRegistryClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	authorizationRulesDisasterRecoveryConfigsClient := authorizationrulesdisasterrecoveryconfigs.NewAuthorizationRulesDisasterRecoveryConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationRulesDisasterRecoveryConfigsClient.Client)

	authorizationRulesEventHubsClient := authorizationruleseventhubs.NewAuthorizationRulesEventHubsClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationRulesEventHubsClient.Client)

	authorizationRulesNamespacesClient := authorizationrulesnamespaces.NewAuthorizationRulesNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&authorizationRulesNamespacesClient.Client)

	checkNameAvailabilityDisasterRecoveryConfigsClient := checknameavailabilitydisasterrecoveryconfigs.NewCheckNameAvailabilityDisasterRecoveryConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&checkNameAvailabilityDisasterRecoveryConfigsClient.Client)

	checkNameAvailabilityNamespacesClient := checknameavailabilitynamespaces.NewCheckNameAvailabilityNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&checkNameAvailabilityNamespacesClient.Client)

	consumerGroupsClient := consumergroups.NewConsumerGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&consumerGroupsClient.Client)

	disasterRecoveryConfigsClient := disasterrecoveryconfigs.NewDisasterRecoveryConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&disasterRecoveryConfigsClient.Client)

	eventHubsClient := eventhubs.NewEventHubsClientWithBaseURI(endpoint)
	configureAuthFunc(&eventHubsClient.Client)

	eventHubsClustersAvailableClusterRegionsClient := eventhubsclustersavailableclusterregions.NewEventHubsClustersAvailableClusterRegionsClientWithBaseURI(endpoint)
	configureAuthFunc(&eventHubsClustersAvailableClusterRegionsClient.Client)

	eventHubsClustersClient := eventhubsclusters.NewEventHubsClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&eventHubsClustersClient.Client)

	eventHubsClustersConfigurationClient := eventhubsclustersconfiguration.NewEventHubsClustersConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&eventHubsClustersConfigurationClient.Client)

	eventHubsClustersNamespaceClient := eventhubsclustersnamespace.NewEventHubsClustersNamespaceClientWithBaseURI(endpoint)
	configureAuthFunc(&eventHubsClustersNamespaceClient.Client)

	namespacesClient := namespaces.NewNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesClient.Client)

	namespacesPrivateEndpointConnectionsClient := namespacesprivateendpointconnections.NewNamespacesPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateEndpointConnectionsClient.Client)

	namespacesPrivateLinkResourcesClient := namespacesprivatelinkresources.NewNamespacesPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateLinkResourcesClient.Client)

	networkRuleSetsClient := networkrulesets.NewNetworkRuleSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&networkRuleSetsClient.Client)

	schemaRegistryClient := schemaregistry.NewSchemaRegistryClientWithBaseURI(endpoint)
	configureAuthFunc(&schemaRegistryClient.Client)

	return Client{
		AuthorizationRulesDisasterRecoveryConfigs:    &authorizationRulesDisasterRecoveryConfigsClient,
		AuthorizationRulesEventHubs:                  &authorizationRulesEventHubsClient,
		AuthorizationRulesNamespaces:                 &authorizationRulesNamespacesClient,
		CheckNameAvailabilityDisasterRecoveryConfigs: &checkNameAvailabilityDisasterRecoveryConfigsClient,
		CheckNameAvailabilityNamespaces:              &checkNameAvailabilityNamespacesClient,
		ConsumerGroups:                               &consumerGroupsClient,
		DisasterRecoveryConfigs:                      &disasterRecoveryConfigsClient,
		EventHubs:                                    &eventHubsClient,
		EventHubsClusters:                            &eventHubsClustersClient,
		EventHubsClustersAvailableClusterRegions:     &eventHubsClustersAvailableClusterRegionsClient,
		EventHubsClustersConfiguration:               &eventHubsClustersConfigurationClient,
		EventHubsClustersNamespace:                   &eventHubsClustersNamespaceClient,
		Namespaces:                                   &namespacesClient,
		NamespacesPrivateEndpointConnections:         &namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:               &namespacesPrivateLinkResourcesClient,
		NetworkRuleSets:                              &networkRuleSetsClient,
		SchemaRegistry:                               &schemaRegistryClient,
	}
}
