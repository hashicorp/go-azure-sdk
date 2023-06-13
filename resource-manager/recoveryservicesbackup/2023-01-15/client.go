package v2023_01_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/aadproperties"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupcrrjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupprotecteditemscrr"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupresourcestorageconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/backupusagesummariescrr"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/crossregionrestore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/crrjobdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/recoverypointscrr"
	"github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/recoverypointsgetaccesstoken"
)

type Client struct {
	AadProperties                *aadproperties.AadPropertiesClient
	BackupCrrJobs                *backupcrrjobs.BackupCrrJobsClient
	BackupProtectedItemsCrr      *backupprotecteditemscrr.BackupProtectedItemsCrrClient
	BackupResourceStorageConfigs *backupresourcestorageconfigs.BackupResourceStorageConfigsClient
	BackupUsageSummariesCRR      *backupusagesummariescrr.BackupUsageSummariesCRRClient
	CrossRegionRestore           *crossregionrestore.CrossRegionRestoreClient
	CrrJobDetails                *crrjobdetails.CrrJobDetailsClient
	RecoveryPointsCrr            *recoverypointscrr.RecoveryPointsCrrClient
	RecoveryPointsGetAccessToken *recoverypointsgetaccesstoken.RecoveryPointsGetAccessTokenClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	aadPropertiesClient := aadproperties.NewAadPropertiesClientWithBaseURI(endpoint)
	configureAuthFunc(&aadPropertiesClient.Client)

	backupCrrJobsClient := backupcrrjobs.NewBackupCrrJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupCrrJobsClient.Client)

	backupProtectedItemsCrrClient := backupprotecteditemscrr.NewBackupProtectedItemsCrrClientWithBaseURI(endpoint)
	configureAuthFunc(&backupProtectedItemsCrrClient.Client)

	backupResourceStorageConfigsClient := backupresourcestorageconfigs.NewBackupResourceStorageConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&backupResourceStorageConfigsClient.Client)

	backupUsageSummariesCRRClient := backupusagesummariescrr.NewBackupUsageSummariesCRRClientWithBaseURI(endpoint)
	configureAuthFunc(&backupUsageSummariesCRRClient.Client)

	crossRegionRestoreClient := crossregionrestore.NewCrossRegionRestoreClientWithBaseURI(endpoint)
	configureAuthFunc(&crossRegionRestoreClient.Client)

	crrJobDetailsClient := crrjobdetails.NewCrrJobDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&crrJobDetailsClient.Client)

	recoveryPointsCrrClient := recoverypointscrr.NewRecoveryPointsCrrClientWithBaseURI(endpoint)
	configureAuthFunc(&recoveryPointsCrrClient.Client)

	recoveryPointsGetAccessTokenClient := recoverypointsgetaccesstoken.NewRecoveryPointsGetAccessTokenClientWithBaseURI(endpoint)
	configureAuthFunc(&recoveryPointsGetAccessTokenClient.Client)

	return Client{
		AadProperties:                &aadPropertiesClient,
		BackupCrrJobs:                &backupCrrJobsClient,
		BackupProtectedItemsCrr:      &backupProtectedItemsCrrClient,
		BackupResourceStorageConfigs: &backupResourceStorageConfigsClient,
		BackupUsageSummariesCRR:      &backupUsageSummariesCRRClient,
		CrossRegionRestore:           &crossRegionRestoreClient,
		CrrJobDetails:                &crrJobDetailsClient,
		RecoveryPointsCrr:            &recoveryPointsCrrClient,
		RecoveryPointsGetAccessToken: &recoveryPointsGetAccessTokenClient,
	}
}
