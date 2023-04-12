package v2021_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/collection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/dicomservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/fhirservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/iotconnectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/proxy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/resource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/workspaceprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/workspaceprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-11-01/workspaces"
)

type Client struct {
	Collection                          *collection.CollectionClient
	DicomServices                       *dicomservices.DicomServicesClient
	FhirServices                        *fhirservices.FhirServicesClient
	IotConnectors                       *iotconnectors.IotConnectorsClient
	PrivateEndpointConnections          *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                *privatelinkresources.PrivateLinkResourcesClient
	Proxy                               *proxy.ProxyClient
	Resource                            *resource.ResourceClient
	WorkspacePrivateEndpointConnections *workspaceprivateendpointconnections.WorkspacePrivateEndpointConnectionsClient
	WorkspacePrivateLinkResources       *workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient
	Workspaces                          *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	collectionClient := collection.NewCollectionClientWithBaseURI(endpoint)
	configureAuthFunc(&collectionClient.Client)

	dicomServicesClient := dicomservices.NewDicomServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&dicomServicesClient.Client)

	fhirServicesClient := fhirservices.NewFhirServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&fhirServicesClient.Client)

	iotConnectorsClient := iotconnectors.NewIotConnectorsClientWithBaseURI(endpoint)
	configureAuthFunc(&iotConnectorsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	proxyClient := proxy.NewProxyClientWithBaseURI(endpoint)
	configureAuthFunc(&proxyClient.Client)

	resourceClient := resource.NewResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceClient.Client)

	workspacePrivateEndpointConnectionsClient := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacePrivateEndpointConnectionsClient.Client)

	workspacePrivateLinkResourcesClient := workspaceprivatelinkresources.NewWorkspacePrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacePrivateLinkResourcesClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		Collection:                          &collectionClient,
		DicomServices:                       &dicomServicesClient,
		FhirServices:                        &fhirServicesClient,
		IotConnectors:                       &iotConnectorsClient,
		PrivateEndpointConnections:          &privateEndpointConnectionsClient,
		PrivateLinkResources:                &privateLinkResourcesClient,
		Proxy:                               &proxyClient,
		Resource:                            &resourceClient,
		WorkspacePrivateEndpointConnections: &workspacePrivateEndpointConnectionsClient,
		WorkspacePrivateLinkResources:       &workspacePrivateLinkResourcesClient,
		Workspaces:                          &workspacesClient,
	}
}
