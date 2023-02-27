// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_12_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupengines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupprotectableitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupprotecteditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupprotectioncontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupprotectionintent"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupresourceencryptionconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupresourcestorageconfigsnoncrr"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupresourcevaultconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupstatus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupusagesummaries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backupworkloaditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/datamove"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/featuresupport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/itemlevelrecoveryconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/jobcancellations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/jobdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/operation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/protectablecontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/protecteditems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/protectioncontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/protectionintent"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/protectionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/recoverypoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/recoverypoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/recoverypointsrecommendedformove"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/resourceguardproxies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/resourceguardproxy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/restores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/securitypins"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/validateoperation"
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
	DataMove                           *datamove.DataMoveClient
	FeatureSupport                     *featuresupport.FeatureSupportClient
	ItemLevelRecoveryConnections       *itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient
	JobCancellations                   *jobcancellations.JobCancellationsClient
	JobDetails                         *jobdetails.JobDetailsClient
	Jobs                               *jobs.JobsClient
	Operation                          *operation.OperationClient
	PrivateEndpointConnection          *privateendpointconnection.PrivateEndpointConnectionClient
	ProtectableContainers              *protectablecontainers.ProtectableContainersClient
	ProtectedItems                     *protecteditems.ProtectedItemsClient
	ProtectionContainers               *protectioncontainers.ProtectionContainersClient
	ProtectionIntent                   *protectionintent.ProtectionIntentClient
	ProtectionPolicies                 *protectionpolicies.ProtectionPoliciesClient
	RecoveryPoint                      *recoverypoint.RecoveryPointClient
	RecoveryPoints                     *recoverypoints.RecoveryPointsClient
	RecoveryPointsRecommendedForMove   *recoverypointsrecommendedformove.RecoveryPointsRecommendedForMoveClient
	ResourceGuardProxies               *resourceguardproxies.ResourceGuardProxiesClient
	ResourceGuardProxy                 *resourceguardproxy.ResourceGuardProxyClient
	Restores                           *restores.RestoresClient
	SecurityPINs                       *securitypins.SecurityPINsClient
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

	dataMoveClient := datamove.NewDataMoveClientWithBaseURI(endpoint)
	configureAuthFunc(&dataMoveClient.Client)

	featureSupportClient := featuresupport.NewFeatureSupportClientWithBaseURI(endpoint)
	configureAuthFunc(&featureSupportClient.Client)

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

	privateEndpointConnectionClient := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionClient.Client)

	protectableContainersClient := protectablecontainers.NewProtectableContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&protectableContainersClient.Client)

	protectedItemsClient := protecteditems.NewProtectedItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&protectedItemsClient.Client)

	protectionContainersClient := protectioncontainers.NewProtectionContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&protectionContainersClient.Client)

	protectionIntentClient := protectionintent.NewProtectionIntentClientWithBaseURI(endpoint)
	configureAuthFunc(&protectionIntentClient.Client)

	protectionPoliciesClient := protectionpolicies.NewProtectionPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&protectionPoliciesClient.Client)

	recoveryPointClient := recoverypoint.NewRecoveryPointClientWithBaseURI(endpoint)
	configureAuthFunc(&recoveryPointClient.Client)

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
		DataMove:                           &dataMoveClient,
		FeatureSupport:                     &featureSupportClient,
		ItemLevelRecoveryConnections:       &itemLevelRecoveryConnectionsClient,
		JobCancellations:                   &jobCancellationsClient,
		JobDetails:                         &jobDetailsClient,
		Jobs:                               &jobsClient,
		Operation:                          &operationClient,
		PrivateEndpointConnection:          &privateEndpointConnectionClient,
		ProtectableContainers:              &protectableContainersClient,
		ProtectedItems:                     &protectedItemsClient,
		ProtectionContainers:               &protectionContainersClient,
		ProtectionIntent:                   &protectionIntentClient,
		ProtectionPolicies:                 &protectionPoliciesClient,
		RecoveryPoint:                      &recoveryPointClient,
		RecoveryPoints:                     &recoveryPointsClient,
		RecoveryPointsRecommendedForMove:   &recoveryPointsRecommendedForMoveClient,
		ResourceGuardProxies:               &resourceGuardProxiesClient,
		ResourceGuardProxy:                 &resourceGuardProxyClient,
		Restores:                           &restoresClient,
		SecurityPINs:                       &securityPINsClient,
		ValidateOperation:                  &validateOperationClient,
	}
}
