// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_03_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/clusterextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/extensionoperationstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/extensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/flux"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/fluxconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/fluxconfigurationoperationstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/operationsinacluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/sourcecontrolconfiguration"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	clusterExtensionsClient := clusterextensions.NewClusterExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&clusterExtensionsClient.Client)

	extensionOperationStatusClient := extensionoperationstatus.NewExtensionOperationStatusClientWithBaseURI(endpoint)
	configureAuthFunc(&extensionOperationStatusClient.Client)

	extensionsClient := extensions.NewExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&extensionsClient.Client)

	fluxClient := flux.NewFluxClientWithBaseURI(endpoint)
	configureAuthFunc(&fluxClient.Client)

	fluxConfigurationClient := fluxconfiguration.NewFluxConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&fluxConfigurationClient.Client)

	fluxConfigurationOperationStatusClient := fluxconfigurationoperationstatus.NewFluxConfigurationOperationStatusClientWithBaseURI(endpoint)
	configureAuthFunc(&fluxConfigurationOperationStatusClient.Client)

	operationsInAClusterClient := operationsinacluster.NewOperationsInAClusterClientWithBaseURI(endpoint)
	configureAuthFunc(&operationsInAClusterClient.Client)

	sourceControlConfigurationClient := sourcecontrolconfiguration.NewSourceControlConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&sourceControlConfigurationClient.Client)

	return Client{
		ClusterExtensions:                &clusterExtensionsClient,
		ExtensionOperationStatus:         &extensionOperationStatusClient,
		Extensions:                       &extensionsClient,
		Flux:                             &fluxClient,
		FluxConfiguration:                &fluxConfigurationClient,
		FluxConfigurationOperationStatus: &fluxConfigurationOperationStatusClient,
		OperationsInACluster:             &operationsInAClusterClient,
		SourceControlConfiguration:       &sourceControlConfigurationClient,
	}
}
