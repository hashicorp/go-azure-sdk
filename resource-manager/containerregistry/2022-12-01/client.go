package v2022_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/containerregistry/2022-12-01/operation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerregistry/2022-12-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerregistry/2022-12-01/registries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerregistry/2022-12-01/replications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerregistry/2022-12-01/scopemaps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerregistry/2022-12-01/tokens"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerregistry/2022-12-01/webhooks"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Operation                  *operation.OperationClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	Registries                 *registries.RegistriesClient
	Replications               *replications.ReplicationsClient
	ScopeMaps                  *scopemaps.ScopeMapsClient
	Tokens                     *tokens.TokensClient
	WebHooks                   *webhooks.WebHooksClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	operationClient, err := operation.NewOperationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Operation client: %+v", err)
	}
	configureFunc(operationClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	registriesClient, err := registries.NewRegistriesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Registries client: %+v", err)
	}
	configureFunc(registriesClient.Client)

	replicationsClient, err := replications.NewReplicationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Replications client: %+v", err)
	}
	configureFunc(replicationsClient.Client)

	scopeMapsClient, err := scopemaps.NewScopeMapsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ScopeMaps client: %+v", err)
	}
	configureFunc(scopeMapsClient.Client)

	tokensClient, err := tokens.NewTokensClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Tokens client: %+v", err)
	}
	configureFunc(tokensClient.Client)

	webHooksClient, err := webhooks.NewWebHooksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building WebHooks client: %+v", err)
	}
	configureFunc(webHooksClient.Client)

	return &Client{
		Operation:                  operationClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		Registries:                 registriesClient,
		Replications:               replicationsClient,
		ScopeMaps:                  scopeMapsClient,
		Tokens:                     tokensClient,
		WebHooks:                   webHooksClient,
	}, nil
}
