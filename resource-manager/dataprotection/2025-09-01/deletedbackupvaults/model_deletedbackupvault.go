package deletedbackupvaults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedBackupVault struct {
	BcdrSecurityLevel               *BCDRSecurityLevel   `json:"bcdrSecurityLevel,omitempty"`
	FeatureSettings                 *FeatureSettings     `json:"featureSettings,omitempty"`
	IsVaultProtectedByResourceGuard *bool                `json:"isVaultProtectedByResourceGuard,omitempty"`
	MonitoringSettings              *MonitoringSettings  `json:"monitoringSettings,omitempty"`
	OriginalBackupVaultId           string               `json:"originalBackupVaultId"`
	OriginalBackupVaultName         string               `json:"originalBackupVaultName"`
	OriginalBackupVaultResourcePath string               `json:"originalBackupVaultResourcePath"`
	ProvisioningState               *ProvisioningState   `json:"provisioningState,omitempty"`
	ReplicatedRegions               *[]string            `json:"replicatedRegions,omitempty"`
	ResourceDeletionInfo            ResourceDeletionInfo `json:"resourceDeletionInfo"`
	ResourceGuardOperationRequests  *[]string            `json:"resourceGuardOperationRequests,omitempty"`
	ResourceMoveDetails             *ResourceMoveDetails `json:"resourceMoveDetails,omitempty"`
	ResourceMoveState               *ResourceMoveState   `json:"resourceMoveState,omitempty"`
	SecureScore                     *SecureScoreLevel    `json:"secureScore,omitempty"`
	SecuritySettings                *SecuritySettings    `json:"securitySettings,omitempty"`
	StorageSettings                 []StorageSetting     `json:"storageSettings"`
}
