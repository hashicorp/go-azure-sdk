package recoverypointsrecommendedformove

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPoint interface {
}

func unmarshalRecoveryPointImplementation(input []byte) (RecoveryPoint, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryPoint into map[string]interface: %+v", err)
	}

	value, ok := temp["objectType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureFileShareRecoveryPoint") {
		var out AzureFileShareRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureFileShareRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadPointInTimeRecoveryPoint") {
		var out AzureWorkloadPointInTimeRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadPointInTimeRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadRecoveryPoint") {
		var out AzureWorkloadRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSAPHanaPointInTimeRecoveryPoint") {
		var out AzureWorkloadSAPHanaPointInTimeRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPHanaPointInTimeRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSAPHanaRecoveryPoint") {
		var out AzureWorkloadSAPHanaRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPHanaRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSQLPointInTimeRecoveryPoint") {
		var out AzureWorkloadSQLPointInTimeRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSQLPointInTimeRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSQLRecoveryPoint") {
		var out AzureWorkloadSQLRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSQLRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GenericRecoveryPoint") {
		var out GenericRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GenericRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IaasVMRecoveryPoint") {
		var out IaasVMRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IaasVMRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	type RawRecoveryPointImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawRecoveryPointImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
