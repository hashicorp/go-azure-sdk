package backupworkloaditems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadItem interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWorkloadItemImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalWorkloadItemImplementation(input []byte) (WorkloadItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkloadItem into map[string]interface: %+v", err)
	}

	value, ok := temp["workloadItemType"].(string)
	if !ok {
		return nil, nil
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

	out := RawWorkloadItemImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
