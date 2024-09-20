package backupworkloaditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadItem = AzureVMWorkloadSAPAseSystemWorkloadItem{}

type AzureVMWorkloadSAPAseSystemWorkloadItem struct {
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

func (s AzureVMWorkloadSAPAseSystemWorkloadItem) WorkloadItem() BaseWorkloadItemImpl {
	return BaseWorkloadItemImpl{
		BackupManagementType: s.BackupManagementType,
		FriendlyName:         s.FriendlyName,
		ProtectionState:      s.ProtectionState,
		WorkloadItemType:     s.WorkloadItemType,
		WorkloadType:         s.WorkloadType,
	}
}

var _ json.Marshaler = AzureVMWorkloadSAPAseSystemWorkloadItem{}

func (s AzureVMWorkloadSAPAseSystemWorkloadItem) MarshalJSON() ([]byte, error) {
	type wrapper AzureVMWorkloadSAPAseSystemWorkloadItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureVMWorkloadSAPAseSystemWorkloadItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureVMWorkloadSAPAseSystemWorkloadItem: %+v", err)
	}

	decoded["workloadItemType"] = "SAPAseSystem"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureVMWorkloadSAPAseSystemWorkloadItem: %+v", err)
	}

	return encoded, nil
}
