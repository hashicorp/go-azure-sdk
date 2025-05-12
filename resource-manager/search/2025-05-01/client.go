package v2025_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2025-05-01/adminkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2025-05-01/networksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2025-05-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2025-05-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2025-05-01/querykeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2025-05-01/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2025-05-01/sharedprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/search/2025-05-01/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdminKeys                              *adminkeys.AdminKeysClient
	NetworkSecurityPerimeterConfigurations *networksecurityperimeterconfigurations.NetworkSecurityPerimeterConfigurationsClient
	PrivateEndpointConnections             *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                   *privatelinkresources.PrivateLinkResourcesClient
	QueryKeys                              *querykeys.QueryKeysClient
	Services                               *services.ServicesClient
	SharedPrivateLinkResources             *sharedprivatelinkresources.SharedPrivateLinkResourcesClient
	Usages                                 *usages.UsagesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	adminKeysClient, err := adminkeys.NewAdminKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdminKeys client: %+v", err)
	}
	configureFunc(adminKeysClient.Client)

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

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	queryKeysClient, err := querykeys.NewQueryKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueryKeys client: %+v", err)
	}
	configureFunc(queryKeysClient.Client)

	servicesClient, err := services.NewServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Services client: %+v", err)
	}
	configureFunc(servicesClient.Client)

	sharedPrivateLinkResourcesClient, err := sharedprivatelinkresources.NewSharedPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SharedPrivateLinkResources client: %+v", err)
	}
	configureFunc(sharedPrivateLinkResourcesClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	return &Client{
		AdminKeys:                              adminKeysClient,
		NetworkSecurityPerimeterConfigurations: networkSecurityPerimeterConfigurationsClient,
		PrivateEndpointConnections:             privateEndpointConnectionsClient,
		PrivateLinkResources:                   privateLinkResourcesClient,
		QueryKeys:                              queryKeysClient,
		Services:                               servicesClient,
		SharedPrivateLinkResources:             sharedPrivateLinkResourcesClient,
		Usages:                                 usagesClient,
	}, nil
}
