package v2021_06_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/collection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/dicomservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/fhirservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/iotconnectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/proxy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/resource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/workspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Collection                 *collection.CollectionClient
	DicomServices              *dicomservices.DicomServicesClient
	FhirServices               *fhirservices.FhirServicesClient
	IotConnectors              *iotconnectors.IotConnectorsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Proxy                      *proxy.ProxyClient
	Resource                   *resource.ResourceClient
	Workspaces                 *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	collectionClient, err := collection.NewCollectionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Collection client: %+v", err)
	}
	configureFunc(collectionClient.Client)

	dicomServicesClient, err := dicomservices.NewDicomServicesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DicomServices client: %+v", err)
	}
	configureFunc(dicomServicesClient.Client)

	fhirServicesClient, err := fhirservices.NewFhirServicesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building FhirServices client: %+v", err)
	}
	configureFunc(fhirServicesClient.Client)

	iotConnectorsClient, err := iotconnectors.NewIotConnectorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building IotConnectors client: %+v", err)
	}
	configureFunc(iotConnectorsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	proxyClient, err := proxy.NewProxyClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Proxy client: %+v", err)
	}
	configureFunc(proxyClient.Client)

	resourceClient, err := resource.NewResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Resource client: %+v", err)
	}
	configureFunc(resourceClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		Collection:                 collectionClient,
		DicomServices:              dicomServicesClient,
		FhirServices:               fhirServicesClient,
		IotConnectors:              iotConnectorsClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		Proxy:                      proxyClient,
		Resource:                   resourceClient,
		Workspaces:                 workspacesClient,
	}, nil
}
