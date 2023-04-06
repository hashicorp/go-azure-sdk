package v2023_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/azurebackupjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/azurebackupjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/backupinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/backuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/backupvaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/deletedbackupinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/dppfeaturesupport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/dppresourceguardproxies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/findrestorabletimeranges"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/recoverypoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-01-01/resourceguards"
)

type Client struct {
	AzureBackupJob           *azurebackupjob.AzureBackupJobClient
	AzureBackupJobs          *azurebackupjobs.AzureBackupJobsClient
	BackupInstances          *backupinstances.BackupInstancesClient
	BackupPolicies           *backuppolicies.BackupPoliciesClient
	BackupVaults             *backupvaults.BackupVaultsClient
	DeletedBackupInstances   *deletedbackupinstances.DeletedBackupInstancesClient
	DppFeatureSupport        *dppfeaturesupport.DppFeatureSupportClient
	DppResourceGuardProxies  *dppresourceguardproxies.DppResourceGuardProxiesClient
	FindRestorableTimeRanges *findrestorabletimeranges.FindRestorableTimeRangesClient
	RecoveryPoint            *recoverypoint.RecoveryPointClient
	ResourceGuards           *resourceguards.ResourceGuardsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	azureBackupJobClient := azurebackupjob.NewAzureBackupJobClientWithBaseURI(endpoint)
	configureAuthFunc(&azureBackupJobClient.Client)

	azureBackupJobsClient := azurebackupjobs.NewAzureBackupJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&azureBackupJobsClient.Client)

	backupInstancesClient := backupinstances.NewBackupInstancesClientWithBaseURI(endpoint)
	configureAuthFunc(&backupInstancesClient.Client)

	backupPoliciesClient := backuppolicies.NewBackupPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&backupPoliciesClient.Client)

	backupVaultsClient := backupvaults.NewBackupVaultsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupVaultsClient.Client)

	deletedBackupInstancesClient := deletedbackupinstances.NewDeletedBackupInstancesClientWithBaseURI(endpoint)
	configureAuthFunc(&deletedBackupInstancesClient.Client)

	dppFeatureSupportClient := dppfeaturesupport.NewDppFeatureSupportClientWithBaseURI(endpoint)
	configureAuthFunc(&dppFeatureSupportClient.Client)

	dppResourceGuardProxiesClient := dppresourceguardproxies.NewDppResourceGuardProxiesClientWithBaseURI(endpoint)
	configureAuthFunc(&dppResourceGuardProxiesClient.Client)

	findRestorableTimeRangesClient := findrestorabletimeranges.NewFindRestorableTimeRangesClientWithBaseURI(endpoint)
	configureAuthFunc(&findRestorableTimeRangesClient.Client)

	recoveryPointClient := recoverypoint.NewRecoveryPointClientWithBaseURI(endpoint)
	configureAuthFunc(&recoveryPointClient.Client)

	resourceGuardsClient := resourceguards.NewResourceGuardsClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceGuardsClient.Client)

	return Client{
		AzureBackupJob:           &azureBackupJobClient,
		AzureBackupJobs:          &azureBackupJobsClient,
		BackupInstances:          &backupInstancesClient,
		BackupPolicies:           &backupPoliciesClient,
		BackupVaults:             &backupVaultsClient,
		DeletedBackupInstances:   &deletedBackupInstancesClient,
		DppFeatureSupport:        &dppFeatureSupportClient,
		DppResourceGuardProxies:  &dppResourceGuardProxiesClient,
		FindRestorableTimeRanges: &findRestorableTimeRangesClient,
		RecoveryPoint:            &recoveryPointClient,
		ResourceGuards:           &resourceGuardsClient,
	}
}
