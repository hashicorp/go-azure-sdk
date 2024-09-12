package v2024_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/mongoclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/privatelinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mongocluster/2024-07-01/replicas"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	FirewallRules              *firewallrules.FirewallRulesClient
	MongoClusters              *mongoclusters.MongoClustersClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinks               *privatelinks.PrivateLinksClient
	Replicas                   *replicas.ReplicasClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	mongoClustersClient, err := mongoclusters.NewMongoClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MongoClusters client: %+v", err)
	}
	configureFunc(mongoClustersClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinksClient, err := privatelinks.NewPrivateLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinks client: %+v", err)
	}
	configureFunc(privateLinksClient.Client)

	replicasClient, err := replicas.NewReplicasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Replicas client: %+v", err)
	}
	configureFunc(replicasClient.Client)

	return &Client{
		FirewallRules:              firewallRulesClient,
		MongoClusters:              mongoClustersClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinks:               privateLinksClient,
		Replicas:                   replicasClient,
	}, nil
}
