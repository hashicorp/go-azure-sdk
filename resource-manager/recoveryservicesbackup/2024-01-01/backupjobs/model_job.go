package backupjobs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Job interface {
	Job() BaseJobImpl
}

var _ Job = BaseJobImpl{}

type BaseJobImpl struct {
	ActivityId           *string               `json:"activityId,omitempty"`
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
	EntityFriendlyName   *string               `json:"entityFriendlyName,omitempty"`
	JobType              string                `json:"jobType"`
	Operation            *string               `json:"operation,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

func (s BaseJobImpl) Job() BaseJobImpl {
	return s
}

var _ Job = RawJobImpl{}

// RawJobImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobImpl struct {
	job    BaseJobImpl
	Type   string
	Values map[string]interface{}
}

func (s RawJobImpl) Job() BaseJobImpl {
	return s.job
}

func UnmarshalJobImplementation(input []byte) (Job, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Job into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["jobType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureIaaSVMJob") {
		var out AzureIaaSVMJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureIaaSVMJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureIaaSVMJobV2") {
		var out AzureIaaSVMJobV2
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureIaaSVMJobV2: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureStorageJob") {
		var out AzureStorageJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureStorageJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadJob") {
		var out AzureWorkloadJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DpmJob") {
		var out DpmJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DpmJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MabJob") {
		var out MabJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MabJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VaultJob") {
		var out VaultJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VaultJob: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobImpl: %+v", err)
	}

	return RawJobImpl{
		job:    parent,
		Type:   value,
		Values: temp,
	}, nil

}
