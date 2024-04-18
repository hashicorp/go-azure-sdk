package v2024_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/azurebackupjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/azurebackupjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/backupinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/backupinstancesextensionrouting"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/backuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/backupvaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/deletedbackupinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/dppfeaturesupport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/dppjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/dppresourceguardproxies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/fetchsecondaryrecoverypoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/findrestorabletimeranges"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/recoverypoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/resourceguards"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AzureBackupJob                  *azurebackupjob.AzureBackupJobClient
	AzureBackupJobs                 *azurebackupjobs.AzureBackupJobsClient
	BackupInstances                 *backupinstances.BackupInstancesClient
	BackupInstancesExtensionRouting *backupinstancesextensionrouting.BackupInstancesExtensionRoutingClient
	BackupPolicies                  *backuppolicies.BackupPoliciesClient
	BackupVaults                    *backupvaults.BackupVaultsClient
	DeletedBackupInstances          *deletedbackupinstances.DeletedBackupInstancesClient
	DppFeatureSupport               *dppfeaturesupport.DppFeatureSupportClient
	DppJob                          *dppjob.DppJobClient
	DppResourceGuardProxies         *dppresourceguardproxies.DppResourceGuardProxiesClient
	FetchSecondaryRecoveryPoints    *fetchsecondaryrecoverypoints.FetchSecondaryRecoveryPointsClient
	FindRestorableTimeRanges        *findrestorabletimeranges.FindRestorableTimeRangesClient
	RecoveryPoint                   *recoverypoint.RecoveryPointClient
	ResourceGuards                  *resourceguards.ResourceGuardsClient
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

	backupInstancesExtensionRoutingClient, err := backupinstancesextensionrouting.NewBackupInstancesExtensionRoutingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupInstancesExtensionRouting client: %+v", err)
	}
	configureFunc(backupInstancesExtensionRoutingClient.Client)

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

	deletedBackupInstancesClient, err := deletedbackupinstances.NewDeletedBackupInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedBackupInstances client: %+v", err)
	}
	configureFunc(deletedBackupInstancesClient.Client)

	dppFeatureSupportClient, err := dppfeaturesupport.NewDppFeatureSupportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DppFeatureSupport client: %+v", err)
	}
	configureFunc(dppFeatureSupportClient.Client)

	dppJobClient, err := dppjob.NewDppJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DppJob client: %+v", err)
	}
	configureFunc(dppJobClient.Client)

	dppResourceGuardProxiesClient, err := dppresourceguardproxies.NewDppResourceGuardProxiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DppResourceGuardProxies client: %+v", err)
	}
	configureFunc(dppResourceGuardProxiesClient.Client)

	fetchSecondaryRecoveryPointsClient, err := fetchsecondaryrecoverypoints.NewFetchSecondaryRecoveryPointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FetchSecondaryRecoveryPoints client: %+v", err)
	}
	configureFunc(fetchSecondaryRecoveryPointsClient.Client)

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
		AzureBackupJob:                  azureBackupJobClient,
		AzureBackupJobs:                 azureBackupJobsClient,
		BackupInstances:                 backupInstancesClient,
		BackupInstancesExtensionRouting: backupInstancesExtensionRoutingClient,
		BackupPolicies:                  backupPoliciesClient,
		BackupVaults:                    backupVaultsClient,
		DeletedBackupInstances:          deletedBackupInstancesClient,
		DppFeatureSupport:               dppFeatureSupportClient,
		DppJob:                          dppJobClient,
		DppResourceGuardProxies:         dppResourceGuardProxiesClient,
		FetchSecondaryRecoveryPoints:    fetchSecondaryRecoveryPointsClient,
		FindRestorableTimeRanges:        findRestorableTimeRangesClient,
		RecoveryPoint:                   recoveryPointClient,
		ResourceGuards:                  resourceGuardsClient,
	}, nil
}
