package operation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreRequest interface {
}

func unmarshalRestoreRequestImplementation(input []byte) (RestoreRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RestoreRequest into map[string]interface: %+v", err)
	}

	value, ok := temp["objectType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureFileShareRestoreRequest") {
		var out AzureFileShareRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureFileShareRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadPointInTimeRestoreRequest") {
		var out AzureWorkloadPointInTimeRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadPointInTimeRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadRestoreRequest") {
		var out AzureWorkloadRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSAPHanaPointInTimeRestoreRequest") {
		var out AzureWorkloadSAPHanaPointInTimeRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPHanaPointInTimeRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest") {
		var out AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSAPHanaRestoreRequest") {
		var out AzureWorkloadSAPHanaRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPHanaRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSAPHanaRestoreWithRehydrateRequest") {
		var out AzureWorkloadSAPHanaRestoreWithRehydrateRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPHanaRestoreWithRehydrateRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSQLPointInTimeRestoreRequest") {
		var out AzureWorkloadSQLPointInTimeRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSQLPointInTimeRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSQLPointInTimeRestoreWithRehydrateRequest") {
		var out AzureWorkloadSQLPointInTimeRestoreWithRehydrateRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSQLPointInTimeRestoreWithRehydrateRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSQLRestoreRequest") {
		var out AzureWorkloadSQLRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSQLRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSQLRestoreWithRehydrateRequest") {
		var out AzureWorkloadSQLRestoreWithRehydrateRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSQLRestoreWithRehydrateRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IaasVMRestoreRequest") {
		var out IaasVMRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IaasVMRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IaasVMRestoreWithRehydrationRequest") {
		var out IaasVMRestoreWithRehydrationRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IaasVMRestoreWithRehydrationRequest: %+v", err)
		}
		return out, nil
	}

	type RawRestoreRequestImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawRestoreRequestImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
