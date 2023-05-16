package v2020_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/actions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/backuprestore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/changedetection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/cloudendpointresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/operationstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/privateendpointconnectionresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/registeredserverresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/serverendpointresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/storagesyncservice"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/storagesyncservicesresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/syncgroupresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/workflowresource"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Actions                           *actions.ActionsClient
	BackupRestore                     *backuprestore.BackupRestoreClient
	ChangeDetection                   *changedetection.ChangeDetectionClient
	CloudEndpointResource             *cloudendpointresource.CloudEndpointResourceClient
	OperationStatus                   *operationstatus.OperationStatusClient
	PrivateEndpointConnectionResource *privateendpointconnectionresource.PrivateEndpointConnectionResourceClient
	PrivateEndpointConnections        *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources              *privatelinkresources.PrivateLinkResourcesClient
	RegisteredServerResource          *registeredserverresource.RegisteredServerResourceClient
	ServerEndpointResource            *serverendpointresource.ServerEndpointResourceClient
	StorageSyncService                *storagesyncservice.StorageSyncServiceClient
	StorageSyncServicesResource       *storagesyncservicesresource.StorageSyncServicesResourceClient
	SyncGroupResource                 *syncgroupresource.SyncGroupResourceClient
	WorkflowResource                  *workflowresource.WorkflowResourceClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	actionsClient, err := actions.NewActionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Actions client: %+v", err)
	}
	configureFunc(actionsClient.Client)

	backupRestoreClient, err := backuprestore.NewBackupRestoreClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BackupRestore client: %+v", err)
	}
	configureFunc(backupRestoreClient.Client)

	changeDetectionClient, err := changedetection.NewChangeDetectionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ChangeDetection client: %+v", err)
	}
	configureFunc(changeDetectionClient.Client)

	cloudEndpointResourceClient, err := cloudendpointresource.NewCloudEndpointResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CloudEndpointResource client: %+v", err)
	}
	configureFunc(cloudEndpointResourceClient.Client)

	operationStatusClient, err := operationstatus.NewOperationStatusClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OperationStatus client: %+v", err)
	}
	configureFunc(operationStatusClient.Client)

	privateEndpointConnectionResourceClient, err := privateendpointconnectionresource.NewPrivateEndpointConnectionResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnectionResource client: %+v", err)
	}
	configureFunc(privateEndpointConnectionResourceClient.Client)

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

	registeredServerResourceClient, err := registeredserverresource.NewRegisteredServerResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RegisteredServerResource client: %+v", err)
	}
	configureFunc(registeredServerResourceClient.Client)

	serverEndpointResourceClient, err := serverendpointresource.NewServerEndpointResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerEndpointResource client: %+v", err)
	}
	configureFunc(serverEndpointResourceClient.Client)

	storageSyncServiceClient, err := storagesyncservice.NewStorageSyncServiceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building StorageSyncService client: %+v", err)
	}
	configureFunc(storageSyncServiceClient.Client)

	storageSyncServicesResourceClient, err := storagesyncservicesresource.NewStorageSyncServicesResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building StorageSyncServicesResource client: %+v", err)
	}
	configureFunc(storageSyncServicesResourceClient.Client)

	syncGroupResourceClient, err := syncgroupresource.NewSyncGroupResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SyncGroupResource client: %+v", err)
	}
	configureFunc(syncGroupResourceClient.Client)

	workflowResourceClient, err := workflowresource.NewWorkflowResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowResource client: %+v", err)
	}
	configureFunc(workflowResourceClient.Client)

	return &Client{
		Actions:                           actionsClient,
		BackupRestore:                     backupRestoreClient,
		ChangeDetection:                   changeDetectionClient,
		CloudEndpointResource:             cloudEndpointResourceClient,
		OperationStatus:                   operationStatusClient,
		PrivateEndpointConnectionResource: privateEndpointConnectionResourceClient,
		PrivateEndpointConnections:        privateEndpointConnectionsClient,
		PrivateLinkResources:              privateLinkResourcesClient,
		RegisteredServerResource:          registeredServerResourceClient,
		ServerEndpointResource:            serverEndpointResourceClient,
		StorageSyncService:                storageSyncServiceClient,
		StorageSyncServicesResource:       storageSyncServicesResourceClient,
		SyncGroupResource:                 syncGroupResourceClient,
		WorkflowResource:                  workflowResourceClient,
	}, nil
}
