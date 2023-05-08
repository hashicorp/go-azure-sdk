package v2021_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

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

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	authorizationRulesDisasterRecoveryConfigsClient, err := authorizationrulesdisasterrecoveryconfigs.NewAuthorizationRulesDisasterRecoveryConfigsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationRulesDisasterRecoveryConfigs client: %+v", err)
	}
	configureFunc(authorizationRulesDisasterRecoveryConfigsClient.Client)

	authorizationRulesEventHubsClient, err := authorizationruleseventhubs.NewAuthorizationRulesEventHubsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationRulesEventHubs client: %+v", err)
	}
	configureFunc(authorizationRulesEventHubsClient.Client)

	authorizationRulesNamespacesClient, err := authorizationrulesnamespaces.NewAuthorizationRulesNamespacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationRulesNamespaces client: %+v", err)
	}
	configureFunc(authorizationRulesNamespacesClient.Client)

	checkNameAvailabilityDisasterRecoveryConfigsClient, err := checknameavailabilitydisasterrecoveryconfigs.NewCheckNameAvailabilityDisasterRecoveryConfigsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailabilityDisasterRecoveryConfigs client: %+v", err)
	}
	configureFunc(checkNameAvailabilityDisasterRecoveryConfigsClient.Client)

	checkNameAvailabilityNamespacesClient, err := checknameavailabilitynamespaces.NewCheckNameAvailabilityNamespacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailabilityNamespaces client: %+v", err)
	}
	configureFunc(checkNameAvailabilityNamespacesClient.Client)

	consumerGroupsClient, err := consumergroups.NewConsumerGroupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ConsumerGroups client: %+v", err)
	}
	configureFunc(consumerGroupsClient.Client)

	disasterRecoveryConfigsClient, err := disasterrecoveryconfigs.NewDisasterRecoveryConfigsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DisasterRecoveryConfigs client: %+v", err)
	}
	configureFunc(disasterRecoveryConfigsClient.Client)

	eventHubsClient, err := eventhubs.NewEventHubsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EventHubs client: %+v", err)
	}
	configureFunc(eventHubsClient.Client)

	eventHubsClustersAvailableClusterRegionsClient, err := eventhubsclustersavailableclusterregions.NewEventHubsClustersAvailableClusterRegionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClustersAvailableClusterRegions client: %+v", err)
	}
	configureFunc(eventHubsClustersAvailableClusterRegionsClient.Client)

	eventHubsClustersClient, err := eventhubsclusters.NewEventHubsClustersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClusters client: %+v", err)
	}
	configureFunc(eventHubsClustersClient.Client)

	eventHubsClustersConfigurationClient, err := eventhubsclustersconfiguration.NewEventHubsClustersConfigurationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClustersConfiguration client: %+v", err)
	}
	configureFunc(eventHubsClustersConfigurationClient.Client)

	eventHubsClustersNamespaceClient, err := eventhubsclustersnamespace.NewEventHubsClustersNamespaceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClustersNamespace client: %+v", err)
	}
	configureFunc(eventHubsClustersNamespaceClient.Client)

	namespacesClient, err := namespaces.NewNamespacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Namespaces client: %+v", err)
	}
	configureFunc(namespacesClient.Client)

	namespacesPrivateEndpointConnectionsClient, err := namespacesprivateendpointconnections.NewNamespacesPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NamespacesPrivateEndpointConnections client: %+v", err)
	}
	configureFunc(namespacesPrivateEndpointConnectionsClient.Client)

	namespacesPrivateLinkResourcesClient, err := namespacesprivatelinkresources.NewNamespacesPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NamespacesPrivateLinkResources client: %+v", err)
	}
	configureFunc(namespacesPrivateLinkResourcesClient.Client)

	networkRuleSetsClient, err := networkrulesets.NewNetworkRuleSetsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NetworkRuleSets client: %+v", err)
	}
	configureFunc(networkRuleSetsClient.Client)

	schemaRegistryClient, err := schemaregistry.NewSchemaRegistryClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SchemaRegistry client: %+v", err)
	}
	configureFunc(schemaRegistryClient.Client)

	return &Client{
		AuthorizationRulesDisasterRecoveryConfigs:    authorizationRulesDisasterRecoveryConfigsClient,
		AuthorizationRulesEventHubs:                  authorizationRulesEventHubsClient,
		AuthorizationRulesNamespaces:                 authorizationRulesNamespacesClient,
		CheckNameAvailabilityDisasterRecoveryConfigs: checkNameAvailabilityDisasterRecoveryConfigsClient,
		CheckNameAvailabilityNamespaces:              checkNameAvailabilityNamespacesClient,
		ConsumerGroups:                               consumerGroupsClient,
		DisasterRecoveryConfigs:                      disasterRecoveryConfigsClient,
		EventHubs:                                    eventHubsClient,
		EventHubsClusters:                            eventHubsClustersClient,
		EventHubsClustersAvailableClusterRegions:     eventHubsClustersAvailableClusterRegionsClient,
		EventHubsClustersConfiguration:               eventHubsClustersConfigurationClient,
		EventHubsClustersNamespace:                   eventHubsClustersNamespaceClient,
		Namespaces:                                   namespacesClient,
		NamespacesPrivateEndpointConnections:         namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:               namespacesPrivateLinkResourcesClient,
		NetworkRuleSets:                              networkRuleSetsClient,
		SchemaRegistry:                               schemaRegistryClient,
	}, nil
}
