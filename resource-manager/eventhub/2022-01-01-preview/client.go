package v2022_01_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/applicationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/authorizationrulesdisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/authorizationruleseventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/authorizationrulesnamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/checknameavailabilitydisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/checknameavailabilitynamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/consumergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/eventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/eventhubsclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/eventhubsclustersavailableclusterregions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/eventhubsclustersconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/eventhubsclustersnamespace"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/namespacesnetworksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/networkrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/schemaregistry"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApplicationGroup                                 *applicationgroup.ApplicationGroupClient
	AuthorizationRulesDisasterRecoveryConfigs        *authorizationrulesdisasterrecoveryconfigs.AuthorizationRulesDisasterRecoveryConfigsClient
	AuthorizationRulesEventHubs                      *authorizationruleseventhubs.AuthorizationRulesEventHubsClient
	AuthorizationRulesNamespaces                     *authorizationrulesnamespaces.AuthorizationRulesNamespacesClient
	CheckNameAvailabilityDisasterRecoveryConfigs     *checknameavailabilitydisasterrecoveryconfigs.CheckNameAvailabilityDisasterRecoveryConfigsClient
	CheckNameAvailabilityNamespaces                  *checknameavailabilitynamespaces.CheckNameAvailabilityNamespacesClient
	ConsumerGroups                                   *consumergroups.ConsumerGroupsClient
	DisasterRecoveryConfigs                          *disasterrecoveryconfigs.DisasterRecoveryConfigsClient
	EventHubs                                        *eventhubs.EventHubsClient
	EventHubsClusters                                *eventhubsclusters.EventHubsClustersClient
	EventHubsClustersAvailableClusterRegions         *eventhubsclustersavailableclusterregions.EventHubsClustersAvailableClusterRegionsClient
	EventHubsClustersConfiguration                   *eventhubsclustersconfiguration.EventHubsClustersConfigurationClient
	EventHubsClustersNamespace                       *eventhubsclustersnamespace.EventHubsClustersNamespaceClient
	Namespaces                                       *namespaces.NamespacesClient
	NamespacesNetworkSecurityPerimeterConfigurations *namespacesnetworksecurityperimeterconfigurations.NamespacesNetworkSecurityPerimeterConfigurationsClient
	NamespacesPrivateEndpointConnections             *namespacesprivateendpointconnections.NamespacesPrivateEndpointConnectionsClient
	NamespacesPrivateLinkResources                   *namespacesprivatelinkresources.NamespacesPrivateLinkResourcesClient
	NetworkRuleSets                                  *networkrulesets.NetworkRuleSetsClient
	SchemaRegistry                                   *schemaregistry.SchemaRegistryClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applicationGroupClient, err := applicationgroup.NewApplicationGroupClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGroup client: %+v", err)
	}
	configureFunc(applicationGroupClient.Client)

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

	namespacesNetworkSecurityPerimeterConfigurationsClient, err := namespacesnetworksecurityperimeterconfigurations.NewNamespacesNetworkSecurityPerimeterConfigurationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NamespacesNetworkSecurityPerimeterConfigurations client: %+v", err)
	}
	configureFunc(namespacesNetworkSecurityPerimeterConfigurationsClient.Client)

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
		ApplicationGroup: applicationGroupClient,
		AuthorizationRulesDisasterRecoveryConfigs:        authorizationRulesDisasterRecoveryConfigsClient,
		AuthorizationRulesEventHubs:                      authorizationRulesEventHubsClient,
		AuthorizationRulesNamespaces:                     authorizationRulesNamespacesClient,
		CheckNameAvailabilityDisasterRecoveryConfigs:     checkNameAvailabilityDisasterRecoveryConfigsClient,
		CheckNameAvailabilityNamespaces:                  checkNameAvailabilityNamespacesClient,
		ConsumerGroups:                                   consumerGroupsClient,
		DisasterRecoveryConfigs:                          disasterRecoveryConfigsClient,
		EventHubs:                                        eventHubsClient,
		EventHubsClusters:                                eventHubsClustersClient,
		EventHubsClustersAvailableClusterRegions:         eventHubsClustersAvailableClusterRegionsClient,
		EventHubsClustersConfiguration:                   eventHubsClustersConfigurationClient,
		EventHubsClustersNamespace:                       eventHubsClustersNamespaceClient,
		Namespaces:                                       namespacesClient,
		NamespacesNetworkSecurityPerimeterConfigurations: namespacesNetworkSecurityPerimeterConfigurationsClient,
		NamespacesPrivateEndpointConnections:             namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:                   namespacesPrivateLinkResourcesClient,
		NetworkRuleSets:                                  networkRuleSetsClient,
		SchemaRegistry:                                   schemaRegistryClient,
	}, nil
}
