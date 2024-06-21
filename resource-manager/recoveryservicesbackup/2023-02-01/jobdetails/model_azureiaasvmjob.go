package jobdetails

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Job = AzureIaaSVMJob{}

type AzureIaaSVMJob struct {
	ActionsInfo           *[]JobSupportedAction       `json:"actionsInfo,omitempty"`
	ContainerName         *string                     `json:"containerName,omitempty"`
	Duration              *string                     `json:"duration,omitempty"`
	ErrorDetails          *[]AzureIaaSVMErrorInfo     `json:"errorDetails,omitempty"`
	ExtendedInfo          *AzureIaaSVMJobExtendedInfo `json:"extendedInfo,omitempty"`
	IsUserTriggered       *bool                       `json:"isUserTriggered,omitempty"`
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

var _ json.Marshaler = AzureIaaSVMJob{}

func (s AzureIaaSVMJob) MarshalJSON() ([]byte, error) {
	type wrapper AzureIaaSVMJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureIaaSVMJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureIaaSVMJob: %+v", err)
	}
	decoded["jobType"] = "AzureIaaSVMJob"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureIaaSVMJob: %+v", err)
	}

	return encoded, nil
}
