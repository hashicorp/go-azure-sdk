package bms

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatusExtendedInfo interface {
	OperationStatusExtendedInfo() BaseOperationStatusExtendedInfoImpl
}

var _ OperationStatusExtendedInfo = BaseOperationStatusExtendedInfoImpl{}

type BaseOperationStatusExtendedInfoImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseOperationStatusExtendedInfoImpl) OperationStatusExtendedInfo() BaseOperationStatusExtendedInfoImpl {
	return s
}

var _ OperationStatusExtendedInfo = RawOperationStatusExtendedInfoImpl{}

// RawOperationStatusExtendedInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOperationStatusExtendedInfoImpl struct {
	operationStatusExtendedInfo BaseOperationStatusExtendedInfoImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawOperationStatusExtendedInfoImpl) OperationStatusExtendedInfo() BaseOperationStatusExtendedInfoImpl {
	return s.operationStatusExtendedInfo
}

func UnmarshalOperationStatusExtendedInfoImplementation(input []byte) (OperationStatusExtendedInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationStatusExtendedInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "OperationStatusJobExtendedInfo") {
		var out OperationStatusJobExtendedInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationStatusJobExtendedInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OperationStatusJobsExtendedInfo") {
		var out OperationStatusJobsExtendedInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationStatusJobsExtendedInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OperationStatusProvisionILRExtendedInfo") {
		var out OperationStatusProvisionILRExtendedInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationStatusProvisionILRExtendedInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OperationStatusValidateOperationExtendedInfo") {
		var out OperationStatusValidateOperationExtendedInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OperationStatusValidateOperationExtendedInfo: %+v", err)
		}
		return out, nil
	}

	var parent BaseOperationStatusExtendedInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOperationStatusExtendedInfoImpl: %+v", err)
	}

	return RawOperationStatusExtendedInfoImpl{
		operationStatusExtendedInfo: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
