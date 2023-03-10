package v2020_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	actionsClient := actions.NewActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&actionsClient.Client)

	backupRestoreClient := backuprestore.NewBackupRestoreClientWithBaseURI(endpoint)
	configureAuthFunc(&backupRestoreClient.Client)

	changeDetectionClient := changedetection.NewChangeDetectionClientWithBaseURI(endpoint)
	configureAuthFunc(&changeDetectionClient.Client)

	cloudEndpointResourceClient := cloudendpointresource.NewCloudEndpointResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&cloudEndpointResourceClient.Client)

	operationStatusClient := operationstatus.NewOperationStatusClientWithBaseURI(endpoint)
	configureAuthFunc(&operationStatusClient.Client)

	privateEndpointConnectionResourceClient := privateendpointconnectionresource.NewPrivateEndpointConnectionResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionResourceClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	registeredServerResourceClient := registeredserverresource.NewRegisteredServerResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&registeredServerResourceClient.Client)

	serverEndpointResourceClient := serverendpointresource.NewServerEndpointResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&serverEndpointResourceClient.Client)

	storageSyncServiceClient := storagesyncservice.NewStorageSyncServiceClientWithBaseURI(endpoint)
	configureAuthFunc(&storageSyncServiceClient.Client)

	storageSyncServicesResourceClient := storagesyncservicesresource.NewStorageSyncServicesResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&storageSyncServicesResourceClient.Client)

	syncGroupResourceClient := syncgroupresource.NewSyncGroupResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&syncGroupResourceClient.Client)

	workflowResourceClient := workflowresource.NewWorkflowResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowResourceClient.Client)

	return Client{
		Actions:                           &actionsClient,
		BackupRestore:                     &backupRestoreClient,
		ChangeDetection:                   &changeDetectionClient,
		CloudEndpointResource:             &cloudEndpointResourceClient,
		OperationStatus:                   &operationStatusClient,
		PrivateEndpointConnectionResource: &privateEndpointConnectionResourceClient,
		PrivateEndpointConnections:        &privateEndpointConnectionsClient,
		PrivateLinkResources:              &privateLinkResourcesClient,
		RegisteredServerResource:          &registeredServerResourceClient,
		ServerEndpointResource:            &serverEndpointResourceClient,
		StorageSyncService:                &storageSyncServiceClient,
		StorageSyncServicesResource:       &storageSyncServicesResourceClient,
		SyncGroupResource:                 &syncGroupResourceClient,
		WorkflowResource:                  &workflowResourceClient,
	}
}
