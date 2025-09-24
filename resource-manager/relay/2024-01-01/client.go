package v2024_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/authorizationrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/hybridconnectionoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/hybridconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/networkrulesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/relaynamespaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/relays"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/wcfrelays"
	"github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/wcfrelaysops"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AuthorizationRules             *authorizationrules.AuthorizationRulesClient
	HybridConnectionOperationGroup *hybridconnectionoperationgroup.HybridConnectionOperationGroupClient
	HybridConnections              *hybridconnections.HybridConnectionsClient
	NetworkRuleSets                *networkrulesets.NetworkRuleSetsClient
	PrivateEndpointConnections     *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources           *privatelinkresources.PrivateLinkResourcesClient
	RelayNamespaces                *relaynamespaces.RelayNamespacesClient
	Relays                         *relays.RelaysClient
	WCFRelays                      *wcfrelays.WCFRelaysClient
	WcfRelaysOps                   *wcfrelaysops.WcfRelaysOpsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	authorizationRulesClient, err := authorizationrules.NewAuthorizationRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizationRules client: %+v", err)
	}
	configureFunc(authorizationRulesClient.Client)

	hybridConnectionOperationGroupClient, err := hybridconnectionoperationgroup.NewHybridConnectionOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridConnectionOperationGroup client: %+v", err)
	}
	configureFunc(hybridConnectionOperationGroupClient.Client)

	hybridConnectionsClient, err := hybridconnections.NewHybridConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridConnections client: %+v", err)
	}
	configureFunc(hybridConnectionsClient.Client)

	networkRuleSetsClient, err := networkrulesets.NewNetworkRuleSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkRuleSets client: %+v", err)
	}
	configureFunc(networkRuleSetsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	relayNamespacesClient, err := relaynamespaces.NewRelayNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RelayNamespaces client: %+v", err)
	}
	configureFunc(relayNamespacesClient.Client)

	relaysClient, err := relays.NewRelaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Relays client: %+v", err)
	}
	configureFunc(relaysClient.Client)

	wCFRelaysClient, err := wcfrelays.NewWCFRelaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WCFRelays client: %+v", err)
	}
	configureFunc(wCFRelaysClient.Client)

	wcfRelaysOpsClient, err := wcfrelaysops.NewWcfRelaysOpsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WcfRelaysOps client: %+v", err)
	}
	configureFunc(wcfRelaysOpsClient.Client)

	return &Client{
		AuthorizationRules:             authorizationRulesClient,
		HybridConnectionOperationGroup: hybridConnectionOperationGroupClient,
		HybridConnections:              hybridConnectionsClient,
		NetworkRuleSets:                networkRuleSetsClient,
		PrivateEndpointConnections:     privateEndpointConnectionsClient,
		PrivateLinkResources:           privateLinkResourcesClient,
		RelayNamespaces:                relayNamespacesClient,
		Relays:                         relaysClient,
		WCFRelays:                      wCFRelaysClient,
		WcfRelaysOps:                   wcfRelaysOpsClient,
	}, nil
}
