package v2025_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupengines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupprotectableitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupprotecteditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupprotectioncontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupprotectionintent"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupresourceencryptionconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupresourcestorageconfigsnoncrr"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupresourcevaultconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupusagesummaries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backupworkloaditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/featuresupport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/fetchtieringcost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/itemlevelrecoveryconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/jobcancellations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/jobdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/operation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/privateendpointconnectionresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/protectablecontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/protecteditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/protectioncontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/protectionintent"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/protectionintentresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/protectionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/recoverypoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/recoverypointsrecommendedformove"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/resourceguardproxies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/resourceguardproxy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/restores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/securitypins"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/softdeletedcontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/validateoperation"
)

type Client struct {
	BackupEngines                      *backupengines.BackupEnginesClient
	BackupJobs                         *backupjobs.BackupJobsClient
	BackupPolicies                     *backuppolicies.BackupPoliciesClient
	BackupProtectableItems             *backupprotectableitems.BackupProtectableItemsClient
	BackupProtectedItems               *backupprotecteditems.BackupProtectedItemsClient
	BackupProtectionContainers         *backupprotectioncontainers.BackupProtectionContainersClient
	BackupProtectionIntent             *backupprotectionintent.BackupProtectionIntentClient
	BackupResourceEncryptionConfigs    *backupresourceencryptionconfigs.BackupResourceEncryptionConfigsClient
	BackupResourceStorageConfigsNonCRR *backupresourcestorageconfigsnoncrr.BackupResourceStorageConfigsNonCRRClient
	BackupResourceVaultConfigs         *backupresourcevaultconfigs.BackupResourceVaultConfigsClient
	BackupStatus                       *backupstatus.BackupStatusClient
	BackupUsageSummaries               *backupusagesummaries.BackupUsageSummariesClient
	BackupWorkloadItems                *backupworkloaditems.BackupWorkloadItemsClient
	Backups                            *backups.BackupsClient
	FeatureSupport                     *featuresupport.FeatureSupportClient
	FetchTieringCost                   *fetchtieringcost.FetchTieringCostClient
	ItemLevelRecoveryConnections       *itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient
	JobCancellations                   *jobcancellations.JobCancellationsClient
	JobDetails                         *jobdetails.JobDetailsClient
	Jobs                               *jobs.JobsClient
	Operation                          *operation.OperationClient
	PrivateEndpointConnectionResources *privateendpointconnectionresources.PrivateEndpointConnectionResourcesClient
	ProtectableContainers              *protectablecontainers.ProtectableContainersClient
	ProtectedItems                     *protecteditems.ProtectedItemsClient
	ProtectionContainers               *protectioncontainers.ProtectionContainersClient
	ProtectionIntent                   *protectionintent.ProtectionIntentClient
	ProtectionIntentResources          *protectionintentresources.ProtectionIntentResourcesClient
	ProtectionPolicies                 *protectionpolicies.ProtectionPoliciesClient
	RecoveryPoints                     *recoverypoints.RecoveryPointsClient
	RecoveryPointsRecommendedForMove   *recoverypointsrecommendedformove.RecoveryPointsRecommendedForMoveClient
	ResourceGuardProxies               *resourceguardproxies.ResourceGuardProxiesClient
	ResourceGuardProxy                 *resourceguardproxy.ResourceGuardProxyClient
	Restores                           *restores.RestoresClient
	SecurityPINs                       *securitypins.SecurityPINsClient
	SoftDeletedContainers              *softdeletedcontainers.SoftDeletedContainersClient
	ValidateOperation                  *validateoperation.ValidateOperationClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	backupEnginesClient := backupengines.NewBackupEnginesClientWithBaseURI(endpoint)
	configureAuthFunc(&backupEnginesClient.Client)

	backupJobsClient := backupjobs.NewBackupJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupJobsClient.Client)

	backupPoliciesClient := backuppolicies.NewBackupPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&backupPoliciesClient.Client)

	backupProtectableItemsClient := backupprotectableitems.NewBackupProtectableItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupProtectableItemsClient.Client)

	backupProtectedItemsClient := backupprotecteditems.NewBackupProtectedItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupProtectedItemsClient.Client)

	backupProtectionContainersClient := backupprotectioncontainers.NewBackupProtectionContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&backupProtectionContainersClient.Client)

	backupProtectionIntentClient := backupprotectionintent.NewBackupProtectionIntentClientWithBaseURI(endpoint)
	configureAuthFunc(&backupProtectionIntentClient.Client)

	backupResourceEncryptionConfigsClient := backupresourceencryptionconfigs.NewBackupResourceEncryptionConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupResourceEncryptionConfigsClient.Client)

	backupResourceStorageConfigsNonCRRClient := backupresourcestorageconfigsnoncrr.NewBackupResourceStorageConfigsNonCRRClientWithBaseURI(endpoint)
	configureAuthFunc(&backupResourceStorageConfigsNonCRRClient.Client)

	backupResourceVaultConfigsClient := backupresourcevaultconfigs.NewBackupResourceVaultConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupResourceVaultConfigsClient.Client)

	backupStatusClient := backupstatus.NewBackupStatusClientWithBaseURI(endpoint)
	configureAuthFunc(&backupStatusClient.Client)

	backupUsageSummariesClient := backupusagesummaries.NewBackupUsageSummariesClientWithBaseURI(endpoint)
	configureAuthFunc(&backupUsageSummariesClient.Client)

	backupWorkloadItemsClient := backupworkloaditems.NewBackupWorkloadItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupWorkloadItemsClient.Client)

	backupsClient := backups.NewBackupsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupsClient.Client)

	featureSupportClient := featuresupport.NewFeatureSupportClientWithBaseURI(endpoint)
	configureAuthFunc(&featureSupportClient.Client)

	fetchTieringCostClient := fetchtieringcost.NewFetchTieringCostClientWithBaseURI(endpoint)
	configureAuthFunc(&fetchTieringCostClient.Client)

	itemLevelRecoveryConnectionsClient := itemlevelrecoveryconnections.NewItemLevelRecoveryConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&itemLevelRecoveryConnectionsClient.Client)

	jobCancellationsClient := jobcancellations.NewJobCancellationsClientWithBaseURI(endpoint)
	configureAuthFunc(&jobCancellationsClient.Client)

	jobDetailsClient := jobdetails.NewJobDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&jobDetailsClient.Client)

	jobsClient := jobs.NewJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&jobsClient.Client)

	operationClient := operation.NewOperationClientWithBaseURI(endpoint)
	configureAuthFunc(&operationClient.Client)

	privateEndpointConnectionResourcesClient := privateendpointconnectionresources.NewPrivateEndpointConnectionResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionResourcesClient.Client)

	protectableContainersClient := protectablecontainers.NewProtectableContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&protectableContainersClient.Client)

	protectedItemsClient := protecteditems.NewProtectedItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&protectedItemsClient.Client)

	protectionContainersClient := protectioncontainers.NewProtectionContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&protectionContainersClient.Client)

	protectionIntentClient := protectionintent.NewProtectionIntentClientWithBaseURI(endpoint)
	configureAuthFunc(&protectionIntentClient.Client)

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

	securityPINsClient := securitypins.NewSecurityPINsClientWithBaseURI(endpoint)
	configureAuthFunc(&securityPINsClient.Client)

	softDeletedContainersClient := softdeletedcontainers.NewSoftDeletedContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&softDeletedContainersClient.Client)

	validateOperationClient := validateoperation.NewValidateOperationClientWithBaseURI(endpoint)
	configureAuthFunc(&validateOperationClient.Client)

	return Client{
		BackupEngines:                      &backupEnginesClient,
		BackupJobs:                         &backupJobsClient,
		BackupPolicies:                     &backupPoliciesClient,
		BackupProtectableItems:             &backupProtectableItemsClient,
		BackupProtectedItems:               &backupProtectedItemsClient,
		BackupProtectionContainers:         &backupProtectionContainersClient,
		BackupProtectionIntent:             &backupProtectionIntentClient,
		BackupResourceEncryptionConfigs:    &backupResourceEncryptionConfigsClient,
		BackupResourceStorageConfigsNonCRR: &backupResourceStorageConfigsNonCRRClient,
		BackupResourceVaultConfigs:         &backupResourceVaultConfigsClient,
		BackupStatus:                       &backupStatusClient,
		BackupUsageSummaries:               &backupUsageSummariesClient,
		BackupWorkloadItems:                &backupWorkloadItemsClient,
		Backups:                            &backupsClient,
		FeatureSupport:                     &featureSupportClient,
		FetchTieringCost:                   &fetchTieringCostClient,
		ItemLevelRecoveryConnections:       &itemLevelRecoveryConnectionsClient,
		JobCancellations:                   &jobCancellationsClient,
		JobDetails:                         &jobDetailsClient,
		Jobs:                               &jobsClient,
		Operation:                          &operationClient,
		PrivateEndpointConnectionResources: &privateEndpointConnectionResourcesClient,
		ProtectableContainers:              &protectableContainersClient,
		ProtectedItems:                     &protectedItemsClient,
		ProtectionContainers:               &protectionContainersClient,
		ProtectionIntent:                   &protectionIntentClient,
		ProtectionIntentResources:          &protectionIntentResourcesClient,
		ProtectionPolicies:                 &protectionPoliciesClient,
		RecoveryPoints:                     &recoveryPointsClient,
		RecoveryPointsRecommendedForMove:   &recoveryPointsRecommendedForMoveClient,
		ResourceGuardProxies:               &resourceGuardProxiesClient,
		ResourceGuardProxy:                 &resourceGuardProxyClient,
		Restores:                           &restoresClient,
		SecurityPINs:                       &securityPINsClient,
		SoftDeletedContainers:              &softDeletedContainersClient,
		ValidateOperation:                  &validateOperationClient,
	}
}
