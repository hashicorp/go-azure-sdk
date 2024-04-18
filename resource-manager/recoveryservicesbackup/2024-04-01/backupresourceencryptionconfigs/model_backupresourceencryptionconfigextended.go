package backupresourceencryptionconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupResourceEncryptionConfigExtended struct {
	EncryptionAtRestType          *EncryptionAtRestType          `json:"encryptionAtRestType,omitempty"`
	InfrastructureEncryptionState *InfrastructureEncryptionState `json:"infrastructureEncryptionState,omitempty"`
	KeyUri                        *string                        `json:"keyUri,omitempty"`
	LastUpdateStatus              *LastUpdateStatus              `json:"lastUpdateStatus,omitempty"`
	SubscriptionId                *string                        `json:"subscriptionId,omitempty"`
	UseSystemAssignedIdentity     *bool                          `json:"useSystemAssignedIdentity,omitempty"`
	UserAssignedIdentity          *string                        `json:"userAssignedIdentity,omitempty"`
}
