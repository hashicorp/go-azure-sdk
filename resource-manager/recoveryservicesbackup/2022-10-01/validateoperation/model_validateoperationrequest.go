package validateoperation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateOperationRequest interface {
}

func unmarshalValidateOperationRequestImplementation(input []byte) (ValidateOperationRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidateOperationRequest into map[string]interface: %+v", err)
	}

	value, ok := temp["objectType"].(string)
	if !ok {
		return nil, nil
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

	type RawValidateOperationRequestImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawValidateOperationRequestImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
