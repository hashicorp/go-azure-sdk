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

	type RawWorkloadItemImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawWorkloadItemImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
