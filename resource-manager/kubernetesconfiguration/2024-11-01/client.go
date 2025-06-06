package v2024_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2024-11-01/clusterextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2024-11-01/extensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2024-11-01/flux"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2024-11-01/fluxconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2024-11-01/operationsinacluster"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ClusterExtensions    *clusterextensions.ClusterExtensionsClient
	Extensions           *extensions.ExtensionsClient
	Flux                 *flux.FluxClient
	FluxConfiguration    *fluxconfiguration.FluxConfigurationClient
	OperationsInACluster *operationsinacluster.OperationsInAClusterClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	clusterExtensionsClient, err := clusterextensions.NewClusterExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClusterExtensions client: %+v", err)
	}
	configureFunc(clusterExtensionsClient.Client)

	extensionsClient, err := extensions.NewExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Extensions client: %+v", err)
	}
	configureFunc(extensionsClient.Client)

	fluxClient, err := flux.NewFluxClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Flux client: %+v", err)
	}
	configureFunc(fluxClient.Client)

	fluxConfigurationClient, err := fluxconfiguration.NewFluxConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FluxConfiguration client: %+v", err)
	}
	configureFunc(fluxConfigurationClient.Client)

	operationsInAClusterClient, err := operationsinacluster.NewOperationsInAClusterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OperationsInACluster client: %+v", err)
	}
	configureFunc(operationsInAClusterClient.Client)

	return &Client{
		ClusterExtensions:    clusterExtensionsClient,
		Extensions:           extensionsClient,
		Flux:                 fluxClient,
		FluxConfiguration:    fluxConfigurationClient,
		OperationsInACluster: operationsInAClusterClient,
	}, nil
}
