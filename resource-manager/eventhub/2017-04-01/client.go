package v2017_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/authorizationrulesdisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/authorizationruleseventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/authorizationrulesnamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/checknameavailabilitydisasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/checknameavailabilitynamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/consumergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/eventhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/messagingplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/networkrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2017-04-01/regions"
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
	MessagingPlan                                *messagingplan.MessagingPlanClient
	Namespaces                                   *namespaces.NamespacesClient
	NetworkRuleSets                              *networkrulesets.NetworkRuleSetsClient
	Regions                                      *regions.RegionsClient
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

	messagingPlanClient := messagingplan.NewMessagingPlanClientWithBaseURI(endpoint)
	configureAuthFunc(&messagingPlanClient.Client)

	namespacesClient := namespaces.NewNamespacesClientWithBaseURI(endpoint)
	configureAuthFunc(&namespacesClient.Client)

	networkRuleSetsClient := networkrulesets.NewNetworkRuleSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&networkRuleSetsClient.Client)

	regionsClient := regions.NewRegionsClientWithBaseURI(endpoint)
	configureAuthFunc(&regionsClient.Client)

	return Client{
		AuthorizationRulesDisasterRecoveryConfigs:    &authorizationRulesDisasterRecoveryConfigsClient,
		AuthorizationRulesEventHubs:                  &authorizationRulesEventHubsClient,
		AuthorizationRulesNamespaces:                 &authorizationRulesNamespacesClient,
		CheckNameAvailabilityDisasterRecoveryConfigs: &checkNameAvailabilityDisasterRecoveryConfigsClient,
		CheckNameAvailabilityNamespaces:              &checkNameAvailabilityNamespacesClient,
		ConsumerGroups:                               &consumerGroupsClient,
		DisasterRecoveryConfigs:                      &disasterRecoveryConfigsClient,
		EventHubs:                                    &eventHubsClient,
		MessagingPlan:                                &messagingPlanClient,
		Namespaces:                                   &namespacesClient,
		NetworkRuleSets:                              &networkRuleSetsClient,
		Regions:                                      &regionsClient,
	}
}
