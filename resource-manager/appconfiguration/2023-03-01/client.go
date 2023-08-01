package v2023_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/configurationstores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/deletedconfigurationstores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/keyvalues"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2023-03-01/replicas"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ConfigurationStores        *configurationstores.ConfigurationStoresClient
	DeletedConfigurationStores *deletedconfigurationstores.DeletedConfigurationStoresClient
	KeyValues                  *keyvalues.KeyValuesClient
	Operations                 *operations.OperationsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Replicas                   *replicas.ReplicasClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	configurationStoresClient, err := configurationstores.NewConfigurationStoresClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationStores client: %+v", err)
	}
	configureFunc(configurationStoresClient.Client)

	deletedConfigurationStoresClient, err := deletedconfigurationstores.NewDeletedConfigurationStoresClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedConfigurationStores client: %+v", err)
	}
	configureFunc(deletedConfigurationStoresClient.Client)

	keyValuesClient, err := keyvalues.NewKeyValuesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building KeyValues client: %+v", err)
	}
	configureFunc(keyValuesClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

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

	replicasClient, err := replicas.NewReplicasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Replicas client: %+v", err)
	}
	configureFunc(replicasClient.Client)

	return &Client{
		ConfigurationStores:        configurationStoresClient,
		DeletedConfigurationStores: deletedConfigurationStoresClient,
		KeyValues:                  keyValuesClient,
		Operations:                 operationsClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		Replicas:                   replicasClient,
	}, nil
}
