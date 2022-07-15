package v2021_01_01_preview

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/authorizationrulesdisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/authorizationruleseventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/authorizationrulesnamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/checknameavailabilitydisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/checknameavailabilitynamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/consumergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/eventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/namespacesprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/namespacesprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-01-01-preview/networkrulesets"
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
	Namespaces                                   *namespaces.NamespacesClient
	NamespacesPrivateEndpointConnections         *namespacesprivateendpointconnections.NamespacesPrivateEndpointConnectionsClient
	NamespacesPrivateLinkResources               *namespacesprivatelinkresources.NamespacesPrivateLinkResourcesClient
	NetworkRuleSets                              *networkrulesets.NetworkRuleSetsClient
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

	namespacesClient := namespaces.NewNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesClient.Client)

	namespacesPrivateEndpointConnectionsClient := namespacesprivateendpointconnections.NewNamespacesPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateEndpointConnectionsClient.Client)

	namespacesPrivateLinkResourcesClient := namespacesprivatelinkresources.NewNamespacesPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesPrivateLinkResourcesClient.Client)

	networkRuleSetsClient := networkrulesets.NewNetworkRuleSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&networkRuleSetsClient.Client)

	return Client{
		AuthorizationRulesDisasterRecoveryConfigs:    &authorizationRulesDisasterRecoveryConfigsClient,
		AuthorizationRulesEventHubs:                  &authorizationRulesEventHubsClient,
		AuthorizationRulesNamespaces:                 &authorizationRulesNamespacesClient,
		CheckNameAvailabilityDisasterRecoveryConfigs: &checkNameAvailabilityDisasterRecoveryConfigsClient,
		CheckNameAvailabilityNamespaces:              &checkNameAvailabilityNamespacesClient,
		ConsumerGroups:                               &consumerGroupsClient,
		DisasterRecoveryConfigs:                      &disasterRecoveryConfigsClient,
		EventHubs:                                    &eventHubsClient,
		Namespaces:                                   &namespacesClient,
		NamespacesPrivateEndpointConnections:         &namespacesPrivateEndpointConnectionsClient,
		NamespacesPrivateLinkResources:               &namespacesPrivateLinkResourcesClient,
		NetworkRuleSets:                              &networkRuleSetsClient,
	}
}
