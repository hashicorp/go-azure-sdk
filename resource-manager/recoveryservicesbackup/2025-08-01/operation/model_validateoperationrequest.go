package operation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateOperationRequest interface {
	ValidateOperationRequest() BaseValidateOperationRequestImpl
}

var _ ValidateOperationRequest = BaseValidateOperationRequestImpl{}

type BaseValidateOperationRequestImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseValidateOperationRequestImpl) ValidateOperationRequest() BaseValidateOperationRequestImpl {
	return s
}

var _ ValidateOperationRequest = RawValidateOperationRequestImpl{}

// RawValidateOperationRequestImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawValidateOperationRequestImpl struct {
	validateOperationRequest BaseValidateOperationRequestImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawValidateOperationRequestImpl) ValidateOperationRequest() BaseValidateOperationRequestImpl {
	return s.validateOperationRequest
}

func (s RawValidateOperationRequestImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalValidateOperationRequestImplementation(input []byte) (ValidateOperationRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidateOperationRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ValidateIaasVMRestoreOperationRequest") {
		var out ValidateIaasVMRestoreOperationRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ValidateIaasVMRestoreOperationRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ValidateRestoreOperationRequest") {
		var out ValidateRestoreOperationRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ValidateRestoreOperationRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseValidateOperationRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseValidateOperationRequestImpl: %+v", err)
	}

	return RawValidateOperationRequestImpl{
		validateOperationRequest: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
