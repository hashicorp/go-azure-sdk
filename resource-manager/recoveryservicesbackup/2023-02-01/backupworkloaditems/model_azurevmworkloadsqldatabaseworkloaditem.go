package backupworkloaditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadItem = AzureVMWorkloadSQLDatabaseWorkloadItem{}

type AzureVMWorkloadSQLDatabaseWorkloadItem struct {
	IsAutoProtectable    *bool   `json:"isAutoProtectable,omitempty"`
	ParentName           *string `json:"parentName,omitempty"`
	ServerName           *string `json:"serverName,omitempty"`
	SubWorkloadItemCount *int64  `json:"subWorkloadItemCount,omitempty"`
	Subinquireditemcount *int64  `json:"subinquireditemcount,omitempty"`

	// Fields inherited from WorkloadItem

	BackupManagementType *string           `json:"backupManagementType,omitempty"`
	FriendlyName         *string           `json:"friendlyName,omitempty"`
	ProtectionState      *ProtectionStatus `json:"protectionState,omitempty"`
	WorkloadItemType     string            `json:"workloadItemType"`
	WorkloadType         *string           `json:"workloadType,omitempty"`
}

func (s AzureVMWorkloadSQLDatabaseWorkloadItem) WorkloadItem() BaseWorkloadItemImpl {
	return BaseWorkloadItemImpl{
		BackupManagementType: s.BackupManagementType,
		FriendlyName:         s.FriendlyName,
		ProtectionState:      s.ProtectionState,
		WorkloadItemType:     s.WorkloadItemType,
		WorkloadType:         s.WorkloadType,
	}
}

var _ json.Marshaler = AzureVMWorkloadSQLDatabaseWorkloadItem{}

func (s AzureVMWorkloadSQLDatabaseWorkloadItem) MarshalJSON() ([]byte, error) {
	type wrapper AzureVMWorkloadSQLDatabaseWorkloadItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureVMWorkloadSQLDatabaseWorkloadItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureVMWorkloadSQLDatabaseWorkloadItem: %+v", err)
	}

	decoded["workloadItemType"] = "SQLDataBase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureVMWorkloadSQLDatabaseWorkloadItem: %+v", err)
	}

	return encoded, nil
}
