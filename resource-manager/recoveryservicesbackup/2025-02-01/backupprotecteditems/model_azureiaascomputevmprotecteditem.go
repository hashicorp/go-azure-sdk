package backupprotecteditems

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectedItem = AzureIaaSComputeVMProtectedItem{}

type AzureIaaSComputeVMProtectedItem struct {
	ExtendedInfo        *AzureIaaSVMProtectedItemExtendedInfo `json:"extendedInfo,omitempty"`
	ExtendedProperties  *ExtendedProperties                   `json:"extendedProperties,omitempty"`
	FriendlyName        *string                               `json:"friendlyName,omitempty"`
	HealthDetails       *[]ResourceHealthDetails              `json:"healthDetails,omitempty"`
	HealthStatus        *HealthStatus                         `json:"healthStatus,omitempty"`
	KpisHealths         *map[string]KPIResourceHealthDetails  `json:"kpisHealths,omitempty"`
	LastBackupStatus    *string                               `json:"lastBackupStatus,omitempty"`
	LastBackupTime      *string                               `json:"lastBackupTime,omitempty"`
	PolicyType          *string                               `json:"policyType,omitempty"`
	ProtectedItemDataId *string                               `json:"protectedItemDataId,omitempty"`
	ProtectionState     *ProtectionState                      `json:"protectionState,omitempty"`
	ProtectionStatus    *string                               `json:"protectionStatus,omitempty"`
	VirtualMachineId    *string                               `json:"virtualMachineId,omitempty"`

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
	ProtectedItemType                string                `json:"protectedItemType"`
	ResourceGuardOperationRequests   *[]string             `json:"resourceGuardOperationRequests,omitempty"`
	SoftDeleteRetentionPeriodInDays  *int64                `json:"softDeleteRetentionPeriodInDays,omitempty"`
	SourceResourceId                 *string               `json:"sourceResourceId,omitempty"`
	VaultId                          *string               `json:"vaultId,omitempty"`
	WorkloadType                     *DataSourceType       `json:"workloadType,omitempty"`
}

func (s AzureIaaSComputeVMProtectedItem) ProtectedItem() BaseProtectedItemImpl {
	return BaseProtectedItemImpl{
		BackupManagementType:             s.BackupManagementType,
		BackupSetName:                    s.BackupSetName,
		ContainerName:                    s.ContainerName,
		CreateMode:                       s.CreateMode,
		DeferredDeleteTimeInUTC:          s.DeferredDeleteTimeInUTC,
		DeferredDeleteTimeRemaining:      s.DeferredDeleteTimeRemaining,
		IsArchiveEnabled:                 s.IsArchiveEnabled,
		IsDeferredDeleteScheduleUpcoming: s.IsDeferredDeleteScheduleUpcoming,
		IsRehydrate:                      s.IsRehydrate,
		IsScheduledForDeferredDelete:     s.IsScheduledForDeferredDelete,
		LastRecoveryPoint:                s.LastRecoveryPoint,
		PolicyId:                         s.PolicyId,
		PolicyName:                       s.PolicyName,
		ProtectedItemType:                s.ProtectedItemType,
		ResourceGuardOperationRequests:   s.ResourceGuardOperationRequests,
		SoftDeleteRetentionPeriodInDays:  s.SoftDeleteRetentionPeriodInDays,
		SourceResourceId:                 s.SourceResourceId,
		VaultId:                          s.VaultId,
		WorkloadType:                     s.WorkloadType,
	}
}

func (o *AzureIaaSComputeVMProtectedItem) GetDeferredDeleteTimeInUTCAsTime() (*time.Time, error) {
	if o.DeferredDeleteTimeInUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeferredDeleteTimeInUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureIaaSComputeVMProtectedItem) SetDeferredDeleteTimeInUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeferredDeleteTimeInUTC = &formatted
}

func (o *AzureIaaSComputeVMProtectedItem) GetLastRecoveryPointAsTime() (*time.Time, error) {
	if o.LastRecoveryPoint == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRecoveryPoint, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureIaaSComputeVMProtectedItem) SetLastRecoveryPointAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRecoveryPoint = &formatted
}

var _ json.Marshaler = AzureIaaSComputeVMProtectedItem{}

func (s AzureIaaSComputeVMProtectedItem) MarshalJSON() ([]byte, error) {
	type wrapper AzureIaaSComputeVMProtectedItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureIaaSComputeVMProtectedItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureIaaSComputeVMProtectedItem: %+v", err)
	}

	decoded["protectedItemType"] = "Microsoft.Compute/virtualMachines"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureIaaSComputeVMProtectedItem: %+v", err)
	}

	return encoded, nil
}
