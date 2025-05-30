package recoverypoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPoint interface {
	RecoveryPoint() BaseRecoveryPointImpl
}

var _ RecoveryPoint = BaseRecoveryPointImpl{}

type BaseRecoveryPointImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseRecoveryPointImpl) RecoveryPoint() BaseRecoveryPointImpl {
	return s
}

var _ RecoveryPoint = RawRecoveryPointImpl{}

// RawRecoveryPointImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRecoveryPointImpl struct {
	recoveryPoint BaseRecoveryPointImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawRecoveryPointImpl) RecoveryPoint() BaseRecoveryPointImpl {
	return s.recoveryPoint
}

func UnmarshalRecoveryPointImplementation(input []byte) (RecoveryPoint, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryPoint into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
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

	if strings.EqualFold(value, "AzureWorkloadSAPAsePointInTimeRecoveryPoint") {
		var out AzureWorkloadSAPAsePointInTimeRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPAsePointInTimeRecoveryPoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadSAPAseRecoveryPoint") {
		var out AzureWorkloadSAPAseRecoveryPoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadSAPAseRecoveryPoint: %+v", err)
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

	var parent BaseRecoveryPointImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRecoveryPointImpl: %+v", err)
	}

	return RawRecoveryPointImpl{
		recoveryPoint: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
