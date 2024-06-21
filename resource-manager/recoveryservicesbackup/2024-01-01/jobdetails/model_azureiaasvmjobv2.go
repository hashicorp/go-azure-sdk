package jobdetails

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Job = AzureIaaSVMJobV2{}

type AzureIaaSVMJobV2 struct {
	ActionsInfo           *[]JobSupportedAction       `json:"actionsInfo,omitempty"`
	ContainerName         *string                     `json:"containerName,omitempty"`
	Duration              *string                     `json:"duration,omitempty"`
	ErrorDetails          *[]AzureIaaSVMErrorInfo     `json:"errorDetails,omitempty"`
	ExtendedInfo          *AzureIaaSVMJobExtendedInfo `json:"extendedInfo,omitempty"`
	VirtualMachineVersion *string                     `json:"virtualMachineVersion,omitempty"`

	// Fields inherited from Job
	ActivityId           *string               `json:"activityId,omitempty"`
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
	EntityFriendlyName   *string               `json:"entityFriendlyName,omitempty"`
	Operation            *string               `json:"operation,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

var _ json.Marshaler = AzureIaaSVMJobV2{}

func (s AzureIaaSVMJobV2) MarshalJSON() ([]byte, error) {
	type wrapper AzureIaaSVMJobV2
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureIaaSVMJobV2: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureIaaSVMJobV2: %+v", err)
	}
	decoded["jobType"] = "AzureIaaSVMJobV2"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureIaaSVMJobV2: %+v", err)
	}

	return encoded, nil
}
