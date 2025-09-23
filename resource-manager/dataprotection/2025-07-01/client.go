package v2025_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/azurebackupjobresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/azurebackuprecoverypointresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/backupinstanceresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/backupinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/backupvaultoperationresults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/backupvaultresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/basebackuppolicyresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/dataprotections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/deletedbackupinstanceresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/dppbaseresourceoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/resourceguardproxybaseresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/resourceguardresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/resourceguards"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AzureBackupJobResources           *azurebackupjobresources.AzureBackupJobResourcesClient
	AzureBackupRecoveryPointResources *azurebackuprecoverypointresources.AzureBackupRecoveryPointResourcesClient
	BackupInstanceResources           *backupinstanceresources.BackupInstanceResourcesClient
	BackupInstances                   *backupinstances.BackupInstancesClient
	BackupVaultOperationResults       *backupvaultoperationresults.BackupVaultOperationResultsClient
	BackupVaultResources              *backupvaultresources.BackupVaultResourcesClient
	BaseBackupPolicyResources         *basebackuppolicyresources.BaseBackupPolicyResourcesClient
	Dataprotections                   *dataprotections.DataprotectionsClient
	DeletedBackupInstanceResources    *deletedbackupinstanceresources.DeletedBackupInstanceResourcesClient
	DppBaseResourceOperationGroup     *dppbaseresourceoperationgroup.DppBaseResourceOperationGroupClient
	ResourceGuardProxyBaseResources   *resourceguardproxybaseresources.ResourceGuardProxyBaseResourcesClient
	ResourceGuardResources            *resourceguardresources.ResourceGuardResourcesClient
	ResourceGuards                    *resourceguards.ResourceGuardsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	azureBackupJobResourcesClient, err := azurebackupjobresources.NewAzureBackupJobResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AzureBackupJobResources client: %+v", err)
	}
	configureFunc(azureBackupJobResourcesClient.Client)

	azureBackupRecoveryPointResourcesClient, err := azurebackuprecoverypointresources.NewAzureBackupRecoveryPointResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AzureBackupRecoveryPointResources client: %+v", err)
	}
	configureFunc(azureBackupRecoveryPointResourcesClient.Client)

	backupInstanceResourcesClient, err := backupinstanceresources.NewBackupInstanceResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupInstanceResources client: %+v", err)
	}
	configureFunc(backupInstanceResourcesClient.Client)

	backupInstancesClient, err := backupinstances.NewBackupInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupInstances client: %+v", err)
	}
	configureFunc(backupInstancesClient.Client)

	backupVaultOperationResultsClient, err := backupvaultoperationresults.NewBackupVaultOperationResultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupVaultOperationResults client: %+v", err)
	}
	configureFunc(backupVaultOperationResultsClient.Client)

	backupVaultResourcesClient, err := backupvaultresources.NewBackupVaultResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupVaultResources client: %+v", err)
	}
	configureFunc(backupVaultResourcesClient.Client)

	baseBackupPolicyResourcesClient, err := basebackuppolicyresources.NewBaseBackupPolicyResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BaseBackupPolicyResources client: %+v", err)
	}
	configureFunc(baseBackupPolicyResourcesClient.Client)

	dataprotectionsClient, err := dataprotections.NewDataprotectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Dataprotections client: %+v", err)
	}
	configureFunc(dataprotectionsClient.Client)

	deletedBackupInstanceResourcesClient, err := deletedbackupinstanceresources.NewDeletedBackupInstanceResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedBackupInstanceResources client: %+v", err)
	}
	configureFunc(deletedBackupInstanceResourcesClient.Client)

	dppBaseResourceOperationGroupClient, err := dppbaseresourceoperationgroup.NewDppBaseResourceOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DppBaseResourceOperationGroup client: %+v", err)
	}
	configureFunc(dppBaseResourceOperationGroupClient.Client)

	resourceGuardProxyBaseResourcesClient, err := resourceguardproxybaseresources.NewResourceGuardProxyBaseResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceGuardProxyBaseResources client: %+v", err)
	}
	configureFunc(resourceGuardProxyBaseResourcesClient.Client)

	resourceGuardResourcesClient, err := resourceguardresources.NewResourceGuardResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceGuardResources client: %+v", err)
	}
	configureFunc(resourceGuardResourcesClient.Client)

	resourceGuardsClient, err := resourceguards.NewResourceGuardsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceGuards client: %+v", err)
	}
	configureFunc(resourceGuardsClient.Client)

	return &Client{
		AzureBackupJobResources:           azureBackupJobResourcesClient,
		AzureBackupRecoveryPointResources: azureBackupRecoveryPointResourcesClient,
		BackupInstanceResources:           backupInstanceResourcesClient,
		BackupInstances:                   backupInstancesClient,
		BackupVaultOperationResults:       backupVaultOperationResultsClient,
		BackupVaultResources:              backupVaultResourcesClient,
		BaseBackupPolicyResources:         baseBackupPolicyResourcesClient,
		Dataprotections:                   dataprotectionsClient,
		DeletedBackupInstanceResources:    deletedBackupInstanceResourcesClient,
		DppBaseResourceOperationGroup:     dppBaseResourceOperationGroupClient,
		ResourceGuardProxyBaseResources:   resourceGuardProxyBaseResourcesClient,
		ResourceGuardResources:            resourceGuardResourcesClient,
		ResourceGuards:                    resourceGuardsClient,
	}, nil
}
