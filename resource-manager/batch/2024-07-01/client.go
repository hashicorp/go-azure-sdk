package v2024_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/applicationpackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/applications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/batchaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/detectorresponses"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/networksecurityperimeterconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/pools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApplicationPackages                    *applicationpackages.ApplicationPackagesClient
	Applications                           *applications.ApplicationsClient
	BatchAccounts                          *batchaccounts.BatchAccountsClient
	Certificates                           *certificates.CertificatesClient
	DetectorResponses                      *detectorresponses.DetectorResponsesClient
	NetworkSecurityPerimeterConfigurations *networksecurityperimeterconfigurations.NetworkSecurityPerimeterConfigurationsClient
	Openapis                               *openapis.OpenapisClient
	Pools                                  *pools.PoolsClient
	PrivateEndpointConnections             *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                   *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applicationPackagesClient, err := applicationpackages.NewApplicationPackagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationPackages client: %+v", err)
	}
	configureFunc(applicationPackagesClient.Client)

	applicationsClient, err := applications.NewApplicationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Applications client: %+v", err)
	}
	configureFunc(applicationsClient.Client)

	batchAccountsClient, err := batchaccounts.NewBatchAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BatchAccounts client: %+v", err)
	}
	configureFunc(batchAccountsClient.Client)

	certificatesClient, err := certificates.NewCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Certificates client: %+v", err)
	}
	configureFunc(certificatesClient.Client)

	detectorResponsesClient, err := detectorresponses.NewDetectorResponsesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DetectorResponses client: %+v", err)
	}
	configureFunc(detectorResponsesClient.Client)

	networkSecurityPerimeterConfigurationsClient, err := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkSecurityPerimeterConfigurations client: %+v", err)
	}
	configureFunc(networkSecurityPerimeterConfigurationsClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

	poolsClient, err := pools.NewPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Pools client: %+v", err)
	}
	configureFunc(poolsClient.Client)

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

	return &Client{
		ApplicationPackages:                    applicationPackagesClient,
		Applications:                           applicationsClient,
		BatchAccounts:                          batchAccountsClient,
		Certificates:                           certificatesClient,
		DetectorResponses:                      detectorResponsesClient,
		NetworkSecurityPerimeterConfigurations: networkSecurityPerimeterConfigurationsClient,
		Openapis:                               openapisClient,
		Pools:                                  poolsClient,
		PrivateEndpointConnections:             privateEndpointConnectionsClient,
		PrivateLinkResources:                   privateLinkResourcesClient,
	}, nil
}
