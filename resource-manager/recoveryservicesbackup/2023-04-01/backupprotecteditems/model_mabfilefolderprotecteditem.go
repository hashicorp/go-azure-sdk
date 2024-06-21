package backupprotecteditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectedItem = MabFileFolderProtectedItem{}

type MabFileFolderProtectedItem struct {
	ComputerName                *string                                 `json:"computerName,omitempty"`
	DeferredDeleteSyncTimeInUTC *int64                                  `json:"deferredDeleteSyncTimeInUTC,omitempty"`
	ExtendedInfo                *MabFileFolderProtectedItemExtendedInfo `json:"extendedInfo,omitempty"`
	FriendlyName                *string                                 `json:"friendlyName,omitempty"`
	LastBackupStatus            *string                                 `json:"lastBackupStatus,omitempty"`
	LastBackupTime              *string                                 `json:"lastBackupTime,omitempty"`
	ProtectionState             *string                                 `json:"protectionState,omitempty"`

	// Fields inherited from ProtectedItem
	BackupManagementType             *BackupManagementType `json:"backupManagementType,omitempty"`
	BackupSetName                    *string               `json:"backupSetName,omitempty"`
	ContainerName                    *string               `json:"containerName,omitempty"`
	CreateMode                       *CreateMode           `json:"createMode,omitempty"`
	DeferredDeleteTimeInUTC          *string               `json:"deferredDeleteTimeInUTC,omitempty"`
	DeferredDeleteTimeRemaining      *string               `json:"deferredDeleteTimeRemaining,omitempty"`
	IsArchiveEnabled                 *bool                 `json:"isArchiveEnabled,omitempty"`
	IsDeferredDeleteScheduleUpcoming *bool                 `json:"isDeferredDeleteScheduleUpcoming,omitempty"`
	IsRehydrate                      *bool                 `json:"isRehydrate,omitempty"`
	IsScheduledForDeferredDelete     *bool                 `json:"isScheduledForDeferredDelete,omitempty"`
	LastRecoveryPoint                *string               `json:"lastRecoveryPoint,omitempty"`
	PolicyId                         *string               `json:"policyId,omitempty"`
	PolicyName                       *string               `json:"policyName,omitempty"`
	ResourceGuardOperationRequests   *[]string             `json:"resourceGuardOperationRequests,omitempty"`
	SoftDeleteRetentionPeriodInDays  *int64                `json:"softDeleteRetentionPeriodInDays,omitempty"`
	SourceResourceId                 *string               `json:"sourceResourceId,omitempty"`
	VaultId                          *string               `json:"vaultId,omitempty"`
	WorkloadType                     *DataSourceType       `json:"workloadType,omitempty"`
}

var _ json.Marshaler = MabFileFolderProtectedItem{}

func (s MabFileFolderProtectedItem) MarshalJSON() ([]byte, error) {
	type wrapper MabFileFolderProtectedItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MabFileFolderProtectedItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MabFileFolderProtectedItem: %+v", err)
	}
	decoded["protectedItemType"] = "MabFileFolderProtectedItem"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MabFileFolderProtectedItem: %+v", err)
	}

	return encoded, nil
}
