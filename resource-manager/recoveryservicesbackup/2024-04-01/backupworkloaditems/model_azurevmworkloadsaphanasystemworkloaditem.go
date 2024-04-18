package backupworkloaditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadItem = AzureVMWorkloadSAPHanaSystemWorkloadItem{}

type AzureVMWorkloadSAPHanaSystemWorkloadItem struct {
	IsAutoProtectable    *bool   `json:"isAutoProtectable,omitempty"`
	ParentName           *string `json:"parentName,omitempty"`
	ServerName           *string `json:"serverName,omitempty"`
	SubWorkloadItemCount *int64  `json:"subWorkloadItemCount,omitempty"`
	Subinquireditemcount *int64  `json:"subinquireditemcount,omitempty"`

	// Fields inherited from WorkloadItem
	BackupManagementType *string           `json:"backupManagementType,omitempty"`
	FriendlyName         *string           `json:"friendlyName,omitempty"`
	ProtectionState      *ProtectionStatus `json:"protectionState,omitempty"`
	WorkloadType         *string           `json:"workloadType,omitempty"`
}

var _ json.Marshaler = AzureVMWorkloadSAPHanaSystemWorkloadItem{}

func (s AzureVMWorkloadSAPHanaSystemWorkloadItem) MarshalJSON() ([]byte, error) {
	type wrapper AzureVMWorkloadSAPHanaSystemWorkloadItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureVMWorkloadSAPHanaSystemWorkloadItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureVMWorkloadSAPHanaSystemWorkloadItem: %+v", err)
	}
	decoded["workloadItemType"] = "SAPHanaSystem"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureVMWorkloadSAPHanaSystemWorkloadItem: %+v", err)
	}

	return encoded, nil
}
