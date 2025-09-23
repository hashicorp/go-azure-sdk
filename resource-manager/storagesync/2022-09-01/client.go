package v2022_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/cloudendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/registeredservers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/serverendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/storagesyncs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/storagesyncservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/syncgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/workflows"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CloudEndpoints             *cloudendpoints.CloudEndpointsClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	RegisteredServers          *registeredservers.RegisteredServersClient
	ServerEndpoints            *serverendpoints.ServerEndpointsClient
	StorageSyncServices        *storagesyncservices.StorageSyncServicesClient
	Storagesyncs               *storagesyncs.StoragesyncsClient
	SyncGroups                 *syncgroups.SyncGroupsClient
	Workflows                  *workflows.WorkflowsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	cloudEndpointsClient, err := cloudendpoints.NewCloudEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudEndpoints client: %+v", err)
	}
	configureFunc(cloudEndpointsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	registeredServersClient, err := registeredservers.NewRegisteredServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegisteredServers client: %+v", err)
	}
	configureFunc(registeredServersClient.Client)

	serverEndpointsClient, err := serverendpoints.NewServerEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerEndpoints client: %+v", err)
	}
	configureFunc(serverEndpointsClient.Client)

	storageSyncServicesClient, err := storagesyncservices.NewStorageSyncServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageSyncServices client: %+v", err)
	}
	configureFunc(storageSyncServicesClient.Client)

	storagesyncsClient, err := storagesyncs.NewStoragesyncsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Storagesyncs client: %+v", err)
	}
	configureFunc(storagesyncsClient.Client)

	syncGroupsClient, err := syncgroups.NewSyncGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SyncGroups client: %+v", err)
	}
	configureFunc(syncGroupsClient.Client)

	workflowsClient, err := workflows.NewWorkflowsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workflows client: %+v", err)
	}
	configureFunc(workflowsClient.Client)

	return &Client{
		CloudEndpoints:             cloudEndpointsClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		RegisteredServers:          registeredServersClient,
		ServerEndpoints:            serverEndpointsClient,
		StorageSyncServices:        storageSyncServicesClient,
		Storagesyncs:               storagesyncsClient,
		SyncGroups:                 syncGroupsClient,
		Workflows:                  workflowsClient,
	}, nil
}
