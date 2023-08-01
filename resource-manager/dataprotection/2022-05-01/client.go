package v2022_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/azurebackupjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/azurebackupjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/backupinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/backuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/backupvaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/dppfeaturesupport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/findrestorabletimeranges"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/recoverypoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2022-05-01/resourceguards"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AzureBackupJob           *azurebackupjob.AzureBackupJobClient
	AzureBackupJobs          *azurebackupjobs.AzureBackupJobsClient
	BackupInstances          *backupinstances.BackupInstancesClient
	BackupPolicies           *backuppolicies.BackupPoliciesClient
	BackupVaults             *backupvaults.BackupVaultsClient
	DppFeatureSupport        *dppfeaturesupport.DppFeatureSupportClient
	FindRestorableTimeRanges *findrestorabletimeranges.FindRestorableTimeRangesClient
	RecoveryPoint            *recoverypoint.RecoveryPointClient
	ResourceGuards           *resourceguards.ResourceGuardsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	azureBackupJobClient, err := azurebackupjob.NewAzureBackupJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AzureBackupJob client: %+v", err)
	}
	configureFunc(azureBackupJobClient.Client)

	azureBackupJobsClient, err := azurebackupjobs.NewAzureBackupJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AzureBackupJobs client: %+v", err)
	}
	configureFunc(azureBackupJobsClient.Client)

	backupInstancesClient, err := backupinstances.NewBackupInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupInstances client: %+v", err)
	}
	configureFunc(backupInstancesClient.Client)

	backupPoliciesClient, err := backuppolicies.NewBackupPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupPolicies client: %+v", err)
	}
	configureFunc(backupPoliciesClient.Client)

	backupVaultsClient, err := backupvaults.NewBackupVaultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupVaults client: %+v", err)
	}
	configureFunc(backupVaultsClient.Client)

	dppFeatureSupportClient, err := dppfeaturesupport.NewDppFeatureSupportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DppFeatureSupport client: %+v", err)
	}
	configureFunc(dppFeatureSupportClient.Client)

	findRestorableTimeRangesClient, err := findrestorabletimeranges.NewFindRestorableTimeRangesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FindRestorableTimeRanges client: %+v", err)
	}
	configureFunc(findRestorableTimeRangesClient.Client)

	recoveryPointClient, err := recoverypoint.NewRecoveryPointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecoveryPoint client: %+v", err)
	}
	configureFunc(recoveryPointClient.Client)

	resourceGuardsClient, err := resourceguards.NewResourceGuardsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceGuards client: %+v", err)
	}
	configureFunc(resourceGuardsClient.Client)

	return &Client{
		AzureBackupJob:           azureBackupJobClient,
		AzureBackupJobs:          azureBackupJobsClient,
		BackupInstances:          backupInstancesClient,
		BackupPolicies:           backupPoliciesClient,
		BackupVaults:             backupVaultsClient,
		DppFeatureSupport:        dppFeatureSupportClient,
		FindRestorableTimeRanges: findRestorableTimeRangesClient,
		RecoveryPoint:            recoveryPointClient,
		ResourceGuards:           resourceGuardsClient,
	}, nil
}
