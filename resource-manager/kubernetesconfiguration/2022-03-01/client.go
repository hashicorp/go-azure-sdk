package v2022_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/clusterextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/extensionoperationstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/extensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/flux"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/fluxconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/fluxconfigurationoperationstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/operationsinacluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/sourcecontrolconfiguration"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ClusterExtensions                *clusterextensions.ClusterExtensionsClient
	ExtensionOperationStatus         *extensionoperationstatus.ExtensionOperationStatusClient
	Extensions                       *extensions.ExtensionsClient
	Flux                             *flux.FluxClient
	FluxConfiguration                *fluxconfiguration.FluxConfigurationClient
	FluxConfigurationOperationStatus *fluxconfigurationoperationstatus.FluxConfigurationOperationStatusClient
	OperationsInACluster             *operationsinacluster.OperationsInAClusterClient
	SourceControlConfiguration       *sourcecontrolconfiguration.SourceControlConfigurationClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	clusterExtensionsClient, err := clusterextensions.NewClusterExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ClusterExtensions client: %+v", err)
	}
	configureFunc(clusterExtensionsClient.Client)

	extensionOperationStatusClient, err := extensionoperationstatus.NewExtensionOperationStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExtensionOperationStatus client: %+v", err)
	}
	configureFunc(extensionOperationStatusClient.Client)

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

	fluxConfigurationOperationStatusClient, err := fluxconfigurationoperationstatus.NewFluxConfigurationOperationStatusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FluxConfigurationOperationStatus client: %+v", err)
	}
	configureFunc(fluxConfigurationOperationStatusClient.Client)

	operationsInAClusterClient, err := operationsinacluster.NewOperationsInAClusterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OperationsInACluster client: %+v", err)
	}
	configureFunc(operationsInAClusterClient.Client)

	sourceControlConfigurationClient, err := sourcecontrolconfiguration.NewSourceControlConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SourceControlConfiguration client: %+v", err)
	}
	configureFunc(sourceControlConfigurationClient.Client)

	return &Client{
		ClusterExtensions:                clusterExtensionsClient,
		ExtensionOperationStatus:         extensionOperationStatusClient,
		Extensions:                       extensionsClient,
		Flux:                             fluxClient,
		FluxConfiguration:                fluxConfigurationClient,
		FluxConfigurationOperationStatus: fluxConfigurationOperationStatusClient,
		OperationsInACluster:             operationsInAClusterClient,
		SourceControlConfiguration:       sourceControlConfigurationClient,
	}, nil
}
