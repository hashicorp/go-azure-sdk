package backupstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupStatusResponse struct {
	ContainerName      *string           `json:"containerName,omitempty"`
	ErrorCode          *string           `json:"errorCode,omitempty"`
	ErrorMessage       *string           `json:"errorMessage,omitempty"`
	FabricName         *FabricName       `json:"fabricName,omitempty"`
	PolicyName         *string           `json:"policyName,omitempty"`
	ProtectedItemName  *string           `json:"protectedItemName,omitempty"`
	ProtectionStatus   *ProtectionStatus `json:"protectionStatus,omitempty"`
	RegistrationStatus *string           `json:"registrationStatus,omitempty"`
	VaultId            *string           `json:"vaultId,omitempty"`
}
