package v2018_01_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/authorizationrulesdisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/authorizationruleseventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/authorizationrulesnamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/checknameavailabilitydisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/checknameavailabilitynamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/consumergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubsclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubsclustersavailableclusterregions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubsclustersconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubsclustersnamespace"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/ipfilterrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/networkrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/regions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/virtualnetworkrules"
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
	IpFilterRules                                *ipfilterrules.IpFilterRulesClient
	Namespaces                                   *namespaces.NamespacesClient
	NamespacesPrivateEndpointConnections         *namespacesprivateendpointconnections.NamespacesPrivateEndpointConnectionsClient
	NamespacesPrivateLinkResources               *namespacesprivatelinkresources.NamespacesPrivateLinkResourcesClient
	NetworkRuleSets                              *networkrulesets.NetworkRuleSetsClient
	Regions                                      *regions.RegionsClient
	VirtualNetworkRules                          *virtualnetworkrules.VirtualNetworkRulesClient
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

	ipFilterRulesClient := ipfilterrules.NewIpFilterRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&ipFilterRulesClient.Client)

	namespacesClient := namespaces.NewNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesClient.Client)

	namespacesPrivateEndpointConnectionsClient := namespacesprivateendpointconnections.NewNamespacesPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateEndpointConnectionsClient.Client)

	namespacesPrivateLinkResourcesClient := namespacesprivatelinkresources.NewNamespacesPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateLinkResourcesClient.Client)

	networkRuleSetsClient := networkrulesets.NewNetworkRuleSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&networkRuleSetsClient.Client)

	regionsClient := regions.NewRegionsClientWithBaseURI(endpoint)
	configureAuthFunc(&regionsClient.Client)

	virtualNetworkRulesClient := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworkRulesClient.Client)

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
		IpFilterRules:                                &ipFilterRulesClient,
		Namespaces:                                   &namespacesClient,
		NamespacesPrivateEndpointConnections:         &namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:               &namespacesPrivateLinkResourcesClient,
		NetworkRuleSets:                              &networkRuleSetsClient,
		Regions:                                      &regionsClient,
		VirtualNetworkRules:                          &virtualNetworkRulesClient,
	}
}
