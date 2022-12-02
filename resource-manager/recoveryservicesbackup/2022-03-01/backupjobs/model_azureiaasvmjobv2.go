package backupjobs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
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

func (o *AzureIaaSVMJobV2) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureIaaSVMJobV2) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *AzureIaaSVMJobV2) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureIaaSVMJobV2) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
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
