package backupworkloaditems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadItem interface {
	WorkloadItem() BaseWorkloadItemImpl
}

var _ WorkloadItem = BaseWorkloadItemImpl{}

type BaseWorkloadItemImpl struct {
	BackupManagementType *string           `json:"backupManagementType,omitempty"`
	FriendlyName         *string           `json:"friendlyName,omitempty"`
	ProtectionState      *ProtectionStatus `json:"protectionState,omitempty"`
	WorkloadItemType     string            `json:"workloadItemType"`
	WorkloadType         *string           `json:"workloadType,omitempty"`
}

func (s BaseWorkloadItemImpl) WorkloadItem() BaseWorkloadItemImpl {
	return s
}

var _ WorkloadItem = RawWorkloadItemImpl{}

// RawWorkloadItemImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawWorkloadItemImpl struct {
	workloadItem BaseWorkloadItemImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawWorkloadItemImpl) WorkloadItem() BaseWorkloadItemImpl {
	return s.workloadItem
}

func (s RawWorkloadItemImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalWorkloadItemImplementation(input []byte) (WorkloadItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["workloadItemType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureVmWorkloadItem") {
		var out AzureVMWorkloadItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMWorkloadItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SAPAseDatabase") {
		var out AzureVMWorkloadSAPAseDatabaseWorkloadItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMWorkloadSAPAseDatabaseWorkloadItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SAPAseSystem") {
		var out AzureVMWorkloadSAPAseSystemWorkloadItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMWorkloadSAPAseSystemWorkloadItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SAPHanaDatabase") {
		var out AzureVMWorkloadSAPHanaDatabaseWorkloadItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMWorkloadSAPHanaDatabaseWorkloadItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SAPHanaSystem") {
		var out AzureVMWorkloadSAPHanaSystemWorkloadItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMWorkloadSAPHanaSystemWorkloadItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SQLDataBase") {
		var out AzureVMWorkloadSQLDatabaseWorkloadItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMWorkloadSQLDatabaseWorkloadItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SQLInstance") {
		var out AzureVMWorkloadSQLInstanceWorkloadItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMWorkloadSQLInstanceWorkloadItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseWorkloadItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWorkloadItemImpl: %+v", err)
	}

	return RawWorkloadItemImpl{
		workloadItem: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
