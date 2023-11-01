package backupworkloaditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadItem = AzureVMWorkloadSAPAseDatabaseWorkloadItem{}

type AzureVMWorkloadSAPAseDatabaseWorkloadItem struct {
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

var _ json.Marshaler = AzureVMWorkloadSAPAseDatabaseWorkloadItem{}

func (s AzureVMWorkloadSAPAseDatabaseWorkloadItem) MarshalJSON() ([]byte, error) {
	type wrapper AzureVMWorkloadSAPAseDatabaseWorkloadItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureVMWorkloadSAPAseDatabaseWorkloadItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureVMWorkloadSAPAseDatabaseWorkloadItem: %+v", err)
	}
	decoded["workloadItemType"] = "SAPAseDatabase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureVMWorkloadSAPAseDatabaseWorkloadItem: %+v", err)
	}

	return encoded, nil
}
