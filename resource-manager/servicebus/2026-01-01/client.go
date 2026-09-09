package v2026_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/armdisasterrecoveries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/disasterrecoveryconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/migrationconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/namespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/namespacesauthorizationrule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/networksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/queues"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/rules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/servicebuses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/topics"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ArmDisasterRecoveries                  *armdisasterrecoveries.ArmDisasterRecoveriesClient
	DisasterRecoveryConfigs                *disasterrecoveryconfigs.DisasterRecoveryConfigsClient
	MigrationConfigs                       *migrationconfigs.MigrationConfigsClient
	Namespaces                             *namespaces.NamespacesClient
	NamespacesAuthorizationRule            *namespacesauthorizationrule.NamespacesAuthorizationRuleClient
	NetworkSecurityPerimeterConfigurations *networksecurityperimeterconfigurations.NetworkSecurityPerimeterConfigurationsClient
	PrivateEndpointConnections             *privateendpointconnections.PrivateEndpointConnectionsClient
	Queues                                 *queues.QueuesClient
	Rules                                  *rules.RulesClient
	Servicebuses                           *servicebuses.ServicebusesClient
	Subscriptions                          *subscriptions.SubscriptionsClient
	Topics                                 *topics.TopicsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	armDisasterRecoveriesClient, err := armdisasterrecoveries.NewArmDisasterRecoveriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ArmDisasterRecoveries client: %+v", err)
	}
	configureFunc(armDisasterRecoveriesClient.Client)

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

	networkSecurityPerimeterConfigurationsClient, err := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterConfigurations client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterConfigurationsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

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

	servicebusesClient, err := servicebuses.NewServicebusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servicebuses client: %+v", err)
	}
	configureFunc(servicebusesClient.Client)

	subscriptionsClient, err := subscriptions.NewSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Subscriptions client: %+v", err)
	}
	configureFunc(subscriptionsClient.Client)

	topicsClient, err := topics.NewTopicsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Topics client: %+v", err)
	}
	configureFunc(topicsClient.Client)

	return &Client{
		ArmDisasterRecoveries:                  armDisasterRecoveriesClient,
		DisasterRecoveryConfigs:                disasterRecoveryConfigsClient,
		MigrationConfigs:                       migrationConfigsClient,
		Namespaces:                             namespacesClient,
		NamespacesAuthorizationRule:            namespacesAuthorizationRuleClient,
		NetworkSecurityPerimeterConfigurations: networkSecurityPerimeterConfigurationsClient,
		PrivateEndpointConnections:             privateEndpointConnectionsClient,
		Queues:                                 queuesClient,
		Rules:                                  rulesClient,
		Servicebuses:                           servicebusesClient,
		Subscriptions:                          subscriptionsClient,
		Topics:                                 topicsClient,
	}, nil
}
