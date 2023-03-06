// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_12_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/azurebackupjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/azurebackupjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/backupinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/backuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/backupvaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/deletedbackupinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/dppfeaturesupport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/findrestorabletimeranges"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/recoverypoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-12-01/resourceguards"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	AzureBackupJob           *azurebackupjob.AzureBackupJobClient
	AzureBackupJobs          *azurebackupjobs.AzureBackupJobsClient
	BackupInstances          *backupinstances.BackupInstancesClient
	BackupPolicies           *backuppolicies.BackupPoliciesClient
	BackupVaults             *backupvaults.BackupVaultsClient
	DeletedBackupInstances   *deletedbackupinstances.DeletedBackupInstancesClient
	DppFeatureSupport        *dppfeaturesupport.DppFeatureSupportClient
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
		FindRestorableTimeRanges: &findRestorableTimeRangesClient,
		RecoveryPoint:            &recoveryPointClient,
		ResourceGuards:           &resourceGuardsClient,
	}
}
