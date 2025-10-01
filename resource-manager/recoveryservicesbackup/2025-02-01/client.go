package v2025_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupengines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupresourceencryptionconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupresourcestorageconfigsnoncrr"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupresourcevaultconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupworkloaditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/bms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/itemlevelrecoveryconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/jobcancellations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/jobdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/privateendpointconnectionresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/protecteditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/protectioncontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/protectionintentresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/protectionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/recoverypoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/recoverypointsrecommendedformove"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/resourceguardproxies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/resourceguardproxy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/restores"
)

type Client struct {
	BackupEngines                      *backupengines.BackupEnginesClient
	BackupJobs                         *backupjobs.BackupJobsClient
	BackupPolicies                     *backuppolicies.BackupPoliciesClient
	BackupResourceEncryptionConfigs    *backupresourceencryptionconfigs.BackupResourceEncryptionConfigsClient
	BackupResourceStorageConfigsNonCRR *backupresourcestorageconfigsnoncrr.BackupResourceStorageConfigsNonCRRClient
	BackupResourceVaultConfigs         *backupresourcevaultconfigs.BackupResourceVaultConfigsClient
	BackupWorkloadItems                *backupworkloaditems.BackupWorkloadItemsClient
	Backups                            *backups.BackupsClient
	Bms                                *bms.BmsClient
	ItemLevelRecoveryConnections       *itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient
	JobCancellations                   *jobcancellations.JobCancellationsClient
	JobDetails                         *jobdetails.JobDetailsClient
	PrivateEndpointConnectionResources *privateendpointconnectionresources.PrivateEndpointConnectionResourcesClient
	ProtectedItems                     *protecteditems.ProtectedItemsClient
	ProtectionContainers               *protectioncontainers.ProtectionContainersClient
	ProtectionIntentResources          *protectionintentresources.ProtectionIntentResourcesClient
	ProtectionPolicies                 *protectionpolicies.ProtectionPoliciesClient
	RecoveryPoints                     *recoverypoints.RecoveryPointsClient
	RecoveryPointsRecommendedForMove   *recoverypointsrecommendedformove.RecoveryPointsRecommendedForMoveClient
	ResourceGuardProxies               *resourceguardproxies.ResourceGuardProxiesClient
	ResourceGuardProxy                 *resourceguardproxy.ResourceGuardProxyClient
	Restores                           *restores.RestoresClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	backupEnginesClient := backupengines.NewBackupEnginesClientWithBaseURI(endpoint)
	configureAuthFunc(&backupEnginesClient.Client)

	backupJobsClient := backupjobs.NewBackupJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupJobsClient.Client)

	backupPoliciesClient := backuppolicies.NewBackupPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&backupPoliciesClient.Client)

	backupResourceEncryptionConfigsClient := backupresourceencryptionconfigs.NewBackupResourceEncryptionConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupResourceEncryptionConfigsClient.Client)

	backupResourceStorageConfigsNonCRRClient := backupresourcestorageconfigsnoncrr.NewBackupResourceStorageConfigsNonCRRClientWithBaseURI(endpoint)
	configureAuthFunc(&backupResourceStorageConfigsNonCRRClient.Client)

	backupResourceVaultConfigsClient := backupresourcevaultconfigs.NewBackupResourceVaultConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupResourceVaultConfigsClient.Client)

	backupWorkloadItemsClient := backupworkloaditems.NewBackupWorkloadItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupWorkloadItemsClient.Client)

	backupsClient := backups.NewBackupsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupsClient.Client)

	bmsClient := bms.NewBmsClientWithBaseURI(endpoint)
	configureAuthFunc(&bmsClient.Client)

	itemLevelRecoveryConnectionsClient := itemlevelrecoveryconnections.NewItemLevelRecoveryConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&itemLevelRecoveryConnectionsClient.Client)

	jobCancellationsClient := jobcancellations.NewJobCancellationsClientWithBaseURI(endpoint)
	configureAuthFunc(&jobCancellationsClient.Client)

	jobDetailsClient := jobdetails.NewJobDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&jobDetailsClient.Client)

	privateEndpointConnectionResourcesClient := privateendpointconnectionresources.NewPrivateEndpointConnectionResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionResourcesClient.Client)

	protectedItemsClient := protecteditems.NewProtectedItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&protectedItemsClient.Client)

	protectionContainersClient := protectioncontainers.NewProtectionContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&protectionContainersClient.Client)

	protectionIntentResourcesClient := protectionintentresources.NewProtectionIntentResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&protectionIntentResourcesClient.Client)

	protectionPoliciesClient := protectionpolicies.NewProtectionPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&protectionPoliciesClient.Client)

	recoveryPointsClient := recoverypoints.NewRecoveryPointsClientWithBaseURI(endpoint)
	configureAuthFunc(&recoveryPointsClient.Client)

	recoveryPointsRecommendedForMoveClient := recoverypointsrecommendedformove.NewRecoveryPointsRecommendedForMoveClientWithBaseURI(endpoint)
	configureAuthFunc(&recoveryPointsRecommendedForMoveClient.Client)

	resourceGuardProxiesClient := resourceguardproxies.NewResourceGuardProxiesClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceGuardProxiesClient.Client)

	resourceGuardProxyClient := resourceguardproxy.NewResourceGuardProxyClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceGuardProxyClient.Client)

	restoresClient := restores.NewRestoresClientWithBaseURI(endpoint)
	configureAuthFunc(&restoresClient.Client)

	return Client{
		BackupEngines:                      &backupEnginesClient,
		BackupJobs:                         &backupJobsClient,
		BackupPolicies:                     &backupPoliciesClient,
		BackupResourceEncryptionConfigs:    &backupResourceEncryptionConfigsClient,
		BackupResourceStorageConfigsNonCRR: &backupResourceStorageConfigsNonCRRClient,
		BackupResourceVaultConfigs:         &backupResourceVaultConfigsClient,
		BackupWorkloadItems:                &backupWorkloadItemsClient,
		Backups:                            &backupsClient,
		Bms:                                &bmsClient,
		ItemLevelRecoveryConnections:       &itemLevelRecoveryConnectionsClient,
		JobCancellations:                   &jobCancellationsClient,
		JobDetails:                         &jobDetailsClient,
		PrivateEndpointConnectionResources: &privateEndpointConnectionResourcesClient,
		ProtectedItems:                     &protectedItemsClient,
		ProtectionContainers:               &protectionContainersClient,
		ProtectionIntentResources:          &protectionIntentResourcesClient,
		ProtectionPolicies:                 &protectionPoliciesClient,
		RecoveryPoints:                     &recoveryPointsClient,
		RecoveryPointsRecommendedForMove:   &recoveryPointsRecommendedForMoveClient,
		ResourceGuardProxies:               &resourceGuardProxiesClient,
		ResourceGuardProxy:                 &resourceGuardProxyClient,
		Restores:                           &restoresClient,
	}
}
