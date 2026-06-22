package operation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreRequest interface {
	RestoreRequest() BaseRestoreRequestImpl
}

var _ RestoreRequest = BaseRestoreRequestImpl{}

type BaseRestoreRequestImpl struct {
	ObjectType                     string    `json:"objectType"`
	ResourceGuardOperationRequests *[]string `json:"resourceGuardOperationRequests,omitempty"`
}

func (s BaseRestoreRequestImpl) RestoreRequest() BaseRestoreRequestImpl {
	return s
}

var _ RestoreRequest = RawRestoreRequestImpl{}

// RawRestoreRequestImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawRestoreRequestImpl struct {
	restoreRequest BaseRestoreRequestImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawRestoreRequestImpl) RestoreRequest() BaseRestoreRequestImpl {
	return s.restoreRequest
}

func (s RawRestoreRequestImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalRestoreRequestImplementation(input []byte) (RestoreRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RestoreRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
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

	if strings.EqualFold(value, "AzureWorkloadSAPAsePointInTimeRestoreRequest") {
		var out AzureWorkloadSAPAsePointInTimeRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPAsePointInTimeRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSAPAseRestoreRequest") {
		var out AzureWorkloadSAPAseRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPAseRestoreRequest: %+v", err)
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

	var parent BaseRestoreRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRestoreRequestImpl: %+v", err)
	}

	return RawRestoreRequestImpl{
		restoreRequest: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
