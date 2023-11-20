package v2024_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/applicationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/authorizationrulesdisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/authorizationruleseventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/authorizationrulesnamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/checknameavailabilitydisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/checknameavailabilitynamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/consumergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclustersavailableclusterregions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclustersconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclustersnamespace"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclustersupgrade"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/namespacesnetworksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/networkrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/schemaregistry"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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
	EventHubsClustersUpgrade                         *eventhubsclustersupgrade.EventHubsClustersUpgradeClient
	Namespaces                                       *namespaces.NamespacesClient
	NamespacesNetworkSecurityPerimeterConfigurations *namespacesnetworksecurityperimeterconfigurations.NamespacesNetworkSecurityPerimeterConfigurationsClient
	NamespacesPrivateEndpointConnections             *namespacesprivateendpointconnections.NamespacesPrivateEndpointConnectionsClient
	NamespacesPrivateLinkResources                   *namespacesprivatelinkresources.NamespacesPrivateLinkResourcesClient
	NetworkRuleSets                                  *networkrulesets.NetworkRuleSetsClient
	SchemaRegistry                                   *schemaregistry.SchemaRegistryClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applicationGroupClient, err := applicationgroup.NewApplicationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGroup client: %+v", err)
	}
	configureFunc(applicationGroupClient.Client)

	authorizationRulesDisasterRecoveryConfigsClient, err := authorizationrulesdisasterrecoveryconfigs.NewAuthorizationRulesDisasterRecoveryConfigsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationRulesDisasterRecoveryConfigs client: %+v", err)
	}
	configureFunc(authorizationRulesDisasterRecoveryConfigsClient.Client)

	authorizationRulesEventHubsClient, err := authorizationruleseventhubs.NewAuthorizationRulesEventHubsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationRulesEventHubs client: %+v", err)
	}
	configureFunc(authorizationRulesEventHubsClient.Client)

	authorizationRulesNamespacesClient, err := authorizationrulesnamespaces.NewAuthorizationRulesNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationRulesNamespaces client: %+v", err)
	}
	configureFunc(authorizationRulesNamespacesClient.Client)

	checkNameAvailabilityDisasterRecoveryConfigsClient, err := checknameavailabilitydisasterrecoveryconfigs.NewCheckNameAvailabilityDisasterRecoveryConfigsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailabilityDisasterRecoveryConfigs client: %+v", err)
	}
	configureFunc(checkNameAvailabilityDisasterRecoveryConfigsClient.Client)

	checkNameAvailabilityNamespacesClient, err := checknameavailabilitynamespaces.NewCheckNameAvailabilityNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailabilityNamespaces client: %+v", err)
	}
	configureFunc(checkNameAvailabilityNamespacesClient.Client)

	consumerGroupsClient, err := consumergroups.NewConsumerGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConsumerGroups client: %+v", err)
	}
	configureFunc(consumerGroupsClient.Client)

	disasterRecoveryConfigsClient, err := disasterrecoveryconfigs.NewDisasterRecoveryConfigsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DisasterRecoveryConfigs client: %+v", err)
	}
	configureFunc(disasterRecoveryConfigsClient.Client)

	eventHubsClient, err := eventhubs.NewEventHubsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventHubs client: %+v", err)
	}
	configureFunc(eventHubsClient.Client)

	eventHubsClustersAvailableClusterRegionsClient, err := eventhubsclustersavailableclusterregions.NewEventHubsClustersAvailableClusterRegionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClustersAvailableClusterRegions client: %+v", err)
	}
	configureFunc(eventHubsClustersAvailableClusterRegionsClient.Client)

	eventHubsClustersClient, err := eventhubsclusters.NewEventHubsClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClusters client: %+v", err)
	}
	configureFunc(eventHubsClustersClient.Client)

	eventHubsClustersConfigurationClient, err := eventhubsclustersconfiguration.NewEventHubsClustersConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClustersConfiguration client: %+v", err)
	}
	configureFunc(eventHubsClustersConfigurationClient.Client)

	eventHubsClustersNamespaceClient, err := eventhubsclustersnamespace.NewEventHubsClustersNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClustersNamespace client: %+v", err)
	}
	configureFunc(eventHubsClustersNamespaceClient.Client)

	eventHubsClustersUpgradeClient, err := eventhubsclustersupgrade.NewEventHubsClustersUpgradeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventHubsClustersUpgrade client: %+v", err)
	}
	configureFunc(eventHubsClustersUpgradeClient.Client)

	namespacesClient, err := namespaces.NewNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Namespaces client: %+v", err)
	}
	configureFunc(namespacesClient.Client)

	namespacesNetworkSecurityPerimeterConfigurationsClient, err := namespacesnetworksecurityperimeterconfigurations.NewNamespacesNetworkSecurityPerimeterConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NamespacesNetworkSecurityPerimeterConfigurations client: %+v", err)
	}
	configureFunc(namespacesNetworkSecurityPerimeterConfigurationsClient.Client)

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

	networkRuleSetsClient, err := networkrulesets.NewNetworkRuleSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkRuleSets client: %+v", err)
	}
	configureFunc(networkRuleSetsClient.Client)

	schemaRegistryClient, err := schemaregistry.NewSchemaRegistryClientWithBaseURI(sdkApi)
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
		EventHubsClustersUpgrade:                         eventHubsClustersUpgradeClient,
		Namespaces:                                       namespacesClient,
		NamespacesNetworkSecurityPerimeterConfigurations: namespacesNetworkSecurityPerimeterConfigurationsClient,
		NamespacesPrivateEndpointConnections:             namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:                   namespacesPrivateLinkResourcesClient,
		NetworkRuleSets:                                  networkRuleSetsClient,
		SchemaRegistry:                                   schemaRegistryClient,
	}, nil
}
